package vimengine

import "unicode/utf8"

type Mode string

const (
	ModeNormal  Mode = "normal"
	ModeInsert  Mode = "insert"
	ModeCommand Mode = "command"
)

const (
	KeyEsc = "esc"
	KeyH   = "h"
	KeyJ   = "j"
	KeyK   = "k"
	KeyL   = "l"
)

type Cursor struct {
	Row        int
	Col        int
	DesiredCol int
}

type State struct {
	Mode   Mode
	Lines  []string
	Cursor Cursor
}

type EventType string

const (
	EventMoved          EventType = "moved"
	EventBoundary       EventType = "boundary"
	EventUnsupportedKey EventType = "unsupported_key"
	EventModeReset      EventType = "mode_reset"
)

type Event struct {
	Type    EventType
	Key     string
	Message string
}

type Result struct {
	State  State
	Events []Event
}

type Engine struct {
	state State
}

func New(lines []string) *Engine {
	return &Engine{state: NewState(lines)}
}

func NewWithState(state State) *Engine {
	return &Engine{state: normalizeState(state)}
}

func NewState(lines []string) State {
	state := State{
		Mode:  ModeNormal,
		Lines: copyLines(lines),
	}
	if len(state.Lines) == 0 {
		state.Lines = []string{""}
	}
	return normalizeState(state)
}

func (e *Engine) State() State {
	return copyState(e.state)
}

func (e *Engine) Apply(key string) Result {
	result := Apply(e.state, key)
	e.state = result.State
	return result
}

func (e *Engine) ApplyKeys(keys []string) Result {
	result := ApplyKeys(e.state, keys)
	e.state = result.State
	return result
}

func ApplyKeys(state State, keys []string) Result {
	next := normalizeState(state)
	events := make([]Event, 0, len(keys))

	for _, key := range keys {
		result := Apply(next, key)
		next = result.State
		events = append(events, result.Events...)
	}

	return Result{
		State:  copyState(next),
		Events: events,
	}
}

func Apply(state State, key string) Result {
	next := normalizeState(state)

	if key == KeyEsc {
		next.Mode = ModeNormal
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventModeReset,
				Key:  key,
			}},
		}
	}

	if next.Mode != ModeNormal {
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "key is not supported outside normal mode",
			}},
		}
	}

	switch key {
	case KeyH:
		return moveHorizontal(next, key, -1)
	case KeyL:
		return moveHorizontal(next, key, 1)
	case KeyJ:
		return moveVertical(next, key, 1)
	case KeyK:
		return moveVertical(next, key, -1)
	default:
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "key is not supported",
			}},
		}
	}
}

func moveHorizontal(state State, key string, delta int) Result {
	next := copyState(state)
	lineMax := maxCursorCol(next.Lines[next.Cursor.Row])
	target := next.Cursor.Col + delta

	if target < 0 || target > lineMax {
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventBoundary,
				Key:  key,
			}},
		}
	}

	next.Cursor.Col = target
	next.Cursor.DesiredCol = target
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func moveVertical(state State, key string, delta int) Result {
	next := copyState(state)
	targetRow := next.Cursor.Row + delta

	if targetRow < 0 || targetRow >= len(next.Lines) {
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventBoundary,
				Key:  key,
			}},
		}
	}

	next.Cursor.Row = targetRow
	next.Cursor.Col = clampCol(next.Cursor.DesiredCol, next.Lines[targetRow])
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func normalizeState(state State) State {
	next := copyState(state)
	if len(next.Lines) == 0 {
		next.Lines = []string{""}
	}
	if next.Mode == "" {
		next.Mode = ModeNormal
	}
	if next.Cursor.Row < 0 {
		next.Cursor.Row = 0
	}
	if next.Cursor.Row >= len(next.Lines) {
		next.Cursor.Row = len(next.Lines) - 1
	}
	next.Cursor.Col = clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	if next.Cursor.DesiredCol < 0 {
		next.Cursor.DesiredCol = next.Cursor.Col
	}
	if next.Cursor.DesiredCol < next.Cursor.Col {
		next.Cursor.DesiredCol = next.Cursor.Col
	}
	return next
}

func clampCol(col int, line string) int {
	if col < 0 {
		return 0
	}
	maxCol := maxCursorCol(line)
	if col > maxCol {
		return maxCol
	}
	return col
}

func maxCursorCol(line string) int {
	length := utf8.RuneCountInString(line)
	if length == 0 {
		return 0
	}
	return length - 1
}

func copyState(state State) State {
	next := state
	next.Lines = copyLines(state.Lines)
	return next
}

func copyLines(lines []string) []string {
	if lines == nil {
		return nil
	}
	next := make([]string, len(lines))
	copy(next, lines)
	return next
}
