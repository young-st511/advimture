# ExecPlan: UI Modal/Hint Post-release Hardening

Slice-ID: UI-MODAL-ACTION-HIERARCHY-002
Created: 2026-06-20
Status: active
Scope-Mode: strict
Allowed-Paths:
- `docs/exec-plans/active/ui-modal-action-hierarchy-002-post-release.md`
- `docs/exec-plans/completed/ui-modal-action-hierarchy-002-post-release.md`
- `docs/gameplay/spec.md`
- `docs/gameplay/tui-screen-contract.md`
- `docs/gameplay/tui-ux-direction.md`
- `docs/verification/spec.md`
- `docs/verification/tui-e2e-loop.md`
- `docs/verification/tui-ui-qa-contract.md`
- `docs/roadmap/PROGRAM.md`
- `docs/roadmap/MIDTERM_TODO.md`
- `docs/roadmap/FORWARD_PLAN.md`
- `docs/roadmap/UX_BACKLOG_001.md`
- `docs/roadmap/CHANGES.md`
- `internal/playable/**`
- `internal/playableview/**`
- `cmd/e2e-runner/**`
- `test/e2e/**`
- `artifacts/e2e/**`

## 영향 도메인

- Gameplay: 실패/성공 modal placement, running hint/quit affordance, 화면 action language를 다룬다.
- Verification: final viewport evidence, wide-width 한글 reconstruction, summary evidence 의미를 다룬다.
- Roadmap: v0.2.0 배포 이후 발견된 UX follow-up을 active slice로 승격한다.

## 공통 제한

- release/tag/push 작업은 하지 않는다.
- `go.mod`, `go.sum` 변경 없음.
- `internal/progress/` 저장 포맷 변경 없음.
- content schema 변경 없음.
- 새 Vim engine capability 추가 없음.
- 새 dependency 추가 없음.
- 기존 exercise target state, optimal keys, constraints 변경 없음.
- Todo마다 기준 고정, focused 검증, 문서 동기화, 커밋을 수행한다.

## 배경

v0.2.0 배포 후 사용자 스크린샷과 SubAgent UX 검토를 같은 기준으로 다시 돌린 결과, 치명적인 진행 불가 P0는 없었다. 다만 실패/성공 modal이 아직 console buffer/status 흐름 뒤에 이어지는 블록처럼 읽히는 장면, running hint/quit utility action이 cue/hint 문장과 같은 줄에서 섞이는 장면, 한글 wide-width final evidence가 중간 공백처럼 보이는 QA blind spot이 확인됐다.

이번 slice는 새 콘텐츠나 schema를 열지 않고, 현재 TUI가 "검증 가능한 Runbook Console"로 읽히는지에 집중한다.

## Todo 1: Plan and Contract Refresh

- 상태: completed
- 목표: 발견된 UX/UI 이슈를 active Todo로 승격하고, 구현 전 수용 기준을 문서화한다.
- 변경 파일:
  - `docs/exec-plans/active/ui-modal-action-hierarchy-002-post-release.md`
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/verification/spec.md`
  - `docs/verification/tui-ui-qa-contract.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/roadmap/CHANGES.md`
- 충족 기준:
  - active slice와 Todo 순서가 문서화된다.
  - modal placement, running utility action, final-frame evidence, action language 정리 기준이 spec/contract에 반영된다.
  - content schema/progress schema/dependency 금지 경계가 명시된다.

## Todo 2: True Console Modal Placement

- 상태: completed
- 목표: failed/succeeded modal을 buffer/status 뒤에 붙은 본문 블록이 아니라 console surface 위 overlay처럼 배치한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `test/e2e/playable_viewport_failure_modal_80x24.yaml`
  - `test/e2e/playable_viewport_success_modal_80x24.yaml`
  - 필요 시 `test/e2e/playable_review_queue.yaml`
- 충족 기준:
  - 80x24 failure/success final viewport에서 modal은 `RUNBOOK CONSOLE` 직후 console surface를 덮으며, buffer/status/grade 뒤에 append된 것처럼 보이지 않는다.
  - final modal 화면에서 `NORMAL · failed`, `NORMAL · succeeded`, `Exercise:`, `Grade:`는 decision focus를 흐리지 않는다.
  - modal action footer는 final viewport에서 잘리지 않는다.
  - `go test ./internal/playableview`와 focused viewport E2E가 통과한다.
- 검증:
  - `go test ./internal/playableview`
  - `playable_viewport_failure_modal_80x24`
  - `playable_viewport_success_modal_80x24`

## Todo 3: Running Utility Action Separation

- 상태: completed
- 목표: running 상태의 `힌트: ?`, `종료: q`를 cue/hint 본문과 물리적으로 분리된 utility action line으로 유지한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `test/e2e/playable_incident_hint_affordance.yaml`
  - 필요 시 `internal/playable/model.go`, `internal/playable/model_test.go`
- 충족 기준:
  - `보조 행동  힌트: ? · 종료: q`는 현재 목표/판단 cue나 `힌트 내용 ... · 등급에 영향`과 같은 wrapping group에 합쳐지지 않는다.
  - `playable_incident_hint_affordance`는 80x24 final viewport assertion으로 hint/action readability를 검증한다.
  - `go test ./internal/playableview ./internal/playable`와 focused hint E2E가 통과한다.
- 검증:
  - `go test ./internal/playableview ./internal/playable`
  - `playable_incident_hint_affordance`

## Todo 4: Final Frame Wide-character Evidence

- 상태: completed
- 목표: `screen_final.txt`가 한글 wide-width continuation cell을 공백처럼 출력해 `힌트 내 용`처럼 보이는 문제를 제거한다.
- 변경 파일:
  - `cmd/e2e-runner/main.go`
  - `cmd/e2e-runner/main_test.go`
  - `docs/verification/spec.md`
  - `docs/verification/tui-e2e-loop.md`
  - `docs/verification/tui-ui-qa-contract.md`
- 충족 기준:
  - terminal cell-grid reconstruction은 double-width continuation cell을 사람이 읽는 문자열에 삽입하지 않는다.
  - 한글 final viewport fixture test가 `힌트 내용`처럼 붙은 텍스트를 검증한다.
  - `go test ./cmd/e2e-runner`와 focused hint E2E가 통과한다.
- 검증:
  - `go test ./cmd/e2e-runner`
  - `playable_incident_hint_affordance`

## Todo 5: Action Language and Evidence Metadata Cleanup

- 상태: pending
- 목표: user-facing action label과 QA evidence 의미를 더 명확히 한다.
- 변경 파일:
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `cmd/e2e-runner/main.go`
  - `cmd/e2e-runner/main_test.go`
  - `test/e2e/**`
  - 관련 docs
- 충족 기준:
  - 사용자 화면의 `next_runbook` label은 `다음 런북: enter`처럼 한국어 product tone을 유지한다.
  - renderer priority marker는 현재 한국어 action label 계약을 기준으로 한다.
  - `summary.json`의 evidence flag는 "raw state 존재"와 "artifact file 저장"을 혼동하지 않도록 문서 또는 필드 의미가 정리된다.
  - 관련 Go test와 representative E2E가 통과한다.

## Todo 6: Evidence Refresh and Close

- 상태: pending
- 목표: 전체 변경을 evidence로 재검증하고 plan을 completed로 이동한다.
- 변경 파일:
  - `artifacts/e2e/**`
  - `docs/roadmap/**`
  - `docs/exec-plans/completed/ui-modal-action-hierarchy-002-post-release.md`
- 충족 기준:
  - `git diff --check` 통과.
  - `go test ./internal/playableview ./internal/playable ./cmd/e2e-runner` 통과.
  - `go test ./...` 통과.
  - focused E2E:
    - `playable_viewport_success_modal_80x24`
    - `playable_viewport_failure_modal_80x24`
    - `playable_incident_hint_affordance`
    - `playable_command_mismatch_feedback`
  - `make e2e-playable` 통과.
  - forbidden boundary 확인: `go.mod`, `go.sum`, `internal/progress/`, content schema 변경 없음.

## 검증 계획

- Todo별 focused Go test
- Todo별 focused E2E
- `go test ./...`
- `make e2e-playable`
- `git diff --check`
- `git diff --name-only -- go.mod go.sum internal/progress`

## 리스크

- modal top placement를 바꾸면 기존 final viewport evidence의 줄 위치가 바뀐다. 문자열 존재보다 "decision focus" 기준을 테스트에 반영해야 한다.
- wide-width cell reconstruction 수정은 E2E evidence 출력에만 영향을 주어야 하며, app runtime 렌더링 의미를 바꾸면 안 된다.
- summary metadata 정리는 기존 artifact 소비자와 충돌할 수 있으므로, 필드 제거보다 의미 문서화 또는 additive field를 우선한다.
