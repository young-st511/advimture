package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewReportsProgressLoadError(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	progressDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(progressDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(progressDir, "progress.json"), []byte("bad"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(progressDir, "progress.json.bak"), []byte("also bad"), 0o644); err != nil {
		t.Fatal(err)
	}

	model := NewWithOptions(Options{
		ContentFS: os.DirFS(filepath.Join("..", "..")),
	})

	view := model.View()
	if !strings.Contains(view, "Playable error:") {
		t.Fatalf("view = %q, want playable error", view)
	}
	if !strings.Contains(view, "진행도 로드 실패") {
		t.Fatalf("view = %q, want progress load failure", view)
	}
}
