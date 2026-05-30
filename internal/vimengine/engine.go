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
	ModeSearch  Mode = "search"
	ModeVisual  Mode = "visual"
)

const (
	KeyEsc         = "esc"
	KeyEnter       = "enter"
	KeyColon       = ":"
	KeySlash       = "/"
	KeyH           = "h"
	KeyJ           = "j"
	KeyK           = "k"
	KeyL           = "l"
	KeyW           = "w"
	KeyB           = "b"
	KeyE           = "e"
	KeyF           = "f"
	KeyT           = "t"
	KeyG           = "g"
	KeyShiftG      = "G"
	KeyZero        = "0"
	KeyDollar      = "$"
	KeyX           = "x"
	KeyR           = "r"
	KeyD           = "d"
	KeyC           = "c"
	KeyY           = "y"
	KeyV           = "v"
	KeyShiftV      = "V"
	KeyN           = "n"
	KeyShiftN      = "N"
	KeyP           = "p"
	KeyShiftP      = "P"
	KeyI           = "i"
	KeyA           = "a"
	KeyShiftA      = "A"
	KeyO           = "o"
	KeyShiftO      = "O"
	KeyDot         = "."
	KeyU           = "u"
	KeyCtrlR       = "ctrl+r"
	KeyDoubleQuote = "\""
	KeySingleQuote = "'"
)

const (
	pendingDeleteInner = KeyD + KeyI
	pendingChangeInner = KeyC + KeyI
	pendingYankInner   = KeyY + KeyI
	pendingDeleteFind  = KeyD + KeyF
	pendingDeleteTill  = KeyD + KeyT
	pendingChangeFind  = KeyC + KeyF
	pendingChangeTill  = KeyC + KeyT
)

type Cursor struct {
	Row        int
	Col        int
	DesiredCol int
}

type State struct {
	Mode              Mode
	Lines             []string
	Cursor            Cursor
	CommandLine       string
	LastCommand       string
	LastSearch        string
	LastSearchForward bool
	PendingKey        string
	Selection         *Selection
	Register          Register
	UndoStack         []Snapshot
	RedoStack         []Snapshot
	LastChange        []string
	Recording         ChangeRecording
}

type Register struct {
	Text     string
	Lines    []string
	Linewise bool
}

type ChangeRecording struct {
	Keys    []string
	Mutated bool
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
	return applyWithOptions(state, key, applyOptions{})
}

type applyOptions struct {
	replaying bool
}

func applyWithOptions(state State, key string, options applyOptions) Result {
	next := normalizeState(state)

	if key == KeyEsc {
		wasInsert := next.Mode == ModeInsert
		next.Mode = ModeNormal
		next.CommandLine = ""
		next.PendingKey = ""
		next.Selection = nil
		next.Cursor.Col = clampCol(next.Cursor.Col, next.Lines[next.Cursor.Row])
		next.Cursor.DesiredCol = next.Cursor.Col
		if wasInsert && !options.replaying {
			next = commitRecording(next, key)
		}
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventModeReset,
				Key:  key,
			}},
		}
	}

	if next.Mode == ModeSearch {
		next.PendingKey = ""
		return applySearchKey(next, key)
	}

	if next.Mode == ModeCommand {
		next.PendingKey = ""
		return applyCommandKey(next, key)
	}

	if next.Mode == ModeInsert {
		next.PendingKey = ""
		result := insertPrintable(next, key)
		if !options.replaying && hasEvent(result, EventChanged) {
			result.State = appendRecording(result.State, key, true)
		}
		return result
	}

	if next.Mode == ModeVisual {
		return applyVisualKey(next, key)
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
		result := applyPendingKey(next, key)
		if !options.replaying {
			result.State = recordPendingChange(next, key, result)
		}
		return result
	}

	switch key {
	case KeyShiftV:
		return enterVisualMode(next, key)
	case KeyV:
		return enterVisualMode(next, key)
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
	case KeySlash:
		next.Mode = ModeSearch
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
	case KeyF, KeyT:
		next.PendingKey = key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventPendingKey,
				Key:  key,
			}},
		}
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
		result := deleteCurrentChar(next, key)
		if !options.replaying && hasEvent(result, EventChanged) {
			result.State = rememberLastChange(result.State, []string{key})
		}
		return result
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
		result := enterInsertMode(next, key, next.Cursor.Col)
		if !options.replaying {
			result.State = startRecording(result.State, []string{key}, false)
		}
		return result
	case KeyA:
		result := enterInsertMode(next, key, next.Cursor.Col+1)
		if !options.replaying {
			result.State = startRecording(result.State, []string{key}, false)
		}
		return result
	case KeyShiftA:
		result := enterInsertMode(next, key, lineRuneLen(next.Lines[next.Cursor.Row]))
		if !options.replaying {
			result.State = startRecording(result.State, []string{key}, false)
		}
		return result
	case KeyO:
		result := openLine(next, key, true)
		if !options.replaying {
			result.State = startRecording(result.State, []string{key}, true)
		}
		return result
	case KeyShiftO:
		result := openLine(next, key, false)
		if !options.replaying {
			result.State = startRecording(result.State, []string{key}, true)
		}
		return result
	case KeyDot:
		return repeatLastChange(next, key)
	case KeyN:
		return repeatSearch(next, key, true)
	case KeyShiftN:
		return repeatSearch(next, key, false)
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

func repeatLastChange(state State, key string) Result {
	if len(state.LastChange) == 0 {
		return boundary(state, key)
	}
	next := copyState(state)
	sequence := copyLines(state.LastChange)
	var events []Event
	for _, replayKey := range sequence {
		result := applyWithOptions(next, replayKey, applyOptions{replaying: true})
		next = result.State
		events = append(events, result.Events...)
	}
	next.LastChange = sequence
	next.Recording = ChangeRecording{}
	if len(events) == 0 {
		events = []Event{{Type: EventChanged, Key: key}}
	}
	return Result{
		State:  copyState(next),
		Events: events,
	}
}

func recordPendingChange(state State, key string, result Result) State {
	if state.PendingKey == KeyR && hasEvent(result, EventChanged) {
		return rememberLastChange(result.State, []string{KeyR, key})
	}
	if !hasEvent(result, EventInsertMode) {
		return result.State
	}
	switch state.PendingKey {
	case KeyC:
		return startRecording(result.State, []string{KeyC, key}, true)
	case pendingChangeInner:
		return startRecording(result.State, []string{KeyC, KeyI, key}, true)
	case pendingChangeFind:
		return startRecording(result.State, []string{KeyC, KeyF, key}, true)
	case pendingChangeTill:
		return startRecording(result.State, []string{KeyC, KeyT, key}, true)
	default:
		return result.State
	}
}

func startRecording(state State, keys []string, mutated bool) State {
	next := copyState(state)
	next.Recording = ChangeRecording{
		Keys:    copyLines(keys),
		Mutated: mutated,
	}
	return next
}

func appendRecording(state State, key string, mutated bool) State {
	next := copyState(state)
	if len(next.Recording.Keys) == 0 {
		return next
	}
	next.Recording.Keys = append(next.Recording.Keys, key)
	next.Recording.Mutated = next.Recording.Mutated || mutated
	return next
}

func commitRecording(state State, key string) State {
	next := copyState(state)
	if len(next.Recording.Keys) == 0 {
		return next
	}
	if next.Recording.Mutated {
		keys := append(copyLines(next.Recording.Keys), key)
		next.LastChange = keys
	}
	next.Recording = ChangeRecording{}
	return next
}

func rememberLastChange(state State, keys []string) State {
	next := copyState(state)
	next.LastChange = copyLines(keys)
	next.Recording = ChangeRecording{}
	return next
}

func hasEvent(result Result, eventType EventType) bool {
	for _, event := range result.Events {
		if event.Type == eventType {
			return true
		}
	}
	return false
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
	if pending == KeyF {
		return moveToCharFind(next, key, false)
	}
	if pending == KeyT {
		return moveToCharFind(next, key, true)
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
	if pending == pendingDeleteInner && isSymmetricQuoteKey(key) {
		return deleteInnerQuote(next, key)
	}
	if pending == pendingChangeInner && isSymmetricQuoteKey(key) {
		return changeInnerQuote(next, key)
	}
	if pending == pendingYankInner && isSymmetricQuoteKey(key) {
		return yankInnerQuote(next, key)
	}
	if pending == pendingDeleteFind {
		return deleteWithCharFind(next, key, false)
	}
	if pending == pendingDeleteTill {
		return deleteWithCharFind(next, key, true)
	}
	if pending == pendingChangeFind {
		return changeWithCharFind(next, key, false)
	}
	if pending == pendingChangeTill {
		return changeWithCharFind(next, key, true)
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
	case KeyF:
		return enterCharFindPending(state, pendingDeleteFind, key)
	case KeyT:
		return enterCharFindPending(state, pendingDeleteTill, key)
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
	case KeyF:
		return enterCharFindPending(state, pendingChangeFind, key)
	case KeyT:
		return enterCharFindPending(state, pendingChangeTill, key)
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

func enterCharFindPending(state State, pending string, key string) Result {
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

func applySearchKey(state State, key string) Result {
	next := copyState(state)
	switch key {
	case KeyEnter:
		query := next.CommandLine
		if query == "" {
			next.Mode = ModeNormal
			next.CommandLine = ""
			return boundary(next, key)
		}
		next.Mode = ModeNormal
		next.CommandLine = ""
		next.LastCommand = KeySlash + query
		next.LastSearch = query
		next.LastSearchForward = true
		return moveToSearchMatch(next, key, query, true)
	case KeySlash:
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "nested search mode is not supported",
			}},
		}
	default:
		if len([]rune(key)) != 1 {
			return Result{
				State: copyState(next),
				Events: []Event{{
					Type:    EventUnsupportedKey,
					Key:     key,
					Message: "search input is not supported",
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

func repeatSearch(state State, key string, sameDirection bool) Result {
	if state.LastSearch == "" {
		return boundary(state, key)
	}
	forward := state.LastSearchForward
	if !sameDirection {
		forward = !forward
	}
	return moveToSearchMatch(state, key, state.LastSearch, forward)
}

func moveToSearchMatch(state State, key string, query string, forward bool) Result {
	if query == "" {
		return boundary(state, key)
	}
	next := copyState(state)
	match, ok := findLiteralMatch(next.Lines, next.Cursor, query, forward)
	if !ok {
		return boundary(next, key)
	}
	next.Cursor.Row = match.row
	next.Cursor.Col = match.col
	next.Cursor.DesiredCol = match.col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func findLiteralMatch(lines []string, cursor Cursor, query string, forward bool) (documentCell, bool) {
	if forward {
		if match, ok := findForwardLiteralMatch(lines, cursor, query, false); ok {
			return match, true
		}
		return findForwardLiteralMatch(lines, cursor, query, true)
	}
	if match, ok := findBackwardLiteralMatch(lines, cursor, query, false); ok {
		return match, true
	}
	return findBackwardLiteralMatch(lines, cursor, query, true)
}

func findForwardLiteralMatch(lines []string, cursor Cursor, query string, wrapped bool) (documentCell, bool) {
	startRow := cursor.Row
	endRow := len(lines) - 1
	if wrapped {
		startRow = 0
		endRow = cursor.Row
	}
	for row := startRow; row <= endRow; row++ {
		startCol := 0
		if row == cursor.Row {
			if wrapped {
				startCol = 0
			} else {
				startCol = cursor.Col + 1
			}
		}
		if matchCol, ok := indexLiteralFrom(lines[row], query, startCol); ok {
			if wrapped && row == cursor.Row && matchCol > cursor.Col {
				continue
			}
			return documentCell{row: row, col: matchCol}, true
		}
	}
	return documentCell{}, false
}

func findBackwardLiteralMatch(lines []string, cursor Cursor, query string, wrapped bool) (documentCell, bool) {
	startRow := cursor.Row
	endRow := 0
	if wrapped {
		startRow = len(lines) - 1
		endRow = cursor.Row
	}
	for row := startRow; row >= endRow; row-- {
		maxCol := lineRuneLen(lines[row])
		if row == cursor.Row {
			if wrapped {
				maxCol = lineRuneLen(lines[row])
			} else {
				maxCol = cursor.Col
			}
		}
		if matchCol, ok := lastIndexLiteralBefore(lines[row], query, maxCol); ok {
			if wrapped && row == cursor.Row && matchCol < cursor.Col {
				continue
			}
			return documentCell{row: row, col: matchCol}, true
		}
	}
	return documentCell{}, false
}

func indexLiteralFrom(line string, query string, startCol int) (int, bool) {
	runes := []rune(line)
	queryRunes := []rune(query)
	if len(queryRunes) == 0 || startCol > len(runes)-len(queryRunes) {
		return 0, false
	}
	if startCol < 0 {
		startCol = 0
	}
	for col := startCol; col <= len(runes)-len(queryRunes); col++ {
		if sameRunes(runes[col:col+len(queryRunes)], queryRunes) {
			return col, true
		}
	}
	return 0, false
}

func lastIndexLiteralBefore(line string, query string, beforeCol int) (int, bool) {
	runes := []rune(line)
	queryRunes := []rune(query)
	if len(queryRunes) == 0 {
		return 0, false
	}
	maxStart := beforeCol - len(queryRunes)
	if maxStart > len(runes)-len(queryRunes) {
		maxStart = len(runes) - len(queryRunes)
	}
	for col := maxStart; col >= 0; col-- {
		if sameRunes(runes[col:col+len(queryRunes)], queryRunes) {
			return col, true
		}
	}
	return 0, false
}

func sameRunes(left []rune, right []rune) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
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

func deleteInnerQuote(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerQuoteRangeForKey(line, state.Cursor.Col, key)
	if !ok {
		return boundary(state, key)
	}
	return deleteLineRange(state, key, start, end)
}

func changeInnerQuote(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerQuoteRangeForKey(line, state.Cursor.Col, key)
	if !ok {
		return boundary(state, key)
	}
	return changeLineRange(state, key, start, end)
}

func yankInnerQuote(state State, key string) Result {
	line := []rune(state.Lines[state.Cursor.Row])
	start, end, ok := innerQuoteRangeForKey(line, state.Cursor.Col, key)
	if !ok {
		return boundary(state, key)
	}
	return yankLineRange(state, key, start, end)
}

func moveToCharFind(state State, key string, till bool) Result {
	targetRunes := []rune(key)
	if len(targetRunes) != 1 {
		return unsupportedCharFindTarget(state, key)
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	match, ok := findCharForward(state.Lines[state.Cursor.Row], start, targetRunes[0])
	if !ok {
		return boundary(state, key)
	}
	target := match
	if till {
		target = match - 1
	}
	if target <= start {
		return boundary(state, key)
	}
	next := copyState(state)
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

func deleteWithCharFind(state State, key string, till bool) Result {
	start, end, ok := charFindRange(state, key, till)
	if !ok {
		return boundary(state, key)
	}
	return deleteLineRange(state, key, start, end)
}

func changeWithCharFind(state State, key string, till bool) Result {
	start, end, ok := charFindRange(state, key, till)
	if !ok {
		return boundary(state, key)
	}
	return changeLineRange(state, key, start, end)
}

func charFindRange(state State, key string, till bool) (int, int, bool) {
	targetRunes := []rune(key)
	if len(targetRunes) != 1 {
		return 0, 0, false
	}
	start := clampCol(state.Cursor.Col, state.Lines[state.Cursor.Row])
	match, ok := findCharForward(state.Lines[state.Cursor.Row], start, targetRunes[0])
	if !ok {
		return 0, 0, false
	}
	end := match + 1
	if till {
		end = match
	}
	if end <= start {
		return 0, 0, false
	}
	return start, end, true
}

func findCharForward(line string, cursorCol int, target rune) (int, bool) {
	runes := []rune(line)
	for col := cursorCol + 1; col < len(runes); col++ {
		if runes[col] == target {
			return col, true
		}
	}
	return 0, false
}

func unsupportedCharFindTarget(state State, key string) Result {
	return Result{
		State: copyState(state),
		Events: []Event{{
			Type:    EventUnsupportedKey,
			Key:     key,
			Message: "char find target must be a single character",
		}},
	}
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

func innerQuoteRange(line []rune, cursorCol int) (int, int, bool) {
	return innerDelimitedRange(line, cursorCol, '"')
}

func innerQuoteRangeForKey(line []rune, cursorCol int, key string) (int, int, bool) {
	if key == KeyDoubleQuote {
		return innerDelimitedRange(line, cursorCol, '"')
	}
	if key == KeySingleQuote {
		return innerDelimitedRange(line, cursorCol, '\'')
	}
	return 0, 0, false
}

func isSymmetricQuoteKey(key string) bool {
	return key == KeyDoubleQuote || key == KeySingleQuote
}

func innerDelimitedRange(line []rune, cursorCol int, delimiter rune) (int, int, bool) {
	if len(line) == 0 {
		return 0, 0, false
	}
	if cursorCol < 0 {
		cursorCol = 0
	}
	if cursorCol >= len(line) {
		cursorCol = len(line) - 1
	}
	if line[cursorCol] == delimiter {
		return 0, 0, false
	}

	left := -1
	for index := cursorCol - 1; index >= 0; index-- {
		if line[index] == delimiter {
			left = index
			break
		}
	}
	if left < 0 {
		return 0, 0, false
	}

	right := -1
	for index := cursorCol + 1; index < len(line); index++ {
		if line[index] == delimiter {
			right = index
			break
		}
	}
	if right < 0 || right <= left+1 {
		return 0, 0, false
	}
	return left + 1, right, true
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
	if next.Mode != ModeCommand && next.Mode != ModeSearch {
		next.CommandLine = ""
	}
	if next.Mode != ModeNormal && next.Mode != ModeVisual {
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
	if next.Mode == ModeVisual {
		if next.Selection == nil || !next.Selection.Active {
			next.Selection = normalizedSelectionForLines(Selection{
				Active: true,
				Kind:   SelectionCharwise,
				Anchor: next.Cursor,
				Head:   next.Cursor,
			}, next.Lines)
		} else {
			selection := *next.Selection
			selection.Anchor = clampSelectionCursor(selection.Anchor, next.Lines)
			selection.Head = clampSelectionCursor(selection.Head, next.Lines)
			next.Selection = normalizedSelectionForLines(selection, next.Lines)
		}
	} else {
		next.Selection = nil
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
	next.LastChange = copyLines(state.LastChange)
	if state.Selection != nil {
		selection := *state.Selection
		next.Selection = &selection
	}
	next.Recording = ChangeRecording{
		Keys:    copyLines(state.Recording.Keys),
		Mutated: state.Recording.Mutated,
	}
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
