package editor

// ExecuteResult holds the result of executing a parsed command.
type ExecuteResult struct {
	SwitchToInsert bool
	Quit           bool
}

// Execute applies a ParseResult to the buffer/cursor/register/undo system.
func Execute(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, pr *ParseResult) ExecuteResult {
	var result ExecuteResult

	// Simple commands
	if pr.SimpleCmd != SimpleCmdNone {
		result = executeSimpleCmd(buf, cur, reg, undo, pr)
		return result
	}

	// Pure motion (no operator)
	if pr.Operator == OperatorNone && pr.Motion != MotionNone {
		ExecuteMotion(buf, cur, pr.Motion, pr.Count, pr.Char)
		return result
	}

	// Linewise operator (dd, yy, cc)
	if pr.IsLinewise && pr.Operator != OperatorNone {
		result = executeLinewise(buf, cur, reg, undo, pr)
		return result
	}

	// Operator + motion
	if pr.Operator != OperatorNone && pr.Motion != MotionNone {
		result = executeOperatorMotion(buf, cur, reg, undo, pr)
		return result
	}

	// Operator + text object
	if pr.Operator != OperatorNone && pr.TextObject != TextObjectNone {
		result = executeOperatorTextObject(buf, cur, reg, undo, pr)
		return result
	}

	return result
}

func executeSimpleCmd(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, pr *ParseResult) ExecuteResult {
	var result ExecuteResult

	switch pr.SimpleCmd {
	case SimpleCmdInsertBefore:
		result.SwitchToInsert = true

	case SimpleCmdInsertAfter:
		lineLen := buf.LineRuneLen(cur.Row)
		if lineLen > 0 {
			cur.Col++
		}
		result.SwitchToInsert = true

	case SimpleCmdInsertLineStart:
		cur.Col = 0
		cur.DesiredCol = 0
		result.SwitchToInsert = true

	case SimpleCmdInsertLineEnd:
		cur.Col = buf.LineRuneLen(cur.Row)
		cur.DesiredCol = cur.Col
		result.SwitchToInsert = true

	case SimpleCmdOpenBelow:
		curRow := cur.Row
		undo.Record(Operation{
			Type: OpInsertLine, Row: curRow + 1, Text: "",
			CursorRow: cur.Row, CursorCol: cur.Col,
		})
		buf.InsertLine(curRow+1, "")
		cur.Row = curRow + 1
		cur.Col = 0
		cur.DesiredCol = 0
		result.SwitchToInsert = true

	case SimpleCmdOpenAbove:
		curRow := cur.Row
		undo.Record(Operation{
			Type: OpInsertLine, Row: curRow, Text: "",
			CursorRow: cur.Row, CursorCol: cur.Col,
		})
		buf.InsertLine(curRow, "")
		cur.Row = curRow
		cur.Col = 0
		cur.DesiredCol = 0
		result.SwitchToInsert = true

	case SimpleCmdDeleteChar:
		var deleted []rune
		var children []Operation
		saveCurRow, saveCurCol := cur.Row, cur.Col
		for i := 0; i < pr.Count; i++ {
			lineLen := buf.LineRuneLen(cur.Row)
			if lineLen == 0 {
				break
			}
			ch := buf.DeleteChar(cur.Row, cur.Col)
			if ch != 0 {
				deleted = append(deleted, ch)
				children = append(children, Operation{
					Type: OpDeleteChar, Row: cur.Row, Col: cur.Col, Char: ch,
				})
			}
			newLen := buf.LineRuneLen(cur.Row)
			if cur.Col >= newLen && newLen > 0 {
				cur.Col = newLen - 1
			}
		}
		if len(children) > 0 {
			if len(children) == 1 {
				children[0].CursorRow = saveCurRow
				children[0].CursorCol = saveCurCol
				undo.Record(children[0])
			} else {
				undo.Record(Operation{
					Type: OpComposite, Children: children,
					CursorRow: saveCurRow, CursorCol: saveCurCol,
				})
			}
			reg.Set(string(deleted), false)
		}
		cur.DesiredCol = cur.Col

	case SimpleCmdReplaceChar:
		lineLen := buf.LineRuneLen(cur.Row)
		if lineLen > 0 && cur.Col < lineLen {
			runes := buf.GetLineRunes(cur.Row)
			runes[cur.Col] = pr.Char
			oldLine := buf.GetLine(cur.Row)
			buf.SetLine(cur.Row, string(runes))
			undo.Record(Operation{
				Type: OpSetLine, Row: cur.Row, Text: string(runes), OldText: oldLine,
				CursorRow: cur.Row, CursorCol: cur.Col,
			})
		}

	case SimpleCmdJoinLine:
		if cur.Row < buf.LineCount()-1 {
			joinCol := buf.LineRuneLen(cur.Row)
			undo.Record(Operation{
				Type: OpJoinLines, Row: cur.Row, Col: joinCol,
				CursorRow: cur.Row, CursorCol: cur.Col,
			})
			buf.JoinLines(cur.Row)
			cur.Col = joinCol
			cur.DesiredCol = cur.Col
		}

	case SimpleCmdPaste:
		executePaste(buf, cur, reg, undo, false)

	case SimpleCmdPasteBefore:
		executePaste(buf, cur, reg, undo, true)

	case SimpleCmdUndo:
		if row, col, ok := undo.Undo(buf); ok {
			cur.MoveTo(row, col)
			cur.ClampToBuffer(buf, false)
		}

	case SimpleCmdRedo:
		if row, col, ok := undo.Redo(buf); ok {
			cur.MoveTo(row, col)
			cur.ClampToBuffer(buf, false)
		}
	}

	return result
}

func executePaste(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, before bool) {
	if reg.Content == "" {
		return
	}

	if reg.Linewise {
		var insertRow int
		if before {
			insertRow = cur.Row
		} else {
			insertRow = cur.Row + 1
		}
		lines := splitLines(reg.Content)
		for i, line := range lines {
			buf.InsertLine(insertRow+i, line)
		}
		var children []Operation
		for i := len(lines) - 1; i >= 0; i-- {
			children = append(children, Operation{
				Type: OpInsertLine, Row: insertRow + i, Text: lines[i],
			})
		}
		undo.Record(Operation{
			Type: OpComposite, Children: children,
			CursorRow: cur.Row, CursorCol: cur.Col,
		})

		cur.Row = insertRow
		cur.Col = 0
		cur.DesiredCol = 0
	} else {
		// Charwise paste
		var insertCol int
		if before {
			insertCol = cur.Col
		} else {
			insertCol = cur.Col + 1
		}
		buf.InsertText(cur.Row, insertCol, reg.Content)

		// Calculate end position for undo
		pastedRunes := []rune(reg.Content)
		endRow := cur.Row
		endCol := insertCol
		for _, r := range pastedRunes {
			if r == '\n' {
				endRow++
				endCol = 0
			} else {
				endCol++
			}
		}

		undo.Record(Operation{
			Type: OpInsertText, Row: cur.Row, Col: insertCol,
			EndRow: endRow, EndCol: endCol,
			Text: reg.Content, CursorRow: cur.Row, CursorCol: cur.Col,
		})

		// Cursor at end of pasted text - 1
		if len(pastedRunes) > 0 {
			cur.Row = endRow
			cur.Col = endCol - 1
			if cur.Col < 0 {
				cur.Col = 0
			}
		}
		cur.DesiredCol = cur.Col
	}
}

func executeLinewise(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, pr *ParseResult) ExecuteResult {
	var result ExecuteResult
	startRow := cur.Row
	endRow := startRow + pr.Count - 1
	if endRow >= buf.LineCount() {
		endRow = buf.LineCount() - 1
	}

	if pr.Operator == OperatorYank {
		// Yank: copy without deleting
		var content string
		for i := startRow; i <= endRow; i++ {
			if i > startRow {
				content += "\n"
			}
			content += buf.GetLine(i)
		}
		reg.Set(content, true)
		return result
	}

	deleted := buf.DeleteLines(startRow, endRow)
	reg.Set(deleted, true)

	undo.Record(Operation{
		Type: OpDeleteLines, Row: startRow, EndRow: endRow,
		Text: deleted, CursorRow: cur.Row, CursorCol: cur.Col,
	})

	// Clamp cursor
	if cur.Row >= buf.LineCount() {
		cur.Row = buf.LineCount() - 1
	}
	cur.Col = 0
	cur.DesiredCol = 0

	if pr.Operator == OperatorChange {
		result.SwitchToInsert = true
	}

	return result
}

func executeOperatorMotion(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, pr *ParseResult) ExecuteResult {
	var result ExecuteResult

	sr, sc, er, ec := MotionRange(buf, cur, pr.Motion, pr.Count, pr.Char)

	if pr.Operator == OperatorYank {
		// Yank: copy the range without deleting
		yanked := getRangeText(buf, sr, sc, er, ec)
		reg.Set(yanked, false)
		return result
	}

	deleted := buf.DeleteRange(sr, sc, er, ec)
	reg.Set(deleted, false)

	undo.Record(Operation{
		Type: OpDeleteRange, Row: sr, Col: sc, EndRow: er, EndCol: ec,
		Text: deleted, CursorRow: cur.Row, CursorCol: cur.Col,
	})

	cur.Row = sr
	cur.Col = sc
	cur.ClampToBuffer(buf, false)
	cur.DesiredCol = cur.Col

	if pr.Operator == OperatorChange {
		result.SwitchToInsert = true
	}

	return result
}

// getRangeText extracts text from the buffer in the given range without deleting.
func getRangeText(buf *Buffer, sr, sc, er, ec int) string {
	if sr == er {
		runes := buf.GetLineRunes(sr)
		if sc > len(runes) {
			sc = len(runes)
		}
		if ec > len(runes) {
			ec = len(runes)
		}
		return string(runes[sc:ec])
	}
	// Multi-line
	var result string
	firstRunes := buf.GetLineRunes(sr)
	if sc > len(firstRunes) {
		sc = len(firstRunes)
	}
	result = string(firstRunes[sc:])
	for i := sr + 1; i < er; i++ {
		result += "\n" + buf.GetLine(i)
	}
	lastRunes := buf.GetLineRunes(er)
	if ec > len(lastRunes) {
		ec = len(lastRunes)
	}
	result += "\n" + string(lastRunes[:ec])
	return result
}

func executeOperatorTextObject(buf *Buffer, cur *Cursor, reg *Register, undo *UndoManager, pr *ParseResult) ExecuteResult {
	var result ExecuteResult

	sr, sc, er, ec, ok := TextObjectRange(buf, cur, pr.TextObject)
	if !ok {
		return result
	}

	if pr.Operator == OperatorYank {
		yanked := getRangeText(buf, sr, sc, er, ec)
		reg.Set(yanked, false)
		return result
	}

	deleted := buf.DeleteRange(sr, sc, er, ec)
	reg.Set(deleted, false)

	undo.Record(Operation{
		Type: OpDeleteRange, Row: sr, Col: sc, EndRow: er, EndCol: ec,
		Text: deleted, CursorRow: cur.Row, CursorCol: cur.Col,
	})

	cur.Row = sr
	cur.Col = sc
	cur.ClampToBuffer(buf, false)
	cur.DesiredCol = cur.Col

	if pr.Operator == OperatorChange {
		result.SwitchToInsert = true
	}

	return result
}

// TextObjectRange returns the range for a text object at the cursor position.
// Returns (startRow, startCol, endRow, endCol, found).
func TextObjectRange(buf *Buffer, cur *Cursor, to TextObjectType) (int, int, int, int, bool) {
	runes := buf.GetLineRunes(cur.Row)
	if len(runes) == 0 {
		return 0, 0, 0, 0, false
	}

	switch to {
	case TextObjectInnerWord:
		return innerWordRange(runes, cur.Row, cur.Col)
	case TextObjectAWord:
		return aWordRange(runes, cur.Row, cur.Col)
	case TextObjectInnerDoubleQuote:
		return innerQuoteRange(runes, cur.Row, '"')
	case TextObjectADoubleQuote:
		return aQuoteRange(runes, cur.Row, '"')
	case TextObjectInnerSingleQuote:
		return innerQuoteRange(runes, cur.Row, '\'')
	case TextObjectASingleQuote:
		return aQuoteRange(runes, cur.Row, '\'')
	case TextObjectInnerParen:
		return innerPairRange(buf, cur, '(', ')')
	case TextObjectAParen:
		return aPairRange(buf, cur, '(', ')')
	case TextObjectInnerBrace:
		return innerPairRange(buf, cur, '{', '}')
	case TextObjectABrace:
		return aPairRange(buf, cur, '{', '}')
	}
	return 0, 0, 0, 0, false
}

func innerWordRange(runes []rune, row, col int) (int, int, int, int, bool) {
	if col >= len(runes) {
		return 0, 0, 0, 0, false
	}
	cc := classifyRune(runes[col])
	start := col
	for start > 0 && classifyRune(runes[start-1]) == cc {
		start--
	}
	end := col
	for end < len(runes)-1 && classifyRune(runes[end+1]) == cc {
		end++
	}
	return row, start, row, end + 1, true
}

func aWordRange(runes []rune, row, col int) (int, int, int, int, bool) {
	_, sc, _, ec, ok := innerWordRange(runes, row, col)
	if !ok {
		return 0, 0, 0, 0, false
	}
	origEc := ec
	// Include trailing whitespace
	for ec < len(runes) && classifyRune(runes[ec]) == classSpace {
		ec++
	}
	// If no trailing whitespace was added, include leading whitespace
	if ec == origEc {
		for sc > 0 && classifyRune(runes[sc-1]) == classSpace {
			sc--
		}
	}
	return row, sc, row, ec, true
}

func innerQuoteRange(runes []rune, row int, quote rune) (int, int, int, int, bool) {
	// Find opening and closing quotes
	first := -1
	second := -1
	for i, r := range runes {
		if r == quote {
			if first == -1 {
				first = i
			} else {
				second = i
				break
			}
		}
	}
	if first == -1 || second == -1 {
		return 0, 0, 0, 0, false
	}
	return row, first + 1, row, second, true
}

func aQuoteRange(runes []rune, row int, quote rune) (int, int, int, int, bool) {
	first := -1
	second := -1
	for i, r := range runes {
		if r == quote {
			if first == -1 {
				first = i
			} else {
				second = i
				break
			}
		}
	}
	if first == -1 || second == -1 {
		return 0, 0, 0, 0, false
	}
	return row, first, row, second + 1, true
}

func innerPairRange(buf *Buffer, cur *Cursor, open, close rune) (int, int, int, int, bool) {
	// Simple single-line implementation for now
	runes := buf.GetLineRunes(cur.Row)
	openIdx := -1
	closeIdx := -1

	// Search backward for open
	for i := cur.Col; i >= 0; i-- {
		if runes[i] == open {
			openIdx = i
			break
		}
	}
	if openIdx == -1 {
		return 0, 0, 0, 0, false
	}

	// Search forward for close
	depth := 0
	for i := openIdx; i < len(runes); i++ {
		if runes[i] == open {
			depth++
		} else if runes[i] == close {
			depth--
			if depth == 0 {
				closeIdx = i
				break
			}
		}
	}
	if closeIdx == -1 {
		return 0, 0, 0, 0, false
	}

	return cur.Row, openIdx + 1, cur.Row, closeIdx, true
}

func aPairRange(buf *Buffer, cur *Cursor, open, close rune) (int, int, int, int, bool) {
	sr, sc, er, ec, ok := innerPairRange(buf, cur, open, close)
	if !ok {
		return 0, 0, 0, 0, false
	}
	return sr, sc - 1, er, ec + 1, true
}
