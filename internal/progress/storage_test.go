package progress

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")

	p := NewProgress()
	p.PlayerName = "tester"
	p.CompleteTutorial("1-1", 10.5, 25)

	if err := SaveToPath(p, path); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if loaded.PlayerName != "tester" {
		t.Errorf("expected 'tester', got '%s'", loaded.PlayerName)
	}
	tp := loaded.Tutorials["1-1"]
	if !tp.Completed {
		t.Error("tutorial 1-1 should be completed")
	}
	if tp.BestTime != 10.5 {
		t.Errorf("expected best time 10.5, got %f", tp.BestTime)
	}
}

func TestLoad_FileNotExist(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nonexistent.json")

	p, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("should not error, got %v", err)
	}
	if p.CompletedTutorialCount() != 0 {
		t.Error("new progress should have 0 completed tutorials")
	}
}

func TestLoad_CorruptedFallbackToBackup(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")
	backupPath := path + ".bak"

	// Write corrupt main file
	os.WriteFile(path, []byte("{invalid json"), 0o644)

	// Write valid backup
	p := NewProgress()
	p.PlayerName = "from-backup"
	SaveToPath(p, backupPath)
	// SaveToPath creates .bak.bak, but backup itself is valid
	// Rewrite backup directly
	os.WriteFile(backupPath, []byte(`{"player_name":"from-backup","created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z","tutorials":{},"missions":{},"total_keystrokes":0}`), 0o644)

	loaded, _ := LoadFromPath(path)
	if loaded.PlayerName != "from-backup" {
		t.Errorf("expected 'from-backup', got '%s'", loaded.PlayerName)
	}
}

func TestLoad_BothCorrupted(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")

	os.WriteFile(path, []byte("bad"), 0o644)
	os.WriteFile(path+".bak", []byte("also bad"), 0o644)

	p, _ := LoadFromPath(path)
	// Should return fresh progress
	if p.CompletedTutorialCount() != 0 {
		t.Error("corrupted files should yield fresh progress")
	}
}

func TestSave_CreatesBackup(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "progress.json")

	// First save
	p := NewProgress()
	p.PlayerName = "v1"
	SaveToPath(p, path)

	// Second save
	p.PlayerName = "v2"
	SaveToPath(p, path)

	// Check backup exists
	backupPath := path + ".bak"
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		t.Error("backup file should exist after second save")
	}
}

func TestRank(t *testing.T) {
	tests := []struct {
		tutorials int
		missions  int
		expected  Rank
	}{
		{0, 0, RankIntern},
		{2, 0, RankIntern},
		{3, 0, RankJunior},
		{6, 0, RankSenior},
		{8, 0, RankStaff},
		{10, 0, RankPrincipal},
		{10, 10, RankVimMaster},
	}

	for _, tt := range tests {
		got := CalculateRank(tt.tutorials, tt.missions)
		if got != tt.expected {
			t.Errorf("CalculateRank(%d, %d) = %s, want %s",
				tt.tutorials, tt.missions, got, tt.expected)
		}
	}
}

func TestProgress_CompleteTutorial(t *testing.T) {
	p := NewProgress()

	p.CompleteTutorial("1-1", 15.0, 30)
	if p.CompletedTutorialCount() != 1 {
		t.Errorf("expected 1 completed, got %d", p.CompletedTutorialCount())
	}

	// Improve time
	p.CompleteTutorial("1-1", 10.0, 20)
	tp := p.Tutorials["1-1"]
	if tp.BestTime != 10.0 {
		t.Errorf("expected best time 10.0, got %f", tp.BestTime)
	}
	if tp.Keystrokes != 20 {
		t.Errorf("expected best keystrokes 20, got %d", tp.Keystrokes)
	}
	if tp.Attempts != 2 {
		t.Errorf("expected 2 attempts, got %d", tp.Attempts)
	}
}

func TestProgress_CurrentRank(t *testing.T) {
	p := NewProgress()
	if p.CurrentRank() != RankIntern {
		t.Errorf("new player should be Intern, got %s", p.CurrentRank())
	}

	for i := 1; i <= 6; i++ {
		p.CompleteTutorial(string(rune('0'+i)), 10, 10)
	}
	if p.CurrentRank() != RankSenior {
		t.Errorf("6 tutorials should be Senior, got %s", p.CurrentRank())
	}
}
