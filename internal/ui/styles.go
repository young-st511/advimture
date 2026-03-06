package ui

import "github.com/charmbracelet/lipgloss"

// Mode colors
var (
	NormalColor      = lipgloss.Color("#5B7FFF") // blue
	InsertColor      = lipgloss.Color("#50C878") // green
	CommandColor     = lipgloss.Color("#FFD700") // yellow
	VisualColor      = lipgloss.Color("#9B59B6") // purple
	OpPendingColor   = lipgloss.Color("#FF8C00") // orange
	MentorColor      = lipgloss.Color("#00CED1") // cyan
	ErrorColor       = lipgloss.Color("#FF4444") // red
	SuccessColor     = lipgloss.Color("#50C878") // green
	DimColor         = lipgloss.Color("#555555") // dim gray
	LineNumberColor  = lipgloss.Color("#888888") // gray
	EmptyLineColor   = lipgloss.Color("#444444") // dark gray
)

// StatusBarStyle returns a lipgloss.Style for the status bar based on the mode name.
func StatusBarStyle(mode string) lipgloss.Style {
	var bg lipgloss.Color
	switch mode {
	case "NORMAL":
		bg = NormalColor
	case "INSERT":
		bg = InsertColor
	case "COMMAND":
		bg = CommandColor
	case "VISUAL":
		bg = VisualColor
	case "OP-PENDING":
		bg = OpPendingColor
	default:
		bg = NormalColor
	}

	return lipgloss.NewStyle().
		Background(bg).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1).
		Bold(true)
}

// Border style for main editor area.
var EditorBorder = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#555555"))

// Mentor message style (Vi 선배).
var MentorStyle = lipgloss.NewStyle().
	Foreground(MentorColor).
	Italic(true)

// Header style.
var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Bold(true).
	Padding(0, 1)

// Dim text.
var DimStyle = lipgloss.NewStyle().
	Foreground(DimColor)

// Line number style.
var LineNumberStyle = lipgloss.NewStyle().
	Foreground(LineNumberColor)

// Empty line (~) style.
var EmptyLineStyle = lipgloss.NewStyle().
	Foreground(EmptyLineColor)
