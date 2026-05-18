# SCENARIO-TONE-001 — Midgame Survival Adventure Tone

## 상태

- 상태: completed
- 시작일: 2026-05-19
- 완료일: 2026-05-19
- 방식: 복수 SubAgent 아이디어 수집 후 종합

## 배경

초반 tutorial path는 movement, survival, navigation, small edits, Ex command까지 연결됐다. 다음 단계는 Vim의 실무 감각을 여는 operator grammar로 넘어가기 전, 중반부 시나리오 톤과 제작 가드레일을 고정하는 것이다.

## 목표

- 중반 시나리오 톤을 “터미널 문제 해결 생존 어드벤처”로 정의한다.
- 생존감이 command 학습을 가리지 않도록 금지 패턴과 검증 기준을 문서화한다.
- 다음 중기 계획을 `operator grammar` 중심으로 정리한다.

## SubAgent 종합

- 시나리오 디자이너: 원격 탐사 기지/고장 난 자동화 시스템을 배경으로, 정확한 Vim 조작이 위기를 낮추는 톤을 제안했다.
- Vim 교육 설계자: 다음 첫 중반 playpack은 `delete-with-motion`, `change-with-motion`이 가장 자연스럽다고 제안했다.
- 검증/비평 에이전트: 생존감은 배경 압력으로만 쓰고, 전투/자원관리/긴 세계관 설명은 금지해야 한다고 제안했다.

## 완료 내용

- `docs/gameplay/scenario-tone.md`를 추가했다.
- `scenario-bank.md` 운영 규칙에 scenario tone 문서를 연결했다.
- `spec.md`에 중반 생존 톤과 다음 operator grammar milestone을 반영했다.
- `vim-curriculum-map.md`에 중반 톤과 첫 operator playpack 후보를 반영했다.
- `docs/README.md`에 새 scenario tone 문서를 등록했다.
- `CHANGES.md`에 중반 톤과 operator grammar 시퀀스 변경을 append-only로 기록했다.
- `MIDTERM_TODO.md`에 다음 중기 플랜 `Operator Grammar Adventure Intro`를 추가했다.
- `PROGRAM.md`의 활성 slice를 `OPERATOR-GAP-001`로 넘겼다.

## 다음 중기 플랜

1. `OPERATOR-GAP-001`: `d/c` operator grammar 구현 범위와 vimengine/runtime gap 결정
2. `VIM-016`: operator pending mode와 `operator + motion` 전이 기반 구현
3. `VIM-017`: `delete-with-motion` (`dw`, `d$`, `dd`) 엔진/테스트 구현
4. `VIM-018`: `change-with-motion` (`cw`, `c$`, `cc`) 엔진/테스트 구현
5. `PLAYPACK-003`: 6문항짜리 operator grammar adventure intro content/E2E 구현
6. `YANK-TEXT-001`: `y/p`와 text object 후보를 다음 playpack으로 설계

## 검증

- `rg "SCENARIO-TONE-001|OPERATOR-GAP-001|delete-with-motion|터미널 문제 해결 생존" docs`: pass
- `git diff --check`: pass
- `go test ./...`: pass
