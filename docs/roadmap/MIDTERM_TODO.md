# Midterm Todo

> 새 엔진 모듈을 실제 게임으로 연결하기 위한 중기 순서다. 각 항목은 독립 ExecPlan으로 열고, 테스트와 커밋까지 닫은 뒤 다음 항목으로 넘어간다.

## 실행 보드

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | E2E-003 | completed | app state summary 기반 buffer/cursor/mode/status/score/progress assertion |
| 2 | PLAY-001 | completed | 첫 playable vertical slice: `hjkl` 목표 이동, 성공, 점수, progress 반영 |
| 3 | LEGACY-001 | completed | 새 playable path가 canonical이 된 뒤 기존 editor/game 격리 |
| 4 | SCENARIO-001 | completed | 첫 5분 시나리오 워크플로우로 content loader 요구사항 발견 |
| 5 | VIM-CURRICULUM-001 | completed | Vim curriculum map과 scenario production harness |
| 6 | CONTENT-001 | completed | 새 content schema용 YAML loader |
| 7 | PLAY-002 | completed | hardcoded playable 제거, file-backed playable |
| 8 | CONTENT-002 | completed | replay/coverage validator |
| 9 | VIM-012 | completed | 다음 command cluster: `w/b/e` word motion |
| 10 | EXERCISE-001 | completed | word motion exercise set |
| 11 | SURVIVAL-001 | completed | `esc`, `:q!`, `:wq` 생존/종료 루프 |
| 12 | NAV-001 | completed | 후반 navigation 확장: `gg`, `G`, line/file motion |
| 13 | EXCMD-001 | completed | `:` 명령어, substitute, range command 기반 |
| 14 | GAMELOOP-001 | completed | 반복 학습 루프, unlock, 자동 저장, 피드백 구조 |

## 루프별 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| VIM-001 | curriculum defaults 결정 완료 | 첫 5분 command cluster approval packet | `rg`, `go test ./...` | completed: 사용자 승인 후 3개 cluster를 `approved`로 승격했다. |
| CONTENT-001 | VIM-001 승인 완료 | root `content/` YAML fixture, loader, validator | `go test ./internal/content/...`, `go test ./...` | completed: draft/planned 콘텐츠는 로드하되 playable 후보에서 제외한다. |
| PLAY-002 | CONTENT-001 통과 | hardcoded playable 제거, file-backed playable | `go test ./internal/playable/...`, `make e2e-smoke` | completed: screen assertion과 app state/progress를 함께 검증했다. |
| CONTENT-002 | file-backed playable 통과 | replay/coverage validator | content replay tests, coverage report | completed: `coverage_required`, `replay_status: pass`, key trace, E2E assertion을 loader gate로 검증한다. |
| VIM-012 | word-motion cluster approved | `w/b/e` engine, oracle fixtures | `go test ./internal/vimengine/...`, oracle comparison | completed: word boundary, 공백, 문장부호, 줄 경계, unsupported mode, DesiredCol 회귀를 고정했다. |
| EXERCISE-001 | VIM-012 통과 | word motion exercise set | replay validator, verifier OK | completed: `w`, `b`, `e`가 각각 approved exercise optimal trace에 등장하고 replay gate를 통과한다. |
| SURVIVAL-001 | command-line scope 승인 | `esc`, `:q!`, `:wq` runtime/app semantics | unit tests, E2E smoke | completed: app exit와 mission success를 분리하고 command-line replay/goal gate를 통과했다. |
| NAV-001 | movement fundamentals 통과 | `gg`, `G`, line/file motion clusters | vimengine/oracle tests | completed: `gg`, `G`, `0`, `$`로 범위를 좁히고 replay/coverage gate를 통과했다. |
| EXCMD-001 | command-line engine 기반 | `:` 명령어, substitute, range command | parser/runtime tests | completed: literal `:s`, `:%s`, `:2,3s`를 buffer target으로 검증한다. |
| GAMELOOP-001 | file-backed multi exercise 가능 | playlist/unlock/retry/hint/autosave loop | E2E first-5-minute | completed: 기존 progress `Missions` map을 재사용해 exercise ID별 autosave를 수행한다. |

## 진행 원칙

- E2E assertion이 부족하면 app wiring을 멈추고 E2E를 먼저 보강한다.
- playable slice는 한 문제만 end-to-end로 완성한다.
- 기존 구현 archive는 새 path가 동작하고 import graph에서 빠진 뒤 진행한다.
- 새 Vim command는 `command catalog -> vimengine -> oracle comparison -> exercise` 순서로 추가한다.
- 각 루프는 ExecPlan을 열고, 해당 루프의 검증과 커밋까지 닫은 뒤 다음 항목으로 넘어간다.
- 후반부 루프일수록 `scenario-production-harness.md`의 Verifier OK, replay gate, coverage gate를 생략하지 않는다.

## 완료된 중기 플랜: 튜토리얼 UX와 학습 제약

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | TUTORIAL-001 | completed | 17개 first tour를 8문항 이하 tutorial episode playlist로 분리 |
| 2 | CONSTRAINT-001 | completed | max input, required command, forbidden input/strategy schema와 runtime 실패 처리 |
| 3 | SCORING-002 | completed | 의도 command 사용 여부를 grade/coaching에 반영 |
| 4 | FAILURE-001 | completed | 실패 화면, 남은 입력 수 UI, `r`/`enter` retry UX 완성 |
| 5 | QA-001 | completed | forbidden input, max input 초과, non-intended route, retry/hint E2E 보강 |
| 6 | CONTENT-003 | completed | Ex command를 중반 고급 튜토리얼로 위치시키고 first tour pacing 재검증 |

## 현재 판단

플랫폼 기반과 첫 튜토리얼 UX는 completed다. 다음 중기 플랜은 콘텐츠를 늘리기 전에 Vim 학습 지도, 엔진 gap, 다음 playpack, 시나리오 톤을 순서대로 고정한다.

## 다음 중기 플랜

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | CURRICULUM-001 | completed | Vim 핵심 커버리지 루브릭과 다음 command cluster 우선순위 확정 |
| 2 | ENGINE-GAP-001 | completed | 다음 playpack에 필요한 vimengine/runtime gap 식별과 구현 순서 결정 |
| 3 | VIM-013 | completed | `x`, `r` buffer mutation engine 구현 |
| 4 | VIM-014 | completed | `i`, `a`, `A` insert mode entry와 printable insertion 구현 |
| 5 | VIM-015 | completed | `u`, `ctrl+r` undo/redo stack 구현 |
| 6 | PLAYPACK-002 | completed | 7문항짜리 “편집 기본기” 튜토리얼 playpack 설계/구현 |
| 7 | SCENARIO-TONE-001 | completed | 중반 생존 어드벤처 톤 원칙과 적용 범위 확정 |

## 다음 중기 플랜: Operator Grammar Adventure Intro

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | OPERATOR-GAP-001 | completed | `d/c` operator grammar 구현 범위와 vimengine/runtime gap 결정 |
| 2 | VIM-016 | completed | operator pending mode와 `operator + motion` 전이 기반 구현 |
| 3 | VIM-017 | completed | `delete-with-motion` (`dw`, `d$`, `dd`) 엔진/테스트 구현 |
| 4 | VIM-018 | completed | `change-with-motion` (`cw`, `c$`, `cc`) 엔진/테스트 구현 |
| 5 | PLAYPACK-003 | completed | 6문항짜리 operator grammar adventure intro content/E2E 구현 |
| 6 | YANK-TEXT-001 | completed | `y/p`와 text object 후보를 다음 playpack으로 설계 |

## 다음 중기 플랜: Yank / Put and Text Object Bridge

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | YANK-TEXT-001 | completed | `y/p`와 text object 후보를 다음 playpack으로 설계 |
| 2 | VIM-019 | completed | unnamed register와 `yw`, `y$`, `yy` yank engine 구현 |
| 3 | VIM-020 | completed | `p`, `P` put engine과 runtime replay smoke 구현 |
| 4 | PLAYPACK-004 | completed | yank/put basic tutorial content/E2E 구현 |
| 5 | TEXT-OBJECT-001 | completed | `iw` 기반 text object gap planning |

## 다음 중기 플랜: Text Object Inner Word

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | TEXT-OBJECT-001 | completed | `iw` 기반 text object 범위와 제외 항목 결정 |
| 2 | VIM-021 | completed | `operator -> i -> w` text object pending/selection 기반 구현 |
| 3 | VIM-022 | completed | `diw`, `ciw`, `yiw` mutation/register semantics 구현 |
| 4 | PLAYPACK-005 | completed | 6문항 text object inner word tutorial content/E2E 구현 |
| 5 | E2E-PLAYPACK-005 | completed | full playlist, forbidden 우회, progress 저장 검증 |

## 다음 중기 플랜: Utility Commands and Long-Run Platform

> 2026-05-21 토론 결과를 반영한다. 우선순위는 콘텐츠 확장과 게임성 강화이며, 단기 구현은 저장 포맷 변경 없이 닫는다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | OPEN-LINE-001 | completed | `o/O` open-line edit 범위와 제외 항목 결정 |
| 2 | VIM-023 | completed | `o`, `O` newline insertion, insert mode entry, undo snapshot 구현 |
| 3 | PLAYPACK-006 | completed | 4~6문항 open-line tutorial content/E2E 구현 |
| 4 | DEBRIEF-001 | completed | 저장 포맷 변경 없는 성공/playlist 완료 debrief와 best record 표시 |
| 5 | REPEAT-GAP-001 | completed | `.` repeat-last-change transaction 범위 결정 |
| 6 | VIM-024 | completed | `.` 최소 subset 구현 |
| 7 | PLAYPACK-007 | completed | repeat-last-change efficiency tutorial content/E2E 구현 |
| 8 | SEARCH-GAP-001 | completed | literal `/`, `n`, `N` search 범위와 `?` hint 충돌 처리 결정 |
| 9 | VIM-025 | completed | literal search state와 next/previous match 구현 |
| 10 | PLAYPACK-008 | completed | search-basic tutorial content/E2E 구현 |
| 11 | PLATFORM-RFC-001 | completed | mastery, spaced review, daily run, progress schema 후보 RFC |

## Utility Commands and Long-Run Platform 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| OPEN-LINE-001 | `open-line-edit`가 curriculum next 후보 | command cluster draft/approval packet, VIM-023/PLAYPACK-006 분리 계획 | 문서 리뷰, `git diff --check` | indentation, auto-comment, count prefix, insert-mode Enter, dot repeat 제외 |
| VIM-023 | OPEN-LINE-001 완료 | completed: `o/O` engine, tuiadapter uppercase mapping, runtime replay smoke | completed: `go test ./internal/vimengine/...`, `go test ./internal/tuiadapter/...`, `go test ./internal/runtime/...` | content/E2E와 섞지 않고 engine contract만 닫았다 |
| PLAYPACK-006 | VIM-023 완료 | completed: `open-line-edit` YAML content, scenario, playlist, full E2E | completed: content replay, `playable_open_line_full.yaml`, forbidden-route E2E, `make e2e-smoke` | 8문항 이하, `o/O` required key와 우회 금지를 고정했다 |
| DEBRIEF-001 | PLAYPACK-006 완료 | completed: 성공/playlist 완료 화면 debrief, 기존 progress 기반 best record 표시 | completed: playable model tests, focused E2E | `internal/progress/` 저장 JSON 구조를 변경하지 않았다 |
| REPEAT-GAP-001 | open-line playpack과 debrief 완료 | completed: last-change transaction RFC, 최소 subset 결정 | completed: docs review | `.` 구현 전 undo/insert/yank/put transaction 경계를 먼저 고정했다 |
| VIM-024 | REPEAT-GAP-001 승인 | completed: `.` 최소 subset engine/runtime 구현 | completed: vimengine/tuiadapter/runtime tests | macro/register 수준으로 확장하지 않았다 |
| PLAYPACK-007 | VIM-024 완료 | completed: repeat-last-change tutorial + efficiency run E2E | completed: content replay, full playlist E2E | 수동 재입력 우회를 constraints로 차단했다 |
| SEARCH-GAP-001 | repeat playpack 완료 | completed: literal search scope, `/` command-line/search state, `?` 보류 결정 | completed: docs review | regex, `?`, highlight 검증을 첫 구현에 넣지 않는다 |
| VIM-025 | SEARCH-GAP-001 승인 | completed: `/`, `n`, `N` literal search engine/runtime 구현 | completed: vimengine/runtime/tuiadapter/playable tests | hint `?` 충돌을 우회 구현하지 않았다 |
| PLAYPACK-008 | VIM-025 완료 | completed: search-basic tutorial + E2E | completed: content replay, search app_state E2E | 검색 결과는 cursor/mode/key trace로 검증했다 |
| PLATFORM-RFC-001 | 최소 3개 utility playpack 완료 | completed: mastery/review/daily run progress 후보 RFC | completed: 저장 포맷 영향 리뷰 | 사용자 승인 없이 progress schema 변경 금지 |

## Utility Commands and Long-Run Platform 완료 판단

2026-05-21 기준 이 중기 플랜은 completed다. `o/O`, `.`, `/ n N`는 각각 gap planning, engine, playable tutorial, E2E까지 연결됐다. 장기 반복 학습 플랫폼은 `docs/roadmap/PLATFORM_RFC_001.md`에 저장 변경 없는 첫 루프와 schema 변경 승인 루프를 분리해 두었다.

## 중간점검 중 즉시 보강한 QA

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | QA-PLATFORM-001 | completed | stale full E2E expectation 복구와 `make e2e-playable` 최신 playpack suite 갱신 |

## 다음 중기 플랜: Review Queue and Incident Runs

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLATFORM-REVIEW-001 | completed | 저장 포맷 변경 없는 review 후보 계산과 `재진단 큐`/`잔류 리스크` 표시 |
| 2 | PLAYLIST-ORDER-001 | completed | tutorial/incident 카테고리 순서를 ID hack 없이 명시적으로 정렬 |
| 3 | INCIDENT-RUN-001 | completed | 기존 command를 섞은 `incident-001` 생존 어드벤처 mixed run |
| 4 | COACHING-001 | completed | schema 변경 없이 strict constraint 전 사전 코칭 UX 개선 |

## Review Queue and Incident Runs 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLATFORM-REVIEW-001 | PLATFORM-RFC-001 완료 | completed: review 순수 계산 패키지, TUI read-only 표시, focused E2E | completed: review/playable tests, focused E2E, `go test ./...`, `make e2e-playable` | `internal/progress/` 저장 JSON 구조를 변경하지 않았다 |
| PLAYLIST-ORDER-001 | review 표시 완료 | completed: playlist ordering contract와 loader/playable 정렬 변경 | completed: content/playable tests, `make e2e-playable` | 기존 progress mission ID를 변경하지 않았다 |
| INCIDENT-RUN-001 | incident ordering 가능 | completed: `incident-001` content/scenario/playlist + full E2E | completed: content replay, incident E2E, `make e2e-playable` | 새 Vim engine 기능을 추가하지 않았다 |
| COACHING-001 | strict constraint UX 리스크 확인 | completed: required key coaching 표시와 `?` hint 노출 개선 | completed: runtime/playable tests, focused E2E, `make e2e-playable` | Practice/Challenge schema 분리 보류 |

## Review Queue and Incident Runs 완료 판단

2026-05-22 기준 이 중기 플랜은 completed다. review queue는 저장 변경 없이 첫 화면과 성공 debrief에 연결됐고, playlist ordering은 tutorial/incident category와 order로 명시화됐다. 첫 incident mixed run은 기존 Vim command 조합으로 완주 가능하며, strict constraint 문항의 coaching panel도 focused E2E와 full playable suite로 검증한다.

## 다음 중기 플랜: Structure Editing and Applied Survival

> 목표는 Vim 실무 효용이 큰 구조 내부 편집을 추가하고, 이를 두 번째 incident run에서 기존 search/substitute와 섞어 적용하는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-002 | completed | 완료된 후보를 정리하고 새 중기 플랜을 roadmap/curriculum에 고정 |
| 2 | TEXT-PAIR-GAP-001 | completed | quote/pair text object 최소 범위와 제외 항목 결정 |
| 3 | VIM-026 | completed | `ci"`, `di"`, `yi"` 중심 quote text object engine 구현 |
| 4 | PLAYPACK-009 | completed | quote/pair text object 튜토리얼 content/E2E 구현 |
| 5 | INCIDENT-RUN-002 | completed | search + substitute + quote/pair를 섞은 두 번째 생존 어드벤처 run |
| 6 | VISUAL-GAP-001 | completed | visual mode 후보 범위와 engine 영향도를 문서로 분리 |

## Structure Editing and Applied Survival 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLAN-REFRESH-002 | Review Queue and Incident Runs 완료 | completed: 새 중기 플랜과 next candidate 갱신 | completed: `git diff --check` | 코드와 content를 건드리지 않는다 |
| TEXT-PAIR-GAP-001 | text-object-inner-word playable 완료 | completed: command cluster approval packet, VIM-026/PLAYPACK-009 분리 계획 | completed: content test, `git diff --check` | nested pair, escaped quote, around object, count prefix, visual selection 제외 |
| VIM-026 | TEXT-PAIR-GAP-001 완료 | completed: quote text object engine/runtime tests | completed: `go test ./internal/vimengine/...`, `go test ./internal/runtime/...` | engine과 content/E2E를 섞지 않았다 |
| PLAYPACK-009 | VIM-026 완료 | completed: quote/pair text object tutorial YAML, scenario, playlist, full E2E | completed: content replay, focused E2E, incident fixture check | 8문항 이하, required key와 우회 금지를 고정했다 |
| INCIDENT-RUN-002 | PLAYPACK-009 완료 | completed: mixed incident content/scenario/playlist + full E2E | completed: content replay, incident E2E | 새 engine 기능을 incident에서 추가하지 않았다 |
| VISUAL-GAP-001 | second incident 완료 | completed: visual mode scope/RFC | completed: `git diff --check` | visual engine 구현은 다음 중기 플랜으로 분리했다 |

## Structure Editing and Applied Survival 완료 판단

2026-05-22 기준 이 중기 플랜은 completed다. quote text object는 gap planning, engine, 4문항 tutorial, full playlist E2E까지 연결됐고, 두 번째 incident run은 search/substitute/quote/dot repeat 조합으로 완주 가능하다. visual mode는 `visual-char-line` draft cluster와 영향도 문서까지만 닫았으며, 실제 visual engine 구현은 다음 중기 플랜에서 별도 gap planning으로 시작한다.

## 다음 중기 플랜: Applied Learning and World Frame

> 목표는 incident를 “종합시험”이 아니라 “하나의 복구 작전”처럼 느끼게 만들고, visual mode 구현 전에 필요한 UX/E2E 계약을 준비하는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-003 | completed | 세계관 프레임과 다음 보강 순서 고정 |
| 2 | INCIDENT-UX-003 | completed | incident 001/002 제목, briefing, feedback, 2단계 hint를 복구 작전으로 정렬 |
| 3 | PROGRESS-LANGUAGE-001 | completed | 저장 포맷 변경 없이 review/best record 문구를 복구국 프레임으로 정렬 |
| 4 | E2E-FIXTURE-001 | completed | 긴 progress fixture 유지보수 완화 전략 수립 |
| 5 | VISUAL-GAP-002 | completed | visual selection state, TUI 표시, app_state assertion 계약 확정 |
| 6 | E2E-007 | completed | selection app_state/content assertion 확장 |
| 7 | VIM-027-TUI-003 | completed | `v` charwise selection foundation과 최소 표시 구현 |

## Applied Learning and World Frame 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLAN-REFRESH-003 | health check와 세계관 후보 선택 완료 | completed: world frame decision, 다음 중기 플랜 | completed: `git diff --check` | command 학습보다 lore를 앞세우지 않는다 |
| INCIDENT-UX-003 | world frame decision 완료 | completed: incident scenario/hint wording 개선, focused E2E 갱신 | completed: content tests, incident E2E, `make e2e-playable` | target_state, optimal_keys, constraints는 변경하지 않았다 |
| PROGRESS-LANGUAGE-001 | incident UX 완료 | review/debrief 문구 정렬 | playable tests, focused E2E | progress 저장 JSON 구조를 변경하지 않는다 |
| E2E-FIXTURE-001 | full E2E fixture 부담 확인 | fixture builder 또는 최소 fixture 전략 | e2e-runner tests, representative E2E | 실제 HOME 사용 금지 |
| VISUAL-GAP-002 | UX 보강 완료 | visual selection contract | docs review, `git diff --check` | visual block, count/register prefix, indentation 제외 |
| E2E-007 | VISUAL-GAP-002 완료 | selection app_state assertion | e2estate/e2e-runner/content tests | 화면 텍스트만으로 selection을 검증하지 않는다 |
| VIM-027-TUI-003 | E2E-007 완료 | completed: charwise `v`, motion selection, `esc` reset, 최소 표시 | completed: vimengine/tui/playable tests, full regression | `d/y` operator application과 playpack은 분리 |

## 2~3개월 주차 계획

## 다음 중기 플랜: Visual Operator and Tutorial

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | VISUAL-OP-001 | completed | charwise visual selection의 `d/y` 적용 범위와 제외 항목 결정 |
| 2 | VIM-028 | completed | 같은 줄 charwise visual selection delete/yank engine 구현 |
| 3 | PLAYPACK-010 | completed | 3문항 visual selection tutorial content/E2E 구현 |

## Visual Operator and Tutorial 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| VISUAL-OP-001 | VIM-027-TUI-003 완료 | completed: `d/y` 범위와 제외 항목 | completed: `git diff --check` | visual operator와 playpack을 섞지 않는다 |
| VIM-028 | VISUAL-OP-001 완료 | completed: charwise same-line visual `d/y` engine | completed: vimengine/runtime tests, full regression | multi-line, `V`, block, count/register prefix 제외 |
| PLAYPACK-010 | VIM-028 완료 | completed: visual tutorial content + focused E2E | completed: content replay, focused E2E, `make e2e-playable` | 4문항 이하, command 학습을 scenario보다 우선 |

## Visual Operator and Tutorial 완료 판단

2026-05-22 기준 이 중기 플랜은 completed다. visual mode는 selection state/render/assertion, 같은 줄 charwise `d/y` engine, 3문항 tutorial, full playlist E2E까지 연결됐다. `V`, visual block, multi-line visual operator, count/register prefix는 후속 hardening 후보로 남긴다.

## 다음 중기 플랜: Post-Visual Applied Mastery and Hardening

> 목표는 visual tutorial 직후 새 command를 무리하게 늘리지 않고, selection 검증 구멍과 engine 분리 압력을 먼저 낮춘 뒤 visual을 실제 incident 적용 런으로 승격하는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-004 | completed | post-visual 리뷰/토론 결과를 종합해 다음 중기 플랜 고정 |
| 2 | QA-SEL-001 | completed | content replay 단계에서 selection assertion mismatch 검증 보강 |
| 3 | ENGINE-SPLIT-001 | completed | `internal/vimengine`의 selection/range/operator helper를 behavior 변화 없이 분리 |
| 4 | VISUAL-HARDEN-001 | completed | charwise visual invariant, undo/register, unsupported edge case 테스트 보강 |
| 5 | INCIDENT-RUN-003 | completed | visual selection을 적용하는 세 번째 Runbook Dispatch incident 구현 |
| 6 | INCIDENT-FLOW-001 | completed | incident beat 연결 문구와 runbook continuity를 저장 변경 없이 보강 |
| 7 | VISUAL-LINE-GAP-001 | completed | linewise `V`와 multi-line visual 범위를 구현 전 gap planning으로 결정 |

## Post-Visual Applied Mastery and Hardening 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLAN-REFRESH-004 | PLAYPACK-010 완료 | completed: 리뷰 종합 문서와 새 중기 플랜 | completed: `git diff --check` | 코드와 content 동작을 변경하지 않는다 |
| QA-SEL-001 | selection app_state assertion 존재 | completed: content replay selection mismatch 테스트, runner/content 문서 동기화 | completed: `go test ./internal/content/...`, `go test ./cmd/e2e-runner/...`, `go test ./...` | E2E 화면 문구만으로 selection을 검증하지 않는다 |
| ENGINE-SPLIT-001 | QA-SEL-001 완료 | completed: selection/range/operator helper 파일 분리 | completed: `go test ./internal/vimengine/...`, `go test ./...` | behavior change와 새 visual 기능 추가 금지 |
| VISUAL-HARDEN-001 | ENGINE-SPLIT-001 완료 | completed: same-line 경계, empty line, toggle/reset, undo/register, unsupported multi-line 테스트 | completed: `go test ./internal/vimengine/...`, `go test ./internal/runtime/...`, `go test ./...` | `V`, visual block, count/register prefix는 넣지 않는다 |
| INCIDENT-RUN-003 | visual tutorial과 hardening 완료 | completed: `릴레이 기지 003: 오염 구간 격리` content/scenario/playlist + full E2E | completed: content replay, focused E2E, `make e2e-playable` | 새 command를 소개하지 않고 기존 command 조합만 사용 |
| INCIDENT-FLOW-001 | incident 003 완료 | completed: incident 001~003 beat 연결 문구, briefing/debrief continuity pass | completed: content tests, incident E2E, `make e2e-playable` | target_state, optimal_keys, constraints 변경 금지 |
| VISUAL-LINE-GAP-001 | charwise visual 적용 런 완료 | completed: `V`, multi-line visual, selection contract 영향도와 제외 항목 문서화 | completed: docs review, `git diff --check` | 구현은 다음 중기 플랜으로 분리 |

2026-05-22 기준 이 중기 플랜은 completed다. selection replay, engine split, visual hardening, visual 적용 incident, incident continuity, linewise visual gap planning을 모두 닫았다. 다음 중기 플랜은 `visual-line-basic` 구현 여부와 command choice drill 또는 platform review loop 중 무엇을 먼저 강화할지 결정하는 단계에서 시작한다.

## Post-Visual 리뷰 판단

Architecture, Learning/Game UX, Scenario/World 리뷰 모두 다음 공통 결론을 냈다. 현재 구조는 건강하지만 visual 이후 복잡도는 selection 검증과 engine 내부 분리 없이는 커질 수 있다. 학습적으로는 새 command보다 “상황에 맞는 도구 선택”과 “incident 적용”이 더 중요하다. 세계관은 충분하므로 lore 확장보다 runbook 작전감과 visual 적용 런을 우선한다.

## 장기 참고 계획

| 기간 | 목표 | 닫는 루프 |
|------|------|----------|
| Week 1-2 | `o/O` 범위 확정과 엔진 구현 | OPEN-LINE-001, VIM-023 |
| Week 3-4 | open-line playable 콘텐츠와 저장 변경 없는 debrief | PLAYPACK-006, DEBRIEF-001 |
| Week 5-6 | `.` repeat-last-change gap planning과 최소 엔진 | REPEAT-GAP-001, VIM-024 |
| Week 7-8 | repeat tutorial과 efficiency run 검증 | PLAYPACK-007 |
| 이후 | literal search와 장기 반복 학습 RFC | SEARCH-GAP-001, VIM-025, PLAYPACK-008, PLATFORM-RFC-001 |

## 다음 중기 플랜: Linewise Visual and Choice Training

> 목표는 charwise visual 이후 바로 복잡한 multi-line charwise나 visual block으로 가지 않고, 실무 체감이 큰 linewise `V`를 테스트 우선으로 구현한 뒤 tutorial과 적용 run으로 승격하는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-005 | completed | linewise visual 중심 중기 플랜 고정 |
| 2 | VISUAL-LINE-001 | completed | linewise `V` scope와 수용 기준 승인 |
| 3 | VIM-029 | completed | linewise `V` foundation과 `d/y` engine/runtime/TUI 구현 |
| 4 | PLAYPACK-011 | completed | 3문항 이하 linewise visual tutorial content/E2E 구현 |
| 5 | INCIDENT-RUN-004 | completed | linewise visual을 적용하는 config block 복구 incident 구현 |
| 6 | COMMAND-CHOICE-001 | completed | 배운 command 중 적절한 도구를 고르는 mixed drill 설계 |
| 7 | PLATFORM-REVIEW-002 | completed | 저장 포맷 변경 없는 review/daily 동기 강화 |

## Linewise Visual and Choice Training 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLAN-REFRESH-005 | Post-Visual plan completed | completed: 새 중기 플랜 | completed: `git diff --check` | 구현 없이 순서와 경계만 고정 |
| VISUAL-LINE-001 | VISUAL-LINE-GAP-001 완료 | completed: linewise `V` acceptance criteria와 ExecPlan | completed: docs review, `git diff --check` | multi-line charwise와 visual block 제외 |
| VIM-029 | VISUAL-LINE-001 승인 | completed: linewise `V`, row motion, `d/y`, linewise register, TUI/app_state | completed: vimengine/tuiadapter/runtime/playable tests, `go test ./...` | 테스트 Red 먼저, progress schema 변경 금지 |
| PLAYPACK-011 | VIM-029 완료 | completed: 3문항 이하 tutorial, full E2E | completed: content replay, focused E2E, `make e2e-playable` | scenario보다 command 차이를 먼저 가르침 |
| INCIDENT-RUN-004 | linewise tutorial 완료 | completed: config block 복구 incident | completed: content replay, incident E2E, `make e2e-playable` | 새 command 소개 금지, 적용만 수행 |
| COMMAND-CHOICE-001 | incident 004 완료 | completed: command choice drill 설계 | completed: docs review, `git diff --check` | 정답 암기보다 도구 선택 이유를 강화 |
| PLATFORM-REVIEW-002 | review queue baseline 존재 | completed: 저장 변경 없는 review/daily UX 개선 | completed: playable/e2estate tests, focused E2E, full regression | progress v2는 별도 승인 전 금지 |

## Linewise Visual and Choice Training 완료 판단

2026-05-23 기준 이 중기 플랜은 completed다. linewise `V`는 scope approval, engine/runtime/TUI, 3문항 tutorial, incident 004 적용 run까지 연결됐다. 이후 command choice drill은 docs-only 적용 레이어로 고정했고, review/daily 동기는 저장 포맷 변경 없이 `오늘의 복구 루트`와 E2E state summary로 보강했다.

## 다음 중기 플랜: TUI Experience Foundation

> 목표는 현재 작동하는 TUI를 제품 UI로 승격하되, 화면 리디자인 전에 LLM이 의미 상태를 안정적으로 확인할 수 있는 QA 계약과 렌더링 경계를 먼저 만든다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | UX-REVIEW-001 | completed | 학습 UX, 콘솔 비주얼, 아키텍처, E2E QA 관점 SubAgent 리뷰 통합 |
| 2 | UX-QA-001 | completed | review/daily route typed app_state assertion과 evidence snapshot |
| 3 | UI-CONTRACT-001 | completed | header/briefing/console/status/action 영역 contract 고정 |
| 4 | UI-RENDER-001 | completed | `internal/playable`에서 순수 renderer 패키지 분리 |
| 5 | UI-HIERARCHY-001 | completed | 현재 exercise 목표를 상위로 올리고 review/daily를 보조 정보로 낮춤 |
| 6 | UI-MODE-001 | completed | tutorial action panel과 incident action panel 언어 분리 |
| 7 | UI-EVIDENCE-001 | completed | final frame 또는 frame timeline evidence 보강 |

## TUI Experience Foundation 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| UX-REVIEW-001 | Linewise Visual and Choice Training 완료 | completed: SubAgent 리뷰와 통합 결론 | completed: docs review | 세계관보다 Vim 학습 목표 우선 |
| UX-QA-001 | review/daily app_state 존재 | completed: typed review assertion, app_state/progress evidence | completed: runner tests, focused E2E | 화면 문자열 contains에 의존하지 않음 |
| UI-CONTRACT-001 | UX-QA-001 완료 | completed: 화면 영역 contract와 첫 렌더링 slice 수용 기준 | completed: docs review, `git diff --check` | 코드 구현 전 정보 위계 고정 |
| UI-RENDER-001 | 영역 contract 완료 | completed: 순수 renderer 패키지와 기존 behavior 보존 | completed: playable/renderer tests, full regression, smoke E2E | Update/game state 변경 금지 |
| UI-HIERARCHY-001 | renderer 분리 완료 | completed: 현재 목표 우선 화면 구조 | completed: renderer/playable tests, smoke E2E | review/daily는 보조 정보 유지 |
| UI-MODE-001 | hierarchy 완료 | completed: tutorial/incident action panel 분리 | completed: playable tests, incident E2E, smoke E2E | incident에서 정답 키 즉시 노출 과다 방지 |
| UI-EVIDENCE-001 | UI 변경 반복 부담 확인 | completed: screen timeline evidence | completed: e2e-runner tests, focused UI QA E2E | timing flake를 늘리지 않음 |

## TUI Experience Foundation 완료 판단

2026-05-23 기준 이 중기 플랜은 completed다. UI 기획 루프를 상세화해 화면 영역 contract를 고정했고, `internal/playableview` renderer 분리, current task first 정보 위계, tutorial/incident action 언어 분리, screen timeline evidence까지 순서대로 연결했다. 다음 UI 루프는 style/theme, viewport parser, modal 안내 같은 더 큰 화면 시스템 변경을 별도 중기 플랜으로 열어 진행한다.

## 다음 중기 플랜: Foundation Product Loop

> 목표는 새 Vim 기능을 더 늘리기보다, 이미 쌓인 엔진/콘텐츠/UI를 첫 플레이 경험과 반복 학습 동기로 묶어 `Vim Learning Foundation`의 제품 출구 조건을 닫는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-006 | completed | Foundation milestone 점검과 다음 2~3주 순서 고정 |
| 2 | FTUE-001 | completed | 첫 5분 canonical route 정의와 evidence 점검 |
| 3 | UI-PLAYTEST-001 | completed | Mission HUD + floating modal 사용성 evidence 기반 polish |
| 4 | CHOICE-PLAY-001 | completed | command-choice drill을 playable content/E2E로 승격 (`incident-005-command-choice`) |
| 5 | DAILY-ROUTE-001 | completed | 저장 포맷 변경 없는 daily/review 반복 동기 강화 |
| 6 | PROGRESS-V2-DECISION-001 | completed | v1 한계 확인 후 progress schema v2 보류 결정 |

## Foundation Product Loop 완료 판단

2026-05-25 기준 이 중기 플랜은 completed다. FTUE는 첫 5분 route와 evidence로 고정됐고, Mission HUD/floating modal은 status line polish까지 마쳤다. `incident-005-command-choice`는 첫 command-choice playable로 승격됐으며, daily route는 progress v1 저장 변경 없이 primary 복구 대상과 이유를 보여준다. Progress schema v2는 아직 실제 병목이 아니므로 보류하고, 다음 중기 플랜은 Foundation playtest/UX polish 또는 content breadth 확장 중에서 선택한다.

## Foundation Product Loop 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| PLAN-REFRESH-006 | UI-HUD-003 완료 | completed: 중기 플랜, active FTUE ExecPlan | completed: `git diff --check` | 새 기능 구현 없이 roadmap만 정리 |
| FTUE-001 | current playable route 존재 | completed: 첫 5분 route, 학습 감정 곡선, representative evidence | completed: focused E2E | 새 command/content schema 추가 금지 |
| UI-PLAYTEST-001 | FTUE route 고정 | completed: 주요 running/failed/succeeded screen evidence와 status line polish | completed: renderer/playable tests, UI-focused E2E | split cockpit으로 바로 확장하지 않음 |
| CHOICE-PLAY-001 | command-choice docs 완료 | 첫 command-choice playable, content replay, focused E2E | content replay, focused E2E, full regression | 새 engine 기능 없이 기존 command만 사용 |
| DAILY-ROUTE-001 | review queue baseline 존재 | no-schema daily route UX와 command summary 후보 | progress/playable tests, focused E2E | progress 저장 포맷 변경 금지 |
| PROGRESS-V2-DECISION-001 | no-schema loop 한계 확인 | schema v2 승인안 또는 보류 결정 | docs review | 승인 전 `internal/progress/` 변경 금지 |

## Foundation 완료 후보 기준

- 첫 실행 5분 안에 `이동 → 생존 → 빠른 이동 → 작은 수정`의 효용이 드러난다.
- 튜토리얼은 command를 가르치고, incident와 command-choice는 command 선택 판단을 훈련한다.
- review/daily loop는 저장 포맷 변경 없이도 다시 플레이할 이유를 준다.
- progress schema v2는 실제 부족함을 evidence로 확인한 뒤 별도 승인한다.

## 다음 중기 플랜: Foundation Playtest / UX Polish

> 목표는 새 Vim 기능이나 저장 포맷을 늘리기 전에, 현재 Foundation이 실제 게임처럼 읽히고 플레이되는지 다지는 것이다.

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAN-REFRESH-007 | completed | Foundation Playtest / UX Polish 순서 고정 |
| 2 | UX-LANGUAGE-001 | completed | 성공 modal의 다음 행동 문구를 tutorial/incident 맥락에 맞게 정리 |
| 3 | PLAYTEST-EVIDENCE-001 | completed | FTUE, incident, command-choice 대표 루트 evidence 점검 |
| 4 | HUD-DENSITY-001 | completed | Mission HUD와 daily/review line의 밀도/줄바꿈 polish |
| 5 | HELP-AFFORDANCE-001 | completed | `?` hint/help affordance와 실제 입력 처리 정합성 검증 |
| 6 | CHOICE-JUDGMENT-001 | completed | command-choice를 key 암기가 아니라 선택 판단 훈련으로 보강 |
| 7 | SUCCESS-MODAL-001 | active | success modal heading과 record density 정리 |
| 8 | UX-BACKLOG-001 | planned | 다음 content breadth로 넘어가기 전 playtest backlog 정리 |
| 9 | CONTENT-BREADTH-001 | planned | 기존 엔진만 사용해 command-choice/applied incident 폭 확장 |

## Foundation Playtest / UX Polish 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| UX-LANGUAGE-001 | Foundation Product Loop 완료 | completed: action copy helper, tutorial/incident E2E 갱신 | model tests, `make e2e-playable` | 저장 포맷과 playlist order 변경 금지 |
| PLAYTEST-EVIDENCE-001 | action copy 정리 완료 | completed: RedTeam representative route evidence review와 backlog | completed: focused E2E evidence | UX 주관 판단을 evidence 없이 문서화하지 않음 |
| HUD-DENSITY-001 | evidence에서 밀도 문제 확인 | completed: HUD line wrapping/copy polish | completed: renderer tests, focused E2E, `make e2e-playable` | console 접근성을 해치지 않음 |
| HELP-AFFORDANCE-001 | hint/help issue 확인 | completed: 첫 `?` hint 응답, retry/quit/mode affordance 정합성 테스트 | completed: runtime/scenario/playable tests, incident hint E2E, `make e2e-playable` | 검색 `?` 구현과 섞지 않음 |
| CHOICE-JUDGMENT-001 | command-choice가 암기처럼 보일 리스크 확인 | completed: 선택 이유 중심 copy와 E2E assertion | completed: content/playable tests, focused E2E, `make e2e-playable` | 새 schema/engine 없이 기존 command 선택 판단만 다룸 |
| SUCCESS-MODAL-001 | success modal 중복 리스크 확인 | 성공 modal heading/density polish | renderer/playable tests, focused E2E | 다음 행동 line이 잘리지 않게 유지 |
| UX-BACKLOG-001 | polish pass 완료 | 다음 중기 플랜 후보와 backlog | docs review | 새 engine 기능을 무계획으로 열지 않음 |
| CONTENT-BREADTH-001 | backlog 우선순위 확정 | 새 applied drill 또는 incident content + E2E | content replay, focused E2E | 기존 엔진으로 가능한 폭만 확장 |

## Foundation Playtest Evidence 판단

2026-05-26 RedTeam 플레이테스트 결과, 대표 route는 모두 통과했지만 HUD 밀도와 줄바꿈이 공통 P1 리스크로 확인됐다. 따라서 다음 slice는 `HUD-DENSITY-001`로 진행한다. Command-choice 판단 강화, success modal 중복 제거, help affordance 검증은 HUD 이후 후속 slice로 분리한다.
