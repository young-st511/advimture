# COMMAND-CHOICE-BREADTH-002 — Bracket Pair Choice Beat

Slice-ID: COMMAND-CHOICE-BREADTH-002
Created: 2026-06-07
Status: completed
Scope-Mode: normal

## 영향 도메인

- Gameplay: `incident-005-command-choice`에 bracket pair scope 판단 beat를 추가해 command-choice breadth를 넓힌다.
- Verification: command-choice full E2E와 content loader/replay gate가 7-beat route를 검증한다.

## 수용 기준 참조

- `docs/roadmap/PROGRAM.md`: 다음 권장 후보는 `COMMAND-CHOICE-BREADTH-002`다.
- `docs/gameplay/spec.md`: command choice drill은 새 command cluster가 아니라 이미 배운 command 선택 적용 레이어다.
- `docs/gameplay/command-choice-drills.md`: bracket pair 후보는 engine hardening 완료 후 command-choice breadth 후보로 승격할 수 있다.

## 목표

`incident-005-command-choice`를 여섯 beat에서 일곱 beat로 확장한다. 새 beat는 `old-value`처럼 단어 단위로는 충분하지 않은 괄호 내부 인자 전체를 보고 `ci(`를 선택하는 판단을 검증한다.

## 포함

- `command-choice-bracket-scope-001` exercise 추가
- 대응 scenario copy 추가
- `incident-005-command-choice` playlist에 seventh beat 추가
- `playable_command_choice_scope` E2E를 7-beat route로 갱신
- content loader/replay expected IDs 갱신
- gameplay/verification/roadmap docs 동기화

## 제외

- progress schema 변경
- content schema 변경
- 새 Vim command/engine capability
- 새 dependency
- release candidate/tag 작업
- nested pair, escaped delimiter, around object, count prefix, multi-line pair hardening

## Step 1: Red 기대값

- 목표: command-choice route가 bracket pair beat까지 요구하도록 테스트/E2E 기대값을 먼저 갱신한다.
- 변경 파일:
  - `internal/content/loader_test.go`
  - `test/e2e/playable_command_choice_scope.yaml`
- 충족 기준: content loader 또는 E2E가 새 exercise/beat 누락으로 실패한다.
- 상세 작업:
  - [x] playable exercise 수와 ID 목록에 `command-choice-bracket-scope-001`을 추가한다.
  - [x] focused E2E route에 bracket pair 입력 `ci(` + replacement + `esc`를 추가한다.

## Step 2: Content 구현

- 목표: 기존 schema와 engine만 사용해 bracket pair choice beat를 playable content로 추가한다.
- 변경 파일:
  - `content/exercises/command-choice.yaml`
  - `content/scenarios/command-choice.yaml`
  - `content/playlists/command-choice.yaml`
- 충족 기준: `command-choice-bracket-scope-001` replay gate가 pass하고 `incident-005-command-choice`가 7-beat route로 닫힌다.
- 상세 작업:
  - [x] `old-value` 괄호 내부 인자를 `up`으로 바꾸는 exercise를 추가한다.
  - [x] `required_keys`로 `c`, `i`, `(`, `esc`를 고정한다.
  - [x] `ciw`, quote object, linewise/visual/substitute 우회를 forbidden key로 막는다.
  - [x] success/failure copy가 command 이름보다 “단어가 아니라 괄호 내부 인자 전체”라는 이유를 먼저 말한다.

## Step 3: 문서/검증/Closeout

- 목표: current docs를 7-beat command-choice route로 동기화하고 검증한다.
- 변경 파일:
  - `docs/gameplay/command-choice-drills.md`
  - `docs/gameplay/spec.md`
  - `docs/verification/spec.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/CHANGES.md`
  - `docs/exec-plans/active/command-choice-breadth-002.md`
  - `docs/exec-plans/completed/command-choice-breadth-002.md`
- 충족 기준: focused Go tests, focused E2E, full Go tests, release-check, diff check가 통과한다.
- 상세 작업:
  - [x] command-choice docs를 7-beat mapping으로 갱신한다.
  - [x] roadmap의 next recommended를 fresh review/hardening 판정으로 넘긴다.
  - [x] 검증 명령을 실행한다.
  - [x] ExecPlan을 completed로 이동한다.

## 완료 결과

- `command-choice-bracket-scope-001`을 추가했다.
- `incident-005-command-choice`는 일곱 beat route가 됐다.
- 새 beat는 hyphenated 괄호 내부 인자 `old-value` 전체를 `ci(`로 `up`으로 교체하는 scope 판단을 검증한다.
- progress schema, content schema, dependency, 새 engine capability는 변경하지 않았다.

## 검증 계획

- `go test ./internal/content ./internal/playable`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`
- `go test ./...`
- `git diff --check`
- `make release-check`
- 금지 경계 확인: `git diff --name-only -- go.mod go.sum internal/progress`

## 실행 규칙

각 Step은 Spec 기반 테스트 작성 -> 구현 -> 검증 -> 문서 동기화 순서로 수행한다. 저장 포맷, content schema, dependency, 새 engine capability 변경은 금지한다.
