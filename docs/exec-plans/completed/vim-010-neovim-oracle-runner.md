# ExecPlan: Neovim oracle runner

Slice-ID: VIM-010
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-010-neovim-oracle-runner.md
- internal/vimoracle/**

## 목표

`exact` compatibility tier command를 검증할 수 있도록, Advimture engine 결과와 외부 oracle 결과를 비교하는 oracle runner 계약을 만든다. Neovim은 런타임이 아니며 optional verification 도구로만 남긴다.

## 범위

- 포함: `internal/vimoracle` 패키지 생성
- 포함: test case, oracle executor interface, comparison result
- 포함: engine key replay와 oracle state 비교
- 포함: oracle error propagation
- 제외: 게임 런타임에 Neovim 연결
- 제외: CI에서 Neovim 설치 요구
- 제외: 모든 Vim command oracle fixture 작성

## 검증 계획

- `go test ./internal/vimoracle`
- `go test ./...`

## 승인 체크

- [x] engine replay 결과와 oracle 결과를 비교한다.
- [x] buffer, cursor, mode mismatch를 보고한다.
- [x] oracle error를 실패로 전파한다.
- [x] Neovim 설치 없이 unit test가 통과한다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: 실제 Neovim CLI 실행기는 이번 slice에서 제외했다. 기본 테스트와 게임 런타임이 Neovim 설치 여부에 의존하면 안 되기 때문이다.
- 2026-05-18: executor interface를 먼저 고정했다. 이후 별도 verification command에서 Neovim headless executor를 붙일 수 있다.
