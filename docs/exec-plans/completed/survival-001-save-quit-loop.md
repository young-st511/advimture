# ExecPlan: Survival save and quit loop

Slice-ID: SURVIVAL-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/vimengine/**
- internal/runtime/**
- internal/content/**
- internal/tuiadapter/**
- internal/playable/**
- internal/e2estate/**
- content/command_clusters/**
- content/exercises/**
- content/scenarios/**
- content/playlists/**
- test/e2e/**
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/active/survival-001-save-quit-loop.md
- docs/exec-plans/completed/survival-001-save-quit-loop.md

## 목표

Vim 초심자의 핵심 생존 루프인 `esc`, `:q!`, `:wq`를 엔진/런타임/content에 연결한다. 앱 종료 단축키 `q`와 Vim command-line 입력의 `q`를 혼동하지 않도록 command mode 경계를 테스트로 고정한다.

## 영향 도메인

- Gameplay: `:q!`, `:wq`는 실제 파일 저장/종료가 아니라 exercise 목표 command로 다룬다.
- Vim engine: command-line mode와 last command state를 순수 전이로 유지한다.
- Verification: app quit과 mission success를 분리해서 검증한다.

## 수용 기준

- `esc`는 insert/command mode에서 Normal mode로 돌아온다.
- `:`는 command-line mode로 진입하고, command mode에서 `q`, `w`, `!`, `enter`는 Vim key 입력으로 처리된다.
- `:q!`, `:wq`는 실행 후 `last_command`로 기록되어 runtime goal로 검증할 수 있다.
- Normal mode의 `q` 앱 종료 단축키는 유지하되, command mode의 `q`는 app quit이 아니다.
- survival exercise는 `esc`, `:q!`, `:wq` coverage를 모두 채우고 replay gate를 통과한다.
- 기존 playable smoke는 여전히 첫 normal motion 문제로 통과한다.

## 범위

- 포함: vimengine command mode와 command execution state
- 포함: runtime/content goal의 command assertion
- 포함: tui input mapping의 mode-aware command handling
- 포함: survival YAML exercise/scenario/playlist beat
- 제외: 실제 파일 write/discard semantics
- 제외: multi-exercise game loop
- 제외: EX command parser 일반화

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/runtime/...`
- `go test ./internal/content/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] command mode와 command goal Red test를 작성한다.
- [x] vimengine command-line state와 `:q!`, `:wq` execution을 구현한다.
- [x] runtime/content goal에 command assertion을 추가한다.
- [x] tui input mapping을 mode-aware로 바꾼다.
- [x] survival exercise/scenario/playlist content를 추가한다.
- [x] replay/content/playable/E2E 검증을 통과시킨다.
- [x] 문서를 completed 상태로 동기화한다.

## 결정 기록

- `:q!`, `:wq`는 이번 slice에서 실제 파일 저장/폐기 semantics를 수행하지 않는다. Vim command-line 입력을 실행하면 `last_command`에 기록하고 runtime command goal로 검증한다.
- survival command coverage는 conceptual command인 `esc`, `:q!`, `:wq`를 `trained_commands`로 확인하고, replay gate는 실제 key trace인 `:`, `q`, `!`, `enter` / `:`, `w`, `q`, `enter`를 재생한다.
- 현재 playable/E2E smoke는 첫 playable 문제를 고정 실행한다. survival command는 unit test, content replay, loader gate, app state command assertion으로 검증하고, 전체 playlist 선택/진행 E2E는 `GAMELOOP-001`에서 다룬다.
