# ExecPlan: Legacy archive pass

Slice-ID: VIM-011
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/LEGACY_INVENTORY.md
- docs/exec-plans/active/vim-011-legacy-archive-pass.md

## 목표

새 엔진/미들웨어/어댑터 경계가 세워진 뒤, 기존 구현을 어떤 기준으로 유지하거나 archive할지 결정한다. 이번 slice에서는 기존 Go 파일을 이동하지 않고 inventory와 archive trigger만 문서화한다.

## 범위

- 포함: legacy package inventory
- 포함: keep/archive/review 기준
- 포함: 실제 archive가 필요한 trigger 정의
- 제외: 기존 `internal/*` 파일 삭제/이동
- 제외: app wiring
- 제외: save schema 변경

## 검증 계획

- `go test ./...`
- `git status --short`

## 승인 체크

- [x] 기존 패키지별 유지/격리 후보가 정리된다.
- [x] archive trigger가 명확하다.
- [x] 이번 slice에서 기존 구현을 이동하지 않는 이유가 기록된다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: 이번 pass에서는 기존 Go 파일을 이동하지 않았다. 현재 기존 app/game/editor가 빌드를 유지하고 있고, 새 app wiring 전에는 이동이 오히려 불필요한 회귀를 만들 수 있기 때문이다.
- 2026-05-18: archive trigger를 `docs/roadmap/LEGACY_INVENTORY.md`에 문서화했다. 이후 충돌이 생기는 slice에서 작은 단위로 이동한다.
