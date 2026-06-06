# LINE-REUSE-APPLIED-001 — Line Reuse Choice Drill

Slice-ID: LINE-REUSE-APPLIED-001
Created: 2026-06-06
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md
- docs/roadmap/CONTENT_QUALITY_PLAN_001.md
- docs/roadmap/POST_POLISH_PLAYTEST_2026-06-06.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/gameplay/vim-curriculum-map.md
- docs/verification/spec.md
- docs/exec-plans/active/line-reuse-applied-001-line-reuse-drill.md
- docs/exec-plans/completed/line-reuse-applied-001-line-reuse-drill.md
- content/exercises/command-choice.yaml
- content/scenarios/command-choice.yaml
- content/playlists/command-choice.yaml
- test/e2e/playable_command_choice_scope.yaml
- internal/content/loader_test.go

## 목표

`incident-005-command-choice`에 line reuse choice beat를 추가한다. 플레이어는 검증된 route 줄 전체를 직접 다시 입력하지 않고 linewise `V` + `y` + `p`로 backup 위치에 재사용해야 한다.

## 포함

- `command-choice-line-reuse-001` exercise
- matching scenario
- `incident-005-command-choice` sixth beat
- focused command-choice E2E update
- loader count/playable list update
- gameplay/verification/roadmap docs 동기화

## 제외

- 새 Vim engine 기능
- 새 content schema
- progress 저장 포맷 변경
- 새 dependency
- search-then-scope, bracket-pair hardening
- 기존 exercise의 target state, optimal keys, constraints 변경

## 수용 기준

- optimal trace는 `V`, `y`, `j`, `p`를 포함하고 `buffer == target && mode == normal`로 clear된다.
- 직접 재입력, charwise visual, delete/change route, Ex command route는 forbidden route로 막는다.
- success/failure copy는 command 이름보다 "검증된 줄 전체를 재입력하지 않고 재사용해야 하는 이유"를 먼저 설명한다.
- `incident-005-command-choice`는 scope/range/inline/quote reuse/repeat-change/line reuse 여섯 판단 질문을 설명한다.
- `playable_command_choice_scope`는 여섯 beat 완료, final/timeline/app_state evidence, progress 저장, key trace를 검증한다.
- `go test ./internal/content`, `go test ./internal/content ./internal/playable`, `go test ./...`, `git diff --check`, focused E2E, `make release-check`를 통과한다.
- 금지 경계 확인: `go.mod`, `go.sum`, `internal/progress/`, content schema 변경 없음.

## Step 1: Content

- [x] exercise 추가
- [x] scenario 추가
- [x] playlist sixth beat 연결

## Step 2: Verification Wiring

- [x] loader expected counts/IDs 갱신
- [x] focused E2E 갱신

## Step 3: Documentation

- [x] gameplay spec 최신화
- [x] command-choice drill mapping 최신화
- [x] evidence bundle 최신화
- [x] roadmap current plan 최신화

## Step 4: Validation

- [x] focused content/playable tests
- [x] focused E2E
- [x] full regression
- [x] release-check

## Step 5: Closeout

- [x] ExecPlan을 completed로 이동
