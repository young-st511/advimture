# SUCCESS-STATUSLINE-CLEANUP-001 — Hide Stale Command After Success

Slice-ID: SUCCESS-STATUSLINE-CLEANUP-001
Created: 2026-06-08
Status: completed
Scope-Mode: normal

## 목표

성공/실패 floating modal이 표시된 뒤 status area에 이전 command/search 입력이 `Command: ...`로 남지 않게 한다. 입력 중인 command/search prompt와 app_state command evidence는 유지한다.

## 포함

- succeeded/failed 상태에서 화면의 `Command: <last>` 표시 숨김
- command/search mode의 입력 중 prompt 유지
- 관련 E2E 기대값 갱신
- TUI screen contract 문서 동기화

## 제외

- app_state `command` 제거
- progress schema 변경
- content schema 변경
- 새 dependency

## Step 1: 테스트 기대값

- 변경 파일:
  - `internal/playable/model_test.go`
  - `test/e2e/playable_substitute_current_line.yaml`
  - `test/e2e/playable_ftue_first_five_route.yaml`
  - `test/e2e/playable_full_first_five_minute.yaml`
  - `test/e2e/playable_command_quit.yaml`
- 상세 작업:
  - [x] succeeded 화면에 `Command:`가 남지 않는 unit expectation 추가
  - [x] E2E screen expectation을 성공 copy/action 중심으로 변경

## Step 2: 구현

- 변경 파일:
  - `internal/playable/model.go`
- 상세 작업:
  - [x] `ShowLastCommand`를 running 상태에서만 켠다.
  - [x] command/search mode의 `ShowCommandLine`은 유지한다.

## Step 3: 문서/검증

- 변경 파일:
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/exec-plans/active/success-statusline-cleanup-001.md`
  - `docs/exec-plans/completed/success-statusline-cleanup-001.md`
- 검증:
  - `go test ./internal/playable ./internal/playableview`
  - focused E2E: command quit, substitute, FTUE route, full first five
  - `go test ./...`
  - `git diff --check`

## 완료 결과

- succeeded/failed floating modal 이후 화면에서 이전 `Command: ...` 표시를 숨겼다.
- command/search mode에서 입력 중인 prompt는 유지했다.
- `app_state.command` evidence는 유지했다.
- Fresh playtest에서 보였던 incident 001/008 성공 화면의 stale command line이 사라졌다.
