# BRACKET-PAIR-HARDEN-001 — Parenthesis/Brace Inner Text Object

Slice-ID: BRACKET-PAIR-HARDEN-001
Created: 2026-06-07
Status: completed
Scope-Mode: normal

## 영향 도메인

- Gameplay: `text-object-quote-pair` cluster를 quote 내부 object에서 bracket pair 내부 object까지 확장한다.
- Verification: engine/runtime/content replay와 full tutorial E2E를 함께 검증한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: `text-object-quote-pair`의 parenthesis/brace scope를 이번 hardening으로 닫는다.
- `docs/gameplay/vim-curriculum-map.md`: `i(`, `i{`를 quote/pair text object hardening 후보에서 완료 coverage로 이동한다.
- `docs/verification/spec.md`: full playpack E2E는 key trace와 app_state evidence를 남긴다.

## 목표

같은 줄의 비중첩 parenthesis/brace 내부 값을 `di(`, `ci(`, `yi(`, `di{`, `ci{`, `yi{`로 삭제/교체/복사할 수 있게 한다.

## 포함

- engine support for `i(` and `i{`
- runtime replay coverage
- `text-object-quote-pair` command cluster coverage update
- separate bracket pair tutorial exercises/scenarios/playlist
- `playable_bracket_pair_full` E2E
- gameplay/verification/roadmap docs 동기화

## 제외

- nested pair matching
- escaped delimiter
- around object (`a(`, `a{`)
- count/register prefix
- visual selection integration
- multi-line pair object
- 새 dependency/progress schema/content schema

## Step 1: Spec 기반 테스트

- 목표: bracket inner object의 engine/runtime/content 기대 동작을 먼저 테스트로 고정한다.
- 변경 파일:
  - `internal/vimengine/engine_test.go`
  - `internal/runtime/session_test.go`
  - `internal/content/loader_test.go`
  - `test/e2e/playable_bracket_pair_full.yaml`
- 충족 기준: 구현 전 `ci(`, `di{`, `yi{` 경로가 실패해야 한다.
- Boundaries 주의: 새 dependency 없음.
- 상세 작업:
  - [x] inner pair range unit test를 추가한다.
  - [x] operator tests for `ci(`, `di{`, `yi{`를 추가한다.
  - [x] runtime replay test를 추가한다.
  - [x] content loader count/coverage 기대값과 E2E route를 확장한다.

## Step 2: Engine/Runtime 구현

- 목표: 기존 inner quote implementation을 확장해 bracket pair 내부 range를 처리한다.
- 변경 파일:
  - `internal/vimengine/engine.go`
- 테스트 파일:
  - `internal/vimengine/engine_test.go`
  - `internal/runtime/session_test.go`
- 충족 기준: `ci(`는 내부 값을 지우고 Insert mode로 진입하며, `di{`는 내부 값만 삭제하고, `yi{`는 내부 값만 charwise register에 담는다.
- Boundaries 주의: nested pair matching은 구현하지 않는다.
- 상세 작업:
  - [x] bracket key constants를 추가한다.
  - [x] pending inner text object dispatch를 bracket pair key로 확장한다.
  - [x] same-line non-nested inner pair range 함수를 추가한다.

## Step 3: Content/E2E 구현

- 목표: parenthesis/brace text object를 별도 bracket pair tutorial에 playable content로 연결한다.
- 변경 파일:
  - `content/command_clusters/text-object-quote-pair.yaml`
  - `content/exercises/text-object-quote-pair.yaml`
  - `content/scenarios/text-object-quote-pair.yaml`
  - `content/playlists/text-object-bracket-pair.yaml`
  - `test/e2e/playable_bracket_pair_full.yaml`
  - `internal/content/loader_test.go`
- 충족 기준: tutorial playlist는 8 beat 이하를 유지하고, coverage_required가 모두 covered로 통과한다.
- Boundaries 주의: 기존 7개 exercise의 target/optimal/constraints는 변경하지 않는다.
- 상세 작업:
  - [x] `di(`, `ci(`, `yi(`, `di{`, `ci{`, `yi{` 6문항을 추가한다.
  - [x] 기존 quote tutorial은 유지하고 새 `tutorial-95-bracket-pair`로 분리한다.
  - [x] focused E2E를 갱신한다.

## Step 4: 문서/검증/Closeout

- 목표: bracket pair hardening을 completed 상태로 정리한다.
- 변경 파일:
  - `docs/gameplay/spec.md`
  - `docs/gameplay/command-catalog.md`
  - `docs/gameplay/vim-curriculum-map.md`
  - `docs/verification/spec.md`
  - `docs/roadmap/PROGRAM.md`
  - `docs/roadmap/MIDTERM_TODO.md`
  - `docs/roadmap/FORWARD_PLAN.md`
  - `docs/roadmap/CHANGES.md`
  - `docs/exec-plans/completed/bracket-pair-harden-001.md`
- 충족 기준: focused Go tests, focused E2E, full Go tests, release-check가 통과한다.
- Boundaries 주의: release/RC/tag 작업은 하지 않는다.
- 상세 작업:
  - [x] 현재 동작 문서를 갱신한다.
  - [x] 검증 명령을 실행한다.
  - [x] ExecPlan을 completed로 이동한다.

## 검증 계획

- `go test ./internal/vimengine ./internal/runtime ./internal/content`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_bracket_pair_full.yaml`
- `go test ./...`
- `make release-check`
- `git diff --check`

## 실행 규칙 (하네스 모드)

각 Step은 Spec 기반 테스트 작성 -> 구현 -> 검증 -> 문서 동기화 순서로 수행한다. 검증 통과 시 다음 Step으로 이동한다. 저장 포맷, content schema, dependency 변경은 금지한다.
