# Changes Log — 시퀀싱/가정 변경 기록

> append-only. 새 항목을 위에 추가하고 기존 항목은 수정하지 않는다.

## 2026-05-18 — 콘텐츠/커리큘럼 기본값 확정

이전 가정: CONTENT-001의 파일 포맷, 위치, draft 파일 정책, 다음 엔진 확장 순서, scenario tone, coverage 예외, 자동 저장 기준이 아직 열려 있었다.

새 가정: 콘텐츠는 YAML 우선이며 repo root `content/`에 둔다. draft 파일은 허용하지만 playable에는 approved/implemented 및 `engine_support: implemented`만 연결한다. 첫 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`이고, 다음 엔진 확장은 `w/b/e`를 우선한다. 후반에는 `gg`, `G`, `:` command, substitute, range command까지 폭넓게 다룬다. 시나리오 톤은 DevOps/터미널 문제 해결이며 과하지 않은 억까 상황을 허용한다. 진행 사항은 자동 저장한다.

이유: CONTENT-001과 이후 game loop 구현 전에 사용자 결정을 명시적으로 고정해 Agent가 임의 선택하지 않게 하기 위함이다.

영향: `docs/roadmap/decisions/0003-content-and-curriculum-defaults.md`, `docs/gameplay/content-requirements.md`, `docs/workflows/scenario-production-harness.md`를 기준으로 다음 루프를 진행한다.

## 2026-05-17 — Harness-first 재기획으로 전환

이전 가정: 기존 `docs/archived/PLAN.md`와 `docs/archived/GAME_DESIGN.md`를 이어서 구현한다.

새 가정: 기존 구현은 참고 자료로 두고, Go 기반 TUI adventure game이라는 큰 방향만 유지한 채 제품 기획과 검증 워크플로우를 다시 세팅한다.

이유: 현재 Advimture의 제품 방향과 재미가 충분하지 않다고 판단했다.

영향: 새 구현 전 `docs/roadmap/`, `docs/exec-plans/`, `docs/gameplay/`, `docs/verification/`를 기준으로 작업한다.
