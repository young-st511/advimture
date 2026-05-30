# Guardrails Guide

> 사람을 위한 안전장치 현황입니다. 지금은 워크플로우 세팅 단계이므로 코드 변경 강제 장치는 제안만 기록합니다.

## 현재 설정된 가드레일

- Go 테스트: `go test ./...`
- Makefile test: `make test`
- Release build: `make build`
- Local release gate: `make release-check`
- 패키지별 테스트: `internal/vimengine`, `internal/runtime`, `internal/content`, `internal/scoring`, `internal/scenario`, `internal/tuiadapter`, `internal/progressadapter`, `internal/vimoracle`, `internal/playable`, `internal/progress`
- E2E smoke: `make e2e-smoke`
- Full playable E2E: `make e2e-playable`
- Makefile: `make run`, `make build`, `make test`, `make release-check`, `make e2e-smoke`, `make e2e-playable`
- `.gitignore`: 바이너리, IDE 파일, OS 파일, `coverage.out`, `artifacts/`

## 현재 비어 있는 부분

- CI 설정 없음
- 전용 lint 설정 없음
- pre-commit hook 없음
- 검증 evidence stage 차단 규칙 없음

## 권장 추가 가드레일

### CI Pipeline

최소 CI는 아래 순서로 구성한다.

1. `go test ./...`
2. `go vet ./...`
3. `make e2e-smoke`

로컬 공개 전 검증은 `make release-check`를 사용한다. 이 target은 `make test`, `make build`, `make e2e-playable`을 순서대로 실행한다.

### TUI E2E QA Loop

웹의 Playwright QA Loop처럼 운영 가능하다. 다만 브라우저 DOM 대신 pseudo terminal과 screen snapshot을 검사한다.

- 앱을 테스트 전용 HOME 또는 임시 progress path로 실행한다.
- pty에 키 입력 시퀀스를 보낸다.
- ANSI escape를 제거하거나 terminal screen buffer로 파싱한다.
- 특정 텍스트, 화면 상태, 종료 코드, progress 파일 변화를 assertion으로 검사한다.
- 실패 시 raw ANSI log, cleaned screen, 키 입력 trace를 `artifacts/e2e/`에 남긴다.

상세 설계는 `docs/verification/tui-e2e-loop.md`를 따른다.

### Ephemeral artifacts 격리

- 검증 산출물은 `artifacts/` 아래에 둔다.
- `artifacts/`는 git에서 제외한다.
- 영구 evidence가 필요한 경우 ExecPlan 본문이나 `docs/roadmap/decisions/`로 승격한다.

### 에스컬레이션 트리거

- 새 의존성 추가
- 저장 포맷 변경
- `~/.advimture/progress.json` 호환성 변경
- 새 content schema ID 변경
- Bubble Tea 앱 구조 대개편
- TUI E2E 러너가 실제 사용자 HOME에 쓰는 변경

### Claude Code Hooks

아직 설정하지 않았다. 나중에 설정할 경우:

- SessionStart: `AGENTS.md`의 Work Start Protocol을 리마인드한다.
- Stop: 변경 파일 기준으로 docs 동기화, 테스트, `git diff` 확인을 리마인드한다.

### 승격 후보

| 규칙 | 현재 강제 수준 | 반복 횟수 | 다음 승격 단계 | 비고 |
|------|--------------|----------|--------------|------|
| 승인 없는 `[draft]` 구현 금지 | 문서 | 0 | hook 리마인더 | 재기획 단계 핵심 규칙 |
| TUI evidence를 `artifacts/`에 격리 | 문서 | 0 | `.gitignore` + pre-commit | 러너 구현 시 승격 |
