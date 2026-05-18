# ExecPlan: Ex substitute and range command foundation

Slice-ID: EXCMD-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/vimengine/**
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
- docs/exec-plans/active/excmd-001-substitute-range.md
- docs/exec-plans/completed/excmd-001-substitute-range.md

## 목표

command-line engine 기반 위에 literal substitute와 range command의 첫 slice를 구현한다. `:` 명령어를 편집 engine state 변화로 다루되, scenario success와 app quit semantics는 분리한다.

## 영향 도메인

- Vim engine: command-line 실행이 buffer를 변경할 수 있다.
- TUI adapter: command mode에서 substitute 입력에 필요한 문자들을 Vim key로 전달한다.
- Content: substitute exercise는 buffer target과 replay gate로 검증한다.

## 수용 기준

- `:s/old/new/`는 현재 줄의 첫 번째 literal match만 바꾼다.
- `:%s/old/new/g`는 전체 buffer의 모든 literal match를 바꾼다.
- `:2,3s/old/new/`는 1-based inclusive range 안의 각 줄에서 첫 번째 literal match만 바꾼다.
- 실행된 Ex command는 `last_command`에 기록된다.
- command mode의 `q`는 계속 app quit이 아니며, 일반 command 문자 입력도 app quit이 아니다.
- substitute exercise는 approved + implemented coverage와 replay gate를 통과한다.

## 범위

- 포함: literal substitute parser
- 포함: `%` all range와 numeric line range
- 포함: current line substitute
- 제외: Vim regex 호환성
- 제외: backreference, flags except `g`
- 제외: interactive confirm flag
- 제외: write/read/file IO Ex command

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
- `go test ./internal/content/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] substitute command Red test를 작성한다.
- [x] command parser와 buffer mutation을 구현한다.
- [x] command mode 입력 매핑을 일반 문자로 확장한다.
- [x] command catalog와 YAML exercise/scenario/playlist를 추가한다.
- [x] content coverage/replay 검증을 통과시킨다.
- [x] 문서를 completed 상태로 동기화한다.
