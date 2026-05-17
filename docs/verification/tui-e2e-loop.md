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
  ├─ assert: 화면 텍스트 + 종료 코드 + progress 파일 확인
  └─ evidence: raw log, cleaned screen, trace 저장
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
| screenshot | cleaned screen snapshot |
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
evidence:
  save_raw_ansi: true
  save_clean_screen: true
  save_key_trace: true
```

## Flake 정책

- 고정 sleep보다 "특정 텍스트가 화면에 나타날 때까지 대기"를 우선한다.
- animation 자체를 검증해야 할 때만 시간 기반 assertion을 사용한다.
- 실패 재시도는 1회 이하로 제한하고, 재시도 전에 evidence를 남긴다.

## 다음 구현 전 결정할 것

- [x] pty 의존성 추가 여부: `github.com/creack/pty` 사용
- [x] scenario 파일 위치: `test/e2e/`
- [x] smoke 실행 명령: `make e2e-smoke`
- [ ] `go test ./...` 안에 포함할 smoke 범위
- [ ] full E2E를 로컬 전용으로 둘지 CI까지 올릴지 여부
