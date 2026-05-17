# ExecPlan: TUI adapter

Slice-ID: VIM-008
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-008-tui-adapter.md
- internal/tuiadapter/**

## 목표

Scenario Runtime을 실제 TUI에 연결하기 전에, 입력과 출력의 adapter 계약을 만든다. 이번 slice는 Bubble Tea app wiring이 아니라, key/action mapping과 render view model 변환을 담당하는 순수 adapter다.

## 범위

- 포함: `internal/tuiadapter` 패키지 생성
- 포함: raw input을 scenario action으로 변환
- 포함: `h/j/k/l`, arrow key를 Vim key로 매핑
- 포함: hint/retry/quit action
- 포함: scenario state를 view model로 변환
- 제외: Bubble Tea model 수정
- 제외: pty E2E 추가
- 제외: progress 저장

## E2E Stop Rule

이번 slice는 실제 Bubble Tea runtime에 연결하지 않는다. 실제 app wiring을 시작하는 순간 현재 smoke E2E가 부족하면 구현을 중단하고 E2E assertion을 먼저 보강한다.

## 검증 계획

- `go test ./internal/tuiadapter`
- `go test ./...`

## 승인 체크

- [x] Vim 이동 키와 arrow key가 같은 key action으로 변환된다.
- [x] hint/retry/quit이 command action으로 변환된다.
- [x] unknown input은 ignored action이 된다.
- [x] scenario state가 stable view model로 변환된다.
- [x] Bubble Tea, filesystem, progress에 의존하지 않는다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: 실제 Bubble Tea app wiring은 이번 slice에서 제외했다. TUI에 연결하는 순간 현재 smoke E2E만으로는 성공/실패 판정이 부족하므로, 먼저 adapter 계약만 단위 테스트로 고정했다.
- 2026-05-18: `r`은 retry command로 예약했다. 이후 실제 Vim `r` command를 가르치는 cluster를 추가할 때는 context-sensitive mapping이나 command palette로 분리해야 한다.
