# Verification Spec

## 개요

Advimture의 테스트와 TUI QA 루프를 정의한다. 웹 Playwright처럼 Agent가 실행-관찰-수정 반복을 돌 수 있게 하되, 대상은 브라우저가 아니라 terminal UI다.

---

## 현재 동작

### Go 테스트

**소스**: `go.mod`, `internal/**/*_test.go`

- 전체 테스트 명령은 `go test ./...`이다.
- 현재 테스트가 있는 주요 패키지는 `internal/vimengine`, `internal/runtime`, `internal/content`, `internal/scoring`, `internal/scenario`, `internal/tuiadapter`, `internal/progressadapter`, `internal/vimoracle`, `internal/playable`, `internal/progress`다.

### TUI E2E Runner

**소스**: `cmd/e2e-runner/main.go`, `test/e2e/*.yaml`, `Makefile`

- E2E runner는 별도 CLI `cmd/e2e-runner`로 제공한다.
- smoke 실행 명령은 `make e2e-smoke`이다.
- 공개 전 로컬 release gate는 `make release-check`이며 `make test`, `make build`, `make e2e-playable`을 순서대로 실행한다.
- Makefile은 runner 실행 시 `artifacts/go-build-cache`를 `GOCACHE`로 사용해 사용자 Go build cache에 쓰지 않는다.
- runner는 scenario YAML의 `setup.home: temp`일 때 테스트 전용 HOME을 만들고 앱을 실행한다.
- runner는 `setup.progress_file` JSON fixture로 테스트 전용 HOME에 `.advimture/progress.json`을 미리 생성할 수 있다.
- runner는 `setup.complete_before: <exercise-id>`로 현재 content의 playable 순서를 읽고 지정 exercise 직전까지 completed progress를 생성할 수 있다. `setup.progress_file`과 `setup.complete_before`는 동시에 사용할 수 없다.
- runner는 pseudo terminal로 앱을 실행하기 위해 `github.com/creack/pty`를 사용한다.
- runner는 키 입력 trace를 전송하고, cleaned screen text, exit code, progress file existence/content, key trace exact match, app state summary를 검증한다.
- visual mode selection assertion은 `docs/verification/selection-app-state-contract.md`의 `selection` object를 기준으로 한다. `internal/e2estate.State`, runner `assert.app_state.selection`, content `e2e_assertions.selection`은 같은 shape를 사용한다.
- review/daily route assertion은 `assert.app_state.review`로 검증한다. `queue_count`, `primary_exercise_id`, `primary_reason`, `daily_route`는 화면 문구 이동과 별개로 stable state로 본다.
- focus panel assertion은 `assert.app_state.ui.focus_panel`로 검증한다. `kind`, `title`, `lines`는 화면 배치가 바뀌어도 현재 UI intent를 stable state로 본다.
- mission/review loop E2E는 success debrief의 `이번 복구`, `최단 복구`, `목표 입력`, `다음 출격` 문구와 `ui.focus_panel.kind/title`을 함께 검증한다.
- command memory UI E2E는 tutorial `기억할 명령`, duplicate coach suppression, incident hint/failure 후 `참고 명령`, 긴 incident hint cue wrapping을 화면과 app_state evidence로 검증한다.
- viewport smoke E2E는 80x24 success/failure floating modal에서 action line과 `ui.focus_panel`이 유지되는지 검증한다.
- runner는 실제 사용자 HOME과 기존 progress file이 보이는 HOME을 기본적으로 거부한다.
- runner는 `summary.json`, raw ANSI log, cleaned screen stream, cleaned final screen, key trace, app state snapshot, progress snapshot을 `artifacts/e2e/{scenario_id}/` 아래에 저장할 수 있다.
- `playable_hjkl_success` smoke scenario는 `l`, `l`, `q` 입력으로 첫 playable exercise를 성공시키고, screen text, progress 파일, key trace, app state summary를 검증한다.
- `playable_ftue_first_five_route` scenario는 FTUE-001 canonical route를 `tutorial-0-movement`부터 `tutorial-3-small-edits` 첫 문항까지 검증하고, screen timeline/final screen/app_state/progress evidence를 남긴다.
- 긴 incident full route, command-choice applied route, 대표 mid tutorial full route는 route 전체 흐름과 마지막 화면을 재검토할 수 있도록 screen timeline/final screen/app_state evidence를 함께 남긴다.
- command-choice applied route는 final/timeline evidence와 app_state evidence를 함께 남긴다.
- `make e2e-playable`은 첫 문제 smoke, constraint/coaching UX, 80x24 success/failure viewport smoke, 80x30 incident hint wrapping smoke, playlist next, progress resume, `:q!` command-line, `:s/api/web/` substitute, first-cut tutorial flow, operator/yank/text-object/open-line/repeat/search/quote text object 7-beat full playpack, incident 001~004, command-choice 5-beat route, incident 006/007 run 경로를 모두 검증한다.
- E2E runner의 `wait_screen_contains`는 이전 wait 이후 새로 출력된 화면만 대상으로 삼아, 반복되는 `Next: enter` 같은 문자열이 과거 화면 로그 때문에 오탐하지 않게 한다.

## 미확인 사항

- [ ] Bubble Tea model-level integration test와 pty-level E2E의 경계를 정해야 한다.
- [ ] full E2E를 CI에 포함할지, 로컬/Agent 전용으로 둘지 결정해야 한다.
- [ ] ANSI strip 기반 text assertion에서 정교한 terminal screen parser로 승격할 시점을 결정해야 한다.
- [x] E2E runner는 app state summary 기반 buffer/cursor/mode/status/score/progress assertion을 지원한다.
- [x] 첫 playable slice에서 실제 앱이 `ADVIMTURE_E2E=1`일 때 app state summary를 쓴다.
- [x] progress fixture 기반 resume E2E와 command-mode E2E를 지원한다.
- [x] first-cut tutorial flow와 최신 full playpack들을 pty E2E로 처음부터 끝까지 검증한다.

---

## 수용 기준

> 이 도메인은 **모드 B**로 운영한다. Agent가 의도를 받아 `[draft]` 초안을 작성하고, 사람이 승인하면 `[draft]`를 제거한다.

### TUI E2E Loop 도입 준비

구현 완료. 현재 동작으로 이동했다.
