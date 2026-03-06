package data

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed tutorials/*.yaml
var tutorialFS embed.FS

// TutorialData represents a tutorial loaded from YAML.
type TutorialData struct {
	ID          string        `yaml:"id"`
	Title       string        `yaml:"title"`
	Description string        `yaml:"description"`
	Substeps    []SubstepData `yaml:"substeps"`
}

// SubstepData represents a single substep within a tutorial.
type SubstepData struct {
	ID          string   `yaml:"id"`
	Title       string   `yaml:"title"`
	Instruction string   `yaml:"instruction"`
	InitialText string   `yaml:"initial_text"`
	CursorRow   int      `yaml:"cursor_row"`
	CursorCol   int      `yaml:"cursor_col"`
	Goal        GoalData `yaml:"goal"`
	Hints       []string `yaml:"hints"`
	AllowedKeys []string `yaml:"allowed_keys,omitempty"` // empty = all keys allowed
	MentorMsg   string   `yaml:"mentor_msg"`
}

// GoalData defines the completion condition for a substep.
type GoalData struct {
	Type  string `yaml:"type"` // cursor_position, cursor_on_char, text_match, save_quit, mode_is, command_used
	Row   int    `yaml:"row,omitempty"`
	Col   int    `yaml:"col,omitempty"`
	Char  string `yaml:"char,omitempty"`
	Text  string `yaml:"text,omitempty"`
	Mode  string `yaml:"mode,omitempty"`
	Cmd   string `yaml:"cmd,omitempty"`
}

// LoadTutorial loads a tutorial YAML file by filename.
func LoadTutorial(filename string) (*TutorialData, error) {
	data, err := tutorialFS.ReadFile("tutorials/" + filename)
	if err != nil {
		return nil, fmt.Errorf("[튜토리얼 파일 로드 실패] %s: %w", filename, err)
	}

	var t TutorialData
	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("[튜토리얼 YAML 파싱 실패] %s: %w", filename, err)
	}

	return &t, nil
}

// LoadAllTutorials loads all tutorial YAML files.
func LoadAllTutorials() ([]*TutorialData, error) {
	entries, err := tutorialFS.ReadDir("tutorials")
	if err != nil {
		return nil, fmt.Errorf("[튜토리얼 디렉토리 읽기 실패] %w", err)
	}

	var tutorials []*TutorialData
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		t, err := LoadTutorial(entry.Name())
		if err != nil {
			return nil, err
		}
		tutorials = append(tutorials, t)
	}
	return tutorials, nil
}
