# ExecPlan: Second Dispatch Pack

Slice-ID: CONTENT-BREADTH-003
Created: 2026-06-27
Status: completed
Scope-Mode: strict

## 영향 도메인

- Gameplay: 기존 engine command만 조합해 Runbook Dispatch incident 009~011을 추가한다.
- Verification: content replay gate, focused E2E, RedTeam E2E, full playable E2E로 새 route와 우회 방지를 검증한다.

## Domain Contract 핵심 제약

- 설계 순서는 `Vim command -> Exercise -> Scenario`다.
- 새 Vim engine capability, 새 dependency, progress schema 변경은 하지 않는다.
- content schema는 유지한다.
- 새 incident는 기존 command cluster만 사용한다.
- TUI E2E는 테스트 전용 HOME을 사용하고 실제 `~/.advimture`를 건드리지 않는다.

## 수용 기준 참조

- `docs/gameplay/spec.md`
  - Incident Run은 새 command 소개가 아니라 이미 배운 command를 조합하는 적용 런이다.
  - command-choice/applied incident는 정답 key sequence보다 선택 이유를 성공/실패 피드백에서 강화한다.
  - approved + implemented exercise는 replay gate와 E2E assertion gate를 통과해야 한다.
- `docs/verification/spec.md`
  - 긴 incident full route는 screen timeline/final screen/app_state evidence를 남긴다.
  - TUI assertion은 화면 텍스트뿐 아니라 key trace, progress, app state를 함께 검증한다.

## 중간 Plan

이번 콘텐츠 확장은 세 개의 짧은 Runbook Dispatch를 순서대로 추가한다.

1. `incident-009-search-inline`: 검색으로 손상 위치를 고정하고, delimiter를 보존하는 inline change를 선택한다.
2. `incident-010-reuse`: 검증된 quote 값과 줄 전체를 직접 다시 입력하지 않고 재사용한다.
3. `incident-011-range-substitute`: 현재 줄 치환과 전체 파일 치환을 상황에 따라 구분한다.
4. `redteam-scope-guard`: 요구 command/range를 만족하지 않고 클리어 가능한지 E2E로 시도한다.

## Step 1: Incident 009 Search Inline Dispatch

- 상태: completed
- 목표: `/`, `n`, `ct,` 조합을 applied incident로 추가한다.
- 변경 파일:
  - `content/playlists/incident-search-inline.yaml`
  - `content/exercises/incident-search-inline.yaml`
  - `content/scenarios/incident-search-inline.yaml`
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_009_full.yaml`
- 테스트 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_009_full.yaml`
- 충족 기준:
  - `incident-009-search-inline`은 order 9 incident playlist다.
  - 각 beat는 기존 `search-basic`/`char-find-line` command만 사용한다.
  - `cf,`, 한 글자 이동, visual/delete 우회를 constraints로 막는다.
  - focused E2E가 final screen, key trace, progress, app_state를 검증한다.
- Boundaries 주의:
  - engine/content schema/progress schema 변경 금지.
- TODO:
  - [x] content loader Red test에 incident 009 expected IDs를 추가한다.
  - [x] exercise/scenario/playlist YAML을 추가한다.
  - [x] focused E2E를 추가한다.
  - [x] focused Go/E2E 검증을 통과시킨다.

## Step 2: Incident 010 Reuse Dispatch

- 상태: completed
- 목표: `yi"`, `P`, `yy`, `p`로 검증된 값/줄을 재사용하는 applied incident를 추가한다.
- 변경 파일:
  - `content/playlists/incident-reuse.yaml`
  - `content/exercises/incident-reuse.yaml`
  - `content/scenarios/incident-reuse.yaml`
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_010_full.yaml`
- 테스트 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_010_full.yaml`
- 충족 기준:
  - `incident-010-reuse`는 order 10 incident playlist다.
  - 직접 재입력, insert/open-line 우회를 forbidden constraints로 막는다.
  - success/failure copy는 "다시 치지 않고 검증된 값/줄을 재사용하는 이유"를 설명한다.
  - focused E2E가 final screen, key trace, progress, app_state를 검증한다.
- Boundaries 주의:
  - 새 register 기능, named register, progress 변경 금지.
- TODO:
  - [x] content loader Red test에 incident 010 expected IDs를 추가한다.
  - [x] exercise/scenario/playlist YAML을 추가한다.
  - [x] focused E2E를 추가한다.
  - [x] focused Go/E2E 검증을 통과시킨다.

## Step 3: Incident 011 Range Substitute Dispatch

- 상태: completed
- 목표: `:s/.../.../g`와 `:%s/.../.../g`의 scope 판단을 applied incident로 추가한다.
- 변경 파일:
  - `content/playlists/incident-range-substitute.yaml`
  - `content/exercises/incident-range-substitute.yaml`
  - `content/scenarios/incident-range-substitute.yaml`
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_011_full.yaml`
- 테스트 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_incident_011_full.yaml`
- 충족 기준:
  - `incident-011-range-substitute`는 order 11 incident playlist다.
  - 첫 beat는 현재 줄 범위만 바꾸고 `%` 우회를 실패시킨다.
  - 둘째 beat는 전체 파일 범위를 요구하고 현재 줄만 바꾸는 route로는 성공하지 못한다.
  - focused E2E가 final screen, key trace, progress, app_state를 검증한다.
- Boundaries 주의:
  - substitute engine capability는 기존 literal substitute 범위 안에서만 사용한다.
- TODO:
  - [x] content loader Red test에 incident 011 expected IDs를 추가한다.
  - [x] exercise/scenario/playlist YAML을 추가한다.
  - [x] focused E2E를 추가한다.
  - [x] focused Go/E2E 검증을 통과시킨다.

## Step 4: RedTeam Scope Guard

- 상태: completed
- 목표: 요구조건을 만족하지 않고 클리어 가능한지 E2E로 공격한다.
- 변경 파일:
  - `test/e2e/playable_incident_011_redteam_scope_guard.yaml`
  - `Makefile`
  - `docs/gameplay/spec.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/CHANGES.md`
- 테스트 파일:
  - `test/e2e/playable_incident_011_redteam_scope_guard.yaml`
- 충족 기준:
  - RedTeam E2E는 forbidden `%` over-scope route가 failed 상태로 막히는지 검증한다.
  - RedTeam E2E는 progress가 저장되지 않고 app_state status가 `failed`인지 검증한다.
  - `make e2e-playable`, `go test ./...`, `git diff --check`가 통과한다.
- Boundaries 주의:
  - RedTeam fixture도 테스트 전용 HOME을 사용한다.
- TODO:
  - [x] RedTeam E2E를 추가한다.
  - [x] Makefile full playable suite에 incident 009~011과 RedTeam guard를 포함한다.
  - [x] docs를 현재 동작으로 동기화한다.
  - [x] full regression을 통과시킨다.

## Verification

- `go test ./internal/content ./internal/playable ./internal/playableview`
- `go test ./...`
- `playable_incident_009_full`
- `playable_incident_010_full`
- `playable_incident_011_full`
- `playable_incident_011_redteam_scope_guard`
- `make e2e-playable`
- `git diff --check`

---

## 실행 규칙 (하네스 모드)

각 Step은 아래 사이클을 순서대로 수행한다.

### 1. Spec 기반 테스트 작성

- ExecPlan의 충족 기준을 content loader test 또는 focused E2E로 먼저 옮긴다.
- 테스트가 구현 전 실패하는 Red 상태를 확인한다.

### 2. 구현

- 테스트를 통과하도록 해당 Step의 content/E2E/docs를 최소 범위로 구현한다.
- 새 schema, 새 dependency, progress 저장 포맷은 변경하지 않는다.

### 3. 검증

- 변경 범위 focused Go test
- focused E2E
- 필요 시 `go test ./...`
- 최종 `make e2e-playable`
- `git diff --check`

### 4. 문서 동기화

- `docs/gameplay/spec.md`, `docs/roadmap/PROGRAM.md`, `docs/roadmap/CHANGES.md`를 현재 동작과 맞춘다.
- 완료 시 ExecPlan을 `docs/exec-plans/completed/`로 이동한다.

### 5. 커밋

AGENTS.md 규칙에 따라 사용자가 명시적으로 요청하기 전까지 커밋하지 않는다.
