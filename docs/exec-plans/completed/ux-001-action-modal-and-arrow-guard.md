# UX-001 — Action Modal And Arrow Guard

## 상태

- 상태: completed
- 시작일: 2026-05-19
- 완료일: 2026-05-19
- 사용자 제보:
  - `enter`로 진행 같은 안내가 하단 일반 텍스트에 묻혀 눈에 잘 들어오지 않는다.
  - 초반 `h/j/k/l` 튜토리얼을 화살표 키로도 깰 수 있다.

## 목표

- 성공/실패/명령 모드/일반 조작 안내를 TUI 안의 박스형 Action panel로 분리해 시각적 우선순위를 높인다.
- 화살표 키 입력은 `h/j/k/l`로 변환하지 않고 `left/right/up/down` 그대로 runtime에 전달해, exercise의 `forbidden_keys` 제약이 즉시 실패로 처리하게 한다.

## 완료 내용

- TUI 하단 안내를 `ACTION` 박스 패널로 분리했다.
- 성공 상태의 `Next: enter`, `Next tutorial: enter`, `Playlist complete` 안내를 Action panel 안에 표시한다.
- 실패 상태의 `Inputs left`, `Attempts`, `Retry: r or enter` 안내를 Action panel 안에 표시한다.
- command-line mode의 `Keys: type command...` 안내도 Action panel 안에 표시한다.
- `left/right/up/down`은 Vim 이동키로 변환하지 않고 그대로 runtime key trace에 남긴다.
- 실제 방향키 escape sequence를 보내는 E2E scenario `playable_arrow_forbidden`을 추가했다.

## 검증

- `go test ./internal/tuiadapter/...`: pass
- `go test ./internal/playable/...`: pass
- `go test ./cmd/e2e-runner/...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_arrow_forbidden.yaml`: pass
- `go test ./...`: pass
- `make e2e-playable`: pass
