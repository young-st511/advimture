# ExecPlan: PM 직접 플레이와 제품 개선 기획

Slice-ID: PM-PLAYTEST-AUDIT-001
Created: 2026-09-29
Status: review-ready
Scope-Mode: strict
Allowed-Paths: docs/exec-plans/, docs/roadmap/PM_PLAYTEST_REVIEW_2026-09-29.md, artifacts/e2e/pm-audit/

## 목표

사용자의 "Product Manager가 되어 제품을 직접 사용하고 버그, 부족한 점, 개선 방향을 기획" 요청을 수행한다. 실제 TUI 입력/화면과 상태 evidence를 기준으로 구현 결함, 의도된 설계의 제품 문제, 추가 사용자 검증이 필요한 가설을 구분한다.

## 범위

- 포함: 독립 HOME의 실제 PTY 플레이, 첫 실행/실패/힌트/재시도/재개/응용/복습 검토, 재현 evidence, 우선순위와 개선안 수용 기준 초안.
- 제외: 게임 코드/콘텐츠/저장 포맷/의존성 수정, 배포, 외부 사용자 실험을 수행했다고 간주하는 것.
- 감사 수행은 사용자 요청 범위다. 후속 제품 변경의 수용 기준은 보고서에서 `[draft]`로 제안하며 이번 작업에서 구현하지 않는다.

## 완료 기준

- [x] 실제 바이너리를 PTY에서 직접 조작하고 첫 실행부터 성공/실패/재시도까지 관찰한다.
- [x] 80x24와 110x34에서 대표 경로를 확인하고 재현 시나리오와 evidence를 남긴다.
- [x] 응용 대안 풀이, 힌트 비용/효용, 저장/재개, 복습의 기대와 실제를 비교한다.
- [x] 발견마다 재현 단계, 영향, 분류, 우선순위, 코드 근거를 기록한다.
- [x] 후속 계획에 순서/완료 기준/사용자 검증 기준을 명시한다.
- [x] 문서 및 evidence를 검토하고 문서 커밋과 PR을 만든다.

## 검증 계획

- `go build`로 앱과 기존 E2E runner를 빌드한다.
- 직접 입력 PTY 로그와 runner의 final/timeline/app_state/progress evidence를 `artifacts/e2e/pm-audit/`에 격리한다.
- 재현 시나리오의 통과는 관찰한 현재 동작의 재현을 뜻하며 제품 품질 통과와 구분한다.
- 전 턴 동일 commit의 `go test ./...` 통과를 기준선으로 사용한다. 코드 변경이 없으면 전체 단위 테스트를 반복하지 않는다.
- `git diff --check`, 변경 경계 확인.

## 의사결정 로그

- 2026-09-29: 기준 commit `aa685ce`, VERSION `0.4.0`, 전용 branch `codex/pm-playtest-audit`. TUI 전용 검증 계약을 따라 브라우저 대신 PTY를 사용한다.
- 2026-09-29: 직접 PTY 세션 3개와 기본 경로 5개/관찰 재현 20개를 수행했다. 모든 관찰 assertion이 충족됐다. 그중 결함을 기대하는 assertion의 통과는 품질 통과와 구별했다.
- 2026-09-29: `docs/roadmap/PM_PLAYTEST_REVIEW_2026-09-29.md`에 12개 발견과 후속 5단계 `[draft]` 기획을 작성했다. 입력/실패 복구를 첫 구현 후보로 권장하며, 기존 PRODUCT/PROGRAM과 제품 구현은 변경하지 않았다.
- 2026-09-29: 문서 커밋 `e341682`, PR https://github.com/young-st511/advimture/pull/1 제출. 감사 작업은 완료했다. `docs/exec-plans/README.md`의 PR merge 이후 이동 규칙에 따라 merge 전까지 review-ready로 active에 둔다. 후속 제품 구현을 승인한 상태는 아니다.

## 한계

- Agent의 전문가 검토다. 실제 초보자의 이해/재미/학습 전이/재방문을 대신 측정하지 않는다.
