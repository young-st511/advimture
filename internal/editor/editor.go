package editor

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Mode represents the editor mode.
type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeCommand
	ModeVisual
	ModeOperatorPending
)

func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	case ModeCommand:
		return "COMMAND"
	case ModeVisual:
		return "VISUAL"
	case ModeOperatorPending:
		return "OP-PENDING"
	default:
		return "NORMAL"
	}
}

// Model is the Bubbletea model for the Vim editor.
type Model struct {
	buf    *Buffer
	cursor *Cursor
	parser *Parser
	reg    *Register
	undo   *UndoManager

	mode     Mode
	width    int
	height   int
	quitting bool

	// Command mode state
	cmdInput string

	// Status message (shown briefly)
	statusMsg   string
	statusError bool

	// Ctrl+C tracking
	ctrlCCount int

	// Keystroke counter
	keystrokes int
}

// New creates a new editor model with sample text.
func New() Model {
	sample := []string{
		"Hello, Advimture!",
		"",
		"This is a Vim practice game.",
		"Use hjkl to move around.",
		"Press i to enter Insert mode.",
		"Press :wq to save and quit.",
	}
	return NewWithLines(sample)
}

// NewWithLines creates a new editor model with the given lines.
func NewWithLines(lines []string) Model {
	return Model{
		buf:    NewBuffer(lines),
		cursor: NewCursor(),
		parser: NewParser(),
		reg:    NewRegister(),
		undo:   NewUndoManager(),
		mode:   ModeNormal,
		width:  80,
		height: 24,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles key input and returns the updated model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		m.ctrlCCount = countCtrlC(msg, m.ctrlCCount)
		if m.ctrlCCount >= 2 {
			m.quitting = true
			return m, nil
		}

		m.keystrokes++
		m.statusMsg = ""

		switch m.mode {
		case ModeNormal, ModeOperatorPending:
			m = m.handleNormalKey(msg)
		case ModeInsert:
			m = m.handleInsertKey(msg)
		case ModeCommand:
			m = m.handleCommandKey(msg)
		}
	}

	return m, nil
}

func countCtrlC(msg tea.KeyMsg, prev int) int {
	if msg.String() == "ctrl+c" {
		return prev + 1
	}
	return 0
}

func (m Model) handleNormalKey(msg tea.KeyMsg) Model {
	key := msg.String()

	// Escape resets parser
	if key == "esc" {
		m.parser.Reset()
		m.mode = ModeNormal
		return m
	}

	result := m.parser.Feed(key)
	if result == nil {
		// Parser still accumulating
		if m.parser.IsOperatorPending() {
			m.mode = ModeOperatorPending
		}
		return m
	}

	// Handle simple commands that affect mode/screen (not buffer operations)
	switch result.SimpleCmd {
	case SimpleCmdEnterCommand:
		m.mode = ModeCommand
		m.cmdInput = ""
		return m
	case SimpleCmdSearchForward:
		m.mode = ModeCommand
		m.cmdInput = ""
		m.statusMsg = "/ 검색은 아직 미구현"
		m.mode = ModeNormal
		return m
	}

	// Execute the parsed command
	execResult := Execute(m.buf, m.cursor, m.reg, m.undo, result)

	if execResult.SwitchToInsert {
		m.mode = ModeInsert
	} else {
		m.mode = ModeNormal
	}

	// Clamp cursor
	m.cursor.ClampToBuffer(m.buf, m.mode == ModeInsert)

	return m
}

func (m Model) handleInsertKey(msg tea.KeyMsg) Model {
	key := msg.String()

	switch key {
	case "esc":
		m.mode = ModeNormal
		// Vim: cursor moves left one when leaving insert mode
		if m.cursor.Col > 0 {
			m.cursor.Col--
		}
		m.cursor.DesiredCol = m.cursor.Col
		m.cursor.ClampToBuffer(m.buf, false)
		return m

	case "backspace":
		if m.cursor.Col > 0 {
			m.cursor.Col--
			ch := m.buf.DeleteChar(m.cursor.Row, m.cursor.Col)
			if ch != 0 {
				m.undo.Record(Operation{
					Type: OpDeleteChar, Row: m.cursor.Row, Col: m.cursor.Col, Char: ch,
					CursorRow: m.cursor.Row, CursorCol: m.cursor.Col + 1,
				})
			}
		} else if m.cursor.Row > 0 {
			// Join with previous line
			prevLen := m.buf.LineRuneLen(m.cursor.Row - 1)
			m.undo.Record(Operation{
				Type: OpJoinLines, Row: m.cursor.Row - 1, Col: prevLen,
				CursorRow: m.cursor.Row, CursorCol: 0,
			})
			m.buf.JoinLines(m.cursor.Row - 1)
			m.cursor.Row--
			m.cursor.Col = prevLen
		}
		return m

	case "enter":
		col := m.cursor.Col
		m.undo.Record(Operation{
			Type: OpSplitLine, Row: m.cursor.Row, Col: col,
			CursorRow: m.cursor.Row, CursorCol: col,
		})
		m.buf.SplitLine(m.cursor.Row, col)
		m.cursor.Row++
		m.cursor.Col = 0
		m.cursor.DesiredCol = 0
		return m

	default:
		// Insert character
		if len(key) == 1 {
			ch := rune(key[0])
			m.undo.Record(Operation{
				Type: OpInsertChar, Row: m.cursor.Row, Col: m.cursor.Col, Char: ch,
				CursorRow: m.cursor.Row, CursorCol: m.cursor.Col,
			})
			m.buf.InsertChar(m.cursor.Row, m.cursor.Col, ch)
			m.cursor.Col++
		} else {
			// Multi-byte rune
			runes := []rune(key)
			if len(runes) == 1 {
				ch := runes[0]
				m.undo.Record(Operation{
					Type: OpInsertChar, Row: m.cursor.Row, Col: m.cursor.Col, Char: ch,
					CursorRow: m.cursor.Row, CursorCol: m.cursor.Col,
				})
				m.buf.InsertChar(m.cursor.Row, m.cursor.Col, ch)
				m.cursor.Col++
			}
		}
		return m
	}
}

func (m Model) handleCommandKey(msg tea.KeyMsg) Model {
	key := msg.String()

	switch key {
	case "esc":
		m.mode = ModeNormal
		m.cmdInput = ""
		return m

	case "enter":
		result := ExecuteCommand(m.cmdInput, m.buf, m.cursor, m.undo)
		m.mode = ModeNormal
		m.cmdInput = ""

		if result.Error != "" {
			m.statusMsg = result.Error
			m.statusError = true
		} else if result.Message != "" {
			m.statusMsg = result.Message
			m.statusError = false
		}

		if result.GotoLine >= 0 {
			m.cursor.Row = result.GotoLine
			m.cursor.Col = 0
			m.cursor.DesiredCol = 0
			m.cursor.ClampToBuffer(m.buf, false)
		}

		if result.Quit {
			m.quitting = true
		}
		return m

	case "backspace":
		if len(m.cmdInput) > 0 {
			m.cmdInput = m.cmdInput[:len(m.cmdInput)-1]
		} else {
			m.mode = ModeNormal
		}
		return m

	default:
		if len(key) == 1 || len([]rune(key)) == 1 {
			m.cmdInput += key
		}
		return m
	}
}

// Quitting returns whether the editor is quitting.
func (m Model) Quitting() bool {
	return m.quitting
}

// Getters for view rendering
func (m Model) GetBuffer() *Buffer     { return m.buf }
func (m Model) GetCursor() *Cursor     { return m.cursor }
func (m Model) GetMode() Mode          { return m.mode }
func (m Model) GetCmdInput() string    { return m.cmdInput }
func (m Model) GetStatusMsg() string   { return m.statusMsg }
func (m Model) IsStatusError() bool    { return m.statusError }
func (m Model) GetKeystrokes() int     { return m.keystrokes }
