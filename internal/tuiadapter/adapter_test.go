package tuiadapter

import (
	"testing"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scenario"
	"github.com/young-st511/advimture/internal/scoring"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestMapInputMapsVimAndArrowKeys(t *testing.T) {
	cases := map[string]string{
		"h":     vimengine.KeyH,
		"left":  vimengine.KeyH,
		"j":     vimengine.KeyJ,
		"down":  vimengine.KeyJ,
		"k":     vimengine.KeyK,
		"up":    vimengine.KeyK,
		"l":     vimengine.KeyL,
		"right": vimengine.KeyL,
	}

	for input, wantKey := range cases {
		action := MapInput(input)
		if action.Type != ActionKey || action.Key != wantKey {
			t.Fatalf("MapInput(%q) = %+v, want key %q", input, action, wantKey)
		}
	}
}

func TestMapInputMapsCommands(t *testing.T) {
	cases := map[string]ActionType{
		"?":      ActionHint,
		"r":      ActionRetry,
		"ctrl+c": ActionQuit,
		"q":      ActionQuit,
	}

	for input, want := range cases {
		action := MapInput(input)
		if action.Type != want {
			t.Fatalf("MapInput(%q) = %q, want %q", input, action.Type, want)
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
	if len(view.BufferLines) != 1 || view.BufferLines[0] != "abc" {
		t.Fatalf("BufferLines = %+v, want [abc]", view.BufferLines)
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
