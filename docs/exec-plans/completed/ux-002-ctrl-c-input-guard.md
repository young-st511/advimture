# UX-002 — Ctrl-C Input Guard

## 상태

- 상태: completed
- 시작일: 2026-05-19
- 완료일: 2026-05-19
- 사용자 제보:
  - `ctrl+c`를 누르면 게임이 바로 종료된다.

## 목표

- `ctrl+c`를 전역 quit shortcut으로 처리하지 않는다.
- `ctrl+c`는 runtime key trace로 전달해 exercise의 `forbidden_keys` 또는 unsupported key handling이 처리하게 한다.
- 앱 종료는 playable 일반 모드의 `q`를 기본으로 유지한다.

## 완료 내용

- `MapInput("ctrl+c")`가 `ActionQuit` 대신 `ActionKey`를 반환하도록 변경했다.
- command-line Action panel에서 `ctrl+c: quit` 안내를 제거했다.
- playable unit test로 `ctrl+c`가 quit command를 반환하지 않는 것을 고정했다.
- E2E `playable_ctrl_c_does_not_quit`을 추가해 실제 Ctrl+C 입력 후 앱이 계속 살아 있고 `q`로 종료되는 흐름을 검증했다.

## 검증

- `go test ./internal/tuiadapter/...`: pass
- `go test ./internal/playable/...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ctrl_c_does_not_quit.yaml`: pass
