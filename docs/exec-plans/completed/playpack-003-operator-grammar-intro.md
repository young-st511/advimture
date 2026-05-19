# PLAYPACK-003 — Operator Grammar Adventure Intro

Slice-ID: PLAYPACK-003
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/playpack-003-operator-grammar-intro.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- content/
- test/e2e/
- internal/content/
- internal/tuiadapter/

## 목표

`dw`, `d$`, `dd`, `cw`, `c$`, `cc`를 각각 한 문항으로 다루는 6문항 operator grammar 입문 tutorial을 만든다. 시나리오는 중반 생존 어드벤처 톤을 사용하되, 설계 순서는 `command -> exercise -> scenario`를 유지한다.

## 범위

- 포함:
  - `delete-with-motion`, `change-with-motion` content cluster를 implemented로 추가
  - 6문항 exercise YAML과 replay/e2e assertions
  - 6개 scenario YAML
  - `tutorial-5-operator-grammar` playlist
  - operator grammar full-play E2E scenario
  - insert mode에서 named `space` 입력을 printable space로 처리하는 TUI adapter 회귀 보강
- 제외:
  - yank/put, text object, count prefix
  - Neovim exact oracle hardening
  - 저장 포맷 변경

## 수용 기준

- 플레이리스트는 8문항 이하이며 6문항으로 구성된다.
- 각 문항은 `trained_commands`와 `optimal_keys`로 대상 command를 명확히 포함한다.
- `dw`, `d$`, `dd`, `cw`, `c$`, `cc`가 각각 coverage report에서 covered로 잡힌다.
- delete 문항은 `max_inputs: 2`로 operator sequence 우회를 줄인다.
- change 문항은 `operator 2키 + replacement + esc`의 exact key count를 `max_inputs`와 `optimal_key_count`에 반영한다.
- `cw` 문항은 현재 pedagogical semantics상 뒤 공백까지 삭제한다는 사실을 힌트/실패 피드백으로 드러낸다.
- TUI E2E는 6문항을 순서대로 완료하고 progress 저장, key trace, final app state를 검증한다.

## 구현 결과

- `content/command_clusters/operator-grammar.yaml`을 추가했다.
- `content/exercises/operator-grammar.yaml`에 6문항을 추가했다.
- `content/scenarios/operator-grammar.yaml`에 중반 생존 톤 scenario를 추가했다.
- `content/playlists/operator-grammar.yaml`에 `tutorial-5-operator-grammar`를 추가했다.
- `test/e2e/playable_operator_grammar_full.yaml`을 추가했다.
- `MapInputForMode("space", insert)`가 printable space를 반환하도록 보강했다.
- content loader 테스트의 root fixture counts, playable IDs, operator coverage를 갱신했다.

## 검증 결과

- `go test ./internal/content/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./internal/vimengine/...`: pass
- `go test ./...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_operator_grammar_full.yaml`: pass
- `make e2e-smoke`: pass
- `git diff --check`: pass

## E2E Evidence

- scenario: `test/e2e/playable_operator_grammar_full.yaml`
- artifact path: `artifacts/e2e/playable_operator_grammar_full/`
- 확인:
  - final buffer: `["job: precheck", "rsync --dry-run ok", "job: report"]`
  - final cursor: `1,17`
  - mode/status: `normal/succeeded`
  - progress mission: `change-with-motion-003`

## 의사결정 로그

- 2026-05-19: SubAgent 아이디어와 검증 체크리스트를 종합해 6문항을 command별 1문항으로 고정했다.
- 2026-05-19: `d$` 문항은 trailing space assertion을 피하기 위해 `tx: SOS###NOISE###` 형태로 설계했다.
- 2026-05-19: TUI E2E는 operator/pending 입력이 paste처럼 합쳐지지 않도록 키 사이에 50ms 간격을 둔다.
