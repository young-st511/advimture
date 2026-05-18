package game

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/data"
)

func TestTutorialQuitGoalCompletesOnQuit(t *testing.T) {
	tutorial := &data.TutorialData{
		ID:    "test-tutorial",
		Title: "test",
		Substeps: []data.SubstepData{
			{
				ID:          "quit-step",
				Title:       "quit",
				Instruction: "quit with :q!",
				InitialText: "hello",
				Goal: data.GoalData{
					Type: "quit",
				},
				AllowedKeys: []string{":", "q", "!", "enter"},
			},
		},
	}

	model := NewTutorial(tutorial)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if model.state != StatePractice {
		t.Fatalf("expected practice state, got %v", model.state)
	}

	for _, msg := range []tea.KeyMsg{
		{Type: tea.KeyRunes, Runes: []rune(":")},
		{Type: tea.KeyRunes, Runes: []rune("q")},
		{Type: tea.KeyRunes, Runes: []rune("!")},
		{Type: tea.KeyEnter},
	} {
		model, _ = model.Update(msg)
	}

	if model.quitting {
		t.Fatal("expected tutorial to stay open after completing quit goal")
	}
	if model.state != StateComplete {
		t.Fatalf("expected complete state, got %v", model.state)
	}
}
