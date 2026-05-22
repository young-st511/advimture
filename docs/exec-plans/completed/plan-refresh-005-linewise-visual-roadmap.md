# PLAN-REFRESH-005 — Linewise Visual Roadmap

Slice-ID: PLAN-REFRESH-005
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/plan-refresh-005-linewise-visual-roadmap.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

`Post-Visual Applied Mastery and Hardening` 완료 후 다음 중기 플랜을 linewise visual 중심으로 고정한다.

## 영향 도메인

- Gameplay: 새 command cluster와 tutorial/playpack 순서를 다룬다. `Vim command -> Exercise -> Scenario` 순서를 유지한다.
- Verification: visual selection은 화면 텍스트만 보지 않고 app_state selection, key trace, progress까지 검증한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: VISUAL-LINE-GAP-001은 다음 visual 후보를 linewise `V` + row motion + `d/y`로 좁혔다.
- `docs/verification/selection-app-state-contract.md`: 후속 linewise visual은 `kind: linewise`와 normalized full-line selection 범위를 사용한다.

## 중기 플랜

| 순서 | ID | 목표 | 검증 |
|------|----|------|------|
| 1 | PLAN-REFRESH-005 | linewise visual 중심 중기 플랜 고정 | `git diff --check` |
| 2 | VISUAL-LINE-001 | linewise `V` scope와 수용 기준 승인 | docs review, `git diff --check` |
| 3 | VIM-029 | linewise `V` foundation과 `d/y` engine/runtime/TUI 구현 | vimengine/tuiadapter/runtime/playable tests, `go test ./...` |
| 4 | PLAYPACK-011 | 3문항 이하 linewise visual tutorial content/E2E | content replay, focused E2E, `make e2e-playable` |
| 5 | INCIDENT-RUN-004 | linewise visual을 적용하는 config block 복구 incident | content replay, incident E2E |
| 6 | COMMAND-CHOICE-001 | 배운 command 중 적절한 도구를 고르는 mixed drill 설계 | docs review, content draft |
| 7 | PLATFORM-REVIEW-002 | 저장 포맷 변경 없는 review/daily 동기 강화 | playable tests, focused E2E |

## 품질 게이트

- VIM-029는 테스트 Red를 먼저 작성한다.
- linewise `d/y`는 unnamed linewise register, cursor landing, selection clear, undo behavior를 unit test로 고정한다.
- TUI/E2E는 `selection.kind: linewise`와 full-line normalized range를 app_state로 검증한다.
- progress 저장 포맷과 `go.mod`/`go.sum`은 변경하지 않는다.

## 이번 루프에서 의도적으로 미루는 것

- multi-line charwise visual operator
- visual block `<C-v>`
- visual `c`, `>`, `<`
- count/register prefix
- dot repeat 연계
- progress schema v2

## 검증 결과

- planned: `git diff --check`
