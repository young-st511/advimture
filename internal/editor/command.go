package editor

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CommandResult represents the outcome of executing a command-mode command.
type CommandResult struct {
	Quit     bool
	Save     bool
	Error    string
	Message  string
	GotoLine int // -1 means no line jump
}

// ExecuteCommand parses and executes an ex command string.
func ExecuteCommand(cmd string, buf *Buffer, cur *Cursor, undo *UndoManager) CommandResult {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return CommandResult{GotoLine: -1}
	}

	// :w, :q, :q!, :wq, :wq!
	switch cmd {
	case "w":
		return CommandResult{Save: true, Message: "저장되었습니다", GotoLine: -1}
	case "q":
		return CommandResult{Quit: true, GotoLine: -1}
	case "q!":
		return CommandResult{Quit: true, GotoLine: -1}
	case "wq", "wq!":
		return CommandResult{Save: true, Quit: true, GotoLine: -1}
	case "x":
		return CommandResult{Save: true, Quit: true, GotoLine: -1}
	}

	// :{number} — go to line
	if n, err := strconv.Atoi(cmd); err == nil && n > 0 {
		return CommandResult{GotoLine: n - 1} // convert to 0-based
	}

	// :{from},{to}d — range delete
	if result, ok := tryRangeDelete(cmd, buf, cur, undo); ok {
		return result
	}

	// :s/old/new/g — substitute
	if result, ok := trySubstitute(cmd, buf, cur, undo); ok {
		return result
	}

	// :%s/old/new/g — global substitute
	if result, ok := tryGlobalSubstitute(cmd, buf, cur, undo); ok {
		return result
	}

	return CommandResult{
		Error:    fmt.Sprintf("알 수 없는 명령어: :%s", cmd),
		GotoLine: -1,
	}
}

var rangeDeleteRe = regexp.MustCompile(`^(\d+),(\d+)d$`)

func tryRangeDelete(cmd string, buf *Buffer, cur *Cursor, undo *UndoManager) (CommandResult, bool) {
	m := rangeDeleteRe.FindStringSubmatch(cmd)
	if m == nil {
		return CommandResult{}, false
	}
	from, _ := strconv.Atoi(m[1])
	to, _ := strconv.Atoi(m[2])
	from-- // 1-based to 0-based
	to--

	if from < 0 {
		from = 0
	}
	if to >= buf.LineCount() {
		to = buf.LineCount() - 1
	}
	if from > to {
		return CommandResult{Error: "잘못된 범위", GotoLine: -1}, true
	}

	deleted := buf.DeleteLines(from, to)
	undo.Record(Operation{
		Type: OpDeleteLines, Row: from, EndRow: to,
		Text: deleted, CursorRow: cur.Row, CursorCol: cur.Col,
	})

	count := to - from + 1
	if cur.Row >= buf.LineCount() {
		cur.Row = buf.LineCount() - 1
	}
	cur.Col = 0
	cur.DesiredCol = 0

	return CommandResult{
		Message:  fmt.Sprintf("%d줄 삭제됨", count),
		GotoLine: -1,
	}, true
}

var substituteRe = regexp.MustCompile(`^s/([^/]*)/([^/]*)(?:/(g?))?$`)

func trySubstitute(cmd string, buf *Buffer, cur *Cursor, undo *UndoManager) (CommandResult, bool) {
	m := substituteRe.FindStringSubmatch(cmd)
	if m == nil {
		return CommandResult{}, false
	}
	pattern := m[1]
	replacement := m[2]
	global := m[3] == "g"

	line := buf.GetLine(cur.Row)
	var newLine string
	var count int

	if global {
		newLine = strings.ReplaceAll(line, pattern, replacement)
		count = strings.Count(line, pattern)
	} else {
		newLine = strings.Replace(line, pattern, replacement, 1)
		if newLine != line {
			count = 1
		}
	}

	if count == 0 {
		return CommandResult{
			Error:    fmt.Sprintf("패턴을 찾을 수 없습니다: %s", pattern),
			GotoLine: -1,
		}, true
	}

	oldLine := buf.GetLine(cur.Row)
	buf.SetLine(cur.Row, newLine)
	undo.Record(Operation{
		Type: OpSetLine, Row: cur.Row, Text: newLine, OldText: oldLine,
		CursorRow: cur.Row, CursorCol: cur.Col,
	})

	return CommandResult{
		Message:  fmt.Sprintf("%d개 치환됨", count),
		GotoLine: -1,
	}, true
}

var globalSubstituteRe = regexp.MustCompile(`^%s/([^/]*)/([^/]*)(?:/(g?))?$`)

func tryGlobalSubstitute(cmd string, buf *Buffer, cur *Cursor, undo *UndoManager) (CommandResult, bool) {
	m := globalSubstituteRe.FindStringSubmatch(cmd)
	if m == nil {
		return CommandResult{}, false
	}
	pattern := m[1]
	replacement := m[2]
	global := m[3] == "g"

	totalCount := 0
	var children []Operation

	for i := 0; i < buf.LineCount(); i++ {
		line := buf.GetLine(i)
		var newLine string
		var lineCount int

		if global {
			lineCount = strings.Count(line, pattern)
			newLine = strings.ReplaceAll(line, pattern, replacement)
		} else {
			if strings.Contains(line, pattern) {
				newLine = strings.Replace(line, pattern, replacement, 1)
				lineCount = 1
			} else {
				continue
			}
		}

		if lineCount > 0 {
			children = append(children, Operation{
				Type: OpSetLine, Row: i, Text: newLine, OldText: line,
			})
			buf.SetLine(i, newLine)
			totalCount += lineCount
		}
	}

	if totalCount == 0 {
		return CommandResult{
			Error:    fmt.Sprintf("패턴을 찾을 수 없습니다: %s", pattern),
			GotoLine: -1,
		}, true
	}

	undo.Record(Operation{
		Type: OpComposite, Children: children,
		CursorRow: cur.Row, CursorCol: cur.Col,
	})

	return CommandResult{
		Message:  fmt.Sprintf("%d개 치환됨", totalCount),
		GotoLine: -1,
	}, true
}
