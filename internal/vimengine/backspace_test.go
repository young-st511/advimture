package vimengine

import "testing"

func TestBackspaceEditsTextModes(t *testing.T) {
	for _, mode := range []Mode{ModeInsert, ModeCommand, ModeSearch} {
		t.Run(string(mode), func(t *testing.T) {
			e := NewWithState(State{Mode: mode, Lines: []string{"한ab"}, Cursor: Cursor{Col: 1}, CommandLine: "한x"})
			state := e.Apply("backspace").State
			if mode == ModeInsert {
				if state.Lines[0] != "ab" || state.Cursor.Col != 0 {
					t.Fatal(state)
				}
			} else if state.CommandLine != "한" {
				t.Fatal(state)
			}
		})
	}
}
