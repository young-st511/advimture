package tuiadapter

import (
	"testing"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scenario"
	"github.com/young-st511/advimture/internal/scoring"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestMapInputMapsVimMovementKeys(t *testing.T) {
	cases := map[string]string{
		"h": vimengine.KeyH,
		"j": vimengine.KeyJ,
		"k": vimengine.KeyK,
		"l": vimengine.KeyL,
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputPreservesArrowKeysForExerciseConstraints(t *testing.T) {
	for _, input := range []string{"left", "down", "up", "right"} {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != input {
			t.Fatalf("MapInput(%q) = %+v, want preserved key %q", input, action, input)
		}
	}
}

func TestMapInputMapsCommands(t *testing.T) {
	cases := map[string]ActionType{
		"?": ActionHint,
		"q": ActionQuit,
	}

	for input, want := range cases {
		action := MapInput(input)
		if action.Type != want {
			t.Fatalf("MapInput(%q) = %q, want %q", input, action.Type, want)
		}
	}
}

func TestMapInputPassesCtrlCToRuntime(t *testing.T) {
	action := MapInput("ctrl+c")

	if action.Type != ActionKey || action.Key != "ctrl+c" {
		t.Fatalf("MapInput(ctrl+c) = %+v, want runtime key", action)
	}
}

func TestMapInputMapsSingleCharEditKeys(t *testing.T) {
	cases := map[string]string{
		"x":      vimengine.KeyX,
		"r":      vimengine.KeyR,
		"i":      vimengine.KeyI,
		"a":      vimengine.KeyA,
		"A":      vimengine.KeyShiftA,
		"u":      vimengine.KeyU,
		"ctrl+r": vimengine.KeyCtrlR,
		"Z":      "Z",
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputMapsOperatorKeys(t *testing.T) {
	cases := map[string]string{
		"d": vimengine.KeyD,
		"c": vimengine.KeyC,
		"y": vimengine.KeyY,
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputInInsertModePassesPrintableCharacters(t *testing.T) {
	for _, input := range []string{"q", "x", "A", "!", "d", "c", " "} {
		action := MapInputForMode(input, vimengine.ModeInsert)
		if action.Type != ActionKey || action.Key != input {
			t.Fatalf("MapInputForMode(%q, insert) = %+v, want key %q", input, action, input)
		}
	}
}

func TestMapInputInInsertModeMapsNamedSpaceToPrintableSpace(t *testing.T) {
	action := MapInputForMode("space", vimengine.ModeInsert)

	if action.Type != ActionKey || action.Key != " " {
		t.Fatalf("MapInputForMode(space, insert) = %+v, want printable space", action)
	}
}

func TestMapInputInInsertModeMapsEscButIgnoresEnter(t *testing.T) {
	action := MapInputForMode("esc", vimengine.ModeInsert)
	if action.Type != ActionKey || action.Key != vimengine.KeyEsc {
		t.Fatalf("esc action = %+v, want esc key", action)
	}

	action = MapInputForMode("enter", vimengine.ModeInsert)
	if action.Type != ActionIgnored {
		t.Fatalf("enter action = %+v, want ignored", action)
	}
}

func TestMapInputInCommandModeTreatsQAsVimKey(t *testing.T) {
	action := MapInputForMode("q", vimengine.ModeCommand)

	if action.Type != ActionKey || action.Key != "q" {
		t.Fatalf("action = %+v, want command-mode q key", action)
	}
}

func TestMapInputMapsCommandLineKeys(t *testing.T) {
	cases := map[string]string{
		":":     vimengine.KeyColon,
		"enter": vimengine.KeyEnter,
		"esc":   vimengine.KeyEsc,
		"w":     vimengine.KeyW,
		"b":     vimengine.KeyB,
		"e":     vimengine.KeyE,
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputInCommandModePassesSubstituteCharacters(t *testing.T) {
	for _, input := range []string{"s", "/", "%", "2", ",", "3", "D"} {
		action := MapInputForMode(input, vimengine.ModeCommand)
		if action.Type != ActionKey || action.Key != input {
			t.Fatalf("MapInputForMode(%q, command) = %+v, want key %q", input, action, input)
		}
	}
}

func TestMapInputMapsNavigationKeys(t *testing.T) {
	cases := map[string]string{
		"g": vimengine.KeyG,
		"G": vimengine.KeyShiftG,
		"0": vimengine.KeyZero,
		"$": vimengine.KeyDollar,
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputIgnoresUnknownInput(t *testing.T) {
	action := MapInput("space")
	if action.Type != ActionIgnored {
		t.Fatalf("action type = %q, want %q", action.Type, ActionIgnored)
	}
}

func TestRenderStateBuildsStableViewModel(t *testing.T) {
	score := scoring.Result{
		Passed:     true,
		ExactKeys:  true,
		Efficiency: 1,
		Grade:      scoring.GradeS,
	}
	state := scenario.State{
		ScenarioID: "door",
		Title:      "Open the door",
		Message:    "Door opened.",
		Status:     exerciseruntime.StatusSucceeded,
		Runtime: exerciseruntime.State{
			ExerciseID: "move-right",
			Status:     exerciseruntime.StatusSucceeded,
			Vim: vimengine.State{
				Mode:  vimengine.ModeNormal,
				Lines: []string{"abc"},
				Cursor: vimengine.Cursor{
					Row: 0,
					Col: 2,
				},
			},
			KeyTrace: []string{vimengine.KeyL, vimengine.KeyL},
			Attempts: 1,
		},
		Score:     &score,
		HintsUsed: 0,
	}

	view := RenderState(state)
	if view.Title != "Open the door" {
		t.Fatalf("Title = %q, want %q", view.Title, "Open the door")
	}
	if view.Message != "Door opened." {
		t.Fatalf("Message = %q, want %q", view.Message, "Door opened.")
	}
	if view.CursorRow != 0 || view.CursorCol != 2 {
		t.Fatalf("cursor = (%d,%d), want (0,2)", view.CursorRow, view.CursorCol)
	}
	if view.Grade != "S" {
		t.Fatalf("Grade = %q, want S", view.Grade)
	}
	if view.CommandLine != "" || view.LastCommand != "" {
		t.Fatalf("command fields = %q/%q, want empty", view.CommandLine, view.LastCommand)
	}
	if len(view.BufferLines) != 1 || view.BufferLines[0] != "abc" {
		t.Fatalf("BufferLines = %+v, want [abc]", view.BufferLines)
	}
}

func TestRenderStateIncludesCommandFields(t *testing.T) {
	state := scenario.State{
		Runtime: exerciseruntime.State{
			Vim: vimengine.State{
				Mode:        vimengine.ModeCommand,
				Lines:       []string{"abc"},
				CommandLine: "q",
				LastCommand: ":q!",
			},
		},
	}

	view := RenderState(state)

	if view.CommandLine != "q" || view.LastCommand != ":q!" {
		t.Fatalf("command fields = %q/%q, want q/:q!", view.CommandLine, view.LastCommand)
	}
}

func TestRenderStateCopiesSlices(t *testing.T) {
	state := scenario.State{
		Runtime: exerciseruntime.State{
			Vim: vimengine.State{
				Lines: []string{"abc"},
			},
			KeyTrace: []string{vimengine.KeyL},
		},
	}

	view := RenderState(state)
	state.Runtime.Vim.Lines[0] = "changed"
	state.Runtime.KeyTrace[0] = "changed"

	if view.BufferLines[0] != "abc" {
		t.Fatalf("BufferLines[0] = %q, want abc", view.BufferLines[0])
	}
	if view.KeyTrace[0] != vimengine.KeyL {
		t.Fatalf("KeyTrace[0] = %q, want %q", view.KeyTrace[0], vimengine.KeyL)
	}
}
