# ExecPlan: Line and file navigation motion

Slice-ID: NAV-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/vimengine/**
- internal/vimoracle/**
- internal/tuiadapter/**
- internal/content/**
- content/command_clusters/**
- content/exercises/**
- content/scenarios/**
- content/playlists/**
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/active/nav-001-line-file-motion.md
- docs/exec-plans/completed/nav-001-line-file-motion.md

## 목표

후반 navigation의 첫 단위로 `gg`, `G`, `0`, `$`를 엔진과 콘텐츠에 연결한다. 짧은 `hjkl` 반복 대신 줄/파일 단위로 목적지에 도달하는 학습 효율을 보여준다.

## 영향 도메인

- Gameplay: line/file navigation은 word motion 이후의 효율 이동 cluster로 다룬다.
- Vim engine: normal mode에서 file start/end와 line start/end 이동을 순수 전이로 구현한다.
- Verification: `gg` prefix 입력, `G` 대문자 입력, line end boundary를 단위 테스트와 replay gate로 고정한다.

## 수용 기준

- `0`은 현재 줄의 첫 column으로 이동하고, `$`는 현재 줄의 마지막 column으로 이동한다.
- `gg`는 현재 buffer의 첫 줄 첫 column으로 이동하고, `G`는 마지막 줄 첫 column으로 이동한다.
- 첫 `g` 입력은 prefix 상태로 남고, 두 번째 `g`에서만 이동한다.
- `gg`, `G`, `0`, `$`는 approved + implemented exercise coverage와 replay gate를 통과한다.
- TUI 입력 매핑은 대문자 `G`와 소문자 `g`를 구분한다.
- 기존 playable smoke는 여전히 첫 normal motion 문제로 통과한다.

## 범위

- 포함: vimengine key constants와 motion implementation
- 포함: tuiadapter input mapping
- 포함: command cluster, exercise, scenario, playlist beat
- 포함: content replay/coverage test 보강
- 제외: numeric count prefix (`10G`, `3gg`)
- 제외: `^`, `%`, mark jump, search jump
- 제외: playlist 기반 multi-exercise playable selection

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/vimoracle/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/content/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] `0`, `$`, `gg`, `G` Red test를 작성한다.
- [x] vimengine navigation motion과 `g` prefix 상태를 구현한다.
- [x] tuiadapter에서 `g`, `G`, `0`, `$` 입력을 매핑한다.
- [x] command catalog와 YAML exercise/scenario/playlist를 추가한다.
- [x] content coverage/replay 검증을 통과시킨다.
- [x] 문서를 completed 상태로 동기화한다.
