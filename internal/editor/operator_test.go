package editor

import "testing"

func setup(lines []string, row, col int) (*Buffer, *Cursor, *Register, *UndoManager) {
	buf := NewBuffer(lines)
	cur := NewCursor()
	cur.MoveTo(row, col)
	reg := NewRegister()
	undo := NewUndoManager()
	return buf, cur, reg, undo
}

func TestExecute_PureMotion(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		Motion: MotionWordForward, Count: 1,
	})
	if cur.Col != 6 {
		t.Errorf("w: expected col 6, got %d", cur.Col)
	}
}

func TestExecute_DeleteWord(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, Motion: MotionWordForward, Count: 1,
	})
	if buf.GetLine(0) != "world" {
		t.Errorf("dw: expected 'world', got '%s'", buf.GetLine(0))
	}
	if reg.Content != "hello " {
		t.Errorf("dw register: expected 'hello ', got '%s'", reg.Content)
	}
	if cur.Col != 0 {
		t.Errorf("dw cursor: expected col 0, got %d", cur.Col)
	}
}

func TestExecute_DeleteToEnd(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 5)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, Motion: MotionLineEnd, Count: 1,
	})
	if buf.GetLine(0) != "hello" {
		t.Errorf("d$: expected 'hello', got '%s'", buf.GetLine(0))
	}
	if reg.Content != " world" {
		t.Errorf("d$ register: expected ' world', got '%s'", reg.Content)
	}
}

func TestExecute_YankWord(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorYank, Motion: MotionWordForward, Count: 1,
	})
	// Yank should not modify buffer
	if buf.GetLine(0) != "hello world" {
		t.Errorf("yw: buffer should not change, got '%s'", buf.GetLine(0))
	}
	if reg.Content != "hello " {
		t.Errorf("yw register: expected 'hello ', got '%s'", reg.Content)
	}
}

func TestExecute_ChangeWord(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 0)
	result := Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorChange, Motion: MotionWordForward, Count: 1,
	})
	if !result.SwitchToInsert {
		t.Error("cw should switch to insert mode")
	}
	if buf.GetLine(0) != "world" {
		t.Errorf("cw: expected 'world', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_DeleteLine(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "bbb", "ccc"}, 1, 1)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, IsLinewise: true, Count: 1,
	})
	if buf.LineCount() != 2 {
		t.Errorf("dd: expected 2 lines, got %d", buf.LineCount())
	}
	if reg.Content != "bbb" || !reg.Linewise {
		t.Errorf("dd register: expected 'bbb' linewise, got '%s' linewise=%v", reg.Content, reg.Linewise)
	}
}

func TestExecute_DeleteMultipleLines(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "bbb", "ccc", "ddd"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, IsLinewise: true, Count: 2,
	})
	if buf.LineCount() != 2 {
		t.Errorf("2dd: expected 2 lines, got %d", buf.LineCount())
	}
	if buf.GetLine(0) != "ccc" {
		t.Errorf("2dd: expected first line 'ccc', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_YankLine(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "bbb"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorYank, IsLinewise: true, Count: 1,
	})
	// Buffer unchanged
	if buf.LineCount() != 2 {
		t.Errorf("yy: should not change buffer")
	}
	if reg.Content != "aaa" || !reg.Linewise {
		t.Errorf("yy register: expected 'aaa' linewise, got '%s'", reg.Content)
	}
}

func TestExecute_ChangeLine(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "bbb"}, 0, 1)
	result := Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorChange, IsLinewise: true, Count: 1,
	})
	if !result.SwitchToInsert {
		t.Error("cc should switch to insert")
	}
}

func TestExecute_InsertBefore(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 2)
	result := Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdInsertBefore,
	})
	if !result.SwitchToInsert {
		t.Error("i should switch to insert")
	}
	if cur.Col != 2 {
		t.Errorf("i: cursor should stay at 2, got %d", cur.Col)
	}
}

func TestExecute_InsertAfter(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 2)
	result := Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdInsertAfter,
	})
	if !result.SwitchToInsert {
		t.Error("a should switch to insert")
	}
	if cur.Col != 3 {
		t.Errorf("a: cursor should be at 3, got %d", cur.Col)
	}
}

func TestExecute_InsertLineStartEnd(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 2)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdInsertLineStart,
	})
	if cur.Col != 0 {
		t.Errorf("I: expected col 0, got %d", cur.Col)
	}

	cur.MoveTo(0, 2)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdInsertLineEnd,
	})
	if cur.Col != 5 {
		t.Errorf("A: expected col 5, got %d", cur.Col)
	}
}

func TestExecute_OpenBelow(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 2)
	result := Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdOpenBelow,
	})
	if !result.SwitchToInsert {
		t.Error("o should switch to insert")
	}
	if buf.LineCount() != 2 {
		t.Errorf("o: expected 2 lines, got %d", buf.LineCount())
	}
	if cur.Row != 1 || cur.Col != 0 {
		t.Errorf("o: expected cursor at (1,0), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestExecute_OpenAbove(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 2)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdOpenAbove,
	})
	if buf.LineCount() != 2 {
		t.Errorf("O: expected 2 lines, got %d", buf.LineCount())
	}
	if cur.Row != 0 || cur.Col != 0 {
		t.Errorf("O: expected cursor at (0,0), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestExecute_DeleteChar(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 1)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdDeleteChar, Count: 1,
	})
	if buf.GetLine(0) != "hllo" {
		t.Errorf("x: expected 'hllo', got '%s'", buf.GetLine(0))
	}
	if reg.Content != "e" {
		t.Errorf("x register: expected 'e', got '%s'", reg.Content)
	}
}

func TestExecute_ReplaceChar(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdReplaceChar, Char: 'H',
	})
	if buf.GetLine(0) != "Hello" {
		t.Errorf("rH: expected 'Hello', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_JoinLine(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello", "world"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdJoinLine,
	})
	if buf.LineCount() != 1 || buf.GetLine(0) != "helloworld" {
		t.Errorf("J: expected 'helloworld', got '%s'", buf.GetLine(0))
	}
	if cur.Col != 5 {
		t.Errorf("J: cursor should be at join point 5, got %d", cur.Col)
	}
}

func TestExecute_PasteLinewise(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "ccc"}, 0, 0)
	reg.Set("bbb", true)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdPaste,
	})
	if buf.LineCount() != 3 || buf.GetLine(1) != "bbb" {
		t.Errorf("p linewise: expected 'bbb' at line 1, got '%s'", buf.GetLine(1))
	}
	if cur.Row != 1 {
		t.Errorf("p linewise: cursor should be at row 1, got %d", cur.Row)
	}
}

func TestExecute_PasteCharwise(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hllo"}, 0, 0)
	reg.Set("e", false)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdPaste,
	})
	if buf.GetLine(0) != "hello" {
		t.Errorf("p charwise: expected 'hello', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_PasteBeforeLinewise(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"bbb"}, 0, 0)
	reg.Set("aaa", true)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdPasteBefore,
	})
	if buf.LineCount() != 2 || buf.GetLine(0) != "aaa" {
		t.Errorf("P linewise: expected 'aaa' at line 0, got '%s'", buf.GetLine(0))
	}
}

func TestExecute_UndoRedo(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello"}, 0, 0)

	// Delete 'h'
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdDeleteChar, Count: 1,
	})
	if buf.GetLine(0) != "ello" {
		t.Fatalf("after x: expected 'ello', got '%s'", buf.GetLine(0))
	}

	// Undo
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdUndo,
	})
	if buf.GetLine(0) != "hello" {
		t.Errorf("after undo: expected 'hello', got '%s'", buf.GetLine(0))
	}

	// Redo
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdRedo,
	})
	if buf.GetLine(0) != "ello" {
		t.Errorf("after redo: expected 'ello', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_DeleteWordUndo(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 0)

	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, Motion: MotionWordForward, Count: 1,
	})
	if buf.GetLine(0) != "world" {
		t.Fatalf("dw: expected 'world', got '%s'", buf.GetLine(0))
	}

	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdUndo,
	})
	if buf.GetLine(0) != "hello world" {
		t.Errorf("undo dw: expected 'hello world', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_DeleteLineUndo(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"aaa", "bbb", "ccc"}, 1, 0)

	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, IsLinewise: true, Count: 1,
	})
	if buf.LineCount() != 2 {
		t.Fatalf("dd: expected 2 lines, got %d", buf.LineCount())
	}

	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdUndo,
	})
	if buf.LineCount() != 3 || buf.GetLine(1) != "bbb" {
		t.Errorf("undo dd: expected 3 lines with 'bbb', got %d lines", buf.LineCount())
	}
}

func TestExecute_TextObjectInnerWord(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 7)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, TextObject: TextObjectInnerWord, Count: 1,
	})
	if buf.GetLine(0) != "hello " {
		t.Errorf("diw: expected 'hello ', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_TextObjectInnerQuote(t *testing.T) {
	buf, cur, reg, undo := setup([]string{`say "hello" end`}, 0, 6)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorChange, TextObject: TextObjectInnerDoubleQuote, Count: 1,
	})
	if buf.GetLine(0) != `say "" end` {
		t.Errorf(`ci": expected 'say "" end', got '%s'`, buf.GetLine(0))
	}
}

func TestExecute_YankTextObject(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"hello world"}, 0, 7)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorYank, TextObject: TextObjectInnerWord, Count: 1,
	})
	if buf.GetLine(0) != "hello world" {
		t.Errorf("yiw: buffer should not change, got '%s'", buf.GetLine(0))
	}
	if reg.Content != "world" {
		t.Errorf("yiw register: expected 'world', got '%s'", reg.Content)
	}
}

func TestExecute_DeleteMultiChar(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"abcde"}, 0, 0)
	Execute(buf, cur, reg, undo, &ParseResult{
		SimpleCmd: SimpleCmdDeleteChar, Count: 3,
	})
	if buf.GetLine(0) != "de" {
		t.Errorf("3x: expected 'de', got '%s'", buf.GetLine(0))
	}
	if reg.Content != "abc" {
		t.Errorf("3x register: expected 'abc', got '%s'", reg.Content)
	}
	// Single undo should restore all 3
	Execute(buf, cur, reg, undo, &ParseResult{SimpleCmd: SimpleCmdUndo})
	if buf.GetLine(0) != "abcde" {
		t.Errorf("undo 3x: expected 'abcde', got '%s'", buf.GetLine(0))
	}
}

func TestExecute_TextObjectInnerParen(t *testing.T) {
	buf, cur, reg, undo := setup([]string{"fn(arg1, arg2)"}, 0, 5)
	Execute(buf, cur, reg, undo, &ParseResult{
		Operator: OperatorDelete, TextObject: TextObjectInnerParen, Count: 1,
	})
	if buf.GetLine(0) != "fn()" {
		t.Errorf("di(: expected 'fn()', got '%s'", buf.GetLine(0))
	}
}
