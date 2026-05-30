# Vim Curriculum Map

> Advimture가 장기적으로 다룰 Vim 기능 지도다. 이 문서는 command catalog를 대체하지 않는다. 새 command cluster를 추가할 때 어느 chapter에 속하는지 판단하는 상위 지도다.

## 운영 규칙

- 이 문서는 넓은 지도이고, 구현 대상은 `command-catalog.md`의 cluster로 승격된 뒤 결정한다.
- 각 chapter는 플레이어가 실제 Vim을 쓰는 능력 단위로 나눈다.
- 각 cluster는 `command -> exercise -> scenario` 워크플로우를 통과해야 한다.
- 초반 콘텐츠는 8문항 이하의 짧은 tutorial episode playlist로 나눈다.
- 초반 튜토리얼은 “첫 투어”이고, 중후반부터 생존 어드벤처와 탐험 비중을 높인다.
- Ex substitute 계열은 초반 튜토리얼에서 빼고 중반 고급 튜토리얼로 분리한다.

## Chapter 0. Survival

목표: Vim에 들어갔을 때 당황하지 않고 모드, 저장, 종료를 통제한다.

후보 cluster:

- `survival-save-quit`: `esc`, `:q!`, `:wq`
- `insert-mode-entry`: `i`, `a`, `A`, `o`, `O`, `esc`
- `command-line-basics`: `:w`, `:q`, `:wq`, `:q!`

Exercise 방향:

- 수정하지 않고 종료하기
- 작은 수정 후 저장 종료하기
- Insert mode에서 Normal mode로 복귀하기

Scenario 방향:

- 낯선 설정 파일에서 안전하게 빠져나오기
- 확인된 변경을 저장하고 세션 닫기

## Chapter 1. Movement

목표: Normal mode에서 위치를 빠르고 정확하게 잡는다.

후보 cluster:

- `normal-motion-basic`: `h`, `j`, `k`, `l`
- `line-motion-basic`: `0`, `^`, `$`
- `word-motion-basic`: `w`, `b`, `e`, `ge`
- `file-motion-basic`: `gg`, `G`, `{count}G`
- `find-in-line`: `f`, `F`, `t`, `T`, `;`, `,`

Exercise 방향:

- 짧은 목표 좌표로 이동하기
- 줄 시작/끝으로 이동하기
- 단어 시작/끝으로 이동하기
- 줄 안 특정 문자로 점프하기

Scenario 방향:

- 로그 줄에서 경고 위치 찾기
- 설정 줄에서 옵션 이름/값 찾기
- 긴 한 줄에서 반복 키 입력을 줄이기

## Chapter 2. Small Edits

목표: 작은 텍스트 수정을 빠르게 수행한다.

후보 cluster:

- `single-char-edit`: `x`, `r`
- `insert-around-cursor`: `i`, `a`, `I`, `A`
- `open-line-edit`: `o`, `O`
- `undo-redo-basic`: `u`, `ctrl+r`
- `repeat-last-change`: `.`

Exercise 방향:

- 잘못된 문자 하나 수정하기
- 줄 끝에 옵션 추가하기
- 같은 수정 반복하기
- 실수 되돌리기

Scenario 방향:

- 설정 오타 하나 고치기
- 환경변수 값 끝에 suffix 추가하기
- 반복된 로그 prefix 정리하기

## Chapter 3. Operator Grammar

목표: Vim의 핵심 문법인 `operator + motion/text object`를 익힌다.

후보 cluster:

- `delete-with-motion`: `d`, `dw`, `d$`, `dd`
- `change-with-motion`: `c`, `cw`, `c$`, `cc`
- `yank-put-basic`: `y`, `yy`, `p`, `P`
- `text-object-inner`: `iw`, `i"`, `i'`, `i(`, `i{`
- `text-object-around`: `aw`, `a"`, `a'`, `a(`, `a{`

Exercise 방향:

- 단어 삭제/변경하기
- 따옴표 안 값 변경하기
- 괄호 안 인자 바꾸기
- 줄 복사 후 붙여넣기

Scenario 방향:

- config value 교체
- 함수 인자 수정
- 반복 설정 블록 복제

## Chapter 4. Search And Replace

목표: 원하는 위치와 패턴을 빠르게 찾고 바꾼다.

후보 cluster:

- `search-basic`: `/`, `?`, `n`, `N`
- `substitute-line`: `:s/foo/bar/`
- `substitute-file`: `:%s/foo/bar/g`
- `range-command-basic`: `:3,8s`, visual range substitute
- `ex-command-navigation`: `:set`, `:nohlsearch`, `:help`, command-line editing basics

Exercise 방향:

- 에러 키워드 찾기
- 한 줄에서 값 치환하기
- 파일 전체의 오래된 이름 바꾸기
- 선택 범위만 치환하기

Scenario 방향:

- 로그에서 장애 원인 찾기
- 서비스 이름 일괄 변경
- 특정 블록만 업데이트

## Chapter 5. Structure And Selection

목표: 보이는 범위와 구조를 이용해 여러 줄/블록을 다룬다.

후보 cluster:

- `visual-char-line`: `v`, `V`
- `visual-block`: `<C-v>`
- `indent-basic`: `>`, `<`, `=`
- `matching-pair`: `%`
- `join-split-basic`: `J`

Exercise 방향:

- 여러 줄 선택 후 들여쓰기
- 컬럼 단위 접두어 추가
- 괄호 짝으로 이동하기
- 줄 합치기

Scenario 방향:

- YAML indentation 정리
- 여러 config row에 prefix 추가
- 괄호 mismatch 찾기

## Chapter 6. Navigation At Scale

목표: 파일, 버퍼, 창 사이를 이동한다.

후보 cluster:

- `buffer-navigation`: `:e`, `:bnext`, `:bprev`, `:bd`
- `window-navigation`: `:split`, `:vsplit`, `<C-w>h/j/k/l`
- `jump-list-basic`: `<C-o>`, `<C-i>`
- `mark-basic`: `m`, `'`, `` ` ``

Exercise 방향:

- 다른 파일로 이동하기
- split 사이 이동하기
- 이전 위치로 돌아가기
- 중요한 위치 mark 후 복귀하기

Scenario 방향:

- config와 log 사이 오가기
- reference 파일 보며 수정하기
- 장애 위치를 mark로 기억하기

## Chapter 7. Automation

목표: 반복 작업을 Vim 문법으로 자동화한다.

후보 cluster:

- `macro-basic`: `q`, `@`, `@@`
- `global-command-basic`: `:g/pattern/command`
- `register-basic`: `"`, named registers
- `numbered-repeat`: `{count}` prefix

Exercise 방향:

- 같은 수정 여러 줄에 반복하기
- 패턴이 있는 줄만 삭제하기
- register로 값 보존하기
- count로 반복 이동/편집 줄이기

Scenario 방향:

- 여러 서비스 블록에 같은 변경 적용
- 특정 로그 라인 정리
- 반복 config migration

## First Curriculum Cut

현재 playable 플랫폼은 아래 tutorial episode 순서를 따른다.

1. Tutorial 0: 커서 감각 회상 — `normal-motion-basic` (`h`, `j`, `k`, `l`)
2. Tutorial 1: Vim 생존 키트 — `survival-save-quit` (`esc`, `:q!`, `:wq`)
3. Tutorial 2: 빠르게 훑기 — `word-motion-basic`, `line-motion-basic`, `file-motion-basic`
4. Tutorial 3: 작은 수정 — `single-char-edit`, `insert-mode-entry`, `undo-redo-basic`
5. Mid tutorial: Ex command 고급 튜토리얼 — `:s`, `:%s`, range substitute
6. Mid tutorial: Operator grammar 입문 — `dw`, `d$`, `dd`, `cw`, `c$`, `cc`
7. Mid tutorial: Yank / put 재사용 — `yy`, `yw`, `y$`, `p`, `P`
8. Mid tutorial: Text object inner word — `diw`, `ciw`, `yiw`
9. Mid tutorial: Literal search — `/`, `n`, `N`
10. Mid tutorial: Quote text object — `ci"`, `di"`, `yi"`
11. Mid tutorial: Visual selection — `v`, charwise `d/y`

최근 playable milestone은 아래 순서로 완료됐다.

1. Structure editing: quote/pair text object — completed, 설정값, JSON 값, 인자 내부를 구조 기준으로 편집
2. Adventure middle: search and replace 응용 — completed, search + substitute + quote/pair 조합
3. Visual selection: completed, 같은 줄 charwise `v` + `d/y` tutorial과 full E2E까지 연결
4. Linewise visual: completed, `V` linewise selection과 applied incident까지 연결
5. Inline target motions: completed, `f/t`, `df/dt/cf/ct` tutorial과 command-choice 적용까지 연결
6. Applied mastery runs: completed, `incident-006`, quote reuse choice, `incident-007` mixed run까지 연결

이 순서는 “첫 투어 -> 안전감 -> 효율 체감 -> 작은 수정 -> Vim 문법 -> 복사/재사용 -> 구조 대상 편집 -> 중반 고급 명령”으로 이어진다.

중반부터는 `docs/gameplay/scenario-tone.md`의 터미널 문제 해결 생존 어드벤처 톤을 따른다. 생존감은 command 학습을 가리지 않는 배경 압력이며, 다음 학습 단위는 `PLAN-REFRESH-009`에서 Foundation exit review 후 고른다.

## Coverage Rubric

- Chapter 0-4는 다음 playable milestone 전에 최소 draft cluster가 있어야 한다.
- approved cluster는 `coverage_required`의 모든 command가 exercise optimal trace에 등장해야 한다.
- 초반 기본 이동 cluster는 방향 감각 자체가 목표이므로 모든 방향 command를 optimal trace에 포함해야 한다.
- 후반 범용 이동/조합 cluster는 모든 Vim 기능을 억지로 한 루프에 넣지 않고, cluster의 `coverage_required`로 핵심 학습 범위를 정의한다.
- exact tier cluster는 implemented 승격 전에 engine unit test 또는 oracle comparison 계획이 있어야 한다.
- chapter가 넓어질수록 scenario보다 replay 가능한 exercise coverage를 먼저 늘린다.
- coverage matrix는 `cluster -> trained commands -> exercise count -> replay status -> oracle status -> e2e status` 순서로 본다.

### Priority Bands

| Band | 의미 | 현재 cluster |
|------|------|--------------|
| foundation | 이미 playable path에 연결되어 다음 콘텐츠의 선행 조건이 됨 | `survival-save-quit`, `normal-motion-basic`, `word-motion-basic`, `whole-file-navigation`, `single-char-edit`, `insert-mode-entry`, `undo-redo-basic`, `vim-ex-command-substitute`, `delete-with-motion`, `change-with-motion`, `yank-put-basic`, `text-object-inner-word`, `open-line-edit`, `repeat-last-change`, `search-basic`, `text-object-quote-pair`, `visual-char-line`, `visual-line-basic`, `char-find-line` |
| decision | 다음 계획에서 먼저 결정할 것 | `PLAN-REFRESH-009` Foundation exit review |
| soon | 다음 milestone 후보이나 review 후 선택 | quote text object hardening, 저장 변경 없는 platform/review loop, command-choice breadth |
| later | 중반 이후 어드벤처나 고급 튜토리얼에서 다룸 | visual block, macro/register/count, buffer/window/navigation-at-scale 계열 |

### Current Planning Candidate

ID: `foundation-exit-review`

목표: foundation 출시 후보 범위와 다음 중기 플랜을 고른다. 바로 새 command를 늘리기보다, 현재 구현된 command coverage가 첫 출시/베타 루프에 충분한지 먼저 판단한다.

선택 후보:

| Candidate | Layer | 이유 |
|-----------|-------|------|
| `quote-pair-hardening` | engine hardening | `ci'`, `ci(`, `ci{`는 실무 config/JSON/함수 인자 편집에 유용하고 기존 quote object의 작은 확장이다. |
| `platform-review-loop` | gameplay system | 장기 반복 학습 플랫폼으로 가려면 저장 변경 전에도 mission/review loop를 더 게임답게 묶어야 한다. |
| `command-choice-breadth` | content/application | 새 engine 없이 이미 배운 command 선택 판단을 더 훈련할 수 있다. |

### Command Choice Layer

`command-choice-drill`은 새 command cluster가 아니라 이미 배운 command 중 상황에 맞는 도구를 고르는 applied layer다.

| Layer | Commands | Engine support | Oracle | 이유 |
|-------|----------|----------------|--------|------|
| `command-choice-drill` | 기존 implemented cluster 조합 | implemented only | not needed | Vim 학습의 다음 병목은 새 key를 외우는 것이 아니라 범위/반복/검색/치환 상황에서 적절한 도구를 고르는 판단이다. |

완료된 beat:

- `scope-choice`: `ciw`, `ci"`, `v...d`, `V...d` 중 편집 범위에 맞는 도구를 고른다.
- `range-choice`: visual/operator 조작과 substitute/range command 중 더 적합한 방법을 고른다.
- `inline-target-choice`: comma/quote/delimiter 보존 여부를 보고 `ct,`와 `cf,` 중 적절한 범위를 고른다.
- `quote-value-reuse`: retype 대신 `yi"` + `P`로 검증된 quote 내부 값을 재사용한다.

후속 후보:

- line reuse: 검증된 줄 전체를 `V` + `y` + `p`로 재사용한다.
- repeat-change reuse: 같은 변경을 `.`로 반복할지 판단한다.
- search-then-act: `/`, `n`, `N`으로 위치를 찾은 뒤 적절한 편집 command를 고른다.

권장 문항 수:

- 첫 playable 승격은 4문항 이하
- 새 command 소개 없이 기존 tutorial과 incident에서 배운 command만 사용

설계 제약:

- command choice는 command cluster가 아니라 적용 레이어로 유지한다.
- 새 schema 승인 전까지 authoring rubric은 문서에 두고, playable enforcement는 기존 constraints로 처리한다.
- scenario는 정답 key sequence보다 선택 이유를 강화한다.
- 자세한 기준은 `docs/gameplay/command-choice-drills.md`를 따른다.

## Known Coverage Gaps

- `open-line-edit`: PLAYPACK-006에서 `o`, `O` 기본 흐름을 다뤘다. indentation, auto-comment, count prefix, dot repeat 연계는 후속 hardening으로 남는다.
- `repeat-last-change`: PLAYPACK-007에서 `.` 반복 기본 흐름을 다뤘다. delete/yank/put/search/macro/register/count prefix는 후속 hardening이다.
- `search-basic`: SEARCH-GAP-001에서 `/`, `n`, `N` literal search로 첫 scope를 고정했고, VIM-025/PLAYPACK-008에서 engine과 tutorial을 연결했다. `?`, regex, highlight, search history는 후속 hardening으로 남는다.
- `platform-review-loop`: mastery, spaced review, daily run은 progress schema 변경 가능성이 있어 RFC와 사용자 승인이 필요하다.
- `text-object-quote-pair`: PLAYPACK-009에서 double quote 내부 object를 연결했다. nested pair, escaped quote, around object, count prefix, visual selection은 후속 hardening이다.
- `visual-char-line`: PLAYPACK-010에서 같은 줄 charwise selection delete/yank tutorial까지 연결했다. multi-line charwise visual, visual block, count/register prefix는 후속 hardening이다.
- `visual-line-basic`: PLAYPACK-011에서 linewise selection delete/yank tutorial까지 연결했다. multi-line charwise visual, visual block, count/register prefix는 후속 hardening이다.
- `command-choice-drill`: COMMAND-CHOICE-001에서 docs-only 설계를 완료했고, incident-005에서 linewise scope, range-choice, inline-target-choice, quote value reuse beat를 playable로 연결했다. 후속 후보는 line reuse, repeat-change reuse, search-then-act다.
- `char-find-line`: CHAR-FIND-GAP-001에서 forward same-line `f/t`와 `df/dt/cf/ct` 첫 scope를 고정했고, VIM-030/PLAYPACK-012에서 engine과 tutorial을 연결했다. `F/T`, `;`, `,`, count prefix, visual mode, yank 결합은 후속 hardening이다.

## Long-Run Platform Direction

Advimture는 단기 데모보다 장기 반복 학습 플랫폼을 목표로 한다. 단, 첫 게임성 강화는 새 저장 포맷이 아니라 기존 score/progress를 읽어 성공 화면과 playlist 완료 화면에서 debrief, best record, retry 동기를 제공하는 수준으로 제한한다.

현재 우선순위:

1. `foundation-exit-review`: 현재 engine/content/UI/E2E가 첫 출시 후보로 충분한지 판단한다.
2. `platform-review-loop`: progress schema 변경 없이 mission/review/daily motivation을 더 게임답게 묶는다.
3. `quote-pair-hardening`: 새 engine을 연다면 기존 quote object의 작은 확장으로 제한한다.
4. `command-choice-breadth`: 이미 배운 command를 섞어 범위/반복/검색/치환/inline target 중 적합한 도구를 고르는 판단을 훈련한다.
5. `progress-schema-v2`: mastery/spaced review/daily run 저장은 실제 병목이 evidence로 확인된 뒤 사용자 승인으로만 연다.

세계관은 `원격 시설 복구국 / Runbook Dispatch`를 유지하되, lore 확장보다 runbook 작전감과 잔류 리스크/재점검 언어를 활용한다. briefing은 `상황 1문장 + Vim 조작 목표 1문장`을 기본으로 유지한다.
