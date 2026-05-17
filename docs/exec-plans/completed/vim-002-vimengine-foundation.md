# ExecPlan: Vim engine foundation

Slice-ID: VIM-002
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/decisions/0002-vim-engine-strategy.md
- docs/gameplay/command-catalog.md
- docs/gameplay/spec.md
- docs/exec-plans/active/vim-002-vimengine-foundation.md
- internal/vimengine/**

## 목표

Advimture의 새 Vim 학습 플랫폼이 의존할 최소 Vim engine을 만든다. 이번 slice는 게임 화면이나 기존 `internal/editor` 교체가 아니라, `State + Key -> State + Events` 형태의 순수 상태 전이를 세우는 첫 루프다.

## 범위

- 포함: `internal/vimengine` 패키지 생성
- 포함: Normal mode의 `h`, `j`, `k`, `l` 이동
- 포함: buffer, cursor, mode를 검증 가능한 상태로 노출
- 포함: 줄/열 경계와 짧은 줄 이동의 기본 처리
- 제외: 기존 `internal/editor` 삭제 또는 교체
- 제외: Bubble Tea 화면 연결
- 제외: 저장/종료 command와 progress 저장
- 제외: Neovim oracle runner 구현

## 연결된 Command Cluster

- `normal-motion-basic`
  - status: draft
  - compatibility_tier: exact
  - commands: `h`, `j`, `k`, `l`

## 구현 계획

1. `internal/vimengine`에 순수 engine 타입을 만든다.
2. engine 상태는 최소 `mode`, `lines`, `cursor`를 가진다.
3. 키 입력 결과는 변경된 상태와 event 목록으로 반환한다.
4. `h`, `j`, `k`, `l`은 Normal mode에서만 처리한다.
5. 지원하지 않는 키는 상태를 바꾸지 않고 event로 보고한다.

## 검증 계획

- `go test ./internal/vimengine`
- 공통 회귀 확인이 필요하면 `go test ./...`
- 변경 종료 전 `git diff -- docs/gameplay/command-catalog.md docs/roadmap/PROGRAM.md docs/gameplay/spec.md docs/exec-plans/active/vim-002-vimengine-foundation.md internal/vimengine`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다. 단, engine 결과를 게임 화면에 연결하는 다음 slice에서는 E2E schema가 부족하면 구현을 멈추고 assertion을 먼저 보강한다.

## 승인 체크

- [x] engine이 Bubble Tea, 파일 시스템, progress 저장에 의존하지 않는다.
- [x] `h`, `j`, `k`, `l`의 정상 이동과 경계 처리가 테스트된다.
- [x] 상태 복사 또는 불변성 경계가 테스트된다.
- [x] unsupported key가 테스트된다.
- [x] 다음 slice에서 exercise grader가 참조할 수 있는 상태 형태가 있다.

## 의사결정 로그

- 2026-05-18: 기존 `internal/editor`를 수정하지 않고 새 `internal/vimengine` 패키지를 추가했다. 기존 구현의 동작 회귀와 새 설계 검증을 분리하기 위함이다.
- 2026-05-18: 이번 slice에서는 TUI E2E를 실행하지 않았다. engine이 아직 Bubble Tea runtime에 연결되지 않았고, Vim semantics는 unit test가 더 직접적인 evidence다.
