# ExecPlan: Playtest Hint RedTeam Report Quality Lift

Slice-ID: QUALITY-LIFT-001
Created: 2026-07-02
Status: completed
Scope-Mode: strict

## 영향 도메인

- Gameplay: Second Dispatch 이후 전체 playable loop의 학습 완성도, hint 공개 방식, success/failure report 언어를 보강한다.
- Verification: fresh playtest review, focused E2E, RedTeam QA matrix를 통해 우회 가능성과 UI 의미 계약을 검증한다.

## Domain Contract 핵심 제약

- progress 저장 포맷은 변경하지 않는다.
- content schema, engine capability, 새 dependency는 변경하지 않는다.
- TUI E2E는 테스트 전용 HOME만 사용한다.
- 새 게임 루프/화면 전환/QA 자동화 변경은 이 ExecPlan 수용 기준과 테스트를 먼저 둔 뒤 진행한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`
  - incident running 기본 화면은 정답 key sequence나 command memory를 기본 노출하지 않는다.
  - hint 요청 또는 실패 후에만 command memory를 점진 공개한다.
  - failed/succeeded modal은 viewport decision surface이며, action footer와 body가 섞이면 안 된다.
  - failed/succeeded 상태는 mode-specific cue보다 우선한다.
- `docs/verification/spec.md`
  - 긴 incident full route는 screen timeline/final screen/app_state evidence를 남긴다.
  - E2E는 화면 텍스트뿐 아니라 key trace, progress, app state를 함께 검증한다.
  - RedTeam/viewport E2E는 누적 stream이 아니라 final viewport와 app_state 의미를 함께 본다.

## 중간 Plan

1. `POST-SECOND-DISPATCH-PLAYTEST-001`: 009~011 이후 전체 playable evidence를 읽고 병목을 문서화한다.
2. `UI-HINT-LADDER-001`: hint를 한 번에 정답처럼 보여주지 않고 단계적으로 공개한다.
3. `REDTEAM-QA-MATRIX-001`: 주요 우회 전략을 matrix로 만들고 representative E2E를 추가한다.
4. `UI-REPORT-001`: success/failure report에서 배운 점, 기록, 다음 행동의 시선 위계를 더 정리한다.

## Step 1: Fresh Playtest Review

- 상태: completed
- 목표: Second Dispatch까지 포함한 fresh evidence를 읽고, 바로 구현할 품질 병목을 명확히 한다.
- 변경 파일:
  - `docs/roadmap/POST_SECOND_DISPATCH_PLAYTEST_REVIEW_2026-07-02.md`
  - `docs/roadmap/PROGRAM.md`
- 테스트 파일:
  - 기존 focused/full E2E evidence
- 충족 기준:
  - 009~011 route와 RedTeam guard의 final/timeline/app_state evidence를 검토한다.
  - hint ladder, RedTeam matrix, report polish의 구체 작업 기준을 리뷰 문서에 정리한다.
  - release/tag/push, progress schema, content schema, 새 dependency는 범위 밖으로 명시한다.
- Boundaries 주의:
  - 문서 리뷰만 수행하고 코드 동작은 변경하지 않는다.
- TODO:
  - [x] fresh evidence를 재생성하거나 최신 E2E artifact를 확인한다.
  - [x] P0/P1/P2 이슈와 다음 구현 순서를 문서화한다.
  - [x] PROGRAM의 active slice를 `QUALITY-LIFT-001`로 갱신한다.

## Step 2: UI-HINT-LADDER-001

- 상태: completed
- 목표: `?` hint 요청이 즉시 exact command spoiler가 되지 않도록 단계적 hint 공개를 구현한다.
- 변경 파일:
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
  - `docs/gameplay/spec.md`
  - 필요 시 focused E2E
- 테스트 파일:
  - `internal/playable/model_test.go`
  - 필요 시 `test/e2e/playable_hint_ladder.yaml`
- 충족 기준:
  - 첫 hint는 판단 관점 또는 목표 범위를 설명한다.
  - 두 번째 hint는 command family 또는 reviewed command 후보를 보여준다.
  - 세 번째 이후에만 기존 exact hint 또는 command memory를 보여준다.
  - failed 상태의 command memory 공개는 기존 계약을 유지한다.
  - progress 저장 포맷, content schema, scoring schema를 변경하지 않는다.
- Boundaries 주의:
  - `internal/progress/`, content YAML schema, `go.mod`, `go.sum` 변경 금지.
- TODO:
  - [x] Red test로 첫/둘째/셋째 hint 공개 단계를 고정한다.
  - [x] Model 내부 상태만으로 hint ladder를 구현한다.
  - [x] focused Go/E2E 검증을 통과시킨다.
  - [x] spec을 현재 동작으로 갱신한다.

## Step 3: REDTEAM-QA-MATRIX-001

- 상태: completed
- 목표: 대표 incident 우회 전략을 matrix로 만들고, 최소 E2E로 검증 가능한 우회 방지망을 확장한다.
- 변경 파일:
  - `docs/verification/redteam-qa-matrix.md`
  - `test/e2e/*redteam*.yaml`
  - `Makefile`
  - `docs/verification/spec.md`
- 테스트 파일:
  - 신규 RedTeam E2E
- 충족 기준:
  - 직접 재입력 우회, delimiter 포함/삭제 실수, range over-scope, under-scope, required key 누락 중 최소 3개 유형을 matrix에 기록한다.
  - 신규 RedTeam E2E는 progress가 저장되지 않고 app_state `status: failed`를 검증한다.
  - `make e2e-playable`에 representative RedTeam E2E를 포함한다.
- Boundaries 주의:
  - E2E는 `setup.home: temp` 또는 `complete_before`만 사용한다.
- TODO:
  - [x] matrix 문서를 추가한다.
  - [x] RedTeam E2E를 2개 이상 추가한다.
  - [x] focused E2E를 통과시킨다.
  - [x] `make e2e-playable`을 통과시킨다.

## Step 4: UI-REPORT-001

- 상태: completed
- 목표: success/failure report를 더 읽기 쉬운 recovery report grammar로 정리한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - 필요 시 `internal/playable/model.go`
  - `docs/gameplay/spec.md`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - viewport focused E2E
- 충족 기준:
  - success report는 `배운 점`, `기록`, `다음 행동`이 80x24에서 명확히 분리된다.
  - failure report는 `실수`, `복구 힌트`, `다시 시도`가 긴 본문과 섞이지 않는다.
  - action footer label은 잘리지 않고, `ui.focus_panel.actions` 의미는 유지한다.
  - 색/style이 없어도 동일한 의미를 읽을 수 있다.
- Boundaries 주의:
  - progress 저장 포맷과 action id 계약은 변경하지 않는다.
- TODO:
  - [x] Red test로 80x24 report priority를 고정한다.
  - [x] renderer를 최소 범위로 보강한다.
  - [x] viewport focused E2E와 Go test를 통과시킨다.
  - [x] spec을 현재 동작으로 갱신한다.

## Step 5: Final Regression and Docs

- 상태: completed
- 목표: 전체 품질 게이트를 통과시키고 ExecPlan을 completed로 이동한다.
- 변경 파일:
  - `docs/roadmap/CHANGES.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/exec-plans/completed/quality-lift-001-playtest-hint-redteam-report.md`
- 테스트 파일:
  - 전체 Go/E2E suite
- 충족 기준:
  - `go test ./...`, `make e2e-playable`, `git diff --check` 통과.
  - `go.mod`, `go.sum`, `internal/progress/` 변경 없음.
  - ExecPlan은 completed로 이동한다.
- Boundaries 주의:
  - 사용자가 명시하기 전까지 commit/tag/release/push는 하지 않는다.
- TODO:
  - [x] full regression을 통과시킨다.
  - [x] docs를 최종 상태로 동기화한다.
  - [x] ExecPlan을 completed로 이동한다.

---

## 실행 규칙 (하네스 모드)

각 Step은 아래 사이클을 순서대로 수행한다.

### 1. Spec 기반 테스트/근거 작성

- 구현 Step은 먼저 테스트 또는 E2E를 작성해 Red 상태를 확인한다.
- 문서 review Step은 fresh evidence와 현재 spec/contract를 먼저 읽고 근거를 문서화한다.

### 2. 구현

- 테스트를 통과하도록 최소 범위로 구현한다.
- 새 schema, 새 dependency, progress 저장 포맷은 변경하지 않는다.

### 3. 검증

- 변경 범위 focused Go test
- focused E2E
- 필요 시 `go test ./...`
- 최종 `make e2e-playable`
- `git diff --check`

### 4. 문서 동기화

- `docs/gameplay/spec.md`, `docs/verification/spec.md`, `docs/roadmap/PROGRAM.md`, `docs/roadmap/CHANGES.md`를 현재 동작과 맞춘다.

### 5. 커밋

AGENTS.md 규칙에 따라 사용자가 명시적으로 요청하기 전까지 커밋하지 않는다.
