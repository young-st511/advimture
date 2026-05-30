# CONTENT-BREADTH-002 — Repeat Change Choice

Slice-ID: CONTENT-BREADTH-002
Created: 2026-05-30
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/verification/spec.md
- docs/exec-plans/active/content-breadth-002-repeat-choice.md
- docs/exec-plans/completed/content-breadth-002-repeat-choice.md
- content/exercises/command-choice.yaml
- content/scenarios/command-choice.yaml
- content/playlists/command-choice.yaml
- test/e2e/playable_command_choice_scope.yaml
- internal/content/loader_test.go

## 목표

`incident-005-command-choice`에 repeat-change choice beat를 추가한다. 플레이어는 같은 단어 교체가 이어지는 상황에서 두 번째 변경을 다시 입력하지 않고 `.`으로 마지막 변경을 반복해야 한다.

## 포함

- `command-choice-repeat-change-001` exercise
- matching scenario
- `incident-005-command-choice` fifth beat
- focused E2E update
- loader count/playable list update
- roadmap/spec 동기화

## 제외

- 새 Vim engine 기능
- 새 content schema
- progress 저장 포맷 변경
- count prefix, macro, register 확장
- line reuse choice

## 수용 기준

- optimal trace는 `ciw`, `esc`, `.`을 포함한다.
- 첫 변경은 직접 수행하되, 두 번째 같은 변경은 `.`으로 반복해야 clear된다.
- 직접 재입력, substitute, yank/put, visual route는 forbidden route로 막는다.
- command-choice E2E는 5-beat incident 완료와 final app_state를 검증한다.
- content replay, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Candidate Decision

- [x] 현재 command-choice 후보와 repeat-last-change engine 범위를 확인한다.
- [x] line reuse보다 repeat-change reuse가 이번 breadth 목표에 더 적합한지 기록한다.

## Step 2: Content

- [x] exercise 추가
- [x] scenario 추가
- [x] playlist fifth beat 연결

## Step 3: Verification Wiring

- [x] loader expected counts/IDs 갱신
- [x] focused E2E 갱신

## Step 4: Validation

- [x] content tests
- [x] focused E2E
- [x] full regression

## Step 5: Closeout

- [x] gameplay/verification/roadmap docs 동기화
- [x] ExecPlan을 completed로 이동
- [x] 커밋/푸시
