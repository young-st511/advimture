package data

import (
	"strings"
)

// GoalChecker holds the state needed to check goals.
type GoalChecker struct {
	LastCommandUsed string
	SaveQuitCalled  bool
}

// NewGoalChecker creates a new goal checker.
func NewGoalChecker() *GoalChecker {
	return &GoalChecker{}
}

// CheckGoal verifies if the goal is met given current state.
func (gc *GoalChecker) CheckGoal(goal GoalData, bufferText string, curRow, curCol int, mode string) bool {
	switch goal.Type {
	case "cursor_position":
		return curRow == goal.Row && curCol == goal.Col

	case "cursor_on_char":
		lines := strings.Split(bufferText, "\n")
		if curRow < 0 || curRow >= len(lines) {
			return false
		}
		runes := []rune(lines[curRow])
		if curCol < 0 || curCol >= len(runes) {
			return false
		}
		return string(runes[curCol]) == goal.Char

	case "text_match":
		return strings.TrimRight(bufferText, "\n") == strings.TrimRight(goal.Text, "\n")

	case "save_quit":
		if !gc.SaveQuitCalled {
			return false
		}
		if goal.Text != "" {
			return strings.TrimRight(bufferText, "\n") == strings.TrimRight(goal.Text, "\n")
		}
		return true

	case "mode_is":
		return strings.EqualFold(mode, goal.Mode)

	case "command_used":
		return gc.LastCommandUsed == goal.Cmd

	default:
		return false
	}
}

// RecordCommand records that a command was used.
func (gc *GoalChecker) RecordCommand(cmd string) {
	gc.LastCommandUsed = cmd
}

// RecordSaveQuit records that :wq was used.
func (gc *GoalChecker) RecordSaveQuit() {
	gc.SaveQuitCalled = true
}

// Reset clears the recorded state for a new substep.
func (gc *GoalChecker) Reset() {
	gc.LastCommandUsed = ""
	gc.SaveQuitCalled = false
}
