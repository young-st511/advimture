# 0002 — Vim Engine Strategy

Status: accepted

## Context

Advimture의 주 목적은 Vim 학습이다. 기존 구현은 Vim 동작이 안정적이지 않고, 게임 진행이 중간에 끝나지 않는 문제가 있었다. 새 플랫폼은 Vim을 실제로 유용하게 쓰는 데 필요한 command를 반복 훈련시키되, 게임 진행, 힌트, 채점, 종료 상태를 안정적으로 제어해야 한다.

검토한 선택지는 다음과 같다.

- Vim C 코드 wrapping/cgo
- Neovim headless/embed RPC를 게임 런타임으로 사용
- 작은 Go Vim-like engine을 직접 구현
- 작은 Go Vim-like engine + Neovim oracle test

## Decision

Advimture는 **작은 Go Vim-like engine을 런타임으로 직접 구현하고, Neovim은 optional oracle test로만 사용한다.**

런타임 구조는 다음 방향을 따른다.

```text
Command Catalog
  -> Exercise Bank
  -> Scenario Bank
  -> Content Validator
  -> Mission Runtime
  -> Small Go Vim Engine
  -> Bubble Tea TUI
  -> Progress Store

Compatibility Tests
  -> Neovim Oracle
```

첫 구현 범위는 command cluster 단위로 제한한다. Vim 전체 구현을 목표로 하지 않는다.

## Consequences

- `go run .`, 기본 단위 테스트, smoke E2E는 Neovim 설치 여부와 무관하게 동작한다.
- Neovim oracle test는 별도 명령 또는 build tag로 분리한다.
- `internal/vimengine`은 순수 상태 전이로 설계한다.
- `internal/runtime`은 mission state, scoring, hint, retry, exit를 책임진다.
- Bubble Tea는 input/output adapter로 유지한다.
- TUI E2E는 Vim semantics를 검증하지 않는다. Vim semantics는 engine unit test, exercise grader, optional Neovim oracle에서 검증한다.

## Design Rules

1. **Oracle is not runtime**
   - Neovim은 검증 도구일 뿐 게임 실행 경로에 들어가지 않는다.

2. **Command cluster scope**
   - 새 엔진 기능은 승인된 command cluster나 active ExecPlan에 연결된 경우에만 추가한다.

3. **Compatibility tier**
   - 각 command cluster는 `exact`, `pedagogical`, `unsupported` 중 하나의 compatibility tier를 가진다.
   - `exact`: Neovim과 결과 일치가 목표다.
   - `pedagogical`: 학습 목적상 단순화를 허용하지만 차이를 문서화한다.
   - `unsupported`: 입력을 차단하고 플레이어에게 설명한다.

4. **Pure transition**
   - 엔진은 `State + Key -> State + Events` 형태로 동작한다.
   - 터미널, 파일 시스템, Bubble Tea, progress 저장에 의존하지 않는다.

5. **Inspectable state**
   - exercise는 최소한 buffer, cursor, mode, exit/submission status, key trace를 기계적으로 검증할 수 있어야 한다.

6. **E2E stop rule**
   - E2E 루프가 부족해서 성공 여부를 신뢰할 수 없으면 구현을 중단하고 E2E evidence/assertion을 먼저 보강한다.

## Initial MVP

첫 플랫폼 MVP는 하나의 완결 미션으로 제한한다.

- command cluster 1개
- exercise 1개
- scenario 1개
- 성공/실패/힌트/재시도/종료/progress 포함
- 이 한 사이클이 안정화되기 전에는 command 범위를 늘리지 않는다.
