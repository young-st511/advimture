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
	CommandLine string
	LastCommand string
	BufferLines []string
	CursorRow   int
	CursorCol   int
	KeyTrace    []string
	Attempts    int
	HintsUsed   int
	Grade       string
}

func MapInput(input string) Action {
	return MapInputForMode(input, vimengine.ModeNormal)
}

func MapInputForMode(input string, mode vimengine.Mode) Action {
	normalized := strings.ToLower(strings.TrimSpace(input))
	if normalized == "ctrl+c" {
		return Action{Type: ActionQuit}
	}
	if mode == vimengine.ModeCommand {
		switch normalized {
		case "enter":
			return Action{Type: ActionKey, Key: vimengine.KeyEnter}
		case "esc":
			return Action{Type: ActionKey, Key: vimengine.KeyEsc}
		case "q", "w", "!":
			return Action{Type: ActionKey, Key: normalized}
		default:
			return Action{Type: ActionIgnored}
		}
	}

	switch normalized {
	case ":":
		return Action{Type: ActionKey, Key: vimengine.KeyColon}
	case "enter":
		return Action{Type: ActionKey, Key: vimengine.KeyEnter}
	case "esc":
		return Action{Type: ActionKey, Key: vimengine.KeyEsc}
	case "h", "left":
		return Action{Type: ActionKey, Key: vimengine.KeyH}
	case "j", "down":
		return Action{Type: ActionKey, Key: vimengine.KeyJ}
	case "k", "up":
		return Action{Type: ActionKey, Key: vimengine.KeyK}
	case "l", "right":
		return Action{Type: ActionKey, Key: vimengine.KeyL}
	case "w":
		return Action{Type: ActionKey, Key: vimengine.KeyW}
	case "b":
		return Action{Type: ActionKey, Key: vimengine.KeyB}
	case "e":
		return Action{Type: ActionKey, Key: vimengine.KeyE}
	case "?":
		return Action{Type: ActionHint}
	case "r":
		return Action{Type: ActionRetry}
	case "q":
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
		CommandLine: state.Runtime.Vim.CommandLine,
		LastCommand: state.Runtime.Vim.LastCommand,
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
