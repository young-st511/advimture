package editor

// MotionType represents a cursor motion command.
type MotionType int

const (
	MotionNone MotionType = iota
	MotionLeft            // h
	MotionDown            // j
	MotionUp              // k
	MotionRight           // l
	MotionWordForward     // w
	MotionWordBackward    // b
	MotionWordEnd         // e
	MotionLineStart       // 0
	MotionLineEnd         // $
	MotionFileTop         // gg
	MotionFileBottom      // G
	MotionFindChar        // f{char}
	MotionTillChar        // t{char}
)

// TextObjectType represents a text object.
type TextObjectType int

const (
	TextObjectNone TextObjectType = iota
	TextObjectInnerWord           // iw
	TextObjectAWord               // aw
	TextObjectInnerDoubleQuote    // i"
	TextObjectADoubleQuote        // a"
	TextObjectInnerSingleQuote    // i'
	TextObjectASingleQuote        // a'
	TextObjectInnerParen          // i(
	TextObjectAParen              // a(
	TextObjectInnerBrace          // i{
	TextObjectABrace              // a{
)

// OperatorType represents an operator command (d, y, c).
type OperatorType int

const (
	OperatorNone   OperatorType = iota
	OperatorDelete              // d
	OperatorYank                // y
	OperatorChange              // c
)

// ParseResult is the parsed result of a Normal mode key sequence.
type ParseResult struct {
	Count      int
	Operator   OperatorType
	Motion     MotionType
	TextObject TextObjectType
	IsLinewise bool
	Char       rune // for f{char}, t{char}, r{char}

	// Simple commands (not operator+motion)
	SimpleCmd SimpleCommand
}

// SimpleCommand represents standalone commands that don't follow operator+motion pattern.
type SimpleCommand int

const (
	SimpleCmdNone SimpleCommand = iota
	SimpleCmdInsertBefore       // i
	SimpleCmdInsertAfter        // a
	SimpleCmdOpenBelow          // o
	SimpleCmdInsertLineStart    // I
	SimpleCmdInsertLineEnd      // A
	SimpleCmdOpenAbove          // O
	SimpleCmdDeleteChar         // x
	SimpleCmdReplaceChar        // r{char}
	SimpleCmdJoinLine           // J
	SimpleCmdDotRepeat          // .
	SimpleCmdUndo               // u
	SimpleCmdRedo               // Ctrl+r
	SimpleCmdPaste              // p
	SimpleCmdPasteBefore        // P
	SimpleCmdEnterCommand       // :
	SimpleCmdSearchForward      // /
	SimpleCmdSearchNext         // n
	SimpleCmdSearchPrev         // N
	SimpleCmdVisual             // v
	SimpleCmdVisualLine         // V
)

// Parser accumulates keystrokes and parses them into commands.
type Parser struct {
	count       int
	secondCount int // count after operator (e.g., d3w → secondCount=3)
	operator    OperatorType
	pending     []rune // accumulated key sequence
	waitChar    bool   // waiting for a character argument (f, t, r)
	waitCmd     rune   // the command waiting for char (f, t, r)
}

// NewParser creates a new parser.
func NewParser() *Parser {
	return &Parser{}
}

// Reset clears the parser state.
func (p *Parser) Reset() {
	p.count = 0
	p.secondCount = 0
	p.operator = OperatorNone
	p.pending = nil
	p.waitChar = false
	p.waitCmd = 0
}

// Feed processes a key and returns a ParseResult if a complete command is recognized.
// Returns nil if more input is needed.
func (p *Parser) Feed(key string) *ParseResult {
	if len(key) == 0 {
		return nil
	}

	// Waiting for character argument (f, t, r)
	if p.waitChar {
		r := []rune(key)
		if len(r) != 1 {
			p.Reset()
			return nil
		}
		result := p.buildCharResult(r[0])
		p.Reset()
		return result
	}

	r := []rune(key)
	if len(r) != 1 {
		// Handle multi-character keys
		return p.handleSpecialKey(key)
	}
	ch := r[0]

	// Accumulate count
	if ch >= '1' && ch <= '9' && p.operator == OperatorNone && p.count == 0 && len(p.pending) == 0 {
		p.count = int(ch - '0')
		return nil
	}
	if ch >= '0' && ch <= '9' && p.count > 0 && p.operator == OperatorNone {
		p.count = p.count*10 + int(ch-'0')
		return nil
	}
	// After operator, accumulate second count (e.g., d3w, d32w, 2d3w)
	if ch >= '1' && ch <= '9' && p.operator != OperatorNone && p.secondCount == 0 && len(p.pending) == 0 {
		p.secondCount = int(ch - '0')
		return nil
	}
	if ch >= '0' && ch <= '9' && p.secondCount > 0 && len(p.pending) == 0 {
		p.secondCount = p.secondCount*10 + int(ch-'0')
		return nil
	}

	// Operator
	if p.operator == OperatorNone {
		switch ch {
		case 'd':
			p.operator = OperatorDelete
			return nil
		case 'y':
			p.operator = OperatorYank
			return nil
		case 'c':
			p.operator = OperatorChange
			return nil
		}
	}

	// Same operator doubled (dd, yy, cc) → linewise
	if p.operator != OperatorNone {
		doubled := false
		switch {
		case p.operator == OperatorDelete && ch == 'd':
			doubled = true
		case p.operator == OperatorYank && ch == 'y':
			doubled = true
		case p.operator == OperatorChange && ch == 'c':
			doubled = true
		}
		if doubled {
			result := &ParseResult{
				Count:    p.effectiveCount(),
				Operator: p.operator,
				IsLinewise: true,
			}
			p.Reset()
			return result
		}
	}

	// Text objects (only valid after operator)
	if p.operator != OperatorNone && (ch == 'i' || ch == 'a') {
		p.pending = append(p.pending, ch)
		return nil
	}

	// Complete text object
	if p.operator != OperatorNone && len(p.pending) == 1 {
		prefix := p.pending[0]
		to := resolveTextObject(prefix, ch)
		if to != TextObjectNone {
			result := &ParseResult{
				Count:      p.effectiveCount(),
				Operator:   p.operator,
				TextObject: to,
			}
			p.Reset()
			return result
		}
		// Not a valid text object, reset
		p.Reset()
		return nil
	}

	// Motion
	motion := resolveMotion(ch)
	if motion != MotionNone {
		result := &ParseResult{
			Count:    p.effectiveCount(),
			Operator: p.operator,
			Motion:   motion,
		}
		p.Reset()
		return result
	}

	// gg (needs two g's)
	if ch == 'g' {
		p.pending = append(p.pending, 'g')
		if len(p.pending) == 2 {
			result := &ParseResult{
				Count:    p.effectiveCount(),
				Operator: p.operator,
				Motion:   MotionFileTop,
			}
			p.Reset()
			return result
		}
		return nil
	}

	// f, t — wait for next char (valid with or without operator)
	if ch == 'f' || ch == 't' {
		p.waitChar = true
		p.waitCmd = ch
		return nil
	}

	// r — wait for next char (only valid without operator)
	if ch == 'r' && p.operator == OperatorNone {
		p.waitChar = true
		p.waitCmd = ch
		return nil
	}

	// Simple commands (only when no operator pending)
	if p.operator == OperatorNone {
		cmd := resolveSimpleCmd(ch)
		if cmd != SimpleCmdNone {
			result := &ParseResult{
				Count:     p.effectiveCount(),
				SimpleCmd: cmd,
			}
			p.Reset()
			return result
		}
	}

	// Unrecognized: reset
	p.Reset()
	return nil
}

func (p *Parser) handleSpecialKey(key string) *ParseResult {
	switch key {
	case "ctrl+r":
		result := &ParseResult{
			Count:     p.effectiveCount(),
			SimpleCmd: SimpleCmdRedo,
		}
		p.Reset()
		return result
	case "ctrl+d", "ctrl+u":
		// Will be handled as scroll commands later
		p.Reset()
		return nil
	}
	p.Reset()
	return nil
}

func (p *Parser) buildCharResult(ch rune) *ParseResult {
	switch p.waitCmd {
	case 'f':
		return &ParseResult{
			Count:    p.effectiveCount(),
			Operator: p.operator,
			Motion:   MotionFindChar,
			Char:     ch,
		}
	case 't':
		return &ParseResult{
			Count:    p.effectiveCount(),
			Operator: p.operator,
			Motion:   MotionTillChar,
			Char:     ch,
		}
	case 'r':
		return &ParseResult{
			Count:     p.effectiveCount(),
			SimpleCmd: SimpleCmdReplaceChar,
			Char:      ch,
		}
	}
	return nil
}

func (p *Parser) effectiveCount() int {
	c := p.count
	if c == 0 {
		c = 1
	}
	if p.secondCount > 0 {
		c *= p.secondCount
	}
	return c
}

// IsOperatorPending returns true if waiting for a motion/text-object after an operator.
func (p *Parser) IsOperatorPending() bool {
	return p.operator != OperatorNone
}

func resolveMotion(ch rune) MotionType {
	switch ch {
	case 'h':
		return MotionLeft
	case 'j':
		return MotionDown
	case 'k':
		return MotionUp
	case 'l':
		return MotionRight
	case 'w':
		return MotionWordForward
	case 'b':
		return MotionWordBackward
	case 'e':
		return MotionWordEnd
	case '0':
		return MotionLineStart
	case '$':
		return MotionLineEnd
	case 'G':
		return MotionFileBottom
	}
	return MotionNone
}

func resolveTextObject(prefix, ch rune) TextObjectType {
	inner := prefix == 'i'
	switch ch {
	case 'w':
		if inner {
			return TextObjectInnerWord
		}
		return TextObjectAWord
	case '"':
		if inner {
			return TextObjectInnerDoubleQuote
		}
		return TextObjectADoubleQuote
	case '\'':
		if inner {
			return TextObjectInnerSingleQuote
		}
		return TextObjectASingleQuote
	case '(', ')':
		if inner {
			return TextObjectInnerParen
		}
		return TextObjectAParen
	case '{', '}':
		if inner {
			return TextObjectInnerBrace
		}
		return TextObjectABrace
	}
	return TextObjectNone
}

func resolveSimpleCmd(ch rune) SimpleCommand {
	switch ch {
	case 'i':
		return SimpleCmdInsertBefore
	case 'a':
		return SimpleCmdInsertAfter
	case 'o':
		return SimpleCmdOpenBelow
	case 'I':
		return SimpleCmdInsertLineStart
	case 'A':
		return SimpleCmdInsertLineEnd
	case 'O':
		return SimpleCmdOpenAbove
	case 'x':
		return SimpleCmdDeleteChar
	case 'J':
		return SimpleCmdJoinLine
	case '.':
		return SimpleCmdDotRepeat
	case 'u':
		return SimpleCmdUndo
	case 'p':
		return SimpleCmdPaste
	case 'P':
		return SimpleCmdPasteBefore
	case ':':
		return SimpleCmdEnterCommand
	case '/':
		return SimpleCmdSearchForward
	case 'n':
		return SimpleCmdSearchNext
	case 'N':
		return SimpleCmdSearchPrev
	case 'v':
		return SimpleCmdVisual
	case 'V':
		return SimpleCmdVisualLine
	}
	return SimpleCmdNone
}
