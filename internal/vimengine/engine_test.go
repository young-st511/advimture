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

func TestVisualModeStartsCharwiseSelection(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcd"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	result := Apply(state, KeyV)

	if result.State.Mode != ModeVisual {
		t.Fatalf("mode = %q, want visual", result.State.Mode)
	}
	selection := result.State.Selection
	if selection == nil || !selection.Active || selection.Kind != SelectionCharwise {
		t.Fatalf("selection = %+v, want active charwise", selection)
	}
	if selection.Anchor.Col != 1 || selection.Head.Col != 1 || selection.Start.Col != 1 || selection.End.Col != 1 {
		t.Fatalf("selection = %+v, want all cols 1", selection)
	}
}

func TestLinewiseVisualModeStartsFullLineSelection(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "beta"},
		Cursor: Cursor{
			Row: 0,
			Col: 2,
		},
	}

	result := Apply(state, KeyShiftV)

	if result.State.Mode != ModeVisual {
		t.Fatalf("mode = %q, want visual", result.State.Mode)
	}
	selection := result.State.Selection
	if selection == nil || !selection.Active || selection.Kind != SelectionLinewise {
		t.Fatalf("selection = %+v, want active linewise", selection)
	}
	if selection.Anchor.Row != 0 || selection.Head.Row != 0 || selection.Start.Row != 0 || selection.Start.Col != 0 || selection.End.Row != 0 || selection.End.Col != 4 {
		t.Fatalf("selection = %+v, want full first line range", selection)
	}
}

func TestLinewiseVisualMotionNormalizesFullLineRange(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "", "gamma"},
		Cursor: Cursor{
			Row: 0,
			Col: 2,
		},
	}

	result := ApplyKeys(state, []string{KeyShiftV, KeyJ, KeyJ})

	selection := result.State.Selection
	if result.State.Mode != ModeVisual || selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual linewise selection", result.State.Mode, selection)
	}
	if selection.Kind != SelectionLinewise || selection.Start.Row != 0 || selection.Start.Col != 0 || selection.End.Row != 2 || selection.End.Col != 4 {
		t.Fatalf("selection = %+v, want linewise rows 0..2", selection)
	}
}

func TestLinewiseVisualMotionNormalizesBackwardRange(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "beta", "gamma"},
		Cursor: Cursor{
			Row: 2,
			Col: 2,
		},
	}

	result := ApplyKeys(state, []string{KeyShiftV, KeyK})

	selection := result.State.Selection
	if selection == nil || selection.Kind != SelectionLinewise || selection.Start.Row != 1 || selection.Start.Col != 0 || selection.End.Row != 2 || selection.End.Col != 4 {
		t.Fatalf("selection = %+v, want linewise rows 1..2", selection)
	}
}

func TestLinewiseVisualYankStoresLinewiseRegisterWithoutChangingBuffer(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
	}

	result := ApplyKeys(state, []string{KeyShiftV, KeyJ, KeyY})

	assertStrings(t, result.State.Lines, []string{"one", "two", "three"})
	if result.State.Mode != ModeNormal || result.State.Selection != nil {
		t.Fatalf("mode/selection = %s/%+v, want normal nil", result.State.Mode, result.State.Selection)
	}
	assertStrings(t, result.State.Register.Lines, []string{"one", "two"})
	if !result.State.Register.Linewise || result.State.Register.Text != "" {
		t.Fatalf("register = %+v, want linewise [one two]", result.State.Register)
	}
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestLinewiseVisualDeleteDeletesRowsStoresRegisterAndUndoRestores(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
	}

	deleted := ApplyKeys(state, []string{KeyShiftV, KeyJ, KeyD})

	assertStrings(t, deleted.State.Lines, []string{"three"})
	assertStrings(t, deleted.State.Register.Lines, []string{"one", "two"})
	if !deleted.State.Register.Linewise || deleted.State.Register.Text != "" {
		t.Fatalf("register = %+v, want linewise deleted rows", deleted.State.Register)
	}
	if deleted.State.Mode != ModeNormal || deleted.State.Selection != nil {
		t.Fatalf("mode/selection = %s/%+v, want normal nil", deleted.State.Mode, deleted.State.Selection)
	}
	if deleted.State.Cursor.Row != 0 || deleted.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (0,0)", deleted.State.Cursor.Row, deleted.State.Cursor.Col)
	}

	restored := Apply(deleted.State, KeyU)
	assertStrings(t, restored.State.Lines, []string{"one", "two", "three"})
}

func TestLinewiseVisualDeleteAllRowsLeavesEmptyFallback(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two"},
	}

	result := ApplyKeys(state, []string{KeyShiftV, KeyShiftG, KeyD})

	assertStrings(t, result.State.Lines, []string{""})
	assertStrings(t, result.State.Register.Lines, []string{"one", "two"})
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestLinewiseVisualSupportsDocumentStartAndEndMotions(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
		Cursor: Cursor{
			Row: 1,
			Col: 1,
		},
	}

	toEnd := ApplyKeys(state, []string{KeyShiftV, KeyShiftG})
	if toEnd.State.Selection == nil || toEnd.State.Selection.Start.Row != 1 || toEnd.State.Selection.End.Row != 2 {
		t.Fatalf("selection after G = %+v, want rows 1..2", toEnd.State.Selection)
	}

	toStart := ApplyKeys(state, []string{KeyShiftV, KeyG, KeyG})
	if toStart.State.Selection == nil || toStart.State.Selection.Start.Row != 0 || toStart.State.Selection.End.Row != 1 {
		t.Fatalf("selection after gg = %+v, want rows 0..1", toStart.State.Selection)
	}
}

func TestVisualModeMotionUpdatesHeadAndRange(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcd"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	result := ApplyKeys(state, []string{KeyV, KeyL, KeyL})

	selection := result.State.Selection
	if result.State.Mode != ModeVisual {
		t.Fatalf("mode = %q, want visual", result.State.Mode)
	}
	if result.State.Cursor.Col != 3 {
		t.Fatalf("cursor col = %d, want 3", result.State.Cursor.Col)
	}
	if selection == nil || selection.Anchor.Col != 1 || selection.Head.Col != 3 || selection.Start.Col != 1 || selection.End.Col != 3 {
		t.Fatalf("selection = %+v, want 1..3", selection)
	}
}

func TestVisualModeNormalizesBackwardSelection(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcd"},
		Cursor: Cursor{
			Row: 0,
			Col: 2,
		},
	}

	result := ApplyKeys(state, []string{KeyV, KeyH, KeyH})

	selection := result.State.Selection
	if selection == nil || selection.Anchor.Col != 2 || selection.Head.Col != 0 || selection.Start.Col != 0 || selection.End.Col != 2 {
		t.Fatalf("selection = %+v, want normalized 0..2 with anchor 2", selection)
	}
}

func TestVisualModeEscAndVResetSelection(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcd"},
	}

	escaped := ApplyKeys(state, []string{KeyV, KeyL, KeyEsc})
	if escaped.State.Mode != ModeNormal || escaped.State.Selection != nil {
		t.Fatalf("escaped state = %s/%+v, want normal nil selection", escaped.State.Mode, escaped.State.Selection)
	}

	toggled := ApplyKeys(state, []string{KeyV, KeyL, KeyV})
	if toggled.State.Mode != ModeNormal || toggled.State.Selection != nil {
		t.Fatalf("toggled state = %s/%+v, want normal nil selection", toggled.State.Mode, toggled.State.Selection)
	}
}

func TestVisualDeleteDeletesInclusiveSelectionAndStoresRegister(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcdef"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	result := ApplyKeys(state, []string{KeyV, KeyL, KeyL, KeyD})

	assertStrings(t, result.State.Lines, []string{"aef"})
	if result.State.Mode != ModeNormal || result.State.Selection != nil {
		t.Fatalf("mode/selection = %s/%+v, want normal nil", result.State.Mode, result.State.Selection)
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want 1", result.State.Cursor.Col)
	}
	if result.State.Register.Text != "bcd" || result.State.Register.Linewise {
		t.Fatalf("register = %+v, want charwise bcd", result.State.Register)
	}
	if !hasEvent(result, EventChanged) {
		t.Fatalf("events = %+v, want changed", result.Events)
	}
}

func TestVisualYankStoresInclusiveSelectionWithoutChangingBuffer(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcdef"},
		Cursor: Cursor{
			Row: 0,
			Col: 3,
		},
	}

	result := ApplyKeys(state, []string{KeyV, KeyH, KeyH, KeyY})

	assertStrings(t, result.State.Lines, []string{"abcdef"})
	if result.State.Mode != ModeNormal || result.State.Selection != nil {
		t.Fatalf("mode/selection = %s/%+v, want normal nil", result.State.Mode, result.State.Selection)
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want normalized start 1", result.State.Cursor.Col)
	}
	if result.State.Register.Text != "bcd" || result.State.Register.Linewise {
		t.Fatalf("register = %+v, want charwise bcd", result.State.Register)
	}
	if !hasEvent(result, EventYanked) {
		t.Fatalf("events = %+v, want yanked", result.Events)
	}
}

func TestVisualDeleteCanBeUndone(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcdef"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	deleted := ApplyKeys(state, []string{KeyV, KeyL, KeyD})
	result := Apply(deleted.State, KeyU)

	assertStrings(t, result.State.Lines, []string{"abcdef"})
	assertEvent(t, result, EventChanged)
}

func TestVisualOperatorRejectsMultiLineSelection(t *testing.T) {
	state := State{
		Mode:  ModeVisual,
		Lines: []string{"abc", "def"},
		Cursor: Cursor{
			Row: 1,
			Col: 1,
		},
		Selection: &Selection{
			Active: true,
			Kind:   SelectionCharwise,
			Anchor: Cursor{Row: 0, Col: 1},
			Head:   Cursor{Row: 1, Col: 1},
			Start:  Cursor{Row: 0, Col: 1},
			End:    Cursor{Row: 1, Col: 1},
		},
	}

	result := Apply(state, KeyD)

	assertStrings(t, result.State.Lines, []string{"abc", "def"})
	if result.State.Mode != ModeVisual || result.State.Selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual selection preserved", result.State.Mode, result.State.Selection)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestVisualBoundaryMotionKeepsSelectionAtCursor(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abc"},
		Cursor: Cursor{
			Row: 0,
			Col: 2,
		},
	}

	result := ApplyKeys(state, []string{KeyV, KeyL})
	selection := result.State.Selection

	if result.State.Mode != ModeVisual || selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual selection", result.State.Mode, selection)
	}
	if result.State.Cursor.Col != 2 || selection.Anchor.Col != 2 || selection.Head.Col != 2 || selection.Start.Col != 2 || selection.End.Col != 2 {
		t.Fatalf("state = cursor %+v selection %+v, want clamped single-cell selection at col 2", result.State.Cursor, selection)
	}
	if !hasEvent(result, EventBoundary) {
		t.Fatalf("events = %+v, want boundary", result.Events)
	}
}

func TestVisualOperatorOnEmptyLinePreservesVisualState(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{""},
	}

	result := ApplyKeys(state, []string{KeyV, KeyD})

	assertStrings(t, result.State.Lines, []string{""})
	if result.State.Mode != ModeVisual || result.State.Selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual selection preserved", result.State.Mode, result.State.Selection)
	}
	if !hasEvent(result, EventUnsupportedKey) {
		t.Fatalf("events = %+v, want unsupported key", result.Events)
	}
}

func TestVisualYankDoesNotCreateUndoSnapshot(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcdef"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	yanked := ApplyKeys(state, []string{KeyV, KeyL, KeyY})
	result := Apply(yanked.State, KeyU)

	assertStrings(t, result.State.Lines, []string{"abcdef"})
	if result.State.Register.Text != "bc" {
		t.Fatalf("register text = %q, want bc", result.State.Register.Text)
	}
	assertEvent(t, result, EventBoundary)
}

func TestVisualDeleteUndoRestoresBufferAndKeepsRegister(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abcdef"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	deleted := ApplyKeys(state, []string{KeyV, KeyL, KeyL, KeyD})
	result := Apply(deleted.State, KeyU)

	assertStrings(t, result.State.Lines, []string{"abcdef"})
	if result.State.Cursor.Col != 3 {
		t.Fatalf("cursor col = %d, want restored visual head col 3", result.State.Cursor.Col)
	}
	if result.State.Register.Text != "bcd" {
		t.Fatalf("register text = %q, want deleted text bcd", result.State.Register.Text)
	}
}

func TestVisualMotionAcrossLinesRejectsOperatorAndPreservesRange(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"abc", "def"},
		Cursor: Cursor{
			Row: 0,
			Col: 1,
		},
	}

	selected := ApplyKeys(state, []string{KeyV, KeyJ})
	result := Apply(selected.State, KeyD)

	assertStrings(t, result.State.Lines, []string{"abc", "def"})
	if result.State.Mode != ModeVisual || result.State.Selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual selection preserved", result.State.Mode, result.State.Selection)
	}
	selection := result.State.Selection
	if selection.Start.Row != 0 || selection.End.Row != 1 {
		t.Fatalf("selection = %+v, want multi-line range preserved", selection)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestNewWithStateClampsVisualSelection(t *testing.T) {
	state := State{
		Mode:  ModeVisual,
		Lines: []string{"ab"},
		Cursor: Cursor{
			Row: 0,
			Col: 99,
		},
		Selection: &Selection{
			Active: true,
			Kind:   SelectionCharwise,
			Anchor: Cursor{Row: -1, Col: 99},
			Head:   Cursor{Row: 99, Col: 99},
		},
	}

	normalized := NewWithState(state).State()
	selection := normalized.Selection

	if normalized.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want 1", normalized.Cursor.Col)
	}
	if selection == nil || selection.Anchor.Row != 0 || selection.Anchor.Col != 1 || selection.Head.Row != 0 || selection.Head.Col != 1 {
		t.Fatalf("selection = %+v, want clamped to row 0 col 1", selection)
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

func TestWordMotionKeyInInsertModeInsertsText(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeInsert,
		Lines: []string{"alpha beta"},
	})

	result := engine.Apply(KeyW)

	if result.State.Lines[0] != "walpha beta" {
		t.Fatalf("line = %q, want walpha beta", result.State.Lines[0])
	}
	assertEvent(t, result, EventChanged)
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

func TestOpenLineBelowEntersInsertMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "omega"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	result := engine.Apply(KeyO)

	assertStrings(t, result.State.Lines, []string{"alpha", "", "omega"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want %q", result.State.Mode, ModeInsert)
	}
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 0 || result.State.Cursor.DesiredCol != 0 {
		t.Fatalf("cursor = (%d,%d,%d), want (1,0,0)", result.State.Cursor.Row, result.State.Cursor.Col, result.State.Cursor.DesiredCol)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestOpenLineAboveEntersInsertMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "omega"},
		Cursor: Cursor{
			Row:        1,
			Col:        2,
			DesiredCol: 2,
		},
	})

	result := engine.Apply(KeyShiftO)

	assertStrings(t, result.State.Lines, []string{"alpha", "", "omega"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want %q", result.State.Mode, ModeInsert)
	}
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 0 || result.State.Cursor.DesiredCol != 0 {
		t.Fatalf("cursor = (%d,%d,%d), want (1,0,0)", result.State.Cursor.Row, result.State.Cursor.Col, result.State.Cursor.DesiredCol)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestOpenLineCanBeUndoneAfterEsc(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha", "omega"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	engine.Apply(KeyO)
	engine.Apply(KeyEsc)
	result := engine.Apply(KeyU)

	assertStrings(t, result.State.Lines, []string{"alpha", "omega"})
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want %q", result.State.Mode, ModeNormal)
	}
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 2 || result.State.Cursor.DesiredCol != 2 {
		t.Fatalf("cursor = (%d,%d,%d), want (0,2,2)", result.State.Cursor.Row, result.State.Cursor.Col, result.State.Cursor.DesiredCol)
	}
	assertEvent(t, result, EventChanged)
}

func TestDotWithoutLastChangeReportsBoundary(t *testing.T) {
	engine := New([]string{"abc"})

	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{"abc"})
	assertEvent(t, result, EventBoundary)
}

func TestDotRepeatsAppendInsertTransaction(t *testing.T) {
	engine := New([]string{"api", "api"})

	for _, key := range []string{KeyShiftA, "!", KeyEsc, KeyJ} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{"api!", "api!"})
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 3 {
		t.Fatalf("cursor = (%d,%d), want (1,3)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestDotRepeatsReplaceCharTransaction(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"bad", "bad"},
		Cursor: Cursor{
			Row:        0,
			Col:        0,
			DesiredCol: 0,
		},
	})

	engine.Apply(KeyR)
	engine.Apply("g")
	engine.Apply(KeyJ)
	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{"gad", "gad"})
	if result.State.LastChange[0] != KeyR || result.State.LastChange[1] != "g" {
		t.Fatalf("last change = %+v, want r g", result.State.LastChange)
	}
}

func TestDotRepeatsChangeInnerWordTransaction(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode=down", "next=down"},
		Cursor: Cursor{
			Row:        0,
			Col:        7,
			DesiredCol: 7,
		},
	})

	for _, key := range []string{KeyC, KeyI, KeyW, "u", "p", KeyEsc, KeyJ} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{"mode=up", "next=up"})
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 6 {
		t.Fatalf("cursor = (%d,%d), want (1,6)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestDotRepeatsOpenLineTransactionWithoutOverwritingLastChange(t *testing.T) {
	engine := New([]string{"alpha", "omega"})

	for _, key := range []string{KeyO, "x", KeyEsc, KeyJ} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{"alpha", "x", "omega", "x"})
	assertStrings(t, result.State.LastChange, []string{KeyO, "x", KeyEsc})
}

func TestLiteralSearchMovesToNextMatch(t *testing.T) {
	engine := New([]string{"info", "warn timeout", "error timeout"})

	for _, key := range []string{KeySlash, "t", "i", "m", "e", "o", "u", "t"} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyEnter)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 5 {
		t.Fatalf("cursor = (%d,%d), want (1,5)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	if result.State.LastSearch != "timeout" {
		t.Fatalf("last search = %q, want timeout", result.State.LastSearch)
	}
	assertEvent(t, result, EventMoved)
}

func TestSearchRepeatNextAndPrevious(t *testing.T) {
	engine := New([]string{"ok timeout", "ok", "next timeout"})

	for _, key := range []string{KeySlash, "t", "i", "m", "e", "o", "u", "t", KeyEnter} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyN)
	if result.State.Cursor.Row != 2 || result.State.Cursor.Col != 5 {
		t.Fatalf("cursor after n = (%d,%d), want (2,5)", result.State.Cursor.Row, result.State.Cursor.Col)
	}

	result = engine.Apply(KeyShiftN)
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 3 {
		t.Fatalf("cursor after N = (%d,%d), want (0,3)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestLiteralSearchWrapsDocument(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"target", "middle", "tail"},
		Cursor: Cursor{
			Row:        2,
			Col:        0,
			DesiredCol: 0,
		},
	})

	for _, key := range []string{KeySlash, "t", "a", "r", "g", "e", "t"} {
		engine.Apply(key)
	}
	result := engine.Apply(KeyEnter)

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want wrapped (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
}

func TestSearchEscCancelsSearchMode(t *testing.T) {
	engine := New([]string{"abc"})

	engine.Apply(KeySlash)
	engine.Apply("a")
	result := engine.Apply(KeyEsc)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.CommandLine != "" || result.State.LastSearch != "" {
		t.Fatalf("command/search = %q/%q, want empty", result.State.CommandLine, result.State.LastSearch)
	}
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

func TestOperatorKeysEnterPendingMode(t *testing.T) {
	for _, key := range []string{KeyD, KeyC} {
		engine := NewWithState(State{
			Mode:  ModeNormal,
			Lines: []string{"alpha beta"},
			Cursor: Cursor{
				Row:        0,
				Col:        3,
				DesiredCol: 3,
			},
		})

		result := engine.Apply(key)

		if result.State.PendingKey != key {
			t.Fatalf("after %q pending key = %q, want %q", key, result.State.PendingKey, key)
		}
		if result.State.Lines[0] != "alpha beta" {
			t.Fatalf("after %q line = %q, want alpha beta", key, result.State.Lines[0])
		}
		if result.State.Cursor.Col != 3 {
			t.Fatalf("after %q cursor col = %d, want 3", key, result.State.Cursor.Col)
		}
		assertEvent(t, result, EventPendingKey)
	}
}

func TestOperatorPendingClearsUnsupportedComboWithoutMutating(t *testing.T) {
	for _, tt := range []struct {
		name     string
		operator string
		nextKey  string
	}{
		{name: "delete unsupported pending", operator: KeyD, nextKey: KeyE},
		{name: "change unsupported pending", operator: KeyC, nextKey: KeyE},
	} {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewWithState(State{
				Mode:  ModeNormal,
				Lines: []string{"alpha beta", "gamma"},
				Cursor: Cursor{
					Row:        0,
					Col:        2,
					DesiredCol: 2,
				},
			})

			assertApply(t, engine, tt.operator, 0, 2, EventPendingKey)
			result := engine.Apply(tt.nextKey)

			assertStrings(t, result.State.Lines, []string{"alpha beta", "gamma"})
			if result.State.PendingKey != "" {
				t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
			}
			if result.State.Mode != ModeNormal {
				t.Fatalf("mode = %q, want normal", result.State.Mode)
			}
			if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 2 {
				t.Fatalf("cursor = (%d,%d), want (0,2)", result.State.Cursor.Row, result.State.Cursor.Col)
			}
			assertEvent(t, result, EventUnsupportedKey)
		})
	}
}

func TestOperatorPendingCanBeCanceledWithEsc(t *testing.T) {
	for _, key := range []string{KeyD, KeyC} {
		engine := New([]string{"alpha beta"})

		assertApply(t, engine, key, 0, 0, EventPendingKey)
		result := engine.Apply(KeyEsc)

		if result.State.PendingKey != "" {
			t.Fatalf("after %q esc pending key = %q, want empty", key, result.State.PendingKey)
		}
		if result.State.Mode != ModeNormal {
			t.Fatalf("after %q esc mode = %q, want normal", key, result.State.Mode)
		}
		assertEvent(t, result, EventModeReset)
	}
}

func TestYankKeyEntersPendingMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        3,
			DesiredCol: 3,
		},
	})

	result := engine.Apply(KeyY)

	if result.State.PendingKey != KeyY {
		t.Fatalf("pending key = %q, want %q", result.State.PendingKey, KeyY)
	}
	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	if result.State.Cursor.Col != 3 {
		t.Fatalf("cursor col = %d, want 3", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventPendingKey)
}

func TestYankWordStoresCharwiseRegisterWithoutMutating(t *testing.T) {
	engine := New([]string{"alpha beta"})

	assertApply(t, engine, KeyY, 0, 0, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	if result.State.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want 0", result.State.Cursor.Col)
	}
	if result.State.Register.Text != "alpha " {
		t.Fatalf("register text = %q, want alpha space", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertEvent(t, result, EventYanked)
}

func TestYankToLineEndStoresCharwiseRegister(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        6,
			DesiredCol: 6,
		},
	})

	engine.Apply(KeyY)
	result := engine.Apply(KeyDollar)

	if result.State.Register.Text != "beta" {
		t.Fatalf("register text = %q, want beta", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	assertEvent(t, result, EventYanked)
}

func TestYankCurrentLineStoresLinewiseRegister(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
		Cursor: Cursor{
			Row:        1,
			Col:        2,
			DesiredCol: 2,
		},
	})

	engine.Apply(KeyY)
	result := engine.Apply(KeyY)

	assertStrings(t, result.State.Register.Lines, []string{"two"})
	if !result.State.Register.Linewise {
		t.Fatal("register linewise = false, want true")
	}
	if result.State.Register.Text != "" {
		t.Fatalf("register text = %q, want empty", result.State.Register.Text)
	}
	assertStrings(t, result.State.Lines, []string{"one", "two", "three"})
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 2 {
		t.Fatalf("cursor = (%d,%d), want (1,2)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventYanked)
}

func TestOperatorInnerTextObjectEntersPendingMode(t *testing.T) {
	for _, tt := range []struct {
		operator string
		pending  string
	}{
		{operator: KeyD, pending: pendingDeleteInner},
		{operator: KeyC, pending: pendingChangeInner},
		{operator: KeyY, pending: pendingYankInner},
	} {
		t.Run(tt.operator, func(t *testing.T) {
			engine := NewWithState(State{
				Mode:  ModeNormal,
				Lines: []string{"alpha beta"},
				Cursor: Cursor{
					Row:        0,
					Col:        2,
					DesiredCol: 2,
				},
			})

			assertApply(t, engine, tt.operator, 0, 2, EventPendingKey)
			result := engine.Apply(KeyI)

			if result.State.PendingKey != tt.pending {
				t.Fatalf("pending key = %q, want %q", result.State.PendingKey, tt.pending)
			}
			assertStrings(t, result.State.Lines, []string{"alpha beta"})
			assertEvent(t, result, EventPendingKey)
		})
	}
}

func TestUnsupportedInnerTextObjectSequenceDoesNotMutate(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	assertApply(t, engine, KeyD, 0, 2, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 2, EventPendingKey)
	result := engine.Apply(KeyE)

	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestDeleteInnerWordDeletesCurrentWordRun(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode=broken"},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyD, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"mode="})
	if result.State.Cursor.Col != 4 {
		t.Fatalf("cursor col = %d, want 4", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestChangeInnerWordEntersInsertModeAtRemovedWordStart(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode=broken"},
		Cursor: Cursor{
			Row:        0,
			Col:        7,
			DesiredCol: 7,
		},
	})

	assertApply(t, engine, KeyC, 0, 7, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 7, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"mode="})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 5 {
		t.Fatalf("cursor col = %d, want 5", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestYankInnerWordStoresCurrentWordRunWithoutMutating(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode=stable"},
		Cursor: Cursor{
			Row:        0,
			Col:        7,
			DesiredCol: 7,
		},
	})

	assertApply(t, engine, KeyY, 0, 7, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 7, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"mode=stable"})
	if result.State.Cursor.Col != 7 {
		t.Fatalf("cursor col = %d, want 7", result.State.Cursor.Col)
	}
	if result.State.Register.Text != "stable" {
		t.Fatalf("register text = %q, want stable", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertEvent(t, result, EventYanked)
}

func TestDeleteInnerQuoteDeletesDoubleQuotedValue(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{`mode="broken"`},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyD, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply(KeyDoubleQuote)

	assertStrings(t, result.State.Lines, []string{`mode=""`})
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestChangeInnerQuoteEntersInsertModeInsideQuotes(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{`mode="broken"`},
		Cursor: Cursor{
			Row:        0,
			Col:        9,
			DesiredCol: 9,
		},
	})

	assertApply(t, engine, KeyC, 0, 9, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 9, EventPendingKey)
	result := engine.Apply(KeyDoubleQuote)

	assertStrings(t, result.State.Lines, []string{`mode=""`})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestYankInnerQuoteStoresValueWithoutMutating(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{`mode="stable"`},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyY, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply(KeyDoubleQuote)

	assertStrings(t, result.State.Lines, []string{`mode="stable"`})
	if result.State.Register.Text != "stable" {
		t.Fatalf("register text = %q, want stable", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertEvent(t, result, EventYanked)
}

func TestDeleteInnerSingleQuoteDeletesValue(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode='broken'"},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyD, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply("'")

	assertStrings(t, result.State.Lines, []string{"mode=''"})
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestChangeInnerSingleQuoteEntersInsertModeInsideQuotes(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode='broken'"},
		Cursor: Cursor{
			Row:        0,
			Col:        9,
			DesiredCol: 9,
		},
	})

	assertApply(t, engine, KeyC, 0, 9, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 9, EventPendingKey)
	result := engine.Apply("'")

	assertStrings(t, result.State.Lines, []string{"mode=''"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestYankInnerSingleQuoteStoresValueWithoutMutating(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode='stable'"},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyY, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply("'")

	assertStrings(t, result.State.Lines, []string{"mode='stable'"})
	if result.State.Register.Text != "stable" {
		t.Fatalf("register text = %q, want stable", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertEvent(t, result, EventYanked)
}

func TestChangeInnerParenEntersInsertModeInsidePair(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"call(old)"},
		Cursor: Cursor{
			Row:        0,
			Col:        6,
			DesiredCol: 6,
		},
	})

	assertApply(t, engine, KeyC, 0, 6, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 6, EventPendingKey)
	result := engine.Apply("(")

	assertStrings(t, result.State.Lines, []string{"call()"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 5 {
		t.Fatalf("cursor col = %d, want 5", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestDeleteInnerBraceDeletesValue(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"cache{DELETE}"},
		Cursor: Cursor{
			Row:        0,
			Col:        8,
			DesiredCol: 8,
		},
	})

	assertApply(t, engine, KeyD, 0, 8, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 8, EventPendingKey)
	result := engine.Apply("{")

	assertStrings(t, result.State.Lines, []string{"cache{}"})
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestYankInnerBraceStoresValueWithoutMutating(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"node{stable}"},
		Cursor: Cursor{
			Row:        0,
			Col:        7,
			DesiredCol: 7,
		},
	})

	assertApply(t, engine, KeyY, 0, 7, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 7, EventPendingKey)
	result := engine.Apply("{")

	assertStrings(t, result.State.Lines, []string{"node{stable}"})
	if result.State.Register.Text != "stable" {
		t.Fatalf("register text = %q, want stable", result.State.Register.Text)
	}
	if result.State.Register.Linewise {
		t.Fatal("register linewise = true, want false")
	}
	assertEvent(t, result, EventYanked)
}

func TestInnerQuoteWithoutPairDoesNotMutate(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{`mode=broken`},
		Cursor: Cursor{
			Row:        0,
			Col:        6,
			DesiredCol: 6,
		},
	})

	assertApply(t, engine, KeyD, 0, 6, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 6, EventPendingKey)
	result := engine.Apply(KeyDoubleQuote)

	assertStrings(t, result.State.Lines, []string{`mode=broken`})
	assertEvent(t, result, EventBoundary)
}

func TestChangeInnerQuoteCanBeRepeated(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{`a="old"`, `b="old"`},
		Cursor: Cursor{
			Row:        0,
			Col:        3,
			DesiredCol: 3,
		},
	})

	engine.Apply(KeyC)
	engine.Apply(KeyI)
	engine.Apply(KeyDoubleQuote)
	engine.Apply("n")
	engine.Apply("e")
	engine.Apply("w")
	engine.Apply(KeyEsc)
	engine.Apply(KeyJ)
	engine.Apply(KeyH)
	result := engine.Apply(KeyDot)

	assertStrings(t, result.State.Lines, []string{`a="new"`, `b="new"`})
	if !hasEvent(result, EventInsertMode) {
		t.Fatalf("events = %+v, want insert mode event", result.Events)
	}
}

func TestInnerWordOnSpaceReportsBoundaryWithoutMutating(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        5,
			DesiredCol: 5,
		},
	})

	assertApply(t, engine, KeyD, 0, 5, EventPendingKey)
	assertApply(t, engine, KeyI, 0, 5, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	assertEvent(t, result, EventBoundary)
}

func TestDeleteInnerWordCanBeUndone(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"mode=broken"},
		Cursor: Cursor{
			Row:        0,
			Col:        7,
			DesiredCol: 7,
		},
	})

	engine.Apply(KeyD)
	engine.Apply(KeyI)
	engine.Apply(KeyW)
	result := engine.Apply(KeyU)

	assertStrings(t, result.State.Lines, []string{"mode=broken"})
	assertEvent(t, result, EventChanged)
}

func TestInnerWordRangeSelectsCurrentWordRun(t *testing.T) {
	for _, tt := range []struct {
		name  string
		line  string
		col   int
		start int
		end   int
		ok    bool
	}{
		{name: "keyword middle", line: "alpha beta", col: 2, start: 0, end: 5, ok: true},
		{name: "keyword end", line: "alpha beta", col: 9, start: 6, end: 10, ok: true},
		{name: "symbol run", line: "a==b", col: 1, start: 1, end: 3, ok: true},
		{name: "space", line: "alpha beta", col: 5, ok: false},
		{name: "empty", line: "", col: 0, ok: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			start, end, ok := innerWordRange([]rune(tt.line), tt.col)
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if !ok {
				return
			}
			if start != tt.start || end != tt.end {
				t.Fatalf("range = (%d,%d), want (%d,%d)", start, end, tt.start, tt.end)
			}
		})
	}
}

func TestInnerQuoteRangeSelectsDoubleQuotedContent(t *testing.T) {
	for _, tt := range []struct {
		name  string
		line  string
		col   int
		start int
		end   int
		ok    bool
	}{
		{name: "middle", line: `mode="stable"`, col: 8, start: 6, end: 12, ok: true},
		{name: "left quote", line: `mode="stable"`, col: 5, ok: false},
		{name: "right quote", line: `mode="stable"`, col: 12, ok: false},
		{name: "outside", line: `mode="stable"`, col: 1, ok: false},
		{name: "missing close", line: `mode="stable`, col: 8, ok: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			start, end, ok := innerQuoteRange([]rune(tt.line), tt.col)
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if !ok {
				return
			}
			if start != tt.start || end != tt.end {
				t.Fatalf("range = (%d,%d), want (%d,%d)", start, end, tt.start, tt.end)
			}
		})
	}
}

func TestInnerPairRangeSelectsParenthesizedContent(t *testing.T) {
	for _, tt := range []struct {
		name  string
		line  string
		col   int
		open  rune
		close rune
		start int
		end   int
		ok    bool
	}{
		{name: "paren middle", line: "call(stable)", col: 6, open: '(', close: ')', start: 5, end: 11, ok: true},
		{name: "brace middle", line: "cache{stable}", col: 8, open: '{', close: '}', start: 6, end: 12, ok: true},
		{name: "on opener", line: "call(stable)", col: 4, open: '(', close: ')', ok: false},
		{name: "on closer", line: "call(stable)", col: 11, open: '(', close: ')', ok: false},
		{name: "outside", line: "call(stable)", col: 1, open: '(', close: ')', ok: false},
		{name: "missing close", line: "call(stable", col: 6, open: '(', close: ')', ok: false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			start, end, ok := innerPairRange([]rune(tt.line), tt.col, tt.open, tt.close)
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if !ok {
				return
			}
			if start != tt.start || end != tt.end {
				t.Fatalf("range = (%d,%d), want (%d,%d)", start, end, tt.start, tt.end)
			}
		})
	}
}

func TestYankUnsupportedComboClearsPendingWithoutChangingRegister(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Register: Register{
			Text: "existing",
		},
	})

	assertApply(t, engine, KeyY, 0, 0, EventPendingKey)
	result := engine.Apply(KeyE)

	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	if result.State.Register.Text != "existing" {
		t.Fatalf("register text = %q, want existing", result.State.Register.Text)
	}
	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	assertEvent(t, result, EventUnsupportedKey)
}

func TestStateCopiesRegisterLines(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two"},
		Register: Register{
			Lines:    []string{"one"},
			Linewise: true,
		},
	})

	state := engine.State()
	state.Register.Lines[0] = "mutated"

	got := engine.State().Register.Lines[0]
	if got != "one" {
		t.Fatalf("register line = %q, want one", got)
	}
}

func TestPutCharwiseAfterCursor(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"ab"},
		Cursor: Cursor{
			Row:        0,
			Col:        0,
			DesiredCol: 0,
		},
		Register: Register{
			Text: "X",
		},
	})

	result := engine.Apply(KeyP)

	assertStrings(t, result.State.Lines, []string{"aXb"})
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want 1", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestPutCharwiseBeforeCursor(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"ab"},
		Cursor: Cursor{
			Row:        0,
			Col:        1,
			DesiredCol: 1,
		},
		Register: Register{
			Text: "XY",
		},
	})

	result := engine.Apply(KeyShiftP)

	assertStrings(t, result.State.Lines, []string{"aXYb"})
	if result.State.Cursor.Col != 2 {
		t.Fatalf("cursor col = %d, want 2", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestPutLinewiseBelowAndAboveCurrentLine(t *testing.T) {
	state := State{
		Mode:  ModeNormal,
		Lines: []string{"one", "three"},
		Cursor: Cursor{
			Row:        0,
			Col:        1,
			DesiredCol: 1,
		},
		Register: Register{
			Lines:    []string{"two"},
			Linewise: true,
		},
	}

	engine := NewWithState(state)
	result := engine.Apply(KeyP)
	assertStrings(t, result.State.Lines, []string{"one", "two", "three"})
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor after p = (%d,%d), want (1,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)

	engine = NewWithState(state)
	result = engine.Apply(KeyShiftP)
	assertStrings(t, result.State.Lines, []string{"two", "one", "three"})
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor after P = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestPutBoundaryWhenRegisterIsEmpty(t *testing.T) {
	engine := New([]string{"ab"})

	result := engine.Apply(KeyP)

	assertStrings(t, result.State.Lines, []string{"ab"})
	assertEvent(t, result, EventBoundary)
}

func TestPutUndoRedo(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"ab"},
		Register: Register{
			Text: "X",
		},
	})

	engine.Apply(KeyP)
	result := engine.Apply(KeyU)
	assertStrings(t, result.State.Lines, []string{"ab"})
	assertEvent(t, result, EventChanged)

	result = engine.Apply(KeyCtrlR)
	assertStrings(t, result.State.Lines, []string{"aXb"})
	assertEvent(t, result, EventChanged)
}

func TestDeleteWordMotionDeletesToNextWordStart(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        0,
			DesiredCol: 0,
		},
	})

	assertApply(t, engine, KeyD, 0, 0, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"beta"})
	if result.State.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want 0", result.State.Cursor.Col)
	}
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteWordMotionFromMiddleOfWord(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	engine.Apply(KeyD)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"albeta"})
	if result.State.Cursor.Col != 2 {
		t.Fatalf("cursor col = %d, want 2", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteToLineEnd(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        6,
			DesiredCol: 6,
		},
	})

	engine.Apply(KeyD)
	result := engine.Apply(KeyDollar)

	assertStrings(t, result.State.Lines, []string{"alpha "})
	if result.State.Cursor.Col != 5 {
		t.Fatalf("cursor col = %d, want 5", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteCurrentLine(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
		Cursor: Cursor{
			Row:        1,
			Col:        2,
			DesiredCol: 2,
		},
	})

	engine.Apply(KeyD)
	result := engine.Apply(KeyD)

	assertStrings(t, result.State.Lines, []string{"one", "three"})
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (1,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteOnlyLineLeavesEmptyBuffer(t *testing.T) {
	engine := New([]string{"one"})

	engine.Apply(KeyD)
	result := engine.Apply(KeyD)

	assertStrings(t, result.State.Lines, []string{""})
	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteWithMotionUndoRedo(t *testing.T) {
	engine := New([]string{"alpha beta"})

	engine.Apply(KeyD)
	engine.Apply(KeyW)
	result := engine.Apply(KeyU)

	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	assertEvent(t, result, EventChanged)

	result = engine.Apply(KeyCtrlR)
	assertStrings(t, result.State.Lines, []string{"beta"})
	assertEvent(t, result, EventChanged)
}

func TestChangeWordMotionEntersInsertMode(t *testing.T) {
	engine := New([]string{"alpha beta"})

	assertApply(t, engine, KeyC, 0, 0, EventPendingKey)
	result := engine.Apply(KeyW)

	assertStrings(t, result.State.Lines, []string{"beta"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want 0", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)

	result = engine.Apply("X")
	assertStrings(t, result.State.Lines, []string{"Xbeta"})
	assertEvent(t, result, EventChanged)
}

func TestChangeToLineEndEntersInsertMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"alpha beta"},
		Cursor: Cursor{
			Row:        0,
			Col:        6,
			DesiredCol: 6,
		},
	})

	engine.Apply(KeyC)
	result := engine.Apply(KeyDollar)

	assertStrings(t, result.State.Lines, []string{"alpha "})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 6 {
		t.Fatalf("cursor col = %d, want insert position 6", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestChangeCurrentLineEntersInsertMode(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"one", "two", "three"},
		Cursor: Cursor{
			Row:        1,
			Col:        2,
			DesiredCol: 2,
		},
	})

	engine.Apply(KeyC)
	result := engine.Apply(KeyC)

	assertStrings(t, result.State.Lines, []string{"one", "", "three"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Row != 1 || result.State.Cursor.Col != 0 {
		t.Fatalf("cursor = (%d,%d), want (1,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventInsertMode)
}

func TestChangeWithMotionUndoRedo(t *testing.T) {
	engine := New([]string{"alpha beta"})

	engine.Apply(KeyC)
	engine.Apply(KeyW)
	engine.Apply("X")
	engine.Apply(KeyEsc)
	result := engine.Apply(KeyU)

	assertStrings(t, result.State.Lines, []string{"beta"})
	assertEvent(t, result, EventChanged)

	result = engine.Apply(KeyU)
	assertStrings(t, result.State.Lines, []string{"alpha beta"})
	assertEvent(t, result, EventChanged)

	result = engine.Apply(KeyCtrlR)
	assertStrings(t, result.State.Lines, []string{"beta"})
	assertEvent(t, result, EventChanged)
}

func TestCharFindMovesToTargetAndTillBeforeTarget(t *testing.T) {
	engine := New([]string{"key=value,backup"})

	assertApply(t, engine, KeyF, 0, 0, EventPendingKey)
	result := engine.Apply("=")
	if result.State.Cursor.Col != 3 {
		t.Fatalf("f= cursor col = %d, want 3", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventMoved)

	engine = New([]string{"key=value,backup"})
	assertApply(t, engine, KeyT, 0, 0, EventPendingKey)
	result = engine.Apply(",")
	if result.State.Cursor.Col != 8 {
		t.Fatalf("t, cursor col = %d, want 8", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventMoved)
}

func TestCharFindReportsBoundaryWhenTargetIsMissing(t *testing.T) {
	engine := New([]string{"key=value"})

	engine.Apply(KeyF)
	result := engine.Apply(",")

	if result.State.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want unchanged 0", result.State.Cursor.Col)
	}
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	assertEvent(t, result, EventBoundary)
}

func TestDeleteWithCharFindIncludesOrExcludesTarget(t *testing.T) {
	engine := New([]string{"alpha,beta"})

	engine.Apply(KeyD)
	engine.Apply(KeyF)
	result := engine.Apply(",")

	assertStrings(t, result.State.Lines, []string{"beta"})
	if result.State.Cursor.Col != 0 {
		t.Fatalf("df, cursor col = %d, want 0", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)

	engine = New([]string{"alpha,beta"})
	engine.Apply(KeyD)
	engine.Apply(KeyT)
	result = engine.Apply(",")

	assertStrings(t, result.State.Lines, []string{",beta"})
	if result.State.Cursor.Col != 0 {
		t.Fatalf("dt, cursor col = %d, want 0", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestChangeWithCharFindEntersInsertModeAndRecordsChange(t *testing.T) {
	engine := New([]string{"alpha,beta"})

	engine.Apply(KeyC)
	engine.Apply(KeyT)
	result := engine.Apply(",")

	assertStrings(t, result.State.Lines, []string{",beta"})
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	assertEvent(t, result, EventInsertMode)

	engine.Apply("X")
	result = engine.Apply(KeyEsc)

	assertStrings(t, result.State.Lines, []string{"X,beta"})
	assertStrings(t, result.State.LastChange, []string{KeyC, KeyT, ",", "X", KeyEsc})
}

func TestUnsupportedKeyDoesNotMove(t *testing.T) {
	engine := New([]string{"abc"})
	result := engine.Apply("z")

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestDeleteCurrentChar(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"abcd"},
		Cursor: Cursor{
			Row:        0,
			Col:        1,
			DesiredCol: 1,
		},
	})

	result := engine.Apply(KeyX)

	if result.State.Lines[0] != "acd" {
		t.Fatalf("line = %q, want acd", result.State.Lines[0])
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want 1", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteCurrentCharClampsCursorAtLineEnd(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"abc"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	result := engine.Apply(KeyX)

	if result.State.Lines[0] != "ab" {
		t.Fatalf("line = %q, want ab", result.State.Lines[0])
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor col = %d, want 1", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestDeleteCurrentCharOnEmptyLineReportsBoundary(t *testing.T) {
	engine := New([]string{""})

	result := engine.Apply(KeyX)

	if result.State.Lines[0] != "" {
		t.Fatalf("line = %q, want empty", result.State.Lines[0])
	}
	assertEvent(t, result, EventBoundary)
}

func TestReplaceCurrentChar(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"abz"},
		Cursor: Cursor{
			Row:        0,
			Col:        2,
			DesiredCol: 2,
		},
	})

	assertApply(t, engine, KeyR, 0, 2, EventPendingKey)
	result := engine.Apply("c")

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	assertEvent(t, result, EventChanged)
}

func TestReplaceCurrentCharCanBeCanceledWithEsc(t *testing.T) {
	engine := New([]string{"abc"})

	assertApply(t, engine, KeyR, 0, 0, EventPendingKey)
	result := engine.Apply(KeyEsc)

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	assertEvent(t, result, EventModeReset)
}

func TestReplaceCurrentCharRejectsMultiCharacterReplacement(t *testing.T) {
	engine := New([]string{"abc"})

	engine.Apply(KeyR)
	result := engine.Apply("enter")

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
	if result.State.PendingKey != "" {
		t.Fatalf("pending key = %q, want empty", result.State.PendingKey)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestInsertBeforeCursor(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"ac"},
		Cursor: Cursor{
			Row:        0,
			Col:        1,
			DesiredCol: 1,
		},
	})

	assertApply(t, engine, KeyI, 0, 1, EventInsertMode)
	result := engine.Apply("b")

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
	if result.State.Mode != ModeInsert {
		t.Fatalf("mode = %q, want insert", result.State.Mode)
	}
	if result.State.Cursor.Col != 2 {
		t.Fatalf("cursor col = %d, want insert position 2", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)
}

func TestAppendAfterCursor(t *testing.T) {
	engine := New([]string{"ac"})

	assertApply(t, engine, KeyA, 0, 1, EventInsertMode)
	result := engine.Apply("b")

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
}

func TestAppendAtLineEnd(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"abc"},
		Cursor: Cursor{
			Row:        0,
			Col:        0,
			DesiredCol: 0,
		},
	})

	assertApply(t, engine, KeyShiftA, 0, 3, EventInsertMode)
	result := engine.Apply("!")

	if result.State.Lines[0] != "abc!" {
		t.Fatalf("line = %q, want abc!", result.State.Lines[0])
	}
	if result.State.Cursor.Col != 4 {
		t.Fatalf("cursor col = %d, want insert position 4", result.State.Cursor.Col)
	}
}

func TestEscFromInsertModeClampsToNormalCursor(t *testing.T) {
	engine := New([]string{"abc"})

	engine.Apply(KeyShiftA)
	engine.Apply("!")
	result := engine.Apply(KeyEsc)

	if result.State.Mode != ModeNormal {
		t.Fatalf("mode = %q, want normal", result.State.Mode)
	}
	if result.State.Cursor.Col != 3 {
		t.Fatalf("cursor col = %d, want 3", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventModeReset)
}

func TestInsertModeRejectsMultiCharacterInput(t *testing.T) {
	engine := New([]string{"abc"})

	engine.Apply(KeyI)
	result := engine.Apply("enter")

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line = %q, want abc", result.State.Lines[0])
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestUndoRedoRestoresSingleCharDelete(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"abc"},
		Cursor: Cursor{
			Row:        0,
			Col:        1,
			DesiredCol: 1,
		},
	})

	engine.Apply(KeyX)
	result := engine.Apply(KeyU)

	if result.State.Lines[0] != "abc" {
		t.Fatalf("line after undo = %q, want abc", result.State.Lines[0])
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("cursor after undo = %d, want 1", result.State.Cursor.Col)
	}
	assertEvent(t, result, EventChanged)

	result = engine.Apply(KeyCtrlR)
	if result.State.Lines[0] != "ac" {
		t.Fatalf("line after redo = %q, want ac", result.State.Lines[0])
	}
	assertEvent(t, result, EventChanged)
}

func TestUndoRedoRestoresInsertText(t *testing.T) {
	engine := New([]string{"ac"})

	engine.Apply(KeyI)
	engine.Apply("b")
	engine.Apply(KeyEsc)
	result := engine.Apply(KeyU)

	if result.State.Lines[0] != "ac" {
		t.Fatalf("line after undo = %q, want ac", result.State.Lines[0])
	}
	if result.State.Mode != ModeNormal {
		t.Fatalf("mode after undo = %q, want normal", result.State.Mode)
	}

	result = engine.Apply(KeyCtrlR)
	if result.State.Lines[0] != "bac" {
		t.Fatalf("line after redo = %q, want bac", result.State.Lines[0])
	}
}

func TestNewChangeClearsRedoStack(t *testing.T) {
	engine := New([]string{"abc"})

	engine.Apply(KeyX)
	engine.Apply(KeyU)
	engine.Apply(KeyR)
	engine.Apply("z")
	result := engine.Apply(KeyCtrlR)

	if result.State.Lines[0] != "zbc" {
		t.Fatalf("line = %q, want zbc", result.State.Lines[0])
	}
	assertEvent(t, result, EventBoundary)
}

func TestUndoRedoBoundaryWhenStackIsEmpty(t *testing.T) {
	engine := New([]string{"abc"})

	result := engine.Apply(KeyU)
	assertEvent(t, result, EventBoundary)

	result = engine.Apply(KeyCtrlR)
	assertEvent(t, result, EventBoundary)
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

func TestCommandLineSubstitutesFirstMatchOnCurrentLine(t *testing.T) {
	engine := NewWithState(State{
		Mode:  ModeNormal,
		Lines: []string{"api api", "api"},
		Cursor: Cursor{
			Row:        0,
			Col:        0,
			DesiredCol: 0,
		},
	})

	result := applyCommandText(engine, "s/api/web/")

	assertStrings(t, result.State.Lines, []string{"web api", "api"})
	if result.State.LastCommand != ":s/api/web/" {
		t.Fatalf("LastCommand = %q, want :s/api/web/", result.State.LastCommand)
	}
	assertEvent(t, result, EventCommandExecuted)
}

func TestCommandLineSubstitutesAllMatchesInWholeFile(t *testing.T) {
	engine := New([]string{"TODO api", "TODO worker"})

	result := applyCommandText(engine, "%s/TODO/DONE/g")

	assertStrings(t, result.State.Lines, []string{"DONE api", "DONE worker"})
	assertEvent(t, result, EventCommandExecuted)
}

func TestCommandLineSubstitutesNumericLineRange(t *testing.T) {
	engine := New([]string{"error one", "error two", "error three"})

	result := applyCommandText(engine, "2,3s/error/ok/")

	assertStrings(t, result.State.Lines, []string{"error one", "ok two", "ok three"})
	assertEvent(t, result, EventCommandExecuted)
}

func TestCommandLineRejectsInvalidSubstituteWithoutMutating(t *testing.T) {
	engine := New([]string{"api"})

	result := applyCommandText(engine, "%s//web/g")

	assertStrings(t, result.State.Lines, []string{"api"})
	if result.State.LastCommand != "" {
		t.Fatalf("LastCommand = %q, want empty", result.State.LastCommand)
	}
	assertEvent(t, result, EventUnsupportedKey)
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

func applyCommandText(engine *Engine, command string) Result {
	engine.Apply(KeyColon)
	for _, r := range command {
		engine.Apply(string(r))
	}
	return engine.Apply(KeyEnter)
}

func assertStrings(t *testing.T, got []string, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d: %+v", len(got), len(want), got)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("value[%d] = %q, want %q", index, got[index], want[index])
		}
	}
}
