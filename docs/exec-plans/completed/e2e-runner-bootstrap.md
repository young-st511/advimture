# ExecPlan: TUI E2E runner bootstrap

Slice-ID: E2E-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths: n/a

## 목표

Advimture를 pseudo terminal에서 실행하고 키 입력, 화면 텍스트, 종료 코드, progress 파일 상태를 검증하는 최소 E2E runner를 만든다. 웹 Playwright QA Loop처럼 Agent가 실패 evidence를 읽고 같은 scenario를 반복 실행할 수 있게 하는 것이 목적이다.

## 범위

- 포함: `cmd/e2e-runner`, `test/e2e` smoke scenario, verification docs, 실행 명령 문서
- 제외: full scenario DSL, CI 통합, screenshot 렌더링, 기존 게임 UX 수정

## 수용 기준

- E2E runner는 테스트 전용 HOME으로 앱을 실행하고 실제 `~/.advimture/progress.json`을 읽거나 쓰지 않는다.
- 각 scenario는 초기 조건, 키 입력 trace, 기대 screen assertion, 기대 exit/progress assertion을 YAML로 가진다.
- 실패 또는 성공 시 raw ANSI log, cleaned final screen, key trace를 `artifacts/e2e/` 아래에 남길 수 있다.
- smoke scenario는 FTUE 첫 화면을 확인하고 `ctrl+c`로 정상 종료되는지 검증한다.

## 검증 계획

- `go test ./cmd/e2e-runner`
- `go run ./cmd/e2e-runner --scenario test/e2e/ftue_ctrl_c_quit.yaml`
- `go test ./...`

## 의사결정 로그

- 2026-05-18: stdin/stdout pipe 방식은 Bubble Tea가 `/dev/tty`를 열지 못해 실패했다. 실제 TUI E2E에는 pty가 필요하므로 `github.com/creack/pty`를 추가한다.
- 2026-05-18: runner는 Go test 내부 helper가 아니라 별도 `cmd/e2e-runner`로 둔다. Agent가 특정 scenario를 반복 실행하기 쉽고, 나중에 CI smoke로도 승격할 수 있기 때문이다.

## 미해결 질문

- full E2E를 CI에 포함할지, 로컬/Agent 전용으로 둘지 결정 필요.
- terminal screen parser를 정교화할지, 당분간 ANSI strip + text assertion으로 갈지 추후 결정.

## 완료 결과

- `cmd/e2e-runner` pty 기반 smoke runner 추가
- `test/e2e/ftue_ctrl_c_quit.yaml` smoke scenario 추가
- `make e2e-smoke` 명령 추가
- verification docs와 guardrails에 실행 방법 반영
- 2026-05-18 LEGACY-001 이후 FTUE smoke는 `docs/archived/legacy-e2e/2026-05-18/`로 이동했고 현재 smoke는 `test/e2e/playable_hjkl_success.yaml`이다.
