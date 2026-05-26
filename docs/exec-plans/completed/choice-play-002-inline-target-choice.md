# CHOICE-PLAY-002 — Inline Target Range Choice

Slice-ID: CHOICE-PLAY-002
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/choice-play-002-inline-target-choice.md
- docs/exec-plans/completed/char-find-applied-001-inline-target-application.md
- content/exercises/command-choice.yaml
- content/scenarios/command-choice.yaml
- content/playlists/command-choice.yaml
- test/e2e/playable_command_choice_scope.yaml
- internal/content/loader_test.go

## 목표

`incident-005-command-choice`에 `ct,` inline target range choice를 추가해, 플레이어가 delimiter를 보존해야 하는 한 줄 설정에서 `cf,`나 `cw`가 아니라 `ct,`를 고르게 만든다.

## 포함

- `command-choice-inline-target-001` exercise
- matching scenario
- `incident-005-command-choice` third beat
- focused E2E update
- loader count/playable list update

## 제외

- 새 engine 기능
- 새 playlist category
- search + inline target 복합 incident
- progress 저장 포맷 변경

## 수용 기준

- optimal trace는 `c`, `t`, `,`, replacement, `esc`를 포함한다.
- `f` route는 forbidden route로 막아 delimiter 삭제 우회를 방지한다.
- command-choice E2E는 3-beat incident 완료와 final app_state를 검증한다.
- `go test ./internal/content`, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Content

- [x] exercise 추가
- [x] scenario 추가
- [x] playlist third beat 연결

## Step 2: Verification Wiring

- [x] loader expected counts/IDs 갱신
- [x] focused E2E 갱신

## Step 3: Validation

- [x] content tests
- [x] focused E2E
- [x] full regression
