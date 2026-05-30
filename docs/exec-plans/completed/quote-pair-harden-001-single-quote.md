# QUOTE-PAIR-HARDEN-001 — Single Quote Text Object

Slice-ID: QUOTE-PAIR-HARDEN-001
Created: 2026-05-30
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md
- docs/verification/spec.md
- docs/exec-plans/active/quote-pair-harden-001-single-quote.md
- docs/exec-plans/completed/quote-pair-harden-001-single-quote.md
- content/command_clusters/text-object-quote-pair.yaml
- content/exercises/text-object-quote-pair.yaml
- content/scenarios/text-object-quote-pair.yaml
- content/playlists/text-object-quote-pair.yaml
- test/e2e/playable_text_object_quote_pair_full.yaml
- internal/vimengine/
- internal/runtime/
- internal/content/loader_test.go

## 목표

`text-object-quote-pair` hardening의 첫 scope로 single quote 내부 object `di'`, `ci'`, `yi'`를 지원한다. 기존 double quote object와 같은 semantics를 유지하되 delimiter만 `'`로 확장한다.

## 포함

- `KeySingleQuote`와 single quote inner text object dispatch
- same-line non-nested single quote 내부 range 탐색
- `di'`, `ci'`, `yi'` engine tests
- runtime replay test
- tutorial-91에 single quote 3문항 추가
- focused E2E 갱신
- docs/spec/roadmap 동기화

## 제외

- `i(`, `i{`
- nested pair
- escaped quote
- around object
- count/register prefix
- visual selection
- progress schema 변경

## 수용 기준

- `di'`는 같은 줄의 single quote 내부 값만 삭제하고 Normal mode를 유지한다.
- `ci'`는 같은 줄의 single quote 내부 값만 삭제하고 Insert mode로 진입한다.
- `yi'`는 같은 줄의 single quote 내부 값을 unnamed charwise register에 저장하고 buffer를 변경하지 않는다.
- single quote pair를 찾지 못하면 buffer/register를 변경하지 않는다.
- `ci'` 변경은 `esc` 이후 `.` 반복 대상으로 기록된다.
- tutorial-91은 8문항 이하를 유지하고 replay gate를 통과한다.
- focused E2E는 single quote 문항까지 포함해 final app_state/progress를 검증한다.
- `go test ./internal/vimengine ./internal/runtime ./internal/content`, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Scope Decision

- [x] SubAgent와 로컬 조사로 `i'`를 첫 scope로 결정한다.
- [x] `i(`, `i{`는 비대칭 pair/nesting 리스크 때문에 후속으로 분리한다.

## Step 2: Engine Red/Green

- [x] vimengine Red tests 추가
- [x] runtime replay Red test 추가
- [x] engine 구현

## Step 3: Content/E2E

- [x] command cluster coverage 갱신
- [x] tutorial exercise/scenario/playlist 추가
- [x] focused E2E 갱신

## Step 4: Validation

- [x] focused Go tests
- [x] focused E2E
- [x] full regression

## Step 5: Closeout

- [x] docs 동기화
- [x] ExecPlan completed 이동
- [x] 커밋/푸시
