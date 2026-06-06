package e2estate

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	Buffer    []string   `json:"buffer"`
	Cursor    Cursor     `json:"cursor"`
	Mode      string     `json:"mode"`
	Command   string     `json:"command,omitempty"`
	Status    string     `json:"status"`
	Score     Score      `json:"score"`
	Progress  Progress   `json:"progress"`
	Review    Review     `json:"review"`
	UI        UI         `json:"ui"`
	Selection *Selection `json:"selection,omitempty"`
}

type Cursor struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Selection struct {
	Active bool   `json:"active"`
	Kind   string `json:"kind"`
	Anchor Cursor `json:"anchor"`
	Head   Cursor `json:"head"`
	Start  Cursor `json:"start"`
	End    Cursor `json:"end"`
}

type Score struct {
	Grade  string `json:"grade"`
	Passed bool   `json:"passed"`
}

type Progress struct {
	MissionID string `json:"mission_id"`
	Completed bool   `json:"completed"`
}

type Review struct {
	QueueCount        int    `json:"queue_count"`
	PrimaryExerciseID string `json:"primary_exercise_id,omitempty"`
	PrimaryReason     string `json:"primary_reason,omitempty"`
	DailyRoute        string `json:"daily_route,omitempty"`
}

type UI struct {
	FocusPanel FocusPanel `json:"focus_panel"`
}

type FocusPanel struct {
	Kind    string       `json:"kind"`
	Title   string       `json:"title"`
	Lines   []string     `json:"lines"`
	Actions []ActionLine `json:"actions,omitempty"`
}

type ActionLine struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func DefaultPath(home string) string {
	return filepath.Join(home, ".advimture", "e2e_state.json")
}

func Write(path string, state State) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}
