package content

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Status string

const (
	StatusDraft       Status = "draft"
	StatusApproved    Status = "approved"
	StatusImplemented Status = "implemented"
	StatusRetired     Status = "retired"
)

type EngineSupport string

const (
	EngineSupportImplemented EngineSupport = "implemented"
	EngineSupportPlanned     EngineSupport = "planned"
	EngineSupportUnsupported EngineSupport = "unsupported"
)

type CommandCluster struct {
	ID                 string        `yaml:"id"`
	Status             Status        `yaml:"status"`
	CompatibilityTier  string        `yaml:"compatibility_tier"`
	EngineSupport      EngineSupport `yaml:"engine_support"`
	CurriculumArea     string        `yaml:"curriculum_area"`
	Title              string        `yaml:"title"`
	Commands           []string      `yaml:"commands"`
	CoverageRequired   []string      `yaml:"coverage_required"`
	Oracle             string        `yaml:"oracle"`
	Purpose            string        `yaml:"purpose"`
	Prerequisite       []string      `yaml:"prerequisite"`
	Difficulty         string        `yaml:"difficulty"`
	UsefulWhen         []string      `yaml:"useful_when"`
	ComboPaths         [][]string    `yaml:"combo_paths"`
	CommonMistakes     []string      `yaml:"common_mistakes"`
	CompatibilityNotes []string      `yaml:"compatibility_notes"`
	DesignNotes        []string      `yaml:"design_notes"`
}

type ExerciseDocument struct {
	ID               string           `yaml:"id"`
	Status           Status           `yaml:"status"`
	CommandCluster   string           `yaml:"command_cluster"`
	EngineSupport    EngineSupport    `yaml:"engine_support"`
	TrainedCommands  []string         `yaml:"trained_commands"`
	ReviewedCommands []string         `yaml:"reviewed_commands"`
	MistakeFocus     []string         `yaml:"mistake_focus"`
	ReplayStatus     string           `yaml:"replay_status"`
	Title            string           `yaml:"title"`
	GoalForPlayer    string           `yaml:"goal_for_player"`
	InitialState     YAMLState        `yaml:"initial_state"`
	TargetState      YAMLGoal         `yaml:"target_state"`
	OptimalKeys      []string         `yaml:"optimal_keys"`
	AllowedKeys      []string         `yaml:"allowed_keys"`
	ForbiddenKeys    []string         `yaml:"forbidden_keys"`
	Hints            []HintSpec       `yaml:"hints"`
	Grading          GradingSpec      `yaml:"grading"`
	E2EAssertions    E2EAssertionSpec `yaml:"e2e_assertions"`
}

type YAMLState struct {
	Mode   string      `yaml:"mode"`
	Cursor *CursorSpec `yaml:"cursor"`
	Buffer string      `yaml:"buffer"`
}

type YAMLGoal struct {
	Mode   string      `yaml:"mode"`
	Cursor *CursorSpec `yaml:"cursor"`
	Buffer string      `yaml:"buffer"`
}

type GradingSpec struct {
	PassCondition   string `yaml:"pass_condition"`
	OptimalKeyCount int    `yaml:"optimal_key_count"`
}

type E2EAssertionSpec struct {
	Buffer []string    `yaml:"buffer"`
	Cursor *CursorSpec `yaml:"cursor"`
	Mode   string      `yaml:"mode"`
	Status string      `yaml:"status"`
}

type ScenarioDocument struct {
	ID                    string        `yaml:"id"`
	Status                Status        `yaml:"status"`
	ExerciseID            string        `yaml:"exercise_id"`
	EngineSupport         EngineSupport `yaml:"engine_support"`
	LearningReinforcement string        `yaml:"learning_reinforcement"`
	DoesNotChange         []string      `yaml:"does_not_change"`
	MissionTitle          string        `yaml:"mission_title"`
	Briefing              string        `yaml:"briefing"`
	ContextRole           string        `yaml:"context_role"`
	MentorSuccess         string        `yaml:"mentor_success"`
	MentorFailure         string        `yaml:"mentor_failure"`
	StoryConstraints      []string      `yaml:"story_constraints"`
}

type PlaylistDocument struct {
	ID               string         `yaml:"id"`
	Status           Status         `yaml:"status"`
	Title            string         `yaml:"title"`
	UnlockPolicy     string         `yaml:"unlock_policy"`
	CompletionPolicy string         `yaml:"completion_policy"`
	Beats            []PlaylistBeat `yaml:"beats"`
}

type PlaylistBeat struct {
	ID             string        `yaml:"id"`
	Role           string        `yaml:"role"`
	CommandCluster string        `yaml:"command_cluster"`
	ExerciseID     string        `yaml:"exercise_id"`
	ScenarioID     string        `yaml:"scenario_id"`
	EngineSupport  EngineSupport `yaml:"engine_support"`
}

type Library struct {
	CommandClusters map[string]CommandCluster
	Exercises       map[string]ExerciseDocument
	Scenarios       map[string]ScenarioDocument
	Playlists       map[string]PlaylistDocument
}

type CoverageReport struct {
	CommandClusterID string
	Required         []string
	Covered          []string
	Missing          []string
}

type commandClusterFile struct {
	CommandClusters []CommandCluster `yaml:"command_clusters"`
}

type exerciseFile struct {
	Exercises []ExerciseDocument `yaml:"exercises"`
}

type scenarioFile struct {
	Scenarios []ScenarioDocument `yaml:"scenarios"`
}

type playlistFile struct {
	Playlists []PlaylistDocument `yaml:"playlists"`
}

func LoadLibrary(root string) (Library, error) {
	lib := Library{
		CommandClusters: make(map[string]CommandCluster),
		Exercises:       make(map[string]ExerciseDocument),
		Scenarios:       make(map[string]ScenarioDocument),
		Playlists:       make(map[string]PlaylistDocument),
	}

	if err := loadYAMLDir(filepath.Join(root, "command_clusters"), func(path string, raw []byte) error {
		var file commandClusterFile
		if err := yaml.Unmarshal(raw, &file); err != nil {
			return err
		}
		for _, cluster := range file.CommandClusters {
			if err := addUnique(lib.CommandClusters, cluster.ID, cluster, "command cluster", path); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return Library{}, err
	}

	if err := loadYAMLDir(filepath.Join(root, "exercises"), func(path string, raw []byte) error {
		var file exerciseFile
		if err := yaml.Unmarshal(raw, &file); err != nil {
			return err
		}
		for _, exercise := range file.Exercises {
			if err := addUnique(lib.Exercises, exercise.ID, exercise, "exercise", path); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return Library{}, err
	}

	if err := loadYAMLDir(filepath.Join(root, "scenarios"), func(path string, raw []byte) error {
		var file scenarioFile
		if err := yaml.Unmarshal(raw, &file); err != nil {
			return err
		}
		for _, scenario := range file.Scenarios {
			if err := addUnique(lib.Scenarios, scenario.ID, scenario, "scenario", path); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return Library{}, err
	}

	if err := loadYAMLDir(filepath.Join(root, "playlists"), func(path string, raw []byte) error {
		var file playlistFile
		if err := yaml.Unmarshal(raw, &file); err != nil {
			return err
		}
		for _, playlist := range file.Playlists {
			if err := addUnique(lib.Playlists, playlist.ID, playlist, "playlist", path); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return Library{}, err
	}

	if err := lib.Validate(); err != nil {
		return Library{}, err
	}
	return lib, nil
}

func (l Library) Validate() error {
	for _, cluster := range l.CommandClusters {
		if strings.TrimSpace(cluster.ID) == "" {
			return fmt.Errorf("command cluster id is required")
		}
		if !validStatus(cluster.Status) {
			return fmt.Errorf("command cluster %q has invalid status %q", cluster.ID, cluster.Status)
		}
		if !validEngineSupport(cluster.EngineSupport) {
			return fmt.Errorf("command cluster %q has invalid engine_support %q", cluster.ID, cluster.EngineSupport)
		}
	}

	for _, exercise := range l.Exercises {
		if err := l.validateExerciseDocument(exercise); err != nil {
			return err
		}
	}

	for _, scenario := range l.Scenarios {
		if err := l.validateScenarioDocument(scenario); err != nil {
			return err
		}
	}

	for _, playlist := range l.Playlists {
		if err := l.validatePlaylistDocument(playlist); err != nil {
			return err
		}
	}
	return nil
}

func (l Library) PlayableExercises() []ExerciseDocument {
	exercises := make([]ExerciseDocument, 0)
	for _, exercise := range l.Exercises {
		if exercise.Status == StatusApproved || exercise.Status == StatusImplemented {
			if exercise.EngineSupport == EngineSupportImplemented {
				exercises = append(exercises, exercise)
			}
		}
	}
	sort.Slice(exercises, func(i, j int) bool {
		return exercises[i].ID < exercises[j].ID
	})
	return exercises
}

func (l Library) CoverageReports() []CoverageReport {
	reports := make([]CoverageReport, 0, len(l.CommandClusters))
	for _, cluster := range l.CommandClusters {
		coveredSet := make(map[string]bool)
		for _, exercise := range l.Exercises {
			if exercise.CommandCluster != cluster.ID || !isApprovedLike(exercise.Status) {
				continue
			}
			for _, key := range exercise.OptimalKeys {
				coveredSet[key] = true
			}
		}

		var covered []string
		var missing []string
		for _, required := range cluster.CoverageRequired {
			if coveredSet[required] {
				covered = append(covered, required)
			} else {
				missing = append(missing, required)
			}
		}
		reports = append(reports, CoverageReport{
			CommandClusterID: cluster.ID,
			Required:         copyStrings(cluster.CoverageRequired),
			Covered:          covered,
			Missing:          missing,
		})
	}
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].CommandClusterID < reports[j].CommandClusterID
	})
	return reports
}

func (l Library) CompileExercise(id string) (CompiledExercise, error) {
	exercise, ok := l.Exercises[id]
	if !ok {
		return CompiledExercise{}, fmt.Errorf("exercise %q not found", id)
	}
	return CompileExercise(exercise.ToExerciseSpec())
}

func (e ExerciseDocument) ToExerciseSpec() ExerciseSpec {
	return ExerciseSpec{
		ID:               e.ID,
		CommandClusterID: e.CommandCluster,
		Title:            e.Title,
		Initial:          e.InitialState.toStateSpec(),
		Goal:             e.TargetState.toGoalSpec(),
		Hints:            copyHints(e.Hints),
		ExpectedKeys:     copyStrings(e.OptimalKeys),
		AllowedKeys:      copyStrings(e.AllowedKeys),
	}
}

func (s YAMLState) toStateSpec() StateSpec {
	return StateSpec{
		Lines:  splitBuffer(s.Buffer),
		Cursor: copyCursor(s.Cursor),
		Mode:   s.Mode,
	}
}

func (g YAMLGoal) toGoalSpec() GoalSpec {
	return GoalSpec{
		Lines:  optionalBuffer(g.Buffer),
		Cursor: copyCursor(g.Cursor),
		Mode:   g.Mode,
	}
}

func (l Library) validateExerciseDocument(exercise ExerciseDocument) error {
	if strings.TrimSpace(exercise.ID) == "" {
		return fmt.Errorf("exercise id is required")
	}
	cluster, ok := l.CommandClusters[exercise.CommandCluster]
	if !ok {
		return fmt.Errorf("exercise %q references missing command cluster %q", exercise.ID, exercise.CommandCluster)
	}
	if !validStatus(exercise.Status) {
		return fmt.Errorf("exercise %q has invalid status %q", exercise.ID, exercise.Status)
	}
	if !validEngineSupport(exercise.EngineSupport) {
		return fmt.Errorf("exercise %q has invalid engine_support %q", exercise.ID, exercise.EngineSupport)
	}
	if isApprovedLike(exercise.Status) && !isApprovedLike(cluster.Status) {
		return fmt.Errorf("approved exercise %q references non-approved command cluster %q", exercise.ID, cluster.ID)
	}
	if exercise.Grading.OptimalKeyCount != len(exercise.OptimalKeys) {
		return fmt.Errorf("exercise %q optimal_key_count = %d, want %d", exercise.ID, exercise.Grading.OptimalKeyCount, len(exercise.OptimalKeys))
	}
	if err := validateKeys(exercise); err != nil {
		return err
	}
	if _, err := CompileExercise(exercise.ToExerciseSpec()); err != nil {
		return fmt.Errorf("exercise %q compile failed: %w", exercise.ID, err)
	}
	return nil
}

func (l Library) validateScenarioDocument(scenario ScenarioDocument) error {
	if strings.TrimSpace(scenario.ID) == "" {
		return fmt.Errorf("scenario id is required")
	}
	exercise, ok := l.Exercises[scenario.ExerciseID]
	if !ok {
		return fmt.Errorf("scenario %q references missing exercise %q", scenario.ID, scenario.ExerciseID)
	}
	if !validStatus(scenario.Status) {
		return fmt.Errorf("scenario %q has invalid status %q", scenario.ID, scenario.Status)
	}
	if !validEngineSupport(scenario.EngineSupport) {
		return fmt.Errorf("scenario %q has invalid engine_support %q", scenario.ID, scenario.EngineSupport)
	}
	if isApprovedLike(scenario.Status) && !isApprovedLike(exercise.Status) {
		return fmt.Errorf("approved scenario %q references non-approved exercise %q", scenario.ID, exercise.ID)
	}
	return nil
}

func (l Library) validatePlaylistDocument(playlist PlaylistDocument) error {
	if strings.TrimSpace(playlist.ID) == "" {
		return fmt.Errorf("playlist id is required")
	}
	for _, beat := range playlist.Beats {
		if _, ok := l.CommandClusters[beat.CommandCluster]; !ok {
			return fmt.Errorf("playlist %q beat %q references missing command cluster %q", playlist.ID, beat.ID, beat.CommandCluster)
		}
		if _, ok := l.Exercises[beat.ExerciseID]; !ok {
			return fmt.Errorf("playlist %q beat %q references missing exercise %q", playlist.ID, beat.ID, beat.ExerciseID)
		}
		if _, ok := l.Scenarios[beat.ScenarioID]; !ok {
			return fmt.Errorf("playlist %q beat %q references missing scenario %q", playlist.ID, beat.ID, beat.ScenarioID)
		}
	}
	return nil
}

func validateKeys(exercise ExerciseDocument) error {
	allowed := make(map[string]bool, len(exercise.AllowedKeys))
	for _, key := range exercise.AllowedKeys {
		allowed[key] = true
	}
	for _, key := range exercise.ForbiddenKeys {
		if allowed[key] {
			return fmt.Errorf("exercise %q key %q is both allowed and forbidden", exercise.ID, key)
		}
	}
	for _, key := range exercise.OptimalKeys {
		if !allowed[key] {
			return fmt.Errorf("exercise %q optimal key %q is not allowed", exercise.ID, key)
		}
	}
	return nil
}

func loadYAMLDir(dir string, visit func(path string, raw []byte) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", dir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !isYAML(entry.Name()) {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		raw, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		if err := visit(path, raw); err != nil {
			return fmt.Errorf("load %s: %w", path, err)
		}
	}
	return nil
}

func addUnique[T any](values map[string]T, id string, value T, kind string, path string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%s id is required in %s", kind, path)
	}
	if _, exists := values[id]; exists {
		return fmt.Errorf("duplicate %s %q", kind, id)
	}
	values[id] = value
	return nil
}

func splitBuffer(buffer string) []string {
	trimmed := strings.TrimSuffix(buffer, "\n")
	if trimmed == "" {
		return []string{""}
	}
	return strings.Split(trimmed, "\n")
}

func optionalBuffer(buffer string) []string {
	if buffer == "" {
		return nil
	}
	return splitBuffer(buffer)
}

func copyCursor(cursor *CursorSpec) *CursorSpec {
	if cursor == nil {
		return nil
	}
	next := *cursor
	return &next
}

func copyHints(hints []HintSpec) []HintSpec {
	if hints == nil {
		return nil
	}
	next := make([]HintSpec, len(hints))
	copy(next, hints)
	return next
}

func validStatus(status Status) bool {
	switch status {
	case StatusDraft, StatusApproved, StatusImplemented, StatusRetired:
		return true
	default:
		return false
	}
}

func validEngineSupport(value EngineSupport) bool {
	switch value {
	case EngineSupportImplemented, EngineSupportPlanned, EngineSupportUnsupported:
		return true
	default:
		return false
	}
}

func isApprovedLike(status Status) bool {
	return status == StatusApproved || status == StatusImplemented
}

func isYAML(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".yaml" || ext == ".yml"
}
