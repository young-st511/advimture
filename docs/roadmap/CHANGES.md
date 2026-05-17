# Changes Log — 시퀀싱/가정 변경 기록

> append-only. 새 항목을 위에 추가하고 기존 항목은 수정하지 않는다.

## 2026-05-17 — Harness-first 재기획으로 전환

이전 가정: 기존 `docs/archived/PLAN.md`와 `docs/archived/GAME_DESIGN.md`를 이어서 구현한다.

새 가정: 기존 구현은 참고 자료로 두고, Go 기반 TUI adventure game이라는 큰 방향만 유지한 채 제품 기획과 검증 워크플로우를 다시 세팅한다.

이유: 현재 Advimture의 제품 방향과 재미가 충분하지 않다고 판단했다.

영향: 새 구현 전 `docs/roadmap/`, `docs/exec-plans/`, `docs/gameplay/`, `docs/verification/`를 기준으로 작업한다.
