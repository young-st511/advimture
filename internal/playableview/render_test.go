package playableview

import (
	"strings"
	"testing"

	"github.com/young-st511/advimture/internal/tuiadapter"
)

func TestRenderLineShowsVisualSelection(t *testing.T) {
	line := RenderLine("abcd", 0, 0, 3, &tuiadapter.SelectionView{
		Active: true,
		Kind:   "charwise",
		Start:  tuiadapter.CursorView{Row: 0, Col: 1},
		End:    tuiadapter.CursorView{Row: 0, Col: 3},
	})

	if line != "> a{b}{c}[d]" {
		t.Fatalf("RenderLine = %q, want visual selection", line)
	}
}

func TestRenderLineShowsLinewiseVisualSelection(t *testing.T) {
	line := RenderLine("abcd", 0, 0, 2, &tuiadapter.SelectionView{
		Active: true,
		Kind:   "linewise",
		Start:  tuiadapter.CursorView{Row: 0, Col: 0},
		End:    tuiadapter.CursorView{Row: 0, Col: 3},
	})

	if line != "> {a}{b}[c]{d}" {
		t.Fatalf("RenderLine = %q, want linewise visual selection", line)
	}
}

func TestRenderScreenIncludesActionPanel(t *testing.T) {
	view := Render(Screen{
		PlaylistTitle: "Tutorial 0",
		Title:         "커서 위치 맞추기",
		Message:       "목표까지 이동하세요.",
		BufferLines:   []string{"abc"},
		Mode:          "normal",
		Status:        "running",
		CursorCol:     1,
		ExerciseTotal: 4,
		ActionLines:   []string{"ACTION", "?: hint  q: quit"},
	})

	for _, want := range []string{"Tutorial 0", "커서 위치 맞추기", "> a[b]c", "Exercise: 1/4", "ACTION"} {
		if !strings.Contains(view, want) {
			t.Fatalf("Render output = %q, want %q", view, want)
		}
	}
}
