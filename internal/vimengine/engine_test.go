package vimengine

import "testing"

func TestNewStateStartsInNormalMode(t *testing.T) {
	state := NewState([]string{"alpha"})

	if state.Mode != ModeNormal {
		t.Fatalf("Mode = %q, want %q", state.Mode, ModeNormal)
	}
	if state.Cursor.Row != 0 || state.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", state.Cursor.Row, state.Cursor.Col)
	}
}

func TestStateCopiesLines(t *testing.T) {
	lines := []string{"alpha"}
	engine := New(lines)
	lines[0] = "changed"

	state := engine.State()
	state.Lines[0] = "mutated"

	got := engine.State().Lines[0]
	if got != "alpha" {
		t.Fatalf("engine line = %q, want %q", got, "alpha")
	}
}

func TestNewWithStateCopiesAndNormalizesState(t *testing.T) {
	source := State{
		Mode: ModeNormal,
		Lines: []string{
			"ab",
			"xyz",
		},
		Cursor: Cursor{
			Row:        99,
			Col:        99,
			DesiredCol: -1,
		},
	}

	engine := NewWithState(source)
	source.Lines[1] = "changed"

	state := engine.State()
	if state.Cursor.Row != 1 || state.Cursor.Col != 2 || state.Cursor.DesiredCol != 2 {
		t.Fatalf("Cursor = (%d,%d,%d), want (1,2,2)", state.Cursor.Row, state.Cursor.Col, state.Cursor.DesiredCol)
	}
	if state.Lines[1] != "xyz" {
		t.Fatalf("line = %q, want %q", state.Lines[1], "xyz")
	}
}

func TestHorizontalMovement(t *testing.T) {
	engine := New([]string{"abc"})

	assertApply(t, engine, KeyL, 0, 1, EventMoved)
	assertApply(t, engine, KeyL, 0, 2, EventMoved)
	assertApply(t, engine, KeyL, 0, 2, EventBoundary)
	assertApply(t, engine, KeyH, 0, 1, EventMoved)
	assertApply(t, engine, KeyH, 0, 0, EventMoved)
	assertApply(t, engine, KeyH, 0, 0, EventBoundary)
}

func TestVerticalMovementClampsToShorterLinesAndRestoresDesiredColumn(t *testing.T) {
	engine := New([]string{"abcd", "x", "wxyz"})

	assertApply(t, engine, KeyL, 0, 1, EventMoved)
	assertApply(t, engine, KeyL, 0, 2, EventMoved)
	assertApply(t, engine, KeyJ, 1, 0, EventMoved)
	assertApply(t, engine, KeyJ, 2, 2, EventMoved)
	assertApply(t, engine, KeyK, 1, 0, EventMoved)
	assertApply(t, engine, KeyK, 0, 2, EventMoved)
}

func TestVerticalMovementHandlesEmptyLines(t *testing.T) {
	engine := New([]string{"abc", "", "xyz"})

	assertApply(t, engine, KeyL, 0, 1, EventMoved)
	assertApply(t, engine, KeyJ, 1, 0, EventMoved)
	assertApply(t, engine, KeyJ, 2, 1, EventMoved)
}

func TestUnsupportedKeyDoesNotMove(t *testing.T) {
	engine := New([]string{"abc"})
	result := engine.Apply("w")

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestApplyIsPureTransition(t *testing.T) {
	state := NewState([]string{"abc"})
	result := Apply(state, KeyL)

	if state.Cursor.Col != 0 {
		t.Fatalf("input state cursor col = %d, want 0", state.Cursor.Col)
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("result cursor col = %d, want 1", result.State.Cursor.Col)
	}
}

func TestUnicodeColumnsUseRunes(t *testing.T) {
	engine := New([]string{"가나"})

	assertApply(t, engine, KeyL, 0, 1, EventMoved)
	assertApply(t, engine, KeyL, 0, 1, EventBoundary)
}

func TestApplyKeysMatchesRepeatedApplyAndPreservesEventOrder(t *testing.T) {
	state := NewState([]string{"abcd", "xy"})
	keys := []string{KeyL, KeyL, KeyJ, KeyH}

	batch := ApplyKeys(state, keys)

	engine := NewWithState(state)
	for _, key := range keys {
		engine.Apply(key)
	}
	repeated := engine.State()

	if batch.State.Cursor != repeated.Cursor {
		t.Fatalf("batch cursor = %+v, want %+v", batch.State.Cursor, repeated.Cursor)
	}
	if len(batch.Events) != len(keys) {
		t.Fatalf("len(batch.Events) = %d, want %d", len(batch.Events), len(keys))
	}
	for index, key := range keys {
		if batch.Events[index].Key != key {
			t.Fatalf("event[%d].Key = %q, want %q", index, batch.Events[index].Key, key)
		}
	}
}

func TestEngineApplyKeysUpdatesEngineState(t *testing.T) {
	engine := New([]string{"abc"})
	result := engine.ApplyKeys([]string{KeyL, KeyL})

	if result.State.Cursor.Col != 2 {
		t.Fatalf("result cursor col = %d, want 2", result.State.Cursor.Col)
	}
	if engine.State().Cursor.Col != 2 {
		t.Fatalf("engine cursor col = %d, want 2", engine.State().Cursor.Col)
	}
}

func assertApply(t *testing.T, engine *Engine, key string, row int, col int, eventType EventType) {
	t.Helper()

	result := engine.Apply(key)
	if result.State.Cursor.Row != row || result.State.Cursor.Col != col {
		t.Fatalf("after %q cursor = (%d,%d), want (%d,%d)", key, result.State.Cursor.Row, result.State.Cursor.Col, row, col)
	}
	assertEvent(t, result, eventType)
}

func assertEvent(t *testing.T, result Result, eventType EventType) {
	t.Helper()

	if len(result.Events) != 1 {
		t.Fatalf("len(Events) = %d, want 1", len(result.Events))
	}
	if result.Events[0].Type != eventType {
		t.Fatalf("event type = %q, want %q", result.Events[0].Type, eventType)
	}
}
