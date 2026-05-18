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

func TestWordMotionMovesByWordStarts(t *testing.T) {
	engine := New([]string{"service api backend enabled"})

	assertApply(t, engine, KeyW, 0, 8, EventMoved)
	assertApply(t, engine, KeyW, 0, 12, EventMoved)
	assertApply(t, engine, KeyB, 0, 8, EventMoved)
	assertApply(t, engine, KeyE, 0, 10, EventMoved)
}

func TestWordMotionTreatsPunctuationAsWords(t *testing.T) {
	engine := New([]string{"foo.bar baz"})

	assertApply(t, engine, KeyW, 0, 3, EventMoved)
	assertApply(t, engine, KeyW, 0, 4, EventMoved)
	assertApply(t, engine, KeyE, 0, 6, EventMoved)
	assertApply(t, engine, KeyW, 0, 8, EventMoved)
}

func TestWordMotionCrossesLineBoundaries(t *testing.T) {
	engine := New([]string{"abc", "", "  def ghi"})

	assertApply(t, engine, KeyW, 2, 2, EventMoved)
	assertApply(t, engine, KeyE, 2, 4, EventMoved)
	assertApply(t, engine, KeyW, 2, 6, EventMoved)
	assertApply(t, engine, KeyB, 2, 2, EventMoved)
	assertApply(t, engine, KeyB, 0, 0, EventMoved)
}

func TestWordMotionReportsBoundaryAtDocumentEdges(t *testing.T) {
	engine := New([]string{"abc"})

	assertApply(t, engine, KeyB, 0, 0, EventBoundary)
	assertApply(t, engine, KeyE, 0, 2, EventMoved)
	assertApply(t, engine, KeyE, 0, 2, EventBoundary)
	assertApply(t, engine, KeyW, 0, 2, EventBoundary)
}

func TestWordMotionUnsupportedOutsideNormalMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeInsert,
		Lines: []string{"alpha beta"},
	})

	result := engine.Apply(KeyW)

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestWordMotionUpdatesDesiredColumnForVerticalMovement(t *testing.T) {
	engine := New([]string{"alpha beta", "xy", "0123456789"})

	assertApply(t, engine, KeyW, 0, 6, EventMoved)
	assertApply(t, engine, KeyJ, 1, 1, EventMoved)
	assertApply(t, engine, KeyJ, 2, 6, EventMoved)
}

func TestLineMotionMovesToStartAndEndOfCurrentLine(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "deploy target"},
		Cursor: Cursor{
			Row:        1,
			Col:        3,
			DesiredCol: 3,
		},
	})

	assertApply(t, engine, KeyZero, 1, 0, EventMoved)
	assertApply(t, engine, KeyDollar, 1, 12, EventMoved)
}

func TestDocumentMotionMovesToFirstAndLastLine(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"top", "middle", "bottom"},
		Cursor: Cursor{
			Row:        1,
			Col:        3,
			DesiredCol: 3,
		},
	})

	assertApply(t, engine, KeyG, 1, 3, EventPendingKey)
	assertApply(t, engine, KeyG, 0, 0, EventMoved)
	assertApply(t, engine, KeyShiftG, 2, 0, EventMoved)
}

func TestDocumentMotionClearsPendingGOnUnsupportedCombo(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"top", "bottom"},
		Cursor: Cursor{
			Row:        1,
			Col:        0,
			DesiredCol: 0,
		},
	})

	assertApply(t, engine, KeyG, 1, 0, EventPendingKey)
	assertApply(t, engine, KeyW, 1, 0, EventUnsupportedKey)
	assertApply(t, engine, KeyG, 1, 0, EventPendingKey)
	assertApply(t, engine, KeyG, 0, 0, EventMoved)
}

func TestUnsupportedKeyDoesNotMove(t *testing.T) {
	engine := New([]string{"abc"})
	result := engine.Apply("x")

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestEscReturnsToNormalMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeInsert,
		Lines: []string{"abc"},
	})

	result := engine.Apply(KeyEsc)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	assertEvent(t, result, EventModeReset)
}

func TestCommandLineExecutesQuitWithoutSave(t *testing.T) {
	engine := New([]string{"scratch"})

	assertApply(t, engine, KeyColon, 0, 0, EventCommandMode)
	assertCommandLine(t, engine.State(), "")
	assertApply(t, engine, "q", 0, 0, EventCommandInput)
	assertCommandLine(t, engine.State(), "q")
	assertApply(t, engine, "!", 0, 0, EventCommandInput)
	result := engine.Apply(KeyEnter)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.LastCommand != ":q!" {
		t.Fatalf("LastCommand = %q, want :q!", result.State.LastCommand)
	}
	assertEvent(t, result, EventCommandExecuted)
}

func TestCommandLineExecutesWriteQuit(t *testing.T) {
	engine := New([]string{"scratch"})

	engine.Apply(KeyColon)
	engine.Apply("w")
	engine.Apply("q")
	result := engine.Apply(KeyEnter)

	if result.State.LastCommand != ":wq" {
		t.Fatalf("LastCommand = %q, want :wq", result.State.LastCommand)
	}
	assertEvent(t, result, EventCommandExecuted)
}

func TestEscClearsCommandLine(t *testing.T) {
	engine := New([]string{"scratch"})

	engine.Apply(KeyColon)
	engine.Apply("q")
	result := engine.Apply(KeyEsc)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.CommandLine != "" {
		t.Fatalf("CommandLine = %q, want empty", result.State.CommandLine)
	}
	if result.State.LastCommand != "" {
		t.Fatalf("LastCommand = %q, want empty", result.State.LastCommand)
	}
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

func assertCommandLine(t *testing.T, state State, want string) {
	t.Helper()

	if state.CommandLine != want {
		t.Fatalf("CommandLine = %q, want %q", state.CommandLine, want)
	}
}
