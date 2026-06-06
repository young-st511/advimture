# JUDGMENT-DRILL-REVIEW-001 — Command Choice as Core Identity

Status: completed
Created: 2026-06-06

## 영향 도메인

- Gameplay content: `incident-005-command-choice`의 기존 scenario copy만 보강한다.
- Gameplay docs: command-choice drill의 현재 구현 beat와 후속 후보를 최신화한다.

## 수용 기준 참조

- `docs/gameplay/command-choice-drills.md`: command choice drill은 새 command 소개가 아니라 도구 선택 판단을 훈련한다.
- `docs/gameplay/domain-contract.md`: 새 content schema, ID/파일명 변경 금지.
- 사용자 제공 중기 Plan: command-choice를 Advimture의 차별점으로 강화한다.

## Step 1: Beat-to-Judgment Review

- 목표: command-choice 5 beat가 scope/range/inline/reuse/repeat-change 판단 질문에 명확히 매핑되게 한다.
- 변경 파일:
  - `content/scenarios/command-choice.yaml`
- 충족 기준:
  - 각 beat의 success/failure copy가 command 이름보다 선택 이유를 우선한다.
  - 새 command, 새 schema, 새 exercise를 추가하지 않는다.
- Boundaries 주의:
  - exercise target state, optimal keys, constraints를 변경하지 않는다.
- 상세 작업:
  - [x] scope choice copy를 검토/보강한다.
  - [x] range choice copy를 검토/보강한다.
  - [x] inline target choice copy를 검토/보강한다.
  - [x] reuse choice copy를 검토/보강한다.
  - [x] repeat-change reuse copy를 검토/보강한다.
  - [x] focused E2E의 exact copy assertion을 최신화한다.

## Step 2: Drill Docs/Evidence Sync

- 목표: 현재 구현 beat와 후속 후보를 문서로 정리한다.
- 변경 파일:
  - `docs/gameplay/command-choice-drills.md`
  - `docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md`
- 충족 기준:
  - `playable_command_choice_scope` final/timeline/app_state evidence가 선택 이유를 확인할 수 있다.
  - 후속 후보는 문서로만 정리한다.
- 상세 작업:
  - [x] 현재 구현 5 beat mapping을 최신화한다.
  - [x] line reuse, search-then-scope, bracket-pair hardening 후보를 문서로만 정리한다.
  - [x] first dispatch + judgment drill evidence bundle을 정리한다.

## 검증 계획

- `go test ./internal/content ./internal/playable`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`
- `git diff --check`

## 완료 기록

- Completed on 2026-06-06.
- `incident-005-command-choice` 5 beat copy를 선택 이유 중심으로 보강했다.
- `docs/gameplay/command-choice-drills.md`와 `docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md`를 최신화했다.
- `playable_command_choice_scope` 통과.
- `go test ./internal/content ./internal/playable` 통과.
