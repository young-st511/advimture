# ExecPlan: Ex command tutorial placement

Slice-ID: CONTENT-003
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/spec.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/content-003-ex-command-pacing.md
- docs/exec-plans/completed/content-003-ex-command-pacing.md
- content/playlists/first-five-minutes.yaml

## 목표

Ex command 학습이 초반 이동/생존 튜토리얼을 압도하지 않도록 중반 고급 튜토리얼로 분리되어 있는지 재검증한다.

## 수용 기준

- `tutorial-0-movement`는 4문항 이하이며 기본 이동 감각만 다룬다.
- `tutorial-1-survival`은 3문항 이하이며 `esc`, `:q!`, `:wq` 생존 흐름만 다룬다.
- `tutorial-2-fast-navigation`은 8문항 이하이며 word/file navigation을 다룬다.
- `tutorial-3-ex-command`는 `Mid Tutorial` 제목을 가지며 substitute 계열만 다룬다.
- `tutorial-3-ex-command`는 default playable path에는 남기되 first tour 초반 튜토리얼과 분리되어 있다.
- 전체 playable E2E가 통과한다.

## 범위

- 포함: content pacing audit
- 포함: docs status update
- 제외: 새 exercise 작성
- 제외: worldbuilding 추가

## 검증 계획

- `go test ./internal/content/...`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 검증 결과

- playlist audit: pass. `tutorial-3-ex-command` is `Mid Tutorial: Ex command 고급 튜토리얼` and contains only substitute exercises.
- `go test ./internal/content/...`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass
- `git diff --check`: pass

## 작업 항목

- [x] playlist 구성을 검토한다.
- [x] content loader test를 실행한다.
- [x] 전체 playable E2E를 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM/MIDTERM을 갱신한다.
