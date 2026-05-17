# Verification Spec

## 개요

Advimture의 테스트와 TUI QA 루프를 정의한다. 웹 Playwright처럼 Agent가 실행-관찰-수정 반복을 돌 수 있게 하되, 대상은 브라우저가 아니라 terminal UI다.

---

## 현재 동작

### Go 테스트

**소스**: `go.mod`, `internal/**/*_test.go`

- 전체 테스트 명령은 `go test ./...`이다.
- 현재 테스트가 있는 패키지는 `internal/editor`, `internal/game`, `internal/data`, `internal/progress`다.

### TUI E2E Runner

**소스**: `cmd/e2e-runner/main.go`, `test/e2e/ftue_ctrl_c_quit.yaml`, `Makefile`

- E2E runner는 별도 CLI `cmd/e2e-runner`로 제공한다.
- smoke 실행 명령은 `go run ./cmd/e2e-runner --scenario test/e2e/ftue_ctrl_c_quit.yaml` 또는 `make e2e-smoke`이다.
- runner는 scenario YAML의 `setup.home: temp`일 때 테스트 전용 HOME을 만들고 앱을 실행한다.
- runner는 pseudo terminal로 앱을 실행하기 위해 `github.com/creack/pty`를 사용한다.
- runner는 키 입력 trace를 전송하고, cleaned screen text, exit code, progress file existence를 검증한다.
- runner는 raw ANSI log, cleaned final screen, key trace를 `artifacts/e2e/{scenario_id}/` 아래에 저장할 수 있다.
- `ftue_ctrl_c_quit` smoke scenario는 FTUE 첫 화면이 보이는지 확인한 뒤 `ctrl+c`로 종료하고 exit code 0과 progress 파일 미생성을 검증한다.

## 미확인 사항

- [ ] Bubble Tea model-level integration test와 pty-level E2E의 경계를 정해야 한다.
- [ ] full E2E를 CI에 포함할지, 로컬/Agent 전용으로 둘지 결정해야 한다.
- [ ] ANSI strip 기반 text assertion에서 정교한 terminal screen parser로 승격할 시점을 결정해야 한다.

---

## 수용 기준

> 이 도메인은 **모드 B**로 운영한다. Agent가 의도를 받아 `[draft]` 초안을 작성하고, 사람이 승인하면 `[draft]`를 제거한다.

### TUI E2E Loop 도입 준비

구현 완료. 현재 동작으로 이동했다.
