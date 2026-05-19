# OPERATOR-GAP-001 — Operator Grammar Engine Gap

## 상태

- 상태: completed
- 시작일: 2026-05-19
- 완료일: 2026-05-19

## 배경

중반 첫 adventure intro는 `docs/gameplay/scenario-tone.md`의 터미널 문제 해결 생존 톤을 따른다. 다음 학습 축은 Vim의 실무 감각을 여는 `operator + motion` 문법이다.

## 결정

첫 operator grammar 범위는 아래로 제한한다.

- `delete-with-motion`: `dw`, `d$`, `dd`
- `change-with-motion`: `cw`, `c$`, `cc`

이번 milestone에서 제외한다.

- count prefix: `2dw`, `3dd`
- visual mode
- text object: `diw`, `ciw`, `ci"`, `di(`
- yank/put: `y`, `yy`, `p`, `P`
- multi-line complex operator: `dG`, `dgg`, range operator
- dot repeat

## 구현 순서

1. `VIM-016`: operator pending 기반 계약
   - `d`, `c`를 누르면 `PendingKey`에 operator를 저장한다.
   - `esc`는 pending operator를 취소한다.
   - unsupported operator combo는 상태 변경 없이 `EventUnsupportedKey`를 낸다.
   - `d`와 `c` key mapping을 TUI adapter에 추가한다.
2. `VIM-017`: `delete-with-motion`
   - `dw`: 현재 커서부터 다음 word boundary 전까지 삭제한다.
   - `d$`: 현재 커서부터 현재 줄 끝까지 삭제한다.
   - `dd`: 현재 줄 전체를 삭제하고 cursor를 다음 유효 줄로 clamp한다.
3. `VIM-018`: `change-with-motion`
   - `cw`: `dw` 범위를 삭제하고 Insert mode로 진입한다.
   - `c$`: `d$` 범위를 삭제하고 Insert mode로 진입한다.
   - `cc`: 현재 줄을 비우고 Insert mode로 진입한다.
4. `PLAYPACK-003`: 6문항 operator grammar adventure intro
   - `dw`, `d$`, `cw`, `c$`, `dd`, `u` 복습을 우선 후보로 둔다.

## Engine Gap

- 현재 `PendingKey`는 `gg`, `r{char}` 같은 짧은 prefix에 쓰인다.
- operator pending도 같은 필드를 재사용할 수 있다. 단 `d`와 `c`는 mutation operator이므로 실제 변경은 후속 motion key까지 기다려야 한다.
- current engine은 motion 함수가 cursor 이동만 반환한다. operator용으로는 “motion target/range” 계산이 필요하다.
- VIM-016에서는 range 삭제를 구현하지 않고 pending contract만 고정한다.
- VIM-017부터 delete/change helper를 추가한다.

## Runtime / Content Gap

- runtime은 key trace와 required key 검증을 이미 지원한다.
- content replay는 `["d", "w"]`, `["c", "w", "o", "p", "e", "n", "esc"]` 같은 trace를 처리할 수 있다.
- content loader는 command cluster 추가와 replay gate를 그대로 사용할 수 있다.
- PLAYPACK-003 전까지는 YAML content를 추가하지 않는다.

## TUI / E2E Gap

- TUI adapter는 아직 `d`, `c`를 명시적으로 Vim key로 매핑하지 않는다.
- VIM-016에서 adapter unit test로 `d`, `c` mapping을 고정한다.
- delete/change가 playable content로 연결되는 PLAYPACK-003에서 E2E를 추가한다.

## 검증

- `rg "OPERATOR-GAP-001|VIM-016|VIM-017|VIM-018|PLAYPACK-003" docs/roadmap docs/exec-plans docs/gameplay`: pass
- `go test ./...`: pass
- `git diff --check`: pass
