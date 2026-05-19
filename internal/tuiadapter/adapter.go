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
	trimmed := strings.TrimSpace(input)
	normalized := strings.ToLower(trimmed)
	if normalized == "ctrl+c" {
		return Action{Type: ActionKey, Key: "ctrl+c"}
	}
	if normalized == "ctrl+r" {
		return Action{Type: ActionKey, Key: vimengine.KeyCtrlR}
	}
	if mode == vimengine.ModeInsert {
		if normalized == "esc" {
			return Action{Type: ActionKey, Key: vimengine.KeyEsc}
		}
		if normalized == "space" {
			return Action{Type: ActionKey, Key: " "}
		}
		if len([]rune(input)) == 1 {
			return Action{Type: ActionKey, Key: input}
		}
		return Action{Type: ActionIgnored}
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
			if len([]rune(trimmed)) == 1 {
				return Action{Type: ActionKey, Key: trimmed}
			}
			return Action{Type: ActionIgnored}
		}
	}
	if trimmed == vimengine.KeyShiftG {
		return Action{Type: ActionKey, Key: vimengine.KeyShiftG}
	}
	if trimmed == vimengine.KeyShiftA {
		return Action{Type: ActionKey, Key: vimengine.KeyShiftA}
	}
	if trimmed == vimengine.KeyShiftP {
		return Action{Type: ActionKey, Key: vimengine.KeyShiftP}
	}

	switch normalized {
	case ":":
		return Action{Type: ActionKey, Key: vimengine.KeyColon}
	case "enter":
		return Action{Type: ActionKey, Key: vimengine.KeyEnter}
	case "esc":
		return Action{Type: ActionKey, Key: vimengine.KeyEsc}
	case "h":
		return Action{Type: ActionKey, Key: vimengine.KeyH}
	case "j":
		return Action{Type: ActionKey, Key: vimengine.KeyJ}
	case "k":
		return Action{Type: ActionKey, Key: vimengine.KeyK}
	case "l":
		return Action{Type: ActionKey, Key: vimengine.KeyL}
	case "left", "down", "up", "right":
		return Action{Type: ActionKey, Key: normalized}
	case "w":
		return Action{Type: ActionKey, Key: vimengine.KeyW}
	case "b":
		return Action{Type: ActionKey, Key: vimengine.KeyB}
	case "e":
		return Action{Type: ActionKey, Key: vimengine.KeyE}
	case "g":
		return Action{Type: ActionKey, Key: vimengine.KeyG}
	case "0":
		return Action{Type: ActionKey, Key: vimengine.KeyZero}
	case "$":
		return Action{Type: ActionKey, Key: vimengine.KeyDollar}
	case "x":
		return Action{Type: ActionKey, Key: vimengine.KeyX}
	case "r":
		return Action{Type: ActionKey, Key: vimengine.KeyR}
	case "d":
		return Action{Type: ActionKey, Key: vimengine.KeyD}
	case "c":
		return Action{Type: ActionKey, Key: vimengine.KeyC}
	case "y":
		return Action{Type: ActionKey, Key: vimengine.KeyY}
	case "p":
		return Action{Type: ActionKey, Key: vimengine.KeyP}
	case "i":
		return Action{Type: ActionKey, Key: vimengine.KeyI}
	case "a":
		return Action{Type: ActionKey, Key: vimengine.KeyA}
	case "u":
		return Action{Type: ActionKey, Key: vimengine.KeyU}
	case "?":
		return Action{Type: ActionHint}
	case "q":
		return Action{Type: ActionQuit}
	default:
		if len([]rune(trimmed)) == 1 {
			return Action{Type: ActionKey, Key: trimmed}
		}
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
