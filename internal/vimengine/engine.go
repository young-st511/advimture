package vimengine

import (
	"unicode"
	"unicode/utf8"
)

type Mode string

const (
	ModeNormal  Mode = "normal"
	ModeInsert  Mode = "insert"
	ModeCommand Mode = "command"
)

const (
	KeyEsc    = "esc"
	KeyEnter  = "enter"
	KeyColon  = ":"
	KeyH      = "h"
	KeyJ      = "j"
	KeyK      = "k"
	KeyL      = "l"
	KeyW      = "w"
	KeyB      = "b"
	KeyE      = "e"
	KeyG      = "g"
	KeyShiftG = "G"
	KeyZero   = "0"
	KeyDollar = "$"
)

type Cursor struct {
	Row        int
	Col        int
	DesiredCol int
}

type State struct {
	Mode        Mode
	Lines       []string
	Cursor      Cursor
	CommandLine string
	LastCommand string
	PendingKey  string
}

type EventType string

const (
	EventMoved           EventType = "moved"
	EventBoundary        EventType = "boundary"
	EventUnsupportedKey  EventType = "unsupported_key"
	EventModeReset       EventType = "mode_reset"
	EventCommandMode     EventType = "command_mode"
	EventCommandInput    EventType = "command_input"
	EventCommandExecuted EventType = "command_executed"
	EventPendingKey      EventType = "pending_key"
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
		next.CommandLine = ""
		next.PendingKey = ""
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventModeReset,
				Key:  key,
			}},
		}
	}

	if next.Mode == ModeCommand {
		next.PendingKey = ""
		return applyCommandKey(next, key)
	}

	if next.Mode != ModeNormal {
		next.PendingKey = ""
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "key is not supported outside normal mode",
			}},
		}
	}

	if next.PendingKey != "" {
		return applyPendingKey(next, key)
	}

	switch key {
	case KeyColon:
		next.Mode = ModeCommand
		next.CommandLine = ""
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventCommandMode,
				Key:  key,
			}},
		}
	case KeyH:
		return moveHorizontal(next, key, -1)
	case KeyL:
		return moveHorizontal(next, key, 1)
	case KeyJ:
		return moveVertical(next, key, 1)
	case KeyK:
		return moveVertical(next, key, -1)
	case KeyW:
		return moveWordForward(next, key)
	case KeyB:
		return moveWordBackward(next, key)
	case KeyE:
		return moveWordEnd(next, key)
	case KeyG:
		next.PendingKey = key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventPendingKey,
				Key:  key,
			}},
		}
	case KeyShiftG:
		return moveDocumentEnd(next, key)
	case KeyZero:
		return moveLineStart(next, key)
	case KeyDollar:
		return moveLineEnd(next, key)
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

func applyPendingKey(state State, key string) Result {
	next := copyState(state)
	pending := next.PendingKey
	next.PendingKey = ""
	if pending == KeyG && key == KeyG {
		return moveDocumentStart(next, key)
	}
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type:    EventUnsupportedKey,
			Key:     key,
			Message: "key sequence is not supported",
		}},
	}
}

func applyCommandKey(state State, key string) Result {
	next := copyState(state)
	switch key {
	case KeyEnter:
		command := ":" + next.CommandLine
		if command != ":q!" && command != ":wq" {
			return Result{
				State: copyState(next),
				Events: []Event{{
					Type:    EventUnsupportedKey,
					Key:     key,
					Message: "command is not supported",
				}},
			}
		}
		next.Mode = ModeNormal
		next.CommandLine = ""
		next.LastCommand = command
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventCommandExecuted,
				Key:  key,
			}},
		}
	case KeyColon:
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "nested command mode is not supported",
			}},
		}
	default:
		if len([]rune(key)) != 1 {
			return Result{
				State: copyState(next),
				Events: []Event{{
					Type:    EventUnsupportedKey,
					Key:     key,
					Message: "command input is not supported",
				}},
			}
		}
		next.CommandLine += key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventCommandInput,
				Key:  key,
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

func moveLineStart(state State, key string) Result {
	next := copyState(state)
	if next.Cursor.Col == 0 {
		return boundary(next, key)
	}
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func moveLineEnd(state State, key string) Result {
	next := copyState(state)
	targetCol := maxCursorCol(next.Lines[next.Cursor.Row])
	if next.Cursor.Col == targetCol {
		return boundary(next, key)
	}
	next.Cursor.Col = targetCol
	next.Cursor.DesiredCol = targetCol
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func moveDocumentStart(state State, key string) Result {
	next := copyState(state)
	if next.Cursor.Row == 0 && next.Cursor.Col == 0 {
		return boundary(next, key)
	}
	next.Cursor.Row = 0
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func moveDocumentEnd(state State, key string) Result {
	next := copyState(state)
	targetRow := len(next.Lines) - 1
	if next.Cursor.Row == targetRow && next.Cursor.Col == 0 {
		return boundary(next, key)
	}
	next.Cursor.Row = targetRow
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

type cellClass int

const (
	cellSpace cellClass = iota
	cellKeyword
	cellSymbol
)

type documentCell struct {
	row   int
	col   int
	class cellClass
}

func moveWordForward(state State, key string) Result {
	cells := documentCells(state.Lines)
	index, exact := cellIndex(cells, state.Cursor)
	if len(cells) == 0 {
		return boundary(state, key)
	}
	if exact && cells[index].class != cellSpace {
		index = endOfWord(cells, index) + 1
	}
	target, ok := nextWordStart(cells, index)
	if !ok {
		return boundary(state, key)
	}
	return moveToCell(state, key, cells[target])
}

func moveWordBackward(state State, key string) Result {
	cells := documentCells(state.Lines)
	index, exact := cellIndex(cells, state.Cursor)
	if len(cells) == 0 {
		return boundary(state, key)
	}
	if exact && cells[index].class != cellSpace && !isWordStart(cells, index) {
		return moveToCell(state, key, cells[startOfWord(cells, index)])
	}
	index--
	for index >= 0 && cells[index].class == cellSpace {
		index--
	}
	if index < 0 {
		return boundary(state, key)
	}
	return moveToCell(state, key, cells[startOfWord(cells, index)])
}

func moveWordEnd(state State, key string) Result {
	cells := documentCells(state.Lines)
	index, exact := cellIndex(cells, state.Cursor)
	if len(cells) == 0 {
		return boundary(state, key)
	}
	if exact && cells[index].class != cellSpace {
		end := endOfWord(cells, index)
		if end > index {
			return moveToCell(state, key, cells[end])
		}
		index = end + 1
	}
	start, ok := nextWordStart(cells, index)
	if !ok {
		return boundary(state, key)
	}
	return moveToCell(state, key, cells[endOfWord(cells, start)])
}

func documentCells(lines []string) []documentCell {
	var cells []documentCell
	for row, line := range lines {
		col := 0
		for _, r := range line {
			cells = append(cells, documentCell{
				row:   row,
				col:   col,
				class: classifyRune(r),
			})
			col++
		}
	}
	return cells
}

func classifyRune(r rune) cellClass {
	if unicode.IsSpace(r) {
		return cellSpace
	}
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return cellKeyword
	}
	return cellSymbol
}

func cellIndex(cells []documentCell, cursor Cursor) (int, bool) {
	for index, cell := range cells {
		if cell.row == cursor.Row && cell.col == cursor.Col {
			return index, true
		}
		if cell.row > cursor.Row || (cell.row == cursor.Row && cell.col > cursor.Col) {
			return index, false
		}
	}
	return len(cells), false
}

func nextWordStart(cells []documentCell, index int) (int, bool) {
	for index < len(cells) {
		if cells[index].class != cellSpace {
			return index, true
		}
		index++
	}
	return 0, false
}

func startOfWord(cells []documentCell, index int) int {
	for index > 0 && sameWord(cells[index-1], cells[index]) {
		index--
	}
	return index
}

func endOfWord(cells []documentCell, index int) int {
	for index+1 < len(cells) && sameWord(cells[index], cells[index+1]) {
		index++
	}
	return index
}

func isWordStart(cells []documentCell, index int) bool {
	return cells[index].class != cellSpace && (index == 0 || !sameWord(cells[index-1], cells[index]))
}

func sameWord(left documentCell, right documentCell) bool {
	if left.class == cellSpace || right.class == cellSpace || left.class != right.class {
		return false
	}
	return left.row == right.row && left.col+1 == right.col
}

func moveToCell(state State, key string, cell documentCell) Result {
	next := copyState(state)
	if next.Cursor.Row == cell.row && next.Cursor.Col == cell.col {
		return boundary(next, key)
	}
	next.Cursor.Row = cell.row
	next.Cursor.Col = cell.col
	next.Cursor.DesiredCol = cell.col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func boundary(state State, key string) Result {
	return Result{
		State: copyState(state),
		Events: []Event{{
			Type: EventBoundary,
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
	if next.Mode != ModeCommand {
		next.CommandLine = ""
	}
	if next.Mode != ModeNormal {
		next.PendingKey = ""
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
