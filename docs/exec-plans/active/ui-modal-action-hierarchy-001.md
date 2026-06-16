# ExecPlan: UI Modal Action Hierarchy Hardening

Slice-ID: UI-MODAL-ACTION-HIERARCHY-001
Created: 2026-06-17
Status: active
Scope-Mode: normal
Allowed-Paths:
- `docs/exec-plans/active/ui-modal-action-hierarchy-001.md`
- `docs/exec-plans/active/release-brew-001-homebrew-tap.md`
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
- `internal/playable/**`
- `internal/playableview/**`
- `internal/tuiadapter/**`
- `internal/e2estate/**`
- `cmd/e2e-runner/**`
- `test/e2e/**`
- `content/scenarios/**`
- `artifacts/e2e/**`

## 영향 도메인

- Gameplay: 플레이어가 보는 실패/성공/힌트/액션 hierarchy를 다룬다. progress 저장 포맷, content schema, 새 dependency는 변경하지 않는다.
- Verification: TUI E2E evidence와 viewport/final-frame 검증을 다룬다. 실제 사용자 HOME을 쓰지 않고 테스트 전용 HOME과 artifact만 사용한다.
- Roadmap: 기존 "P0/P1 UX blocker 없음" 판단을 사용자 스크린샷 기반의 P1 release-blocker 판단으로 갱신한다.

## 수용 기준 참조

- `docs/gameplay/spec.md` > TUI Modal/Action Hierarchy
- `docs/verification/spec.md` > TUI final-frame/modal evidence
- `docs/gameplay/tui-screen-contract.md` > FocusPanel / Floating Feedback Modal
- `docs/verification/tui-ui-qa-contract.md` > Evidence Snapshot / Backlog 승격 기준

## 목표

실패/성공/debrief 화면을 "본문 뒤에 붙은 안내 블록"이 아니라, 다음 행동을 즉시 읽을 수 있는 Runbook Console modal decision surface로 만든다. Running 상태의 hint/quit affordance도 일반 문구가 아니라 조작 가능한 action으로 보이게 하며, QA가 이 차이를 문자열 존재만으로 통과시키지 않도록 evidence를 보강한다.

## 공통 제한

- release/tag/push 작업은 하지 않는다.
- `go.mod`, `go.sum` 변경 없음.
- progress 저장 포맷 변경 없음.
- content schema 변경 없음.
- 새 Vim engine capability 추가 없음.
- 새 dependency 추가 없음.
- Todo마다 검증 후 커밋한다.

## Todo 1: UX Plan Refresh

- 상태: completed
- 목표: 사용자 스크린샷 기반 UX blocker 판단을 current docs에 반영하고 release plan을 보류한다.
- 변경 파일:
  - `docs/exec-plans/active/ui-modal-action-hierarchy-001.md`
  - `docs/exec-plans/active/release-brew-001-homebrew-tap.md`
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/verification/spec.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/UX_BACKLOG_001.md`
  - `docs/roadmap/CHANGES.md`
- 충족 기준:
  - UX issue는 P1 release-blocker로 기록된다.
  - Brew release plan은 UI hardening 완료 전까지 paused/on-hold로 기록된다.
  - 자동 수행 Todo 순서와 금지 경계가 문서화된다.

## Todo 2: Final Frame QA

- 상태: pending
- 목표: modal 위치와 action footer가 실제 final viewport에서 보이는지 검증할 수 있게 한다.
- 변경 파일:
  - `cmd/e2e-runner/**`
  - `test/e2e/**`
  - `docs/verification/**`
- 테스트 파일:
  - `cmd/e2e-runner/main_test.go`
  - focused viewport E2E scenarios
- 충족 기준:
  - `screen_final.txt` 또는 새 final-frame evidence가 terminal height 기준 단일 final frame을 표현한다.
  - 80x24 success/failure E2E가 primary action footer 존재와 clipping 방지를 검증한다.
  - 기존 substring-only 검증만으로 modal UX를 통과시키지 않는다.

## Todo 3: Modal Overlay + Action Footer

- 상태: pending
- 목표: failed/succeeded/debrief 화면을 viewport 기준 modal decision surface로 렌더링하고 primary action을 강조한다.
- 변경 파일:
  - `internal/playableview/**`
  - `internal/playable/**`
  - `test/e2e/**`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - `internal/playable/model_test.go`
  - `playable_viewport_success_modal_80x24`
  - `playable_viewport_failure_modal_80x24`
- 충족 기준:
  - failed/succeeded modal은 buffer 흐름에 단순 append된 것처럼 보이지 않는다.
  - modal에는 primary action footer와 secondary action footer가 분리되어 표시된다.
  - 실패/성공 상태는 visual/insert/command/search mode cue보다 우선한다.
  - `NORMAL · failed`, `Exercise`, `Grade`는 decision focus를 해치지 않는 낮은 위계로 정리된다.

## Todo 4: Running Action Cue

- 상태: pending
- 목표: running 상태의 `힌트: ?`, `종료: q`를 조작 가능한 affordance로 구조화한다.
- 변경 파일:
  - `internal/playable/**`
  - `internal/playableview/**`
  - `internal/e2estate/**` (필요 시 internal DTO만)
  - `test/e2e/**`
- 테스트 파일:
  - `internal/playable/model_test.go`
  - `internal/playableview/render_test.go`
  - `playable_incident_hint_affordance`
  - `playable_coaching_panel`
- 충족 기준:
  - running focus panel의 현재 목표/판단 cue와 utility actions가 같은 위계로 섞이지 않는다.
  - E2E는 running hint/quit affordance를 action 의미로 검증할 수 있다.

## Todo 5: Hint Affordance

- 상태: pending
- 목표: hint revealed 상태가 도움말처럼 보이고, `Hint:`/`힌트:` 중복과 hint 비용 불명확성을 줄인다.
- 변경 파일:
  - `internal/playable/**`
  - `internal/playableview/**`
  - `test/e2e/**`
  - `docs/gameplay/**`
- 테스트 파일:
  - `internal/playable/model_test.go`
  - `playable_incident_hint_affordance`
- 충족 기준:
  - 힌트 내용과 힌트 호출 action이 구분된다.
  - hint 사용이 grade에 영향을 줄 수 있음을 짧고 과하지 않게 보여준다.
  - command/search/insert/visual mode에는 실제 입력 처리와 맞지 않는 일반 hint/quit 안내를 섞지 않는다.

## Todo 6: Incident Progressive Briefing

- 상태: pending
- 목표: incident briefing에서 exact command 과노출을 줄이고 판단 목표를 먼저 보이게 한다.
- 변경 파일:
  - `content/scenarios/**`
  - `docs/gameplay/**`
  - focused E2E scenarios
- 테스트 파일:
  - `internal/content/**`
  - affected focused E2E
- 충족 기준:
  - incident briefing은 상황 1문장 + 판단 목표 1문장을 유지한다.
  - exact command는 기본 briefing보다 hint/failure/command memory에서 노출된다.
  - target state, optimal keys, constraints는 변경하지 않는다.

## Todo 7: Evidence Refresh + Close

- 상태: pending
- 목표: 대표 UX evidence를 현행 UI로 재생성하고 stale artifact 판단을 분리한 뒤 plan을 완료한다.
- 변경 파일:
  - `artifacts/e2e/**`
  - `docs/roadmap/**`
  - `docs/exec-plans/active/ui-modal-action-hierarchy-001.md`
  - `docs/exec-plans/completed/ui-modal-action-hierarchy-001.md`
- 충족 기준:
  - 대표 success/failure/hint/incident/review evidence가 legacy `ACTION`/English action label에 오염되지 않는다.
  - `go test ./...`, focused E2E, `make e2e-playable`, `git diff --check`가 통과한다.
  - ExecPlan은 completed로 이동한다.

## 실행 규칙

각 Todo는 아래 순서로 진행한다.

1. Spec 또는 ExecPlan 충족 기준을 먼저 테스트/문서 기준으로 고정한다.
2. 구현한다.
3. focused Go test와 관련 focused E2E를 실행한다.
4. 필요 시 SubAgent 2개를 병렬로 사용해 plan 정합성과 코드 리뷰를 받는다.
5. 문서를 동기화한다.
6. Todo 단위로 커밋한다.
7. 다음 Todo로 이동한다.

## 검증 계획

- `gofmt` 대상 Go 파일
- `go test ./internal/playableview ./internal/playable ./cmd/e2e-runner`
- `go test ./...`
- focused E2E:
  - `playable_viewport_success_modal_80x24`
  - `playable_viewport_failure_modal_80x24`
  - `playable_incident_hint_affordance`
  - `playable_command_mismatch_feedback`
  - `playable_review_queue`
- `make e2e-playable`
- `git diff --check`
- 금지 경계 확인:
  - `go.mod`, `go.sum` 변경 없음
  - `internal/progress/` 저장 포맷 변경 없음
  - content schema 변경 없음

## 의사결정 로그

- 2026-06-17: 사용자가 실패 화면 스크린샷을 근거로 modal과 action affordance 문제를 제기했다.
- 2026-06-17: SubAgent 감사와 별도 세션 감사 결과를 종합해 진행 불가 P0는 아니지만 출시 가능한 수준 기준의 P1 release-blocker로 판정했다.
- 2026-06-17: 사용자는 추천안 전체를 승인했다. Release/tag/push는 보류하고, overlay + action footer + running action 구조화 + evidence hardening을 Todo별 커밋으로 수행한다.
- 2026-06-17: Todo 1 완료. Active UX hardening plan을 만들고, release plan을 paused로 전환했으며, roadmap/backlog/spec/verification 문서가 P1 modal/action blocker를 가리키도록 정리했다.

## 미해결 질문

- 없음. 구현 중 progress schema, content schema, dependency, release/tag/push 필요성이 생기면 즉시 중단하고 별도 승인을 받는다.
