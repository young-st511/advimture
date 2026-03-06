package editor

import "testing"

func feedKeys(p *Parser, keys ...string) *ParseResult {
	var result *ParseResult
	for _, k := range keys {
		result = p.Feed(k)
		if result != nil {
			return result
		}
	}
	return result
}

func TestParser_SimpleMotion(t *testing.T) {
	tests := []struct {
		name   string
		keys   []string
		motion MotionType
		count  int
	}{
		{"h", []string{"h"}, MotionLeft, 1},
		{"j", []string{"j"}, MotionDown, 1},
		{"k", []string{"k"}, MotionUp, 1},
		{"l", []string{"l"}, MotionRight, 1},
		{"w", []string{"w"}, MotionWordForward, 1},
		{"b", []string{"b"}, MotionWordBackward, 1},
		{"e", []string{"e"}, MotionWordEnd, 1},
		{"0", []string{"0"}, MotionLineStart, 1},
		{"$", []string{"$"}, MotionLineEnd, 1},
		{"G", []string{"G"}, MotionFileBottom, 1},
		{"gg", []string{"g", "g"}, MotionFileTop, 1},
		{"5j", []string{"5", "j"}, MotionDown, 5},
		{"3w", []string{"3", "w"}, MotionWordForward, 3},
		{"12l", []string{"1", "2", "l"}, MotionRight, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result := feedKeys(p, tt.keys...)
			if result == nil {
				t.Fatal("expected result, got nil")
			}
			if result.Motion != tt.motion {
				t.Errorf("expected motion %d, got %d", tt.motion, result.Motion)
			}
			if result.Count != tt.count {
				t.Errorf("expected count %d, got %d", tt.count, result.Count)
			}
			if result.Operator != OperatorNone {
				t.Errorf("expected no operator, got %d", result.Operator)
			}
		})
	}
}

func TestParser_OperatorMotion(t *testing.T) {
	tests := []struct {
		name     string
		keys     []string
		operator OperatorType
		motion   MotionType
		count    int
	}{
		{"dw", []string{"d", "w"}, OperatorDelete, MotionWordForward, 1},
		{"d$", []string{"d", "$"}, OperatorDelete, MotionLineEnd, 1},
		{"yw", []string{"y", "w"}, OperatorYank, MotionWordForward, 1},
		{"cw", []string{"c", "w"}, OperatorChange, MotionWordForward, 1},
		{"3dw", []string{"3", "d", "w"}, OperatorDelete, MotionWordForward, 3},
		{"d3w", []string{"d", "3", "w"}, OperatorDelete, MotionWordForward, 3},
		{"d32w", []string{"d", "3", "2", "w"}, OperatorDelete, MotionWordForward, 32},
		{"2d3w", []string{"2", "d", "3", "w"}, OperatorDelete, MotionWordForward, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result := feedKeys(p, tt.keys...)
			if result == nil {
				t.Fatal("expected result, got nil")
			}
			if result.Operator != tt.operator {
				t.Errorf("expected operator %d, got %d", tt.operator, result.Operator)
			}
			if result.Motion != tt.motion {
				t.Errorf("expected motion %d, got %d", tt.motion, result.Motion)
			}
			if result.Count != tt.count {
				t.Errorf("expected count %d, got %d", tt.count, result.Count)
			}
		})
	}
}

func TestParser_LinewiseOperator(t *testing.T) {
	tests := []struct {
		name     string
		keys     []string
		operator OperatorType
		count    int
	}{
		{"dd", []string{"d", "d"}, OperatorDelete, 1},
		{"yy", []string{"y", "y"}, OperatorYank, 1},
		{"cc", []string{"c", "c"}, OperatorChange, 1},
		{"3dd", []string{"3", "d", "d"}, OperatorDelete, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result := feedKeys(p, tt.keys...)
			if result == nil {
				t.Fatal("expected result, got nil")
			}
			if result.Operator != tt.operator {
				t.Errorf("expected operator %d, got %d", tt.operator, result.Operator)
			}
			if !result.IsLinewise {
				t.Error("expected linewise")
			}
			if result.Count != tt.count {
				t.Errorf("expected count %d, got %d", tt.count, result.Count)
			}
		})
	}
}

func TestParser_TextObject(t *testing.T) {
	tests := []struct {
		name       string
		keys       []string
		operator   OperatorType
		textObject TextObjectType
	}{
		{"ciw", []string{"c", "i", "w"}, OperatorChange, TextObjectInnerWord},
		{"diw", []string{"d", "i", "w"}, OperatorDelete, TextObjectInnerWord},
		{"yi\"", []string{"y", "i", "\""}, OperatorYank, TextObjectInnerDoubleQuote},
		{"ci(", []string{"c", "i", "("}, OperatorChange, TextObjectInnerParen},
		{"da{", []string{"d", "a", "{"}, OperatorDelete, TextObjectABrace},
		{"ci'", []string{"c", "i", "'"}, OperatorChange, TextObjectInnerSingleQuote},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result := feedKeys(p, tt.keys...)
			if result == nil {
				t.Fatal("expected result, got nil")
			}
			if result.Operator != tt.operator {
				t.Errorf("expected operator %d, got %d", tt.operator, result.Operator)
			}
			if result.TextObject != tt.textObject {
				t.Errorf("expected text object %d, got %d", tt.textObject, result.TextObject)
			}
		})
	}
}

func TestParser_SimpleCommands(t *testing.T) {
	tests := []struct {
		name string
		key  string
		cmd  SimpleCommand
	}{
		{"i", "i", SimpleCmdInsertBefore},
		{"a", "a", SimpleCmdInsertAfter},
		{"o", "o", SimpleCmdOpenBelow},
		{"I", "I", SimpleCmdInsertLineStart},
		{"A", "A", SimpleCmdInsertLineEnd},
		{"O", "O", SimpleCmdOpenAbove},
		{"x", "x", SimpleCmdDeleteChar},
		{"J", "J", SimpleCmdJoinLine},
		{".", ".", SimpleCmdDotRepeat},
		{"u", "u", SimpleCmdUndo},
		{"p", "p", SimpleCmdPaste},
		{"P", "P", SimpleCmdPasteBefore},
		{":", ":", SimpleCmdEnterCommand},
		{"/", "/", SimpleCmdSearchForward},
		{"n", "n", SimpleCmdSearchNext},
		{"N", "N", SimpleCmdSearchPrev},
		{"v", "v", SimpleCmdVisual},
		{"V", "V", SimpleCmdVisualLine},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			result := p.Feed(tt.key)
			if result == nil {
				t.Fatal("expected result, got nil")
			}
			if result.SimpleCmd != tt.cmd {
				t.Errorf("expected cmd %d, got %d", tt.cmd, result.SimpleCmd)
			}
		})
	}
}

func TestParser_FindChar(t *testing.T) {
	p := NewParser()
	r := p.Feed("f")
	if r != nil {
		t.Fatal("f should not complete immediately")
	}
	r = p.Feed("x")
	if r == nil {
		t.Fatal("fx should complete")
	}
	if r.Motion != MotionFindChar || r.Char != 'x' {
		t.Errorf("expected FindChar 'x', got motion=%d char='%c'", r.Motion, r.Char)
	}
}

func TestParser_TillChar(t *testing.T) {
	p := NewParser()
	feedKeys(p, "t")
	r := p.Feed(".")
	if r == nil {
		t.Fatal("t. should complete")
	}
	if r.Motion != MotionTillChar || r.Char != '.' {
		t.Errorf("expected TillChar '.', got motion=%d char='%c'", r.Motion, r.Char)
	}
}

func TestParser_ReplaceChar(t *testing.T) {
	p := NewParser()
	feedKeys(p, "r")
	r := p.Feed("z")
	if r == nil {
		t.Fatal("rz should complete")
	}
	if r.SimpleCmd != SimpleCmdReplaceChar || r.Char != 'z' {
		t.Errorf("expected ReplaceChar 'z', got cmd=%d char='%c'", r.SimpleCmd, r.Char)
	}
}

func TestParser_OperatorFindChar(t *testing.T) {
	p := NewParser()
	r := feedKeys(p, "d", "f", "x")
	if r == nil {
		t.Fatal("dfx should complete")
	}
	if r.Operator != OperatorDelete || r.Motion != MotionFindChar || r.Char != 'x' {
		t.Errorf("expected delete+FindChar 'x', got op=%d motion=%d char='%c'",
			r.Operator, r.Motion, r.Char)
	}
}

func TestParser_OperatorRejectReplace(t *testing.T) {
	p := NewParser()
	feedKeys(p, "d", "r")
	// 'r' after operator should not enter waitChar; parser resets
	if p.IsOperatorPending() {
		t.Error("should not be pending after invalid 'dr'")
	}
	// Parser should work normally after reset
	r := p.Feed("j")
	if r == nil || r.Motion != MotionDown {
		t.Error("expected 'j' to work after 'dr' reset")
	}
}

func TestParser_CtrlR(t *testing.T) {
	p := NewParser()
	r := p.Feed("ctrl+r")
	if r == nil {
		t.Fatal("ctrl+r should complete")
	}
	if r.SimpleCmd != SimpleCmdRedo {
		t.Errorf("expected Redo, got %d", r.SimpleCmd)
	}
}

func TestParser_IsOperatorPending(t *testing.T) {
	p := NewParser()
	if p.IsOperatorPending() {
		t.Error("should not be pending initially")
	}
	p.Feed("d")
	if !p.IsOperatorPending() {
		t.Error("should be pending after 'd'")
	}
	p.Feed("w")
	if p.IsOperatorPending() {
		t.Error("should not be pending after 'dw' completes")
	}
}

func TestParser_Reset(t *testing.T) {
	p := NewParser()
	p.Feed("3")
	p.Feed("d")
	p.Reset()
	if p.IsOperatorPending() {
		t.Error("should not be pending after reset")
	}
	// Should work normally after reset
	r := p.Feed("j")
	if r == nil || r.Motion != MotionDown {
		t.Error("expected 'j' to work after reset")
	}
}
