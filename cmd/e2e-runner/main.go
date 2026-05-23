package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/progress"
	"gopkg.in/yaml.v3"
)

type scenario struct {
	ID       string          `yaml:"id"`
	Command  []string        `yaml:"command"`
	Timeout  int             `yaml:"timeout_ms"`
	Terminal terminalConfig  `yaml:"terminal"`
	Setup    setupConfig     `yaml:"setup"`
	Steps    []step          `yaml:"steps"`
	Assert   assertionConfig `yaml:"assert"`
	Evidence evidenceConfig  `yaml:"evidence"`
}

type terminalConfig struct {
	Width  uint16 `yaml:"width"`
	Height uint16 `yaml:"height"`
}

type setupConfig struct {
	Home            string `yaml:"home"`
	AllowUnsafeHome bool   `yaml:"allow_unsafe_home"`
	ProgressFile    string `yaml:"progress_file"`
	CompleteBefore  string `yaml:"complete_before"`
	ContentRoot     string `yaml:"content_root"`
}

type step struct {
	Send               string `yaml:"send"`
	WaitMs             int    `yaml:"wait_ms"`
	WaitScreenContains string `yaml:"wait_screen_contains"`
}

type assertionConfig struct {
	ScreenContains       []string          `yaml:"screen_contains"`
	ExitCode             *int              `yaml:"exit_code"`
	ProgressFileExists   *bool             `yaml:"progress_file_exists"`
	ProgressFileContains []string          `yaml:"progress_file_contains"`
	KeyTrace             []string          `yaml:"key_trace"`
	AppState             appStateAssertion `yaml:"app_state"`
}

type appStateAssertion struct {
	Path      string              `yaml:"path"`
	Buffer    []string            `yaml:"buffer"`
	Cursor    *cursorAssertion    `yaml:"cursor"`
	Mode      string              `yaml:"mode"`
	Command   string              `yaml:"command"`
	Status    string              `yaml:"status"`
	Score     *scoreAssertion     `yaml:"score"`
	Progress  *progressAssertion  `yaml:"progress"`
	Review    *reviewAssertion    `yaml:"review"`
	UI        *uiAssertion        `yaml:"ui"`
	Selection *selectionAssertion `yaml:"selection"`
	Contains  map[string]string   `yaml:"contains"`
}

type cursorAssertion struct {
	Row *int `yaml:"row"`
	Col *int `yaml:"col"`
}

type scoreAssertion struct {
	Grade  string `yaml:"grade"`
	Passed *bool  `yaml:"passed"`
}

type progressAssertion struct {
	MissionID string `yaml:"mission_id"`
	Completed *bool  `yaml:"completed"`
}

type reviewAssertion struct {
	QueueCount        *int   `yaml:"queue_count"`
	PrimaryExerciseID string `yaml:"primary_exercise_id"`
	PrimaryReason     string `yaml:"primary_reason"`
	DailyRoute        string `yaml:"daily_route"`
}

type uiAssertion struct {
	FocusPanel *focusPanelAssertion `yaml:"focus_panel"`
}

type focusPanelAssertion struct {
	Kind  string   `yaml:"kind"`
	Title string   `yaml:"title"`
	Lines []string `yaml:"lines"`
}

type selectionAssertion struct {
	Active *bool            `yaml:"active"`
	Kind   string           `yaml:"kind"`
	Anchor *cursorAssertion `yaml:"anchor"`
	Head   *cursorAssertion `yaml:"head"`
	Start  *cursorAssertion `yaml:"start"`
	End    *cursorAssertion `yaml:"end"`
}

type appStateSummary struct {
	Buffer    []string           `json:"buffer"`
	Cursor    appStateCursor     `json:"cursor"`
	Mode      string             `json:"mode"`
	Command   string             `json:"command"`
	Status    string             `json:"status"`
	Score     appStateScore      `json:"score"`
	Progress  appStateProgress   `json:"progress"`
	Review    appStateReview     `json:"review"`
	UI        appStateUI         `json:"ui"`
	Selection *appStateSelection `json:"selection"`
	Extra     map[string]any     `json:"-"`
}

type appStateCursor struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type appStateScore struct {
	Grade  string `json:"grade"`
	Passed bool   `json:"passed"`
}

type appStateProgress struct {
	MissionID string `json:"mission_id"`
	Completed bool   `json:"completed"`
}

type appStateReview struct {
	QueueCount        int    `json:"queue_count"`
	PrimaryExerciseID string `json:"primary_exercise_id"`
	PrimaryReason     string `json:"primary_reason"`
	DailyRoute        string `json:"daily_route"`
}

type appStateUI struct {
	FocusPanel appStateFocusPanel `json:"focus_panel"`
}

type appStateFocusPanel struct {
	Kind  string   `json:"kind"`
	Title string   `json:"title"`
	Lines []string `json:"lines"`
}

type appStateSelection struct {
	Active bool           `json:"active"`
	Kind   string         `json:"kind"`
	Anchor appStateCursor `json:"anchor"`
	Head   appStateCursor `json:"head"`
	Start  appStateCursor `json:"start"`
	End    appStateCursor `json:"end"`
}

type evidenceConfig struct {
	SaveRawANSI        bool `yaml:"save_raw_ansi"`
	SaveCleanScreen    bool `yaml:"save_clean_screen"`
	SaveScreenTimeline bool `yaml:"save_screen_timeline"`
	SaveKeyTrace       bool `yaml:"save_key_trace"`
	SaveSummary        bool `yaml:"save_summary"`
	SaveAppState       bool `yaml:"save_app_state"`
	SaveProgress       bool `yaml:"save_progress"`
}

type runResult struct {
	raw                []byte
	clean              string
	exitCode           int
	homeDir            string
	trace              []string
	progressFileExists bool
	progressRaw        []byte
	appStateExists     bool
	appStateRaw        []byte
}

type summaryEvidence struct {
	ScenarioID         string   `json:"scenario_id"`
	Passed             bool     `json:"passed"`
	Error              string   `json:"error,omitempty"`
	ExitCode           int      `json:"exit_code"`
	HomeDir            string   `json:"home_dir"`
	KeyTrace           []string `json:"key_trace"`
	ScreenBytes        int      `json:"screen_bytes"`
	ScreenTimeline     bool     `json:"screen_timeline_evidence"`
	ProgressFileExists bool     `json:"progress_file_exists"`
	ProgressEvidence   bool     `json:"progress_evidence"`
	AppStatePath       string   `json:"app_state_path,omitempty"`
	AppStateExists     bool     `json:"app_state_exists"`
	AppStateEvidence   bool     `json:"app_state_evidence"`
}

func main() {
	scenarioPath := flag.String("scenario", "", "path to an E2E scenario YAML file")
	artifactRoot := flag.String("artifacts", "artifacts/e2e", "directory for E2E evidence")
	flag.Parse()

	if *scenarioPath == "" {
		fmt.Fprintln(os.Stderr, "missing required --scenario")
		os.Exit(2)
	}

	sc, err := loadScenario(*scenarioPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scenario load failed: %v\n", err)
		os.Exit(2)
	}

	result, err := runScenario(sc)
	if writeErr := writeEvidence(*artifactRoot, sc, result, err); writeErr != nil {
		fmt.Fprintf(os.Stderr, "evidence write failed: %v\n", writeErr)
		if err == nil {
			err = writeErr
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "scenario failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("scenario passed: %s\n", sc.ID)
}

func loadScenario(path string) (scenario, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return scenario{}, err
	}

	var sc scenario
	if err := yaml.Unmarshal(raw, &sc); err != nil {
		return scenario{}, err
	}
	if sc.ID == "" {
		return scenario{}, errors.New("id is required")
	}
	if len(sc.Command) == 0 {
		sc.Command = []string{"go", "run", "."}
	}
	if sc.Timeout <= 0 {
		sc.Timeout = 5000
	}
	if sc.Terminal.Width == 0 {
		sc.Terminal.Width = 100
	}
	if sc.Terminal.Height == 0 {
		sc.Terminal.Height = 30
	}
	if !sc.Evidence.SaveRawANSI && !sc.Evidence.SaveCleanScreen && !sc.Evidence.SaveKeyTrace && !sc.Evidence.SaveSummary && !sc.Evidence.SaveAppState && !sc.Evidence.SaveProgress {
		sc.Evidence.SaveSummary = true
	}
	return sc, nil
}

func runScenario(sc scenario) (runResult, error) {
	homeDir, cleanup, err := setupHome(sc)
	if err != nil {
		return runResult{}, err
	}
	defer cleanup()

	cmd := exec.Command(sc.Command[0], sc.Command[1:]...)
	cmd.Env = append(os.Environ(),
		"HOME="+homeDir,
		"TERM=xterm-256color",
		"ADVIMTURE_E2E=1",
	)
	cmd.Env = append(cmd.Env, goToolEnv(homeDir)...)

	ptmx, err := pty.StartWithSize(cmd, &pty.Winsize{
		Rows: sc.Terminal.Height,
		Cols: sc.Terminal.Width,
	})
	if err != nil {
		return runResult{homeDir: homeDir}, err
	}
	defer ptmx.Close()

	var mu sync.Mutex
	var raw bytes.Buffer
	doneReading := make(chan struct{})
	go func() {
		defer close(doneReading)
		_, _ = io.Copy(writerFunc(func(p []byte) (int, error) {
			mu.Lock()
			defer mu.Unlock()
			n, err := raw.Write(p)
			respondTerminalQueries(ptmx, p)
			return n, err
		}), ptmx)
	}()

	var trace []string
	deadline := time.Now().Add(time.Duration(sc.Timeout) * time.Millisecond)
	waitOffset := 0

	for _, st := range sc.Steps {
		if st.WaitScreenContains != "" {
			nextOffset, err := waitForScreen(&mu, &raw, st.WaitScreenContains, deadline, waitOffset)
			if err != nil {
				terminate(cmd)
				return collectResult(sc, &mu, &raw, homeDir, trace, exitCode(cmd)), err
			}
			waitOffset = nextOffset
		}
		if st.WaitMs > 0 {
			time.Sleep(time.Duration(st.WaitMs) * time.Millisecond)
		}
		if st.Send != "" {
			trace = append(trace, st.Send)
			if _, err := ptmx.Write([]byte(keyBytes(st.Send))); err != nil {
				terminate(cmd)
				return collectResult(sc, &mu, &raw, homeDir, trace, exitCode(cmd)), err
			}
		}
	}

	waitErr := waitWithTimeout(cmd, time.Until(deadline))
	if waitErr != nil && sc.Assert.ExitCode != nil {
		terminate(cmd)
		return collectResult(sc, &mu, &raw, homeDir, trace, exitCode(cmd)), waitErr
	}
	if waitErr != nil {
		terminate(cmd)
	}

	_ = ptmx.Close()
	<-doneReading

	result := collectResult(sc, &mu, &raw, homeDir, trace, exitCode(cmd))
	if err := assertScenario(sc, result); err != nil {
		return result, err
	}
	return result, nil
}

func setupHome(sc scenario) (string, func(), error) {
	if sc.Setup.Home == "" || sc.Setup.Home == "temp" {
		dir, err := os.MkdirTemp("", "advimture-e2e-home-*")
		if err != nil {
			return "", func() {}, err
		}
		if err := writeProgressFixture(dir, sc.Setup); err != nil {
			_ = os.RemoveAll(dir)
			return "", func() {}, err
		}
		return dir, func() { _ = os.RemoveAll(dir) }, nil
	}
	abs, err := filepath.Abs(sc.Setup.Home)
	if err != nil {
		return "", func() {}, err
	}
	if err := guardHome(abs, sc.Setup.AllowUnsafeHome); err != nil {
		return "", func() {}, err
	}
	if err := writeProgressFixture(abs, sc.Setup); err != nil {
		return "", func() {}, err
	}
	return abs, func() {}, nil
}

func writeProgressFixture(homeDir string, setup setupConfig) error {
	raw, err := progressFixtureJSON(setup)
	if err != nil {
		return err
	}
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	dir := filepath.Join(homeDir, ".advimture")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("progress fixture dir: %w", err)
	}
	if !json.Valid([]byte(raw)) {
		return fmt.Errorf("progress fixture must be valid JSON")
	}
	path := filepath.Join(dir, "progress.json")
	if err := os.WriteFile(path, []byte(raw), 0o644); err != nil {
		return fmt.Errorf("progress fixture write: %w", err)
	}
	return nil
}

func progressFixtureJSON(setup setupConfig) (string, error) {
	hasInline := strings.TrimSpace(setup.ProgressFile) != ""
	hasBuilder := strings.TrimSpace(setup.CompleteBefore) != ""
	if hasInline && hasBuilder {
		return "", fmt.Errorf("setup.progress_file and setup.complete_before cannot be used together")
	}
	if hasInline {
		return setup.ProgressFile, nil
	}
	if !hasBuilder {
		return "", nil
	}
	progressState, err := buildCompletedBeforeProgress(setup.ContentRoot, setup.CompleteBefore)
	if err != nil {
		return "", err
	}
	raw, err := json.MarshalIndent(progressState, "", "  ")
	if err != nil {
		return "", fmt.Errorf("progress fixture marshal: %w", err)
	}
	return string(raw), nil
}

func buildCompletedBeforeProgress(root string, exerciseID string) (*progress.Progress, error) {
	if strings.TrimSpace(exerciseID) == "" {
		return nil, fmt.Errorf("complete_before exercise id is required")
	}
	if root == "" {
		root = "content"
	}
	ids, err := playableExerciseIDs(root)
	if err != nil {
		return nil, err
	}

	progressState := progress.NewProgress()
	found := false
	for _, id := range ids {
		if id == exerciseID {
			found = true
			break
		}
		progressState.Missions[id] = progress.MissionProgress{
			Completed: true,
			BestGrade: "S",
			Attempts:  1,
		}
	}
	if !found {
		return nil, fmt.Errorf("complete_before exercise %q was not found in playable sequence", exerciseID)
	}
	return progressState, nil
}

func playableExerciseIDs(root string) ([]string, error) {
	library, err := content.LoadLibrary(root)
	if err != nil {
		return nil, fmt.Errorf("load content for progress fixture: %w", err)
	}

	var ids []string
	for _, playlist := range library.PlayablePlaylists() {
		for _, beat := range playlist.Beats {
			exercise, ok := library.Exercises[beat.ExerciseID]
			if !ok || !isE2EPlayableExercise(exercise) {
				continue
			}
			scenarioDoc, ok := library.Scenarios[beat.ScenarioID]
			if !ok || !isE2EPlayableScenario(scenarioDoc) {
				continue
			}
			ids = append(ids, beat.ExerciseID)
		}
	}
	return ids, nil
}

func isE2EPlayableExercise(exercise content.ExerciseDocument) bool {
	return (exercise.Status == content.StatusApproved || exercise.Status == content.StatusImplemented) &&
		exercise.EngineSupport == content.EngineSupportImplemented &&
		exercise.ReplayStatus == content.ReplayStatusPass
}

func isE2EPlayableScenario(scenarioDoc content.ScenarioDocument) bool {
	return (scenarioDoc.Status == content.StatusApproved || scenarioDoc.Status == content.StatusImplemented) &&
		scenarioDoc.EngineSupport == content.EngineSupportImplemented
}

func guardHome(path string, allowUnsafe bool) error {
	if allowUnsafe {
		return nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	absHome, err := filepath.Abs(home)
	if err != nil {
		return err
	}
	cleanPath := filepath.Clean(path)
	cleanHome := filepath.Clean(absHome)
	if cleanPath == cleanHome {
		return fmt.Errorf("unsafe home %q: use setup.home: temp or set allow_unsafe_home explicitly", path)
	}
	progressPath := filepath.Join(cleanPath, ".advimture", "progress.json")
	if _, err := os.Stat(progressPath); err == nil {
		return fmt.Errorf("unsafe home %q: existing progress file would be visible to E2E", path)
	}
	return nil
}

func waitForScreen(mu *sync.Mutex, raw *bytes.Buffer, want string, deadline time.Time, offset int) (int, error) {
	for time.Now().Before(deadline) {
		mu.Lock()
		snapshot := append([]byte(nil), raw.Bytes()...)
		mu.Unlock()
		if offset < 0 || offset > len(snapshot) {
			offset = 0
		}
		clean := cleanTerminal(snapshot[offset:])
		if strings.Contains(clean, want) {
			return len(snapshot), nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return offset, fmt.Errorf("timed out waiting for screen to contain %q", want)
}

func waitWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = time.Millisecond
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("process did not exit before timeout")
	}
}

func assertScenario(sc scenario, result runResult) error {
	for _, want := range sc.Assert.ScreenContains {
		if !strings.Contains(result.clean, want) {
			return fmt.Errorf("screen does not contain %q", want)
		}
	}
	if sc.Assert.ExitCode != nil && result.exitCode != *sc.Assert.ExitCode {
		return fmt.Errorf("exit code: got %d, want %d", result.exitCode, *sc.Assert.ExitCode)
	}
	if sc.Assert.ProgressFileExists != nil {
		_, err := os.Stat(progressPath(result.homeDir))
		exists := err == nil
		if exists != *sc.Assert.ProgressFileExists {
			return fmt.Errorf("progress file exists: got %v, want %v", exists, *sc.Assert.ProgressFileExists)
		}
	}
	if len(sc.Assert.ProgressFileContains) > 0 {
		raw, err := os.ReadFile(progressPath(result.homeDir))
		if err != nil {
			return fmt.Errorf("progress file read failed: %w", err)
		}
		text := string(raw)
		for _, want := range sc.Assert.ProgressFileContains {
			if !strings.Contains(text, want) {
				return fmt.Errorf("progress file does not contain %q", want)
			}
		}
	}
	if len(sc.Assert.KeyTrace) > 0 && !sameStrings(result.trace, sc.Assert.KeyTrace) {
		return fmt.Errorf("key trace: got %v, want %v", result.trace, sc.Assert.KeyTrace)
	}
	if wantsAppStateAssertion(sc.Assert.AppState) {
		state, raw, err := loadAppStateSummary(result.homeDir, sc.Assert.AppState.Path)
		if err != nil {
			return err
		}
		if err := assertAppState(sc.Assert.AppState, state, raw); err != nil {
			return err
		}
	}
	return nil
}

func collectResult(sc scenario, mu *sync.Mutex, raw *bytes.Buffer, homeDir string, trace []string, code int) runResult {
	mu.Lock()
	defer mu.Unlock()
	rawBytes := append([]byte(nil), raw.Bytes()...)
	return runResult{
		raw:                rawBytes,
		clean:              cleanTerminal(rawBytes),
		exitCode:           code,
		homeDir:            homeDir,
		trace:              trace,
		progressFileExists: progressFileExists(homeDir),
		progressRaw:        readOptionalFile(progressPath(homeDir)),
		appStateExists:     appStateExists(homeDir, sc.Assert.AppState.Path),
		appStateRaw:        readOptionalFile(appStatePath(homeDir, sc.Assert.AppState.Path)),
	}
}

func writeEvidence(root string, sc scenario, result runResult, runErr error) error {
	if root == "" || sc.ID == "" {
		return nil
	}
	dir := filepath.Join(root, sc.ID)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if sc.Evidence.SaveRawANSI {
		if err := os.WriteFile(filepath.Join(dir, "raw.log"), result.raw, 0o644); err != nil {
			return err
		}
	}
	if sc.Evidence.SaveCleanScreen {
		if err := os.WriteFile(filepath.Join(dir, "screen.txt"), []byte(result.clean), 0o644); err != nil {
			return err
		}
	}
	if sc.Evidence.SaveScreenTimeline {
		if err := os.WriteFile(filepath.Join(dir, "screen_timeline.txt"), []byte(result.clean), 0o644); err != nil {
			return err
		}
	}
	if sc.Evidence.SaveKeyTrace {
		if err := os.WriteFile(filepath.Join(dir, "key_trace.txt"), []byte(strings.Join(result.trace, "\n")), 0o644); err != nil {
			return err
		}
	}
	if sc.Evidence.SaveAppState {
		if len(result.appStateRaw) == 0 {
			return fmt.Errorf("app state evidence requested but app state was not captured")
		}
		if err := os.WriteFile(filepath.Join(dir, "app_state.json"), result.appStateRaw, 0o644); err != nil {
			return err
		}
	}
	if sc.Evidence.SaveProgress {
		if len(result.progressRaw) == 0 {
			return fmt.Errorf("progress evidence requested but progress was not captured")
		}
		if err := os.WriteFile(filepath.Join(dir, "progress.json"), result.progressRaw, 0o644); err != nil {
			return err
		}
	}
	summary := buildSummary(sc, result, runErr)
	raw, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "summary.json"), raw, 0o644); err != nil {
		return err
	}
	return nil
}

func buildSummary(sc scenario, result runResult, runErr error) summaryEvidence {
	summary := summaryEvidence{
		ScenarioID:         sc.ID,
		Passed:             runErr == nil,
		ExitCode:           result.exitCode,
		HomeDir:            result.homeDir,
		KeyTrace:           append([]string(nil), result.trace...),
		ScreenBytes:        len(result.clean),
		ScreenTimeline:     sc.Evidence.SaveScreenTimeline,
		ProgressFileExists: result.progressFileExists || progressFileExists(result.homeDir),
		ProgressEvidence:   len(result.progressRaw) > 0,
		AppStatePath:       appStatePath(result.homeDir, sc.Assert.AppState.Path),
		AppStateExists:     result.appStateExists || appStateExists(result.homeDir, sc.Assert.AppState.Path),
		AppStateEvidence:   len(result.appStateRaw) > 0,
	}
	if runErr != nil {
		summary.Error = runErr.Error()
	}
	return summary
}

func progressFileExists(homeDir string) bool {
	if homeDir == "" {
		return false
	}
	_, err := os.Stat(progressPath(homeDir))
	return err == nil
}

func progressPath(homeDir string) string {
	if homeDir == "" {
		return ""
	}
	return filepath.Join(homeDir, ".advimture", "progress.json")
}

func readOptionalFile(path string) []byte {
	if path == "" {
		return nil
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return raw
}

func wantsAppStateAssertion(assertion appStateAssertion) bool {
	return assertion.Path != "" ||
		len(assertion.Buffer) > 0 ||
		assertion.Cursor != nil ||
		assertion.Mode != "" ||
		assertion.Command != "" ||
		assertion.Status != "" ||
		assertion.Score != nil ||
		assertion.Progress != nil ||
		assertion.Review != nil ||
		assertion.UI != nil ||
		assertion.Selection != nil ||
		len(assertion.Contains) > 0
}

func loadAppStateSummary(homeDir string, configuredPath string) (appStateSummary, []byte, error) {
	path := appStatePath(homeDir, configuredPath)
	raw, err := os.ReadFile(path)
	if err != nil {
		return appStateSummary{}, nil, fmt.Errorf("app state summary read failed: %w", err)
	}
	var state appStateSummary
	if err := json.Unmarshal(raw, &state); err != nil {
		return appStateSummary{}, nil, fmt.Errorf("app state summary parse failed: %w", err)
	}
	return state, raw, nil
}

func appStatePath(homeDir string, configuredPath string) string {
	if configuredPath != "" {
		if filepath.IsAbs(configuredPath) {
			return configuredPath
		}
		return filepath.Join(homeDir, configuredPath)
	}
	if homeDir == "" {
		return ""
	}
	return filepath.Join(homeDir, ".advimture", "e2e_state.json")
}

func appStateExists(homeDir string, configuredPath string) bool {
	path := appStatePath(homeDir, configuredPath)
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

func assertAppState(assertion appStateAssertion, state appStateSummary, raw []byte) error {
	if len(assertion.Buffer) > 0 && !sameStrings(state.Buffer, assertion.Buffer) {
		return fmt.Errorf("app state buffer: got %v, want %v", state.Buffer, assertion.Buffer)
	}
	if assertion.Cursor != nil {
		if assertion.Cursor.Row != nil && state.Cursor.Row != *assertion.Cursor.Row {
			return fmt.Errorf("app state cursor row: got %d, want %d", state.Cursor.Row, *assertion.Cursor.Row)
		}
		if assertion.Cursor.Col != nil && state.Cursor.Col != *assertion.Cursor.Col {
			return fmt.Errorf("app state cursor col: got %d, want %d", state.Cursor.Col, *assertion.Cursor.Col)
		}
	}
	if assertion.Mode != "" && state.Mode != assertion.Mode {
		return fmt.Errorf("app state mode: got %q, want %q", state.Mode, assertion.Mode)
	}
	if assertion.Command != "" && state.Command != assertion.Command {
		return fmt.Errorf("app state command: got %q, want %q", state.Command, assertion.Command)
	}
	if assertion.Status != "" && state.Status != assertion.Status {
		return fmt.Errorf("app state status: got %q, want %q", state.Status, assertion.Status)
	}
	if assertion.Score != nil {
		if assertion.Score.Grade != "" && state.Score.Grade != assertion.Score.Grade {
			return fmt.Errorf("app state score grade: got %q, want %q", state.Score.Grade, assertion.Score.Grade)
		}
		if assertion.Score.Passed != nil && state.Score.Passed != *assertion.Score.Passed {
			return fmt.Errorf("app state score passed: got %v, want %v", state.Score.Passed, *assertion.Score.Passed)
		}
	}
	if assertion.Progress != nil {
		if assertion.Progress.MissionID != "" && state.Progress.MissionID != assertion.Progress.MissionID {
			return fmt.Errorf("app state progress mission_id: got %q, want %q", state.Progress.MissionID, assertion.Progress.MissionID)
		}
		if assertion.Progress.Completed != nil && state.Progress.Completed != *assertion.Progress.Completed {
			return fmt.Errorf("app state progress completed: got %v, want %v", state.Progress.Completed, *assertion.Progress.Completed)
		}
	}
	if assertion.Review != nil {
		if err := assertReview(*assertion.Review, state.Review); err != nil {
			return err
		}
	}
	if assertion.UI != nil {
		if err := assertUI(*assertion.UI, state.UI); err != nil {
			return err
		}
	}
	if assertion.Selection != nil {
		if state.Selection == nil {
			return fmt.Errorf("app state selection: got nil, want assertion")
		}
		if err := assertSelection(*assertion.Selection, *state.Selection); err != nil {
			return err
		}
	}
	text := string(raw)
	for key, want := range assertion.Contains {
		if !strings.Contains(text, want) {
			return fmt.Errorf("app state contains %q: missing %q", key, want)
		}
	}
	return nil
}

func assertUI(assertion uiAssertion, state appStateUI) error {
	if assertion.FocusPanel != nil {
		if err := assertFocusPanel(*assertion.FocusPanel, state.FocusPanel); err != nil {
			return err
		}
	}
	return nil
}

func assertFocusPanel(assertion focusPanelAssertion, state appStateFocusPanel) error {
	if assertion.Kind != "" && state.Kind != assertion.Kind {
		return fmt.Errorf("app state ui focus_panel kind: got %q, want %q", state.Kind, assertion.Kind)
	}
	if assertion.Title != "" && state.Title != assertion.Title {
		return fmt.Errorf("app state ui focus_panel title: got %q, want %q", state.Title, assertion.Title)
	}
	if len(assertion.Lines) > 0 && !sameStrings(state.Lines, assertion.Lines) {
		return fmt.Errorf("app state ui focus_panel lines: got %v, want %v", state.Lines, assertion.Lines)
	}
	return nil
}

func assertReview(assertion reviewAssertion, state appStateReview) error {
	if assertion.QueueCount != nil && state.QueueCount != *assertion.QueueCount {
		return fmt.Errorf("app state review queue_count: got %d, want %d", state.QueueCount, *assertion.QueueCount)
	}
	if assertion.PrimaryExerciseID != "" && state.PrimaryExerciseID != assertion.PrimaryExerciseID {
		return fmt.Errorf("app state review primary_exercise_id: got %q, want %q", state.PrimaryExerciseID, assertion.PrimaryExerciseID)
	}
	if assertion.PrimaryReason != "" && state.PrimaryReason != assertion.PrimaryReason {
		return fmt.Errorf("app state review primary_reason: got %q, want %q", state.PrimaryReason, assertion.PrimaryReason)
	}
	if assertion.DailyRoute != "" && state.DailyRoute != assertion.DailyRoute {
		return fmt.Errorf("app state review daily_route: got %q, want %q", state.DailyRoute, assertion.DailyRoute)
	}
	return nil
}

func assertSelection(assertion selectionAssertion, state appStateSelection) error {
	if assertion.Active != nil && state.Active != *assertion.Active {
		return fmt.Errorf("app state selection active: got %v, want %v", state.Active, *assertion.Active)
	}
	if assertion.Kind != "" && state.Kind != assertion.Kind {
		return fmt.Errorf("app state selection kind: got %q, want %q", state.Kind, assertion.Kind)
	}
	if assertion.Anchor != nil {
		if err := assertSelectionCursor("anchor", *assertion.Anchor, state.Anchor); err != nil {
			return err
		}
	}
	if assertion.Head != nil {
		if err := assertSelectionCursor("head", *assertion.Head, state.Head); err != nil {
			return err
		}
	}
	if assertion.Start != nil {
		if err := assertSelectionCursor("start", *assertion.Start, state.Start); err != nil {
			return err
		}
	}
	if assertion.End != nil {
		if err := assertSelectionCursor("end", *assertion.End, state.End); err != nil {
			return err
		}
	}
	return nil
}

func assertSelectionCursor(name string, assertion cursorAssertion, state appStateCursor) error {
	if assertion.Row != nil && state.Row != *assertion.Row {
		return fmt.Errorf("app state selection %s row: got %d, want %d", name, state.Row, *assertion.Row)
	}
	if assertion.Col != nil && state.Col != *assertion.Col {
		return fmt.Errorf("app state selection %s col: got %d, want %d", name, state.Col, *assertion.Col)
	}
	return nil
}

func keyBytes(key string) string {
	switch key {
	case "enter":
		return "\r"
	case "esc":
		return "\x1b"
	case "ctrl+c":
		return "\x03"
	case "ctrl+r":
		return "\x12"
	case "tab":
		return "\t"
	case "space":
		return " "
	case "up":
		return "\x1b[A"
	case "down":
		return "\x1b[B"
	case "right":
		return "\x1b[C"
	case "left":
		return "\x1b[D"
	default:
		return key
	}
}

var ansiPattern = regexp.MustCompile(`\x1b\][^\x1b]*(\x1b\\|\x07)|\x1b\[[0-?]*[ -/]*[@-~]|\x1b[=>?]?`)

func cleanTerminal(raw []byte) string {
	s := ansiPattern.ReplaceAllString(string(raw), "")
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	var b strings.Builder
	for _, r := range s {
		if r == '\n' || r == '\t' || r >= 0x20 {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func exitCode(cmd *exec.Cmd) int {
	if cmd.ProcessState == nil {
		return -1
	}
	return cmd.ProcessState.ExitCode()
}

func terminate(cmd *exec.Cmd) {
	if cmd.Process != nil && cmd.ProcessState == nil {
		_ = cmd.Process.Signal(syscall.SIGTERM)
		time.Sleep(100 * time.Millisecond)
		if cmd.ProcessState == nil {
			_ = cmd.Process.Kill()
		}
	}
}

func goToolEnv(homeDir string) []string {
	var env []string
	if os.Getenv("GOCACHE") == "" {
		env = append(env, "GOCACHE="+filepath.Join(homeDir, ".cache", "go-build"))
	}
	for _, key := range []string{"GOPATH", "GOMODCACHE"} {
		if os.Getenv(key) != "" {
			continue
		}
		value, err := lookupGoEnv(key)
		if err != nil {
			continue
		}
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			env = append(env, key+"="+trimmed)
		}
	}
	return env
}

var lookupGoEnv = func(key string) (string, error) {
	value, err := exec.Command("go", "env", key).Output()
	return string(value), err
}

func respondTerminalQueries(w io.Writer, p []byte) {
	if bytes.Contains(p, []byte("\x1b[6n")) {
		_, _ = w.Write([]byte("\x1b[1;1R"))
	}
	if bytes.Contains(p, []byte("\x1b]11;?\x1b\\")) || bytes.Contains(p, []byte("\x1b]11;?\x07")) {
		_, _ = w.Write([]byte("\x1b]11;rgb:0000/0000/0000\x1b\\"))
	}
}

func sameStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

type writerFunc func([]byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) {
	return f(p)
}
