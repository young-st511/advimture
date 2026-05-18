package e2estate

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	Buffer   []string `json:"buffer"`
	Cursor   Cursor   `json:"cursor"`
	Mode     string   `json:"mode"`
	Command  string   `json:"command,omitempty"`
	Status   string   `json:"status"`
	Score    Score    `json:"score"`
	Progress Progress `json:"progress"`
}

type Cursor struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Score struct {
	Grade  string `json:"grade"`
	Passed bool   `json:"passed"`
}

type Progress struct {
	MissionID string `json:"mission_id"`
	Completed bool   `json:"completed"`
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
