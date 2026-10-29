package playable

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

var firstRecoveryExercises = []string{"normal-motion-basic-001", "word-motion-basic-001", "search-basic-001", "text-object-inner-word-003", "incident-hotfix-001", "incident-hotfix-003"}

func (m Model) routeMenuView() string {
	if m.chapterMenu {
		return m.chapterView()
	}
	return "ADVIMTURE | 출격 경로\n\n[1] 첫 복구     조작 4개 → 원인 추적 → 상태 복구\n[2] 전체 학습   미완료 학습부터 이어가기\n[3] 응용 도전   복구 범위 판별\n[4] 복습        완료한 문항 다시 연습\n[5] 이어하기   현재 경로\n[6] 학습 목록   에피소드 선택\n\n[esc] 돌아가기   [q] 종료\n" + m.routeNotice
}

func (m *Model) selectRoute(input string) tea.Cmd {
	if m.chapterMenu {
		starts := m.chapterStarts()
		switch input {
		case "j", "down":
			if m.chapterIndex+1 < len(starts) {
				m.chapterIndex++
			}
		case "k", "up":
			if m.chapterIndex > 0 {
				m.chapterIndex--
			}
		case "esc":
			m.chapterMenu = false
		case "q":
			return tea.Quit
		case "enter":
			if len(starts) > 0 {
				entry := m.entries[starts[m.chapterIndex]]
				ids := []string{}
				for _, e := range m.entries {
					if e.PlaylistID == entry.PlaylistID {
						ids = append(ids, e.ExerciseID)
					}
				}
				m.startRoute(entry.PlaylistTitle, ids)
				m.chapterMenu = false
			}
		}
		return nil
	}
	switch input {
	case "q":
		return tea.Quit
	case "esc", "5":
		m.routeMenu = false
	case "6":
		m.chapterMenu = true
		m.chapterIndex = 0
	case "1":
		m.startRoute("첫 복구", firstRecoveryExercises)
	case "2":
		m.route = nil
		m.routeName = ""
		m.reviewMode = false
		m.routeComplete = false
		m.routeMenu = false
		m.jumpToEntry(firstIncompleteIndex(m.entries, m.progress))
	case "3":
		m.startRoute("응용 도전", []string{"command-choice-scope-001", "command-choice-inline-target-001", "command-choice-line-reuse-001"})
	case "4":
		ids := []string{}
		for _, entry := range m.entries {
			if m.progress.Missions[entry.ExerciseID].Completed {
				ids = append(ids, entry.ExerciseID)
			}
		}
		if len(ids) == 0 {
			m.routeNotice = "아직 완료한 문항이 없습니다. 첫 복구부터 시작하세요."
			break
		}
		m.startRoute("복습", ids)
		m.reviewMode = true
		// Focus practice on remaining review candidates; otherwise allow a replay.
		if index, ok := m.reviewDispatchTargetIndex(); ok {
			m.jumpToEntry(index)
		}
	}
	_ = m.writeE2EState()
	return nil
}

func (m Model) chapterStarts() []int {
	starts := []int{}
	for i, e := range m.entries {
		if i == 0 || e.PlaylistID != m.entries[i-1].PlaylistID {
			starts = append(starts, i)
		}
	}
	return starts
}

func (m Model) chapterView() string {
	starts := m.chapterStarts()
	var b strings.Builder
	b.WriteString("ADVIMTURE | 학습 목록\n\n")
	start := m.chapterIndex / 8 * 8
	for i := start; i < len(starts) && i < start+8; i++ {
		marker := "  "
		if i == m.chapterIndex {
			marker = "> "
		}
		b.WriteString(marker + m.entries[starts[i]].PlaylistTitle + "\n")
	}
	b.WriteString("\n[j/↓] 다음  [k/↑] 이전  [enter] 시작  [esc] 경로 선택\n")
	return b.String()
}

func (m *Model) startRoute(name string, ids []string) {
	indices := []int{}
	for _, id := range ids {
		for i, entry := range m.entries {
			if entry.ExerciseID == id {
				indices = append(indices, i)
				break
			}
		}
	}
	if len(indices) == 0 {
		m.routeNotice = "이 경로에 실행 가능한 문항이 없습니다."
		return
	}
	m.route = indices
	m.routePos = 0
	m.routeName = name
	m.routeComplete = false
	m.reviewMode = false
	m.routeMenu = false
	m.routeNotice = ""
	m.jumpToEntry(indices[0])
}

func (m Model) routeFooter() string {
	if m.run == nil {
		return ""
	}
	if m.inputMode() != "normal" && m.run.State().Status == "running" {
		return ""
	}
	text := "경로  [m] 선택 · 첫 복구 / 전체 학습 / 응용 / 복습"
	if m.routeName != "" {
		text = fmt.Sprintf("경로  [m] 선택 · %s %d/%d", m.routeName, m.routePos+1, len(m.route))
	}
	if m.routeComplete {
		text = "경로 완료 · [m] 다음 출격 선택 · [q] 종료"
	}
	return strings.TrimSpace(text)
}
