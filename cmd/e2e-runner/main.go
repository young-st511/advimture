package main

import (
	"bytes"
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
	Home string `yaml:"home"`
}

type step struct {
	Send               string `yaml:"send"`
	WaitMs             int    `yaml:"wait_ms"`
	WaitScreenContains string `yaml:"wait_screen_contains"`
}

type assertionConfig struct {
	ScreenContains     []string `yaml:"screen_contains"`
	ExitCode           *int     `yaml:"exit_code"`
	ProgressFileExists *bool    `yaml:"progress_file_exists"`
}

type evidenceConfig struct {
	SaveRawANSI     bool `yaml:"save_raw_ansi"`
	SaveCleanScreen bool `yaml:"save_clean_screen"`
	SaveKeyTrace    bool `yaml:"save_key_trace"`
}

type runResult struct {
	raw      []byte
	clean    string
	exitCode int
	homeDir  string
	trace    []string
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
	if writeErr := writeEvidence(*artifactRoot, sc, result); writeErr != nil {
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
	cmd.Env = append(cmd.Env, goCacheEnv()...)

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

	for _, st := range sc.Steps {
		if st.WaitScreenContains != "" {
			if err := waitForScreen(&mu, &raw, st.WaitScreenContains, deadline); err != nil {
				terminate(cmd)
				return collectResult(&mu, &raw, homeDir, trace, exitCode(cmd)), err
			}
		}
		if st.WaitMs > 0 {
			time.Sleep(time.Duration(st.WaitMs) * time.Millisecond)
		}
		if st.Send != "" {
			trace = append(trace, st.Send)
			if _, err := ptmx.Write([]byte(keyBytes(st.Send))); err != nil {
				terminate(cmd)
				return collectResult(&mu, &raw, homeDir, trace, exitCode(cmd)), err
			}
		}
	}

	waitErr := waitWithTimeout(cmd, time.Until(deadline))
	if waitErr != nil && sc.Assert.ExitCode != nil {
		terminate(cmd)
		return collectResult(&mu, &raw, homeDir, trace, exitCode(cmd)), waitErr
	}
	if waitErr != nil {
		terminate(cmd)
	}

	_ = ptmx.Close()
	<-doneReading

	result := collectResult(&mu, &raw, homeDir, trace, exitCode(cmd))
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
		return dir, func() { _ = os.RemoveAll(dir) }, nil
	}
	abs, err := filepath.Abs(sc.Setup.Home)
	if err != nil {
		return "", func() {}, err
	}
	return abs, func() {}, nil
}

func waitForScreen(mu *sync.Mutex, raw *bytes.Buffer, want string, deadline time.Time) error {
	for time.Now().Before(deadline) {
		mu.Lock()
		clean := cleanTerminal(raw.Bytes())
		mu.Unlock()
		if strings.Contains(clean, want) {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("timed out waiting for screen to contain %q", want)
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
		progressPath := filepath.Join(result.homeDir, ".advimture", "progress.json")
		_, err := os.Stat(progressPath)
		exists := err == nil
		if exists != *sc.Assert.ProgressFileExists {
			return fmt.Errorf("progress file exists: got %v, want %v", exists, *sc.Assert.ProgressFileExists)
		}
	}
	return nil
}

func collectResult(mu *sync.Mutex, raw *bytes.Buffer, homeDir string, trace []string, code int) runResult {
	mu.Lock()
	defer mu.Unlock()
	rawBytes := append([]byte(nil), raw.Bytes()...)
	return runResult{
		raw:      rawBytes,
		clean:    cleanTerminal(rawBytes),
		exitCode: code,
		homeDir:  homeDir,
		trace:    trace,
	}
}

func writeEvidence(root string, sc scenario, result runResult) error {
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
	if sc.Evidence.SaveKeyTrace {
		if err := os.WriteFile(filepath.Join(dir, "key_trace.txt"), []byte(strings.Join(result.trace, "\n")), 0o644); err != nil {
			return err
		}
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
	case "tab":
		return "\t"
	case "space":
		return " "
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

func goCacheEnv() []string {
	var env []string
	for _, key := range []string{"GOCACHE", "GOMODCACHE"} {
		if os.Getenv(key) != "" {
			continue
		}
		value, err := exec.Command("go", "env", key).Output()
		if err != nil {
			continue
		}
		if trimmed := strings.TrimSpace(string(value)); trimmed != "" {
			env = append(env, key+"="+trimmed)
		}
	}
	return env
}

func respondTerminalQueries(w io.Writer, p []byte) {
	if bytes.Contains(p, []byte("\x1b[6n")) {
		_, _ = w.Write([]byte("\x1b[1;1R"))
	}
	if bytes.Contains(p, []byte("\x1b]11;?\x1b\\")) || bytes.Contains(p, []byte("\x1b]11;?\x07")) {
		_, _ = w.Write([]byte("\x1b]11;rgb:0000/0000/0000\x1b\\"))
	}
}

type writerFunc func([]byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) {
	return f(p)
}
