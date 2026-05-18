# Vim Curriculum Map

> Advimture가 장기적으로 다룰 Vim 기능 지도다. 이 문서는 command catalog를 대체하지 않는다. 새 command cluster를 추가할 때 어느 chapter에 속하는지 판단하는 상위 지도다.

## 운영 규칙

- 이 문서는 넓은 지도이고, 구현 대상은 `command-catalog.md`의 cluster로 승격된 뒤 결정한다.
- 각 chapter는 플레이어가 실제 Vim을 쓰는 능력 단위로 나눈다.
- 각 cluster는 `command -> exercise -> scenario` 워크플로우를 통과해야 한다.
- 첫 5분 콘텐츠는 Chapter 0-1의 극초반에만 해당한다.

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
- `undo-redo-basic`: `u`, `<C-r>`
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

초기 playable 플랫폼은 아래 순서를 우선한다.

1. `normal-motion-basic`
2. `survival-save-quit`
3. `word-motion-basic`
4. `line-motion-basic`
5. `single-char-edit`
6. `insert-mode-entry`
7. `delete-with-motion`
8. `search-basic`

이 순서는 “첫 성공 경험 -> 안전감 -> 효율 체감 -> 작은 수정 -> Vim 문법”으로 이어진다.

## Coverage Rubric

- Chapter 0-3은 첫 playable milestone 전에 최소 draft cluster가 있어야 한다.
- approved cluster는 `coverage_required`의 모든 command가 exercise optimal trace에 등장해야 한다.
- 초반 기본 이동 cluster는 방향 감각 자체가 목표이므로 모든 방향 command를 optimal trace에 포함해야 한다.
- 후반 범용 이동/조합 cluster는 모든 Vim 기능을 억지로 한 루프에 넣지 않고, cluster의 `coverage_required`로 핵심 학습 범위를 정의한다.
- exact tier cluster는 implemented 승격 전에 engine unit test 또는 oracle comparison 계획이 있어야 한다.
- chapter가 넓어질수록 scenario보다 replay 가능한 exercise coverage를 먼저 늘린다.
- coverage matrix는 `cluster -> trained commands -> exercise count -> replay status -> oracle status -> e2e status` 순서로 본다.

## Known Coverage Gaps

- `normal-motion-basic`: 현재 draft는 `l`, `j` 중심이다. `h`, `k` optimal trace exercise가 필요하다.
- `survival-save-quit`: command-line mode와 app exit semantics가 필요하다.
- `word-motion-basic`: `w/b/e` engine support와 word boundary oracle fixture가 필요하다.
- `file-motion-basic`: `gg`, `G`, `{count}G`는 후반 navigation-at-scale 전에 필요하다.
- `ex-command-navigation`: `:` command-line 입력, range, substitute를 폭넓게 다루기 위한 별도 engine/runtime 설계가 필요하다.
