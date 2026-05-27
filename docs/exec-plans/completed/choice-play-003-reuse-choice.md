# CHOICE-PLAY-003 — Quote Value Reuse Choice

Slice-ID: CHOICE-PLAY-003
Created: 2026-05-28
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/choice-play-003-reuse-choice.md
- docs/exec-plans/completed/reuse-choice-001-reuse-judgment.md
- content/exercises/command-choice.yaml
- content/scenarios/command-choice.yaml
- content/playlists/command-choice.yaml
- test/e2e/playable_command_choice_scope.yaml
- internal/content/loader_test.go

## 목표

`incident-005-command-choice`에 quote value reuse-choice beat를 추가한다. 플레이어는 검증된 quote 내부 token을 직접 다시 입력하지 않고 `yi"` + `P`로 빈 quote 위치에 재사용해야 한다.

## 포함

- `command-choice-quote-reuse-001` exercise
- matching scenario
- `incident-005-command-choice` fourth beat
- focused E2E update
- loader count/playable list update

## 제외

- 새 Vim engine 기능
- 새 schema
- linewise reuse-choice
- dot repeat-choice
- progress 저장 포맷 변경

## 수용 기준

- optimal trace는 `y`, `i`, `"`, `P`를 포함한다.
- 직접 재입력/수정 우회는 forbidden route로 막는다.
- command-choice E2E는 4-beat incident 완료와 final app_state를 검증한다.
- content replay, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Content

- [x] exercise 추가
- [x] scenario 추가
- [x] playlist fourth beat 연결

## Step 2: Verification Wiring

- [x] loader expected counts/IDs 갱신
- [x] focused E2E 갱신

## Step 3: Validation

- [x] content tests
- [x] focused E2E
- [x] full regression
