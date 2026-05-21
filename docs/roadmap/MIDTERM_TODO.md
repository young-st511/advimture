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
| 6 | VIM-024 | active | `.` 최소 subset 구현 |
| 7 | PLAYPACK-007 | pending | repeat-last-change efficiency tutorial content/E2E 구현 |
| 8 | SEARCH-GAP-001 | pending | literal `/`, `n`, `N` search 범위와 `?` hint 충돌 처리 결정 |
| 9 | VIM-025 | pending | literal search state와 next/previous match 구현 |
| 10 | PLAYPACK-008 | pending | search-basic tutorial content/E2E 구현 |
| 11 | PLATFORM-RFC-001 | pending | mastery, spaced review, daily run, progress schema 후보 RFC |

## Utility Commands and Long-Run Platform 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| OPEN-LINE-001 | `open-line-edit`가 curriculum next 후보 | command cluster draft/approval packet, VIM-023/PLAYPACK-006 분리 계획 | 문서 리뷰, `git diff --check` | indentation, auto-comment, count prefix, insert-mode Enter, dot repeat 제외 |
| VIM-023 | OPEN-LINE-001 완료 | completed: `o/O` engine, tuiadapter uppercase mapping, runtime replay smoke | completed: `go test ./internal/vimengine/...`, `go test ./internal/tuiadapter/...`, `go test ./internal/runtime/...` | content/E2E와 섞지 않고 engine contract만 닫았다 |
| PLAYPACK-006 | VIM-023 완료 | completed: `open-line-edit` YAML content, scenario, playlist, full E2E | completed: content replay, `playable_open_line_full.yaml`, forbidden-route E2E, `make e2e-smoke` | 8문항 이하, `o/O` required key와 우회 금지를 고정했다 |
| DEBRIEF-001 | PLAYPACK-006 완료 | completed: 성공/playlist 완료 화면 debrief, 기존 progress 기반 best record 표시 | completed: playable model tests, focused E2E | `internal/progress/` 저장 JSON 구조를 변경하지 않았다 |
| REPEAT-GAP-001 | open-line playpack과 debrief 완료 | completed: last-change transaction RFC, 최소 subset 결정 | completed: docs review | `.` 구현 전 undo/insert/yank/put transaction 경계를 먼저 고정했다 |
| VIM-024 | REPEAT-GAP-001 승인 | `.` 최소 subset engine/runtime 구현 | vimengine/runtime tests | macro/register 수준으로 확장하지 않는다 |
| PLAYPACK-007 | VIM-024 완료 | repeat-last-change tutorial + efficiency run E2E | content replay, full playlist E2E | 수동 재입력 우회를 constraints로 차단 |
| SEARCH-GAP-001 | repeat playpack 완료 | literal search scope, `/` command-line/search state, `?` 보류 결정 | docs review | regex, `?`, highlight 검증을 첫 구현에 넣지 않는다 |
| VIM-025 | SEARCH-GAP-001 승인 | `/`, `n`, `N` literal search engine/runtime 구현 | vimengine/runtime/tuiadapter tests | hint `?` 충돌을 우회 구현하지 않는다 |
| PLAYPACK-008 | VIM-025 완료 | search-basic tutorial + E2E | content replay, search app_state E2E | 검색 결과는 cursor/mode/key trace로 검증한다 |
| PLATFORM-RFC-001 | 최소 3개 utility playpack 완료 | mastery/review/daily run progress 후보 RFC | 저장 포맷 영향 리뷰 | 사용자 승인 없이 progress schema 변경 금지 |

## 2~3개월 주차 계획

| 기간 | 목표 | 닫는 루프 |
|------|------|----------|
| Week 1-2 | `o/O` 범위 확정과 엔진 구현 | OPEN-LINE-001, VIM-023 |
| Week 3-4 | open-line playable 콘텐츠와 저장 변경 없는 debrief | PLAYPACK-006, DEBRIEF-001 |
| Week 5-6 | `.` repeat-last-change gap planning과 최소 엔진 | REPEAT-GAP-001, VIM-024 |
| Week 7-8 | repeat tutorial과 efficiency run 검증 | PLAYPACK-007 |
| 이후 | literal search와 장기 반복 학습 RFC | SEARCH-GAP-001, VIM-025, PLAYPACK-008, PLATFORM-RFC-001 |
