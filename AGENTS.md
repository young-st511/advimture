# AGENTS.md

## 한 줄 요약

Advimture는 Go + Bubble Tea로 만든 Vim 학습용 TUI 게임이다. 에이전트는 작은 범위로 수정하고, 관련 패키지 테스트를 우선 검증한다.

## 먼저 볼 파일

- `main.go`: 앱 시작점
- `internal/app/app.go`: 화면 전환
- `internal/editor/`: Vim 동작 핵심
- `internal/game/`: FTUE, 튜토리얼, 미션
- `internal/data/`: 튜토리얼/미션 YAML 로더
- `internal/progress/`: 진행 저장
- `PLAN.md`, `GAME_DESIGN.md`: 기획과 구현 의도

## 작업 원칙

- 수정 전 관련 파일과 테스트를 먼저 읽는다.
- 요청 범위를 벗어난 리팩터링은 하지 않는다.
- 더티 워크트리의 사용자 변경은 덮어쓰지 않는다.
- 사용자 노출 문구는 한국어 톤을 우선 유지한다.
- 새 의존성 추가나 저장 포맷 변경은 필요성과 영향이 분명할 때만 한다.
- 동작 변경 시 가능한 한 같은 패키지에 테스트를 함께 보강한다.

## 파일별 주의점

- `internal/editor/`: motion, parser, operator, undo 변경 시 기존 `*_test.go`부터 확인한다.
- `internal/data/tutorials`, `internal/data/missions`: YAML ID와 파일명은 명시적 요구가 없으면 유지한다.
- `internal/progress/`: 실제 저장 경로는 `~/.advimture/progress.json`이므로 테스트에서는 경로 주입 helper를 우선 사용한다.
- `internal/app/`, `internal/ui/`, `internal/game/`: 화면 전환이나 TUI 변경 시 수동 실행 확인이 필요할 수 있다.

## 기본 명령

- 실행: `go run .`
- 빌드: `go build .`
- 전체 테스트: `go test ./...`
- 에디터 테스트: `go test ./internal/editor/...`
- 게임 테스트: `go test ./internal/game/...`
- 데이터 테스트: `go test ./internal/data/...`
- 진행 저장 테스트: `go test ./internal/progress/...`

## 검증 원칙

- 가능하면 변경한 패키지 범위만 먼저 검증한다.
- 공통 동작에 영향이 있으면 `go test ./...`까지 확인한다.
- 작업 마무리 전 `git diff -- <changed files>`로 변경 범위를 다시 본다.

## 금지 또는 사전 확인 필요

- 불필요한 전역 리네임, 대규모 구조 변경
- 호환성 없는 save schema 변경
- 광범위한 파일 삭제/이동
- destructive git 명령

## 보고 방식

- 최종 보고 순서는 항상 `변경 내용 / 이유 / 검증 결과`로 한다.
