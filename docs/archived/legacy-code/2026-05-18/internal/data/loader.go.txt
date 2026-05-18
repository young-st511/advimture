package data

import (
	"embed"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed tutorials/*.yaml
var tutorialFS embed.FS

//go:embed missions/*.yaml
var missionFS embed.FS

// ─── Tutorial types ───────────────────────────────────────────────────────────

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
	Type string `yaml:"type"` // cursor_position, cursor_on_char, text_match, save_quit, quit, mode_is, command_used
	Row  int    `yaml:"row,omitempty"`
	Col  int    `yaml:"col,omitempty"`
	Char string `yaml:"char,omitempty"`
	Text string `yaml:"text,omitempty"`
	Mode string `yaml:"mode,omitempty"`
	Cmd  string `yaml:"cmd,omitempty"`
}

// LoadTutorial loads a tutorial YAML file by filename.
func LoadTutorial(filename string) (*TutorialData, error) {
	raw, err := tutorialFS.ReadFile("tutorials/" + filename)
	if err != nil {
		return nil, fmt.Errorf("[튜토리얼 파일 로드 실패] %s: %w", filename, err)
	}

	var t TutorialData
	if err := yaml.Unmarshal(raw, &t); err != nil {
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

// ─── Mission types ─────────────────────────────────────────────────────────────

// MissionData represents a mission loaded from YAML.
type MissionData struct {
	ID                string            `yaml:"id"`
	Title             string            `yaml:"title"`
	Difficulty        int               `yaml:"difficulty"` // 1~3
	Story             string            `yaml:"story"`
	InitialText       string            `yaml:"initial_text"`
	ExpectedText      string            `yaml:"expected_text"`
	CursorStart       MissionCursorPos  `yaml:"cursor_start"`
	OptimalKeystrokes int               `yaml:"optimal_keystrokes"`
	GoalType          string            `yaml:"goal_type,omitempty"`    // "text_match" (default) or "cursor_on_line"
	GoalPattern       string            `yaml:"goal_pattern,omitempty"` // pattern for cursor_on_line goal
	RequiredTutorials []string          `yaml:"required_tutorials,omitempty"`
	OptimalSolutions  []OptimalSolution `yaml:"optimal_solutions,omitempty"`
	Tips              []MissionTip      `yaml:"tips,omitempty"`
}

// MissionCursorPos holds the starting cursor position for a mission.
type MissionCursorPos struct {
	Row int `yaml:"row"`
	Col int `yaml:"col"`
}

// OptimalSolution describes one optimal approach for a mission.
type OptimalSolution struct {
	Description string `yaml:"description"`
	Keys        string `yaml:"keys"`
	Count       int    `yaml:"count"`
}

// MissionTip is a contextual hint shown when a trigger condition is met.
type MissionTip struct {
	Trigger string `yaml:"trigger"`
	Message string `yaml:"message"`
}

// DiffLine represents a single line difference between expected and actual text.
type DiffLine struct {
	LineNum  int
	Expected string
	Yours    string
}

// LoadMission loads a mission YAML file by filename.
func LoadMission(filename string) (*MissionData, error) {
	raw, err := missionFS.ReadFile("missions/" + filename)
	if err != nil {
		return nil, fmt.Errorf("[미션 파일 로드 실패] %s: %w", filename, err)
	}

	var m MissionData
	if err := yaml.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("[미션 YAML 파싱 실패] %s: %w", filename, err)
	}

	return &m, nil
}

// LoadAllMissions loads all mission YAML files (m00_test.yaml excluded).
func LoadAllMissions() ([]*MissionData, error) {
	entries, err := missionFS.ReadDir("missions")
	if err != nil {
		return nil, fmt.Errorf("[미션 디렉토리 읽기 실패] %w", err)
	}

	var missions []*MissionData
	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == "m00_test.yaml" {
			continue
		}
		m, err := LoadMission(entry.Name())
		if err != nil {
			return nil, err
		}
		missions = append(missions, m)
	}
	return missions, nil
}

// ValidateMission checks that a MissionData has all required fields.
func ValidateMission(m *MissionData) error {
	if m.ID == "" {
		return fmt.Errorf("mission ID가 비어있습니다")
	}
	if m.Title == "" {
		return fmt.Errorf("mission Title이 비어있습니다")
	}
	if m.InitialText == "" {
		return fmt.Errorf("mission InitialText가 비어있습니다")
	}
	if m.GoalType == "" || m.GoalType == "text_match" {
		if m.ExpectedText == "" {
			return fmt.Errorf("text_match 미션에 ExpectedText가 필요합니다")
		}
	}
	if m.OptimalKeystrokes <= 0 {
		return fmt.Errorf("OptimalKeystrokes는 0보다 커야 합니다")
	}
	return nil
}

// CompareText compares expected and actual text line by line, ignoring trailing whitespace.
// Returns (match bool, diffs []DiffLine, totalDiff int).
// diffs contains at most 3 entries; totalDiff is the full count of differing lines.
func CompareText(expected, actual string) (bool, []DiffLine, int) {
	expectedLines := strings.Split(strings.TrimRight(expected, "\n"), "\n")
	actualLines := strings.Split(strings.TrimRight(actual, "\n"), "\n")

	if len(expectedLines) != len(actualLines) {
		diff := DiffLine{
			LineNum:  0,
			Expected: fmt.Sprintf("(%d줄)", len(expectedLines)),
			Yours:    fmt.Sprintf("(%d줄)", len(actualLines)),
		}
		return false, []DiffLine{diff}, 1
	}

	var diffs []DiffLine
	totalDiff := 0
	for i, exp := range expectedLines {
		expTrimmed := strings.TrimRight(exp, " \t")
		actTrimmed := strings.TrimRight(actualLines[i], " \t")
		if expTrimmed != actTrimmed {
			totalDiff++
			if len(diffs) < 3 {
				diffs = append(diffs, DiffLine{
					LineNum:  i + 1,
					Expected: expTrimmed,
					Yours:    actTrimmed,
				})
			}
		}
	}

	return totalDiff == 0, diffs, totalDiff
}
