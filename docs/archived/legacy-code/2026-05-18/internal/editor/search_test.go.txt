package editor

import "testing"

func TestFindMatches_Basic(t *testing.T) {
	lines := []string{"hello world", "world hello", "foo bar"}
	matches := FindMatches(lines, "world")
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
	if matches[0].Row != 0 || matches[0].Col != 6 {
		t.Errorf("first match: expected (0,6), got (%d,%d)", matches[0].Row, matches[0].Col)
	}
	if matches[1].Row != 1 || matches[1].Col != 0 {
		t.Errorf("second match: expected (1,0), got (%d,%d)", matches[1].Row, matches[1].Col)
	}
}

func TestFindMatches_EmptyPattern(t *testing.T) {
	lines := []string{"hello"}
	matches := FindMatches(lines, "")
	if len(matches) != 0 {
		t.Errorf("empty pattern should return no matches, got %d", len(matches))
	}
}

func TestFindMatches_NoMatch(t *testing.T) {
	lines := []string{"hello", "world"}
	matches := FindMatches(lines, "xyz")
	if len(matches) != 0 {
		t.Errorf("expected 0 matches, got %d", len(matches))
	}
}

func TestFindMatches_LineBoundary(t *testing.T) {
	lines := []string{"abc", "def"}
	// "cd" spans across lines — should NOT match
	matches := FindMatches(lines, "cd")
	if len(matches) != 0 {
		t.Errorf("cross-line pattern should not match, got %d", len(matches))
	}
}

func TestFindMatches_Unicode(t *testing.T) {
	lines := []string{"안녕하세요", "hello 안녕"}
	matches := FindMatches(lines, "안녕")
	if len(matches) != 2 {
		t.Fatalf("expected 2 unicode matches, got %d", len(matches))
	}
	if matches[0].Row != 0 || matches[0].Col != 0 {
		t.Errorf("expected (0,0), got (%d,%d)", matches[0].Row, matches[0].Col)
	}
	if matches[1].Row != 1 || matches[1].Col != 6 {
		t.Errorf("expected (1,6), got (%d,%d)", matches[1].Row, matches[1].Col)
	}
}

func TestFindMatches_MatchLength(t *testing.T) {
	lines := []string{"hello"}
	matches := FindMatches(lines, "ell")
	if len(matches) != 1 {
		t.Fatalf("expected 1 match, got %d", len(matches))
	}
	if matches[0].Len != 3 {
		t.Errorf("expected len 3, got %d", matches[0].Len)
	}
}

func TestFindMatches_MultipleOnOneLine(t *testing.T) {
	lines := []string{"aaa", "aa"}
	matches := FindMatches(lines, "aa")
	// "aaa" has "aa" starting at col 0 and col 1
	if len(matches) < 2 {
		t.Fatalf("expected at least 2 matches, got %d", len(matches))
	}
	if matches[0].Row != 0 || matches[0].Col != 0 {
		t.Errorf("first match: expected (0,0), got (%d,%d)", matches[0].Row, matches[0].Col)
	}
	if matches[1].Row != 0 || matches[1].Col != 1 {
		t.Errorf("second match: expected (0,1), got (%d,%d)", matches[1].Row, matches[1].Col)
	}
}

func TestFindMatches_SingleChar(t *testing.T) {
	lines := []string{"ERROR: something", "info: ok", "ERROR: another"}
	matches := FindMatches(lines, "ERROR")
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
	if matches[0].Row != 0 {
		t.Errorf("expected row 0, got %d", matches[0].Row)
	}
	if matches[1].Row != 2 {
		t.Errorf("expected row 2, got %d", matches[1].Row)
	}
}
