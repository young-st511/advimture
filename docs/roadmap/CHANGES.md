# Changes Log — 시퀀싱/가정 변경 기록

> append-only. 새 항목을 위에 추가하고 기존 항목은 수정하지 않는다.

## 2026-05-22 — Structure editing 완료 후 visual readiness와 incident UX 우선

이전 가정: `text-object-quote-pair`와 두 번째 incident run을 닫은 뒤 다음 후보인 `visual-char-line` 구현으로 바로 넘어갈 수 있다.

새 가정: visual mode는 다음 큰 Vim capability 후보가 맞지만, 구현 전에 selection state contract, TUI rendering, content/E2E app_state assertion을 먼저 닫는다. 또한 incident run은 기능적으로 통과했지만 아직 종합시험 감각이 남아 있으므로, 새 command 추가 전에 incident game-feel과 2단계 힌트 보강을 검토한다.

이유: Architecture Review와 Learning/Game UX Review 모두 visual mode의 표시/검증 계약 선행 필요성과 incident 적용 런의 게임성 보강 필요를 지적했다.

영향: 다음 중기 플랜 후보는 `INCIDENT-UX-003`, `E2E-FIXTURE-001`, `VISUAL-GAP-002`, `E2E-007`, `VIM-027/TUI-003` 순서로 검토한다.

## 2026-05-21 — Utility command tutorial 완료와 장기 반복 학습 RFC 분리

이전 가정: `o/O`, `.`, `/ n N`를 단기 utility command 확장으로 다루되, 이후 장기 반복 학습 플랫폼의 저장 구조는 아직 정리되지 않았다.

새 가정: `open-line-edit`, `repeat-last-change`, `search-basic`은 각각 playable tutorial과 E2E까지 연결된 foundation cluster로 승격한다. 장기 반복 학습은 `PLATFORM-RFC-001`에서 저장 변경 없는 review 후보와 progress schema v2 후보를 분리한다.

이유: Vim 학습 플랫폼은 명령어 추가만으로 완성되지 않고, 복습과 숙련도 추적이 필요하다. 다만 사용자 로컬 저장 포맷은 영향이 크므로 schema 변경은 별도 승인 게이트 뒤에 둔다.

영향: 다음 루프는 저장 변경 없는 `PLATFORM-REVIEW-001`로 시작하거나, 사용자 승인 후 `PROGRESS-SCHEMA-001`로 progress v2를 설계한다.

## 2026-05-19 — 중반 생존 어드벤처 톤과 operator grammar 순서 확정

이전 가정: 중반부터 생존 어드벤처 비중을 높인다는 방향은 있었지만, 생존감의 범위와 다음 콘텐츠 시퀀스가 구체적으로 고정되지 않았다.

새 가정: 중반 톤은 `docs/gameplay/scenario-tone.md`의 “터미널 문제 해결 생존 어드벤처”를 따른다. 생존감은 전투, 체력, 인벤토리가 아니라 로그, 설정, 통신, 임시 패치, 저장/폐기 판단 같은 텍스트 조작 압력에서 온다. 다음 중기 플랜은 `delete-with-motion`, `change-with-motion` 중심의 operator grammar adventure intro로 진행한다.

이유: 작은 수정 튜토리얼 이후 Vim이 실무 도구처럼 느껴지려면 `operator + motion` 문법으로 넘어가는 것이 자연스럽다. 시나리오가 command 학습을 가리지 않도록 톤 가드레일을 먼저 고정했다.

영향: `docs/gameplay/scenario-tone.md`, `docs/gameplay/vim-curriculum-map.md`, `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/PROGRAM.md`를 기준으로 다음 루프를 진행한다.

## 2026-05-18 — 튜토리얼 페이싱과 제약 설계 원칙 확정

이전 가정: 첫 playable pack은 17개 exercise를 한 번에 완주하는 vertical slice이며, 초반/중반 튜토리얼 분리와 실패/재시도/입력 제약 기준은 아직 열려 있었다.

새 가정: 초반은 8문항 이하의 짧은 튜토리얼 에피소드 묶음으로 운영한다. 초반은 “첫 투어” 느낌으로 command를 넓게 맛보게 하고, 중후반부터 생존 어드벤처와 탐험 비중을 높인다. `:s`, `:%s`, range substitute는 초반에서 빼고 중반 고급 튜토리얼로 분리한다. 최대 입력 수 초과와 금지 입력/금지 우회 전략 사용은 즉시 실패이며, 초반 재시도는 무제한 기본이다. 재시도는 `r`과 `enter`를 모두 허용한다.

이유: 첫 pack이 E2E vertical slice로는 유효하지만 실제 학습 단위로는 넓기 때문에, 플레이어가 부담 없이 익히는 짧은 episode 구조가 필요하다. 또한 Vim 학습 게임으로서 “목표에 도착했는가”뿐 아니라 “의도 command를 썼는가”를 강제해야 한다.

영향: `docs/roadmap/decisions/0004-tutorial-pacing-and-constraints.md`, `docs/gameplay/spec.md`, `docs/gameplay/vim-curriculum-map.md`, `docs/workflows/scenario-flow-workbench.md`를 기준으로 다음 구현 루프를 진행한다.

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
