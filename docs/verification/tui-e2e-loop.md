# TUI E2E QA Loop

## 결론

가능하다. Playwright가 브라우저를 열고 DOM/screenshot을 검사하듯이, Advimture는 pseudo terminal을 열고 키 입력, terminal screen, progress state를 검사하면 된다.

## 권장 구조

```text
E2E scenario
  ├─ setup: 임시 HOME, progress fixture, terminal size
  ├─ launch: go run . 또는 빌드된 advimture binary
  ├─ drive: 키 입력 trace 전송
  ├─ observe: ANSI screen buffer 수집
  ├─ assert: 화면 텍스트 + 종료 코드 + key trace + progress 파일 확인
  └─ evidence: summary JSON, raw log, cleaned screen, trace 저장
```

## Agent Loop

1. E2E scenario를 실행한다.
2. 실패하면 `artifacts/e2e/<scenario>/`의 evidence를 읽는다.
3. 실패 원인을 spec 또는 구현 중 하나로 분류한다.
4. 구현 문제면 작은 patch를 만든다.
5. 같은 scenario를 다시 실행한다.
6. 통과하면 관련 Go 패키지 테스트와 `go test ./...`를 실행한다.

## Playwright와 다른 점

| 웹 Playwright | TUI Advimture |
|---------------|---------------|
| Browser page | pseudo terminal |
| DOM selector | screen text / cursor / ANSI buffer |
| click/type | key trace |
| screenshot | cleaned screen snapshot / summary JSON |
| localStorage/cookies | temporary progress file |
| network wait | prompt/state text wait |

## 구현 후보

### 1. Go test 내부 pty 러너

- `go test` 안에서 binary를 띄우고 pty로 키를 주입한다.
- 장점: Go 테스트와 CI 통합이 쉽다.
- 단점: pty/screen parser 의존성을 검토해야 한다.

### 2. 별도 `cmd/e2e-runner` — 현재 채택

- scenario YAML/JSON을 읽어 TUI를 실행한다.
- 장점: Agent가 특정 scenario만 반복 실행하기 쉽다.
- 단점: 별도 CLI 유지보수가 필요하다.

### 3. Bubble Tea model-level integration test

- `tea.KeyMsg`, `tea.WindowSizeMsg`를 직접 `Update`에 넣는다.
- 장점: 빠르고 안정적이다.
- 단점: 실제 terminal rendering, alt screen, ANSI 문제를 놓칠 수 있다.

현재는 **pty smoke E2E**를 먼저 세팅했다. 이후 권장 순서는 **model-level integration test 보강 → full scenario E2E 확장**이다.

## Scenario 포맷 초안

```yaml
id: ftue_quit_success
terminal:
  width: 100
  height: 30
setup:
  home: temp
  progress_file: |
    {"missions":{}}
steps:
  - send: "esc"
  - send: ":"
  - send: "q"
  - send: "!"
  - send: "enter"
assert:
  screen_contains:
    - "Tutorial"
  progress:
    file_exists: true
  key_trace:
    - "esc"
    - ":"
    - "q"
    - "!"
    - "enter"
evidence:
  save_raw_ansi: true
  save_clean_screen: true
  save_key_trace: true
  save_summary: true
```

## 현재 Assertion

- `screen_contains`: cleaned terminal text 포함 여부
- `exit_code`: process exit code
- `progress_file_exists`: test HOME의 `.advimture/progress.json` 존재 여부
- `key_trace`: runner가 실제 전송한 key trace와 기대 trace의 exact match
- `progress_file_contains`: progress JSON 텍스트에 포함되어야 하는 문자열
- `app_state`: test HOME 내부 app state summary JSON 검증
- `app_state.selection`: visual mode의 selection active/kind/anchor/head/start/end 검증
- `app_state.review`: review queue count, primary exercise/reason, daily route 문구 검증
- `setup.progress_file`: test HOME의 `.advimture/progress.json`에 미리 쓸 JSON fixture
- `playable_small_edits_full`: 7개 small edits exercise와 `ctrl+r` 입력을 검증하는 집중 E2E
- `playable_arrow_forbidden`: 실제 방향키 escape sequence가 `h/j/k/l`로 우회되지 않고 forbidden input으로 실패하는지 검증하는 E2E
- `playable_ctrl_c_does_not_quit`: 실제 Ctrl+C가 앱을 즉시 종료하지 않고 key trace에 남는지 검증하는 E2E
- `playable_coaching_panel`: strict constraint 문항의 required key coaching과 `?` hint 패널 표시를 검증하는 E2E
- `playable_full_first_five_minute`: movement/survival/navigation/small-edit/Ex command까지의 first-cut tutorial flow E2E
- `playable_open_line_full`, `playable_repeat_last_change_full`, `playable_search_basic_full`, `playable_text_object_quote_pair_full`: utility/structure command full playpack E2E
- `playable_visual_selection_full`: charwise visual delete/yank-put/backward selection tutorial E2E
- `playable_incident_001_full`, `playable_incident_002_full`: 기존 command를 섞은 incident run E2E

`wait_screen_contains`는 이전 wait 이후 새로 출력된 화면만 기다린다. 반복되는 UI 문구를 사용할 수는 있지만, 같은 키를 빠르게 반복 입력할 때는 Bubble Tea가 입력을 묶어 받을 수 있으므로 사람 입력에 가까운 짧은 `wait_ms`를 둔다.

### `assert.app_state`

기본 파일 위치는 test HOME 기준 `.advimture/e2e_state.json`이다. 실제 앱은 첫 playable slice부터 `ADVIMTURE_E2E=1`일 때 이 파일을 쓸 수 있다.

```yaml
assert:
  app_state:
    buffer: ["abc"]
    cursor:
      row: 0
      col: 2
    mode: "normal"
    status: "succeeded"
    score:
      grade: "S"
      passed: true
    progress:
      mission_id: "normal-motion-basic-001"
      completed: true
    review:
      queue_count: 3
      primary_exercise_id: normal-motion-basic-002
      primary_reason: incomplete
      daily_route: "오늘의 복구 루트: 3건 대기"
    selection:
      active: true
      kind: "charwise"
      anchor: {row: 0, col: 0}
      head: {row: 0, col: 2}
      start: {row: 0, col: 0}
      end: {row: 0, col: 2}
```

Content replay gate도 `e2e_assertions.selection`이 선언된 exercise에 한해 optimal key replay 결과의 selection을 비교한다. 따라서 visual 관련 content는 화면 문구만으로 통과하지 않고, content load 단계와 TUI E2E 단계에서 같은 selection shape를 검증할 수 있다.

## 현재 Evidence

- `summary.json`: pass/fail, error, exit code, HOME, key trace, screen byte 수, progress file/app state 존재 여부
- `raw.log`: raw ANSI stream
- `screen.txt`: cleaned terminal text
- `screen_timeline.txt`: cleaned terminal text를 UI QA용 누적 흐름 evidence로 명시 저장한 파일
- `key_trace.txt`: 전송한 key trace
- `app_state.json`: test HOME의 `.advimture/e2e_state.json` snapshot
- `progress.json`: test HOME의 `.advimture/progress.json` snapshot

## Safety Guard

- `setup.home: temp`가 기본 권장값이다.
- resume 검증은 `setup.progress_file` fixture를 사용해 test HOME에만 progress를 만든다.
- 실제 사용자 HOME은 기본적으로 거부한다.
- 기존 progress file이 보이는 HOME도 기본적으로 거부한다.
- 예외가 꼭 필요할 때만 `setup.allow_unsafe_home: true`를 명시한다.

## Flake 정책

- 고정 sleep보다 "특정 텍스트가 화면에 나타날 때까지 대기"를 우선한다.
- animation 자체를 검증해야 할 때만 시간 기반 assertion을 사용한다.
- 실패 재시도는 1회 이하로 제한하고, 재시도 전에 evidence를 남긴다.

## 다음 구현 전 결정할 것

- [x] pty 의존성 추가 여부: `github.com/creack/pty` 사용
- [x] scenario 파일 위치: `test/e2e/`
- [x] smoke 실행 명령: `make e2e-smoke`
- [x] summary JSON evidence
- [x] key trace assertion
- [x] unsafe HOME guard
- [x] app state summary 기반 buffer/cursor/mode/status/score/progress assertion
- [ ] `go test ./...` 안에 포함할 smoke 범위
- [ ] full E2E를 로컬 전용으로 둘지 CI까지 올릴지 여부
