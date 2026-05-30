# ExecPlan: RELEASE-READINESS-001 First Release Readiness

Status: completed
Date: 2026-05-30

## Goal

첫 공개 또는 내부 배포 전에 사람이 바로 실행하고 검증할 수 있도록 README, Makefile, release gate 문서를 현재 구현 상태와 맞춘다.

## Context

- `UI-POLISH-002` 이후 tutorial/incident command memory cue까지 연결됐다.
- README가 재기획/워크플로우 온보딩 중심이라 플레이 가능한 앱의 실행, 저장 위치, 검증 명령, 한계를 충분히 설명하지 못했다.
- 기능 확장보다 release readiness 문서와 검증 루프의 정합성을 먼저 닫았다.

## Scope

포함:

- README를 현재 플레이 가능한 제품 상태 중심으로 갱신한다.
- 설치/실행/테스트/E2E/release build 명령을 명확히 한다.
- 로컬 progress 저장 위치와 E2E temp HOME 안전장치를 문서화한다.
- 현재 콘텐츠 범위와 known limitations를 공개 전 기준으로 정리한다.
- Makefile에 release 검증을 위한 작은 target을 추가한다.
- roadmap/guardrail 문서를 release readiness 완료 상태와 맞춘다.

제외:

- 새 Vim engine 기능
- 새 content/playpack 추가
- 저장 포맷 변경
- content schema 변경
- CI 추가
- UI 레이아웃 추가 개편

## Acceptance Criteria

- [x] README 첫 화면이 "재기획 세팅 단계"가 아니라 현재 실행 가능한 Vim adventure game을 설명한다.
- [x] README에 요구 Go 버전, 실행 명령, 저장 위치, reset 방법, 테스트 명령, E2E 명령, release check 명령이 있다.
- [x] README에 현재 플레이 가능한 command/content 범위와 known limitations가 있다.
- [x] Makefile에 release check 또는 build target이 문서와 일치한다.
- [x] guardrails/roadmap 문서에서 full E2E suite 관련 stale 표현을 정리한다.
- [x] `go test ./...`, `go build .`, `make e2e-playable`, `git diff --check`가 통과한다.

## Verification

- `env PATH=/opt/homebrew/bin:$PATH GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`: pass
- `env PATH=/opt/homebrew/bin:$PATH make release-check`: pass
  - includes `make test`
  - includes `make build`
  - includes `make e2e-playable`
- `git diff --check`: pass

참고: Codex sandbox에서는 `go build` 중 사용자 Go module stat cache 쓰기 경고가 한 줄 출력됐지만 exit code는 0이었다. 로컬 사용자 환경의 release gate 자체는 `make release-check`로 유지한다.

## Result

- README를 player/developer-facing quick guide로 갱신했다.
- Makefile에 `build`, `test`, `release-check` target을 추가했다.
- `docs/guardrails.md`, `docs/verification/spec.md`, roadmap 문서를 release readiness 상태와 맞췄다.
- 다음 후보는 기능 확장이 아니라 fresh playtest 기반의 release candidate review로 둔다.
