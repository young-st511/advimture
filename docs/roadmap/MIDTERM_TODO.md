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
| 3 | VIM-013 | planning | `x`, `r` buffer mutation engine 구현 |
| 4 | VIM-014 | pending | `i`, `a`, `A` insert mode entry와 printable insertion 구현 |
| 5 | VIM-015 | pending | `u`, `<C-r>` undo/redo stack 구현 |
| 6 | PLAYPACK-002 | pending | 6~8문항짜리 “편집 기본기” 튜토리얼 playpack 설계/구현 |
| 7 | SCENARIO-TONE-001 | pending | 중반 생존 어드벤처 톤 원칙과 적용 범위 확정 |
