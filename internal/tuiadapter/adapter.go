package tuiadapter

import (
	"strings"

	"github.com/young-st511/advimture/internal/scenario"
	"github.com/young-st511/advimture/internal/vimengine"
)

type ActionType string

const (
	ActionKey     ActionType = "key"
	ActionHint    ActionType = "hint"
	ActionRetry   ActionType = "retry"
	ActionQuit    ActionType = "quit"
	ActionIgnored ActionType = "ignored"
)

type Action struct {
	Type ActionType
	Key  string
}

type ViewModel struct {
	ScenarioID  string
	Title       string
	Message     string
	Status      string
	Mode        string
	BufferLines []string
	CursorRow   int
	CursorCol   int
	KeyTrace    []string
	Attempts    int
	HintsUsed   int
	Grade       string
}

func MapInput(input string) Action {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "h", "left":
		return Action{Type: ActionKey, Key: vimengine.KeyH}
	case "j", "down":
		return Action{Type: ActionKey, Key: vimengine.KeyJ}
	case "k", "up":
		return Action{Type: ActionKey, Key: vimengine.KeyK}
	case "l", "right":
		return Action{Type: ActionKey, Key: vimengine.KeyL}
	case "?":
		return Action{Type: ActionHint}
	case "r":
		return Action{Type: ActionRetry}
	case "q", "ctrl+c":
		return Action{Type: ActionQuit}
	default:
		return Action{Type: ActionIgnored}
	}
}

func RenderState(state scenario.State) ViewModel {
	view := ViewModel{
		ScenarioID:  state.ScenarioID,
		Title:       state.Title,
		Message:     state.Message,
		Status:      string(state.Status),
		Mode:        string(state.Runtime.Vim.Mode),
		BufferLines: copyStrings(state.Runtime.Vim.Lines),
		CursorRow:   state.Runtime.Vim.Cursor.Row,
		CursorCol:   state.Runtime.Vim.Cursor.Col,
		KeyTrace:    copyStrings(state.Runtime.KeyTrace),
		Attempts:    state.Runtime.Attempts,
		HintsUsed:   state.HintsUsed,
	}
	if state.Score != nil {
		view.Grade = string(state.Score.Grade)
	}
	return view
}

func copyStrings(values []string) []string {
	if values == nil {
		return nil
	}
	next := make([]string, len(values))
	copy(next, values)
	return next
}
