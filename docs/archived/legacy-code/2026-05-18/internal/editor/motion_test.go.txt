package editor

import "testing"

func makeBufCur(lines []string, row, col int) (*Buffer, *Cursor) {
	buf := NewBuffer(lines)
	cur := NewCursor()
	cur.MoveTo(row, col)
	return buf, cur
}

func TestMotion_LeftRight(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello"}, 0, 2)

	ExecuteMotion(buf, cur, MotionLeft, 1, 0)
	if cur.Col != 1 {
		t.Errorf("h: expected col 1, got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionLeft, 1, 0)
	ExecuteMotion(buf, cur, MotionLeft, 1, 0) // at 0, should stay
	if cur.Col != 0 {
		t.Errorf("h at 0: expected col 0, got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionRight, 1, 0)
	if cur.Col != 1 {
		t.Errorf("l: expected col 1, got %d", cur.Col)
	}

	// Right at end of line
	cur.MoveTo(0, 4)
	ExecuteMotion(buf, cur, MotionRight, 1, 0)
	if cur.Col != 4 {
		t.Errorf("l at end: expected col 4, got %d", cur.Col)
	}
}

func TestMotion_LeftRight_EmptyLine(t *testing.T) {
	buf, cur := makeBufCur([]string{""}, 0, 0)
	ExecuteMotion(buf, cur, MotionLeft, 1, 0)
	if cur.Col != 0 {
		t.Errorf("h on empty: expected 0, got %d", cur.Col)
	}
	ExecuteMotion(buf, cur, MotionRight, 1, 0)
	if cur.Col != 0 {
		t.Errorf("l on empty: expected 0, got %d", cur.Col)
	}
}

func TestMotion_DownUp(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello", "hi", "world"}, 0, 3)

	ExecuteMotion(buf, cur, MotionDown, 1, 0)
	if cur.Row != 1 || cur.Col != 1 {
		t.Errorf("j: expected (1,1), got (%d,%d)", cur.Row, cur.Col)
	}
	if cur.DesiredCol != 3 {
		t.Errorf("j: expected DesiredCol 3, got %d", cur.DesiredCol)
	}

	ExecuteMotion(buf, cur, MotionDown, 1, 0)
	if cur.Row != 2 || cur.Col != 3 {
		t.Errorf("j again: expected (2,3), got (%d,%d)", cur.Row, cur.Col)
	}

	// Up
	ExecuteMotion(buf, cur, MotionUp, 1, 0)
	if cur.Row != 1 || cur.Col != 1 {
		t.Errorf("k: expected (1,1), got (%d,%d)", cur.Row, cur.Col)
	}

	// Up at top
	ExecuteMotion(buf, cur, MotionUp, 1, 0)
	ExecuteMotion(buf, cur, MotionUp, 1, 0) // stays at 0
	if cur.Row != 0 {
		t.Errorf("k at top: expected row 0, got %d", cur.Row)
	}

	// Down at bottom
	cur.MoveTo(2, 0)
	ExecuteMotion(buf, cur, MotionDown, 1, 0)
	if cur.Row != 2 {
		t.Errorf("j at bottom: expected row 2, got %d", cur.Row)
	}
}

func TestMotion_DownUp_DesiredCol(t *testing.T) {
	buf, cur := makeBufCur([]string{
		"long line here",
		"hi",
		"another long one",
	}, 0, 10)

	ExecuteMotion(buf, cur, MotionDown, 1, 0)
	if cur.Col != 1 || cur.DesiredCol != 10 {
		t.Errorf("expected col=1 desired=10, got col=%d desired=%d", cur.Col, cur.DesiredCol)
	}

	ExecuteMotion(buf, cur, MotionDown, 1, 0)
	if cur.Col != 10 {
		t.Errorf("expected col restored to 10, got %d", cur.Col)
	}
}

func TestMotion_DownUp_Count(t *testing.T) {
	buf, cur := makeBufCur([]string{"a", "b", "c", "d", "e"}, 0, 0)

	ExecuteMotion(buf, cur, MotionDown, 3, 0)
	if cur.Row != 3 {
		t.Errorf("3j: expected row 3, got %d", cur.Row)
	}

	ExecuteMotion(buf, cur, MotionUp, 2, 0)
	if cur.Row != 1 {
		t.Errorf("2k: expected row 1, got %d", cur.Row)
	}
}

func TestMotion_LineStartEnd(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 5)

	ExecuteMotion(buf, cur, MotionLineStart, 1, 0)
	if cur.Col != 0 {
		t.Errorf("0: expected col 0, got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionLineEnd, 1, 0)
	if cur.Col != 10 { // "hello world" has 11 runes, last index = 10
		t.Errorf("$: expected col 10, got %d", cur.Col)
	}
}

func TestMotion_LineEnd_EmptyLine(t *testing.T) {
	buf, cur := makeBufCur([]string{""}, 0, 0)
	ExecuteMotion(buf, cur, MotionLineEnd, 1, 0)
	if cur.Col != 0 {
		t.Errorf("$ on empty: expected 0, got %d", cur.Col)
	}
}

func TestMotion_FileTopBottom(t *testing.T) {
	buf, cur := makeBufCur([]string{"aaa", "bbb", "ccc"}, 1, 2)

	ExecuteMotion(buf, cur, MotionFileTop, 1, 0)
	if cur.Row != 0 || cur.Col != 0 {
		t.Errorf("gg: expected (0,0), got (%d,%d)", cur.Row, cur.Col)
	}

	ExecuteMotion(buf, cur, MotionFileBottom, 1, 0)
	if cur.Row != 2 {
		t.Errorf("G: expected row 2, got %d", cur.Row)
	}
}

func TestMotion_WordForward(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world foo"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Col != 6 {
		t.Errorf("w from 'hello': expected col 6 ('w'), got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Col != 12 {
		t.Errorf("w from 'world': expected col 12 ('f'), got %d", cur.Col)
	}
}

func TestMotion_WordForward_MixedClasses(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello.world"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Col != 5 {
		t.Errorf("w: expected col 5 ('.'), got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Col != 6 {
		t.Errorf("w: expected col 6 ('w'), got %d", cur.Col)
	}
}

func TestMotion_WordForward_CrossLine(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello", "world"}, 0, 3)

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Row != 1 || cur.Col != 0 {
		t.Errorf("w cross line: expected (1,0), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestMotion_WordForward_EmptyLine(t *testing.T) {
	buf, cur := makeBufCur([]string{"", "hello"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Row != 1 || cur.Col != 0 {
		t.Errorf("w on empty: expected (1,0), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestMotion_WordForward_Count(t *testing.T) {
	buf, cur := makeBufCur([]string{"one two three four"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordForward, 3, 0)
	if cur.Col != 14 {
		t.Errorf("3w: expected col 14 ('four'), got %d", cur.Col)
	}
}

func TestMotion_WordBackward(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world foo"}, 0, 12)

	ExecuteMotion(buf, cur, MotionWordBackward, 1, 0)
	if cur.Col != 6 {
		t.Errorf("b from 'foo': expected col 6 ('w'), got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionWordBackward, 1, 0)
	if cur.Col != 0 {
		t.Errorf("b from 'world': expected col 0 ('h'), got %d", cur.Col)
	}
}

func TestMotion_WordBackward_CrossLine(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello", "world"}, 1, 0)

	ExecuteMotion(buf, cur, MotionWordBackward, 1, 0)
	if cur.Row != 0 || cur.Col != 0 {
		t.Errorf("b cross line: expected (0,0), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestMotion_WordBackward_AtStart(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordBackward, 1, 0)
	if cur.Col != 0 {
		t.Errorf("b at start: expected 0, got %d", cur.Col)
	}
}

func TestMotion_WordEnd(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordEnd, 1, 0)
	if cur.Col != 4 {
		t.Errorf("e: expected col 4 ('o'), got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionWordEnd, 1, 0)
	if cur.Col != 10 {
		t.Errorf("e: expected col 10 ('d'), got %d", cur.Col)
	}
}

func TestMotion_WordEnd_CrossLine(t *testing.T) {
	buf, cur := makeBufCur([]string{"hi", "world"}, 0, 1)

	ExecuteMotion(buf, cur, MotionWordEnd, 1, 0)
	if cur.Row != 1 || cur.Col != 4 {
		t.Errorf("e cross line: expected (1,4), got (%d,%d)", cur.Row, cur.Col)
	}
}

func TestMotion_FindChar(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	ExecuteMotion(buf, cur, MotionFindChar, 1, 'o')
	if cur.Col != 4 {
		t.Errorf("fo: expected col 4, got %d", cur.Col)
	}

	// Not found: stays
	ExecuteMotion(buf, cur, MotionFindChar, 1, 'z')
	if cur.Col != 4 {
		t.Errorf("fz: expected col 4 (unchanged), got %d", cur.Col)
	}
}

func TestMotion_TillChar(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	ExecuteMotion(buf, cur, MotionTillChar, 1, 'o')
	if cur.Col != 3 {
		t.Errorf("to: expected col 3, got %d", cur.Col)
	}
}

func TestMotion_FindChar_Count(t *testing.T) {
	buf, cur := makeBufCur([]string{"aXbXcXd"}, 0, 0)

	ExecuteMotion(buf, cur, MotionFindChar, 3, 'X')
	if cur.Col != 5 {
		t.Errorf("3fX: expected col 5 (third X), got %d", cur.Col)
	}
}

func TestMotion_TillChar_Count(t *testing.T) {
	buf, cur := makeBufCur([]string{"aXbXcXd"}, 0, 0)

	ExecuteMotion(buf, cur, MotionTillChar, 2, 'X')
	if cur.Col != 2 {
		t.Errorf("2tX: expected col 2 (before second X), got %d", cur.Col)
	}
}

func TestMotion_FindChar_Unicode(t *testing.T) {
	buf, cur := makeBufCur([]string{"안녕하세요"}, 0, 0)

	ExecuteMotion(buf, cur, MotionFindChar, 1, '세')
	if cur.Col != 3 {
		t.Errorf("f세: expected col 3, got %d", cur.Col)
	}
}

func TestMotionRange_ForwardExclusive(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	// w is exclusive: "hello " (cols 0-5), endCol=6 exclusive
	sr, sc, er, ec := MotionRange(buf, cur, MotionWordForward, 1, 0)
	if sr != 0 || sc != 0 || er != 0 || ec != 6 {
		t.Errorf("MotionRange(w): expected (0,0,0,6), got (%d,%d,%d,%d)", sr, sc, er, ec)
	}
}

func TestMotionRange_ForwardInclusive(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	// e is inclusive: "hello" (cols 0-4), endCol=5 exclusive
	sr, sc, er, ec := MotionRange(buf, cur, MotionWordEnd, 1, 0)
	if sr != 0 || sc != 0 || er != 0 || ec != 5 {
		t.Errorf("MotionRange(e): expected (0,0,0,5), got (%d,%d,%d,%d)", sr, sc, er, ec)
	}
}

func TestMotionRange_Backward(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 6)

	// b is exclusive backward: from col 6 ('w'), b goes to col 0 ('h')
	// Range should be [0, 6) — does NOT include original cursor char
	sr, sc, er, ec := MotionRange(buf, cur, MotionWordBackward, 1, 0)
	if sr != 0 || sc != 0 || er != 0 || ec != 6 {
		t.Errorf("MotionRange(b): expected (0,0,0,6), got (%d,%d,%d,%d)", sr, sc, er, ec)
	}
}

func TestMotionRange_FindChar(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 0)

	// f is inclusive: fo finds 'o' at col 4, range [0, 5)
	sr, sc, er, ec := MotionRange(buf, cur, MotionFindChar, 1, 'o')
	if sr != 0 || sc != 0 || er != 0 || ec != 5 {
		t.Errorf("MotionRange(fo): expected (0,0,0,5), got (%d,%d,%d,%d)", sr, sc, er, ec)
	}
}

func TestMotion_LeftRight_Count(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello world"}, 0, 5)

	ExecuteMotion(buf, cur, MotionLeft, 3, 0)
	if cur.Col != 2 {
		t.Errorf("3h from 5: expected col 2, got %d", cur.Col)
	}

	ExecuteMotion(buf, cur, MotionRight, 5, 0)
	if cur.Col != 7 {
		t.Errorf("5l from 2: expected col 7, got %d", cur.Col)
	}
}

func TestMotion_WordForward_LastLine(t *testing.T) {
	buf, cur := makeBufCur([]string{"hello"}, 0, 3)

	ExecuteMotion(buf, cur, MotionWordForward, 1, 0)
	if cur.Col != 4 {
		t.Errorf("w on last line past word: expected col 4, got %d", cur.Col)
	}
}

func TestMotion_WordEnd_SingleChar(t *testing.T) {
	buf, cur := makeBufCur([]string{"a b c"}, 0, 0)

	ExecuteMotion(buf, cur, MotionWordEnd, 1, 0)
	if cur.Col != 2 {
		t.Errorf("e from 'a': expected col 2 ('b'), got %d", cur.Col)
	}
}
