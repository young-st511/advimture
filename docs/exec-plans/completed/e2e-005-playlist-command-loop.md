# ExecPlan: Playlist and command-mode E2E hardening

Slice-ID: E2E-005
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- cmd/e2e-runner/**
- test/e2e/**
- Makefile
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/e2e-005-playlist-command-loop.md
- docs/exec-plans/completed/e2e-005-playlist-command-loop.md

## 목표

GAMELOOP-001 이후 넓어진 TUI 경로를 E2E에서 직접 검증한다. 첫 문제 smoke만 보던 상태에서 playlist next, progress resume, command-line mission, substitute mission까지 pty 기반으로 확인한다.

## 영향 도메인

- Verification: E2E scenario setup에서 테스트 전용 HOME에 progress fixture를 심을 수 있어야 한다.
- Gameplay: 실제 TUI 입력으로 `enter` next, `:q!`, `:s/old/new/`가 동작해야 한다.
- Safety: 실제 `~/.advimture`는 계속 사용하지 않는다.

## 수용 기준

- E2E runner는 `setup.home: temp`와 `setup.progress_file` 조합으로 테스트 progress를 생성한다.
- E2E는 첫 문제 성공 후 `enter`로 다음 survival exercise에 진입하고 `esc` 성공을 검증한다.
- E2E는 progress fixture가 있을 때 첫 미완료 exercise에서 시작하는 resume 동작을 검증한다.
- E2E는 `:q!` command-line exercise를 실제 TUI 입력으로 성공시킨다.
- E2E는 `:s/api/web/` substitute exercise를 실제 TUI 입력으로 성공시킨다.
- `make e2e-playable`은 보강된 playable E2E suite를 실행한다.

## 범위

- 포함: runner setup fixture
- 포함: E2E scenario YAML 추가
- 포함: Makefile suite target 보강
- 제외: CI workflow 추가
- 제외: screenshot/image 기반 assertion
- 제외: full 14-exercise complete run

## 검증 계획

- `go test ./cmd/e2e-runner`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] progress fixture setup Red test를 작성한다.
- [x] runner setup에 progress fixture write를 추가한다.
- [x] playlist next/resume/command/substitute E2E scenarios를 추가한다.
- [x] Makefile E2E suite를 보강한다.
- [x] 반복 실행으로 flake와 assertion 누락을 보강한다.
- [x] 문서를 completed 상태로 동기화한다.

## 발견 및 보강

- E2E가 첫 문제 이후 다음 문제를 저장할 때 이전 완료 progress가 사라지는 버그를 잡았다. 원인은 playable model의 progress 갱신 receiver가 값 receiver였기 때문이며, pointer receiver로 고정했다.
- command mode 문자 입력은 pty 환경에서 짧은 안정 대기가 필요했다. command-mode scenario는 화면 도달 대기와 최소 wait를 섞어 flake를 줄인다.
