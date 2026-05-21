package vimengine

import (
	"strconv"
	"strings"
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
	KeyX      = "x"
	KeyR      = "r"
	KeyD      = "d"
	KeyC      = "c"
	KeyY      = "y"
	KeyP      = "p"
	KeyShiftP = "P"
	KeyI      = "i"
	KeyA      = "a"
	KeyShiftA = "A"
	KeyO      = "o"
	KeyShiftO = "O"
	KeyU      = "u"
	KeyCtrlR  = "ctrl+r"
)

const (
	pendingDeleteInner = KeyD + KeyI
	pendingChangeInner = KeyC + KeyI
	pendingYankInner   = KeyY + KeyI
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
	Register    Register
	UndoStack   []Snapshot
	RedoStack   []Snapshot
}

type Register struct {
	Text     string
	Lines    []string
	Linewise bool
}

type Snapshot struct {
	Lines  []string
	Cursor Cursor
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
	EventChanged         EventType = "changed"
	EventInsertMode      EventType = "insert_mode"
	EventYanked          EventType = "yanked"
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
		next.Cursor.Col = clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
		next.Cursor.DesiredCol = next.Cursor.Col
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

	if next.Mode == ModeInsert {
		next.PendingKey = ""
		return insertPrintable(next, key)
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
	case KeyX:
		return deleteCurrentChar(next, key)
	case KeyR:
		next.PendingKey = key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventPendingKey,
				Key:  key,
			}},
		}
	case KeyD, KeyC, KeyY:
		next.PendingKey = key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventPendingKey,
				Key:  key,
			}},
		}
	case KeyP:
		return putRegister(next, key, true)
	case KeyShiftP:
		return putRegister(next, key, false)
	case KeyI:
		return enterInsertMode(next, key, next.Cursor.Col)
	case KeyA:
		return enterInsertMode(next, key, next.Cursor.Col+1)
	case KeyShiftA:
		return enterInsertMode(next, key, lineRuneLen(next.Lines[next.Cursor.Row]))
	case KeyO:
		return openLine(next, key, true)
	case KeyShiftO:
		return openLine(next, key, false)
	case KeyU:
		return undoLastChange(next, key)
	case KeyCtrlR:
		return redoLastChange(next, key)
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

func enterInsertMode(state State, key string, insertCol int) Result {
	next := copyState(state)
	next.Mode = ModeInsert
	next.Cursor.Col = clampInsertCol(insertCol, next.Lines[next.Cursor.Row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventInsertMode,
			Key:  key,
		}},
	}
}

func openLine(state State, key string, below bool) Result {
	next := pushUndo(state)
	insertRow := next.Cursor.Row
	if below {
		insertRow++
	}
	next.Lines = append(next.Lines[:insertRow], append([]string{""}, next.Lines[insertRow:]...)...)
	next.Mode = ModeInsert
	next.Cursor.Row = insertRow
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventInsertMode,
			Key:  key,
		}},
	}
}

func insertPrintable(state State, key string) Result {
	next := pushUndo(state)
	inserted := []rune(key)
	if len(inserted) != 1 {
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "insert requires one character",
			}},
		}
	}
	line := []rune(next.Lines[next.Cursor.Row])
	col := clampInsertCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	line = append(line[:col], append(inserted, line[col:]...)...)
	next.Lines[next.Cursor.Row] = string(line)
	next.Cursor.Col = col + 1
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func applyPendingKey(state State, key string) Result {
	next := copyState(state)
	pending := next.PendingKey
	next.PendingKey = ""
	if pending == KeyG && key == KeyG {
		return moveDocumentStart(next, key)
	}
	if pending == KeyR {
		return replaceCurrentChar(next, key)
	}
	if pending == KeyD {
		return deleteWithMotion(next, key)
	}
	if pending == KeyC {
		return changeWithMotion(next, key)
	}
	if pending == KeyY {
		return yankWithMotion(next, key)
	}
	if pending == pendingDeleteInner && key == KeyW {
		return deleteInnerWord(next, key)
	}
	if pending == pendingChangeInner && key == KeyW {
		return changeInnerWord(next, key)
	}
	if pending == pendingYankInner && key == KeyW {
		return yankInnerWord(next, key)
	}
	if pending == pendingDeleteInner || pending == pendingChangeInner || pending == pendingYankInner {
		return unsupportedTextObject(next, key)
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

func deleteWithMotion(state State, key string) Result {
	switch key {
	case KeyI:
		return enterInnerTextObjectPending(state, pendingDeleteInner, key)
	case KeyW:
		return deleteWordForward(state, key)
	case KeyDollar:
		return deleteToLineEnd(state, key)
	case KeyD:
		return deleteCurrentLine(state, key)
	default:
		return Result{
			State: copyState(state),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "delete sequence is not supported",
			}},
		}
	}
}

func yankWithMotion(state State, key string) Result {
	switch key {
	case KeyI:
		return enterInnerTextObjectPending(state, pendingYankInner, key)
	case KeyW:
		return yankWordForward(state, key)
	case KeyDollar:
		return yankToLineEnd(state, key)
	case KeyY:
		return yankCurrentLine(state, key)
	default:
		return Result{
			State: copyState(state),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "yank sequence is not supported",
			}},
		}
	}
}

func changeWithMotion(state State, key string) Result {
	switch key {
	case KeyI:
		return enterInnerTextObjectPending(state, pendingChangeInner, key)
	case KeyW:
		return changeWordForward(state, key)
	case KeyDollar:
		return changeToLineEnd(state, key)
	case KeyC:
		return changeCurrentLine(state, key)
	default:
		return Result{
			State: copyState(state),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "change sequence is not supported",
			}},
		}
	}
}

func enterInnerTextObjectPending(state State, pending string, key string) Result {
	next := copyState(state)
	next.PendingKey = pending
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventPendingKey,
			Key:  key,
		}},
	}
}

func unsupportedTextObject(state State, key string) Result {
	return Result{
		State: copyState(state),
		Events: []Event{{
			Type:    EventUnsupportedKey,
			Key:     key,
			Message: "text object sequence is not supported",
		}},
	}
}

func applyCommandKey(state State, key string) Result {
	next := copyState(state)
	switch key {
	case KeyEnter:
		command := ":" + next.CommandLine
		if command == ":q!" || command == ":wq" {
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
		}
		substituted, ok := applySubstituteCommand(next, command)
		if !ok {
			return Result{
				State: copyState(next),
				Events: []Event{{
					Type:    EventUnsupportedKey,
					Key:     key,
					Message: "command is not supported",
				}},
			}
		}
		next = substituted
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

func deleteWordForward(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	end, ok := deleteWordForwardEnd(line, start)
	if !ok || end <= start {
		return boundary(state, key)
	}
	return deleteLineRange(state, key, start, end)
}

func deleteWordForwardEnd(line []rune, start int) (int, bool) {
	if start >= len(line) {
		return len(line), false
	}
	index := start
	currentClass := classifyRune(line[index])
	if currentClass != cellSpace {
		for index < len(line) && classifyRune(line[index]) == currentClass {
			index++
		}
	}
	for index < len(line) && classifyRune(line[index]) == cellSpace {
		index++
	}
	if index > start {
		return index, true
	}
	return len(line), len(line) > start
}

func yankWordForward(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	end, ok := deleteWordForwardEnd(line, start)
	if !ok || end <= start {
		return boundary(state, key)
	}
	return yankLineRange(state, key, start, end)
}

func changeWordForward(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	end, ok := deleteWordForwardEnd(line, start)
	if !ok || end <= start {
		return boundary(state, key)
	}
	return changeLineRange(state, key, start, end)
}

func deleteInnerWord(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerWordRange(line, state.Cursor.Col)
	if !ok {
		return boundary(state, key)
	}
	return deleteLineRange(state, key, start, end)
}

func changeInnerWord(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerWordRange(line, state.Cursor.Col)
	if !ok {
		return boundary(state, key)
	}
	return changeLineRange(state, key, start, end)
}

func yankInnerWord(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerWordRange(line, state.Cursor.Col)
	if !ok {
		return boundary(state, key)
	}
	return yankLineRange(state, key, start, end)
}

func innerWordRange(line []rune, cursorCol int) (int, int, bool) {
	if len(line) == 0 {
		return 0, 0, false
	}
	if cursorCol < 0 {
		cursorCol = 0
	}
	if cursorCol >= len(line) {
		cursorCol = len(line) - 1
	}
	class := classifyRune(line[cursorCol])
	if class == cellSpace {
		return 0, 0, false
	}
	start := cursorCol
	for start > 0 && classifyRune(line[start-1]) == class {
		start--
	}
	end := cursorCol + 1
	for end < len(line) && classifyRune(line[end]) == class {
		end++
	}
	return start, end, true
}

func deleteToLineEnd(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	return deleteLineRange(state, key, start, len(line))
}

func yankToLineEnd(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	return yankLineRange(state, key, start, len(line))
}

func changeToLineEnd(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	if len(line) == 0 {
		return boundary(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	return changeLineRange(state, key, start, len(line))
}

func deleteLineRange(state State, key string, start int, end int) Result {
	next := pushUndo(state)
	line := []rune(next.Lines[next.Cursor.Row])
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	if start >= end {
		return boundary(state, key)
	}
	line = append(line[:start], line[end:]...)
	next.Lines[next.Cursor.Row] = string(line)
	next.Cursor.Col = clampCol(start, next.Lines[next.Cursor.Row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func yankLineRange(state State, key string, start int, end int) Result {
	next := copyState(state)
	line := []rune(next.Lines[next.Cursor.Row])
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	if start >= end {
		return boundary(state, key)
	}
	next.Register = Register{
		Text: string(line[start:end]),
	}
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventYanked,
			Key:  key,
		}},
	}
}

func putRegister(state State, key string, after bool) Result {
	if state.Register.Linewise {
		return putLinewiseRegister(state, key, after)
	}
	return putCharwiseRegister(state, key, after)
}

func putCharwiseRegister(state State, key string, after bool) Result {
	if state.Register.Text == "" {
		return boundary(state, key)
	}
	next := pushUndo(state)
	line := []rune(next.Lines[next.Cursor.Row])
	inserted := []rune(next.Register.Text)
	col := clampInsertCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	insertCol := col
	if after && len(line) > 0 {
		insertCol = col + 1
	}
	if insertCol > len(line) {
		insertCol = len(line)
	}
	line = append(line[:insertCol], append(inserted, line[insertCol:]...)...)
	next.Lines[next.Cursor.Row] = string(line)
	next.Cursor.Col = insertCol + len(inserted) - 1
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func putLinewiseRegister(state State, key string, after bool) Result {
	if len(state.Register.Lines) == 0 {
		return boundary(state, key)
	}
	next := pushUndo(state)
	insertRow := next.Cursor.Row
	if after {
		insertRow++
	}
	lines := copyLines(next.Register.Lines)
	next.Lines = append(next.Lines[:insertRow], append(lines, next.Lines[insertRow:]...)...)
	next.Cursor.Row = insertRow
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func changeLineRange(state State, key string, start int, end int) Result {
	next := pushUndo(state)
	line := []rune(next.Lines[next.Cursor.Row])
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	if start >= end {
		return boundary(state, key)
	}
	line = append(line[:start], line[end:]...)
	next.Lines[next.Cursor.Row] = string(line)
	next.Mode = ModeInsert
	next.Cursor.Col = clampInsertCol(start, next.Lines[next.Cursor.Row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventInsertMode,
			Key:  key,
		}},
	}
}

func deleteCurrentLine(state State, key string) Result {
	next := pushUndo(state)
	if len(next.Lines) == 1 {
		next.Lines[0] = ""
		next.Cursor.Row = 0
		next.Cursor.Col = 0
		next.Cursor.DesiredCol = 0
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventChanged,
				Key:  key,
			}},
		}
	}
	row := next.Cursor.Row
	next.Lines = append(next.Lines[:row], next.Lines[row+1:]...)
	if row >= len(next.Lines) {
		row = len(next.Lines) - 1
	}
	next.Cursor.Row = row
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func yankCurrentLine(state State, key string) Result {
	next := copyState(state)
	next.Register = Register{
		Lines:    []string{next.Lines[next.Cursor.Row]},
		Linewise: true,
	}
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventYanked,
			Key:  key,
		}},
	}
}

func changeCurrentLine(state State, key string) Result {
	next := pushUndo(state)
	next.Lines[next.Cursor.Row] = ""
	next.Mode = ModeInsert
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventInsertMode,
			Key:  key,
		}},
	}
}

func deleteCurrentChar(state State, key string) Result {
	next := pushUndo(state)
	runes := []rune(next.Lines[next.Cursor.Row])
	if len(runes) == 0 {
		return boundary(state, key)
	}
	col := clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	runes = append(runes[:col], runes[col+1:]...)
	next.Lines[next.Cursor.Row] = string(runes)
	next.Cursor.Col = clampCol(col, next.Lines[next.Cursor.Row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func replaceCurrentChar(state State, key string) Result {
	next := pushUndo(state)
	replacement := []rune(key)
	if len(replacement) != 1 {
		return Result{
			State: copyState(state),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "replace requires one character",
			}},
		}
	}
	runes := []rune(next.Lines[next.Cursor.Row])
	if len(runes) == 0 {
		return boundary(state, key)
	}
	col := clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	runes[col] = replacement[0]
	next.Lines[next.Cursor.Row] = string(runes)
	next.Cursor.Col = col
	next.Cursor.DesiredCol = col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

type substituteCommand struct {
	startRow int
	endRow   int
	old      string
	new      string
	global   bool
}

func applySubstituteCommand(state State, command string) (State, bool) {
	parsed, ok := parseSubstituteCommand(state, command)
	if !ok {
		return State{}, false
	}
	next := pushUndo(state)
	for row := parsed.startRow; row <= parsed.endRow; row++ {
		if parsed.global {
			next.Lines[row] = strings.ReplaceAll(next.Lines[row], parsed.old, parsed.new)
			continue
		}
		next.Lines[row] = strings.Replace(next.Lines[row], parsed.old, parsed.new, 1)
	}
	next.Mode = ModeNormal
	next.CommandLine = ""
	next.LastCommand = command
	next.Cursor.Col = clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return next, true
}

func undoLastChange(state State, key string) Result {
	if len(state.UndoStack) == 0 {
		return boundary(state, key)
	}
	next := copyState(state)
	snapshot := next.UndoStack[len(next.UndoStack)-1]
	next.UndoStack = next.UndoStack[:len(next.UndoStack)-1]
	next.RedoStack = append(next.RedoStack, snapshotState(state))
	applySnapshot(&next, snapshot)
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func redoLastChange(state State, key string) Result {
	if len(state.RedoStack) == 0 {
		return boundary(state, key)
	}
	next := copyState(state)
	snapshot := next.RedoStack[len(next.RedoStack)-1]
	next.RedoStack = next.RedoStack[:len(next.RedoStack)-1]
	next.UndoStack = append(next.UndoStack, snapshotState(state))
	applySnapshot(&next, snapshot)
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func pushUndo(state State) State {
	next := copyState(state)
	next.UndoStack = append(next.UndoStack, snapshotState(state))
	next.RedoStack = nil
	return next
}

func snapshotState(state State) Snapshot {
	return Snapshot{
		Lines:  copyLines(state.Lines),
		Cursor: state.Cursor,
	}
}

func applySnapshot(state *State, snapshot Snapshot) {
	state.Mode = ModeNormal
	state.CommandLine = ""
	state.PendingKey = ""
	state.Lines = copyLines(snapshot.Lines)
	if len(state.Lines) == 0 {
		state.Lines = []string{""}
	}
	state.Cursor = snapshot.Cursor
	state.Cursor.Col = clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	state.Cursor.DesiredCol = state.Cursor.Col
}

func parseSubstituteCommand(state State, command string) (substituteCommand, bool) {
	if !strings.HasPrefix(command, ":") {
		return substituteCommand{}, false
	}
	body := strings.TrimPrefix(command, ":")
	substituteIndex := strings.Index(body, "s/")
	if substituteIndex < 0 {
		return substituteCommand{}, false
	}
	rangeSpec := body[:substituteIndex]
	rest := body[substituteIndex+2:]
	parts := strings.Split(rest, "/")
	if len(parts) < 2 {
		return substituteCommand{}, false
	}
	old := parts[0]
	if old == "" {
		return substituteCommand{}, false
	}
	newValue := parts[1]
	flags := ""
	if len(parts) > 2 {
		flags = parts[2]
	}
	startRow, endRow, ok := parseSubstituteRange(state, rangeSpec)
	if !ok {
		return substituteCommand{}, false
	}
	return substituteCommand{
		startRow: startRow,
		endRow:   endRow,
		old:      old,
		new:      newValue,
		global:   strings.Contains(flags, "g"),
	}, true
}

func parseSubstituteRange(state State, rangeSpec string) (int, int, bool) {
	switch rangeSpec {
	case "":
		return state.Cursor.Row, state.Cursor.Row, true
	case "%":
		return 0, len(state.Lines) - 1, true
	default:
		parts := strings.Split(rangeSpec, ",")
		if len(parts) != 2 {
			return 0, 0, false
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, false
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, false
		}
		start--
		end--
		if start < 0 || end < start || end >= len(state.Lines) {
			return 0, 0, false
		}
		return start, end, true
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
	if next.Mode == ModeInsert {
		next.Cursor.Col = clampInsertCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	} else {
		next.Cursor.Col = clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
	}
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

func clampInsertCol(col int, line string) int {
	if col < 0 {
		return 0
	}
	maxCol := lineRuneLen(line)
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

func lineRuneLen(line string) int {
	return utf8.RuneCountInString(line)
}

func copyState(state State) State {
	next := state
	next.Lines = copyLines(state.Lines)
	next.Register = copyRegister(state.Register)
	next.UndoStack = copySnapshots(state.UndoStack)
	next.RedoStack = copySnapshots(state.RedoStack)
	return next
}

func copyRegister(register Register) Register {
	return Register{
		Text:     register.Text,
		Lines:    copyLines(register.Lines),
		Linewise: register.Linewise,
	}
}

func copySnapshots(snapshots []Snapshot) []Snapshot {
	if snapshots == nil {
		return nil
	}
	next := make([]Snapshot, len(snapshots))
	for index, snapshot := range snapshots {
		next[index] = Snapshot{
			Lines:  copyLines(snapshot.Lines),
			Cursor: snapshot.Cursor,
		}
	}
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
