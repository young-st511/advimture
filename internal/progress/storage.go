package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirName      = ".advimture"
	fileName     = "progress.json"
	backupSuffix = ".bak"
)

// progressDir returns the path to ~/.advimture/
func progressDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("[홈 디렉토리 확인 실패] %w", err)
	}
	return filepath.Join(home, dirName), nil
}

// progressPath returns the full path to the progress file.
func progressPath() (string, error) {
	dir, err := progressDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fileName), nil
}

// Load reads progress from disk. Returns new progress if file doesn't exist.
func Load() (*Progress, error) {
	path, err := progressPath()
	if err != nil {
		return nil, err
	}
	return loadFromPath(path)
}

func loadBackup(mainPath string) (*Progress, error) {
	return readProgressFile(mainPath + backupSuffix)
}

// Save writes progress to disk using atomic write (temp file + rename).
func Save(p *Progress) error {
	dir, err := progressDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("[디렉토리 생성 실패] %w", err)
	}

	path := filepath.Join(dir, fileName)

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("[JSON 직렬화 실패] %w", err)
	}

	// Backup existing file
	if _, err := os.Stat(path); err == nil {
		_ = os.Rename(path, path+backupSuffix)
	}

	// Atomic write: temp file then rename
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0o644); err != nil {
		return fmt.Errorf("[임시 파일 쓰기 실패] %w", err)
	}

	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("[파일 이동 실패] %w", err)
	}

	return nil
}

// SaveToPath writes progress to a specific path (for testing).
func SaveToPath(p *Progress, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	// Backup
	if _, err := os.Stat(path); err == nil {
		_ = os.Rename(path, path+backupSuffix)
	}

	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

// LoadFromPath reads progress from a specific path (for testing).
func LoadFromPath(path string) (*Progress, error) {
	return loadFromPath(path)
}

func loadFromPath(path string) (*Progress, error) {
	p, err := readProgressFile(path)
	if err == nil {
		return p, nil
	}
	if os.IsNotExist(err) {
		return NewProgress(), nil
	}

	backup, backupErr := loadBackup(path)
	if backupErr == nil {
		return backup, nil
	}
	return nil, fmt.Errorf("[진행도 로드 실패] %s: %w; backup: %v", path, err, backupErr)
}

func readProgressFile(path string) (*Progress, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	if p.Tutorials == nil {
		p.Tutorials = make(map[string]TutorialProgress)
	}
	if p.Missions == nil {
		p.Missions = make(map[string]MissionProgress)
	}
	return &p, nil
}
