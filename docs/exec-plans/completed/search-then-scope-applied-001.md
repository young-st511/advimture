# SEARCH-THEN-SCOPE-APPLIED-001 — Search Then Scope Drill

Slice-ID: SEARCH-THEN-SCOPE-APPLIED-001
Created: 2026-06-07
Status: completed
Scope-Mode: normal

## 영향 도메인

- Gameplay: 새 incident content를 추가하되 schema/progress/dependency는 변경하지 않는다.
- Verification: TUI E2E는 `setup.home: temp`와 app_state/key_trace/progress evidence를 사용한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: Incident Run은 tutorial이 아니라 기존 command를 조합하는 적용 런이다.
- `docs/gameplay/command-choice-drills.md`: `search-then-act` 후보는 먼저 marker를 찾고 주변 scope를 판단한다.
- `docs/verification/spec.md`: 긴 incident/applied route는 final/timeline/app_state evidence를 남긴다.

## 목표

`/marker`로 문제 위치를 찾은 뒤 그 자리에서 linewise scope를 판단해 줄 묶음을 제거하는 작은 applied incident를 추가한다.

## 선택 결정

`incident-005-command-choice`는 이미 6 beat라 더 붙이지 않는다. 이번 slice는 별도 `incident-008-search-scope` 1-beat run으로 열어 route 길이를 낮추고, search와 linewise scope의 조합 자체를 한 문제 안에서 검증한다.

## 포함

- `incident-search-scope-001` exercise/scenario/playlist
- `playable_incident_008_full` focused E2E
- `make e2e-playable` 편입
- content loader count/list update
- gameplay/verification/roadmap docs 동기화

## 제외

- 새 Vim engine 기능
- 새 content schema
- progress 저장 포맷 변경
- 새 dependency
- `incident-005-command-choice` beat 추가

## Step 1: Spec 기반 테스트

- 목표: 새 incident가 library와 E2E suite에 포함되어야 함을 먼저 검증한다.
- 변경 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_008_full.yaml`
  - `Makefile`
- 충족 기준: loader가 새 exercise/scenario/playlist 수와 ID를 기대하고, focused E2E가 search -> linewise delete trace를 검증한다.
- Boundaries 주의: 실제 HOME을 쓰지 않는다.
- 상세 작업:
  - [x] loader expected count/list에 새 incident를 추가한다.
  - [x] focused E2E fixture를 추가한다.
  - [x] `make e2e-playable`에 focused E2E를 추가한다.

## Step 2: Content 구현

- 목표: 기존 `search-basic` + `visual-line-basic` engine만으로 playable content를 추가한다.
- 변경 파일:
  - `content/exercises/incident-search-scope.yaml`
  - `content/scenarios/incident-search-scope.yaml`
  - `content/playlists/incident-search-scope.yaml`
- 테스트 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_008_full.yaml`
- 충족 기준: optimal trace는 `/breach enter V j d`를 포함하고, target buffer/cursor/mode가 app_state로 검증된다.
- Boundaries 주의: content schema와 기존 exercise ID는 변경하지 않는다.
- 상세 작업:
  - [x] exercise를 추가한다.
  - [x] scenario copy는 상황 1문장 + Vim 판단 목표 1문장으로 유지한다.
  - [x] playlist order 8로 연결한다.

## Step 3: 문서/검증/Closeout

- 목표: 현재 roadmap과 verification spec을 새 incident 기준으로 맞춘다.
- 변경 파일:
  - `docs/gameplay/spec.md`
  - `docs/gameplay/command-choice-drills.md`
  - `docs/gameplay/vim-curriculum-map.md`
  - `docs/verification/spec.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/CHANGES.md`
  - `docs/exec-plans/active/search-then-scope-applied-001.md`
  - `docs/exec-plans/completed/search-then-scope-applied-001.md`
- 충족 기준: `SEARCH-THEN-SCOPE-APPLIED-001`은 completed로 이동하고, 다음 active 후보가 bracket pair hardening으로 정렬된다.
- Boundaries 주의: release/RC/tag 문서는 열지 않는다.
- 상세 작업:
  - [x] 현재 동작 문서를 갱신한다.
  - [x] focused tests/E2E/full tests를 실행한다.
  - [x] ExecPlan을 completed로 이동한다.

## 검증 계획

- `go test ./internal/content`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_008_full.yaml`
- `go test ./...`
- `git diff --check`

## 실행 규칙 (하네스 모드)

각 Step은 Spec 기반 테스트 작성 -> 구현 -> 검증 -> 문서 동기화 순서로 수행한다. 검증 통과 시 다음 Step으로 이동한다. 저장 포맷, content schema, dependency 변경은 금지한다.
