# First 5-Minute Route

> FTUE-001 기준의 첫 실행 루프다. 목적은 처음 실행한 플레이어가 짧은 시간 안에 “Vim으로 문제를 수습하는 감각”을 얻는 것이다.

## Route

| 순서 | Playlist | Exercises | 학습 감각 |
|------|----------|-----------|-----------|
| 1 | `tutorial-0-movement` | 4 | `h/j/k/l`로 홈 포지션 이동 감각을 만든다. |
| 2 | `tutorial-1-survival` | 3 | `esc`, `:q!`, `:wq`로 터미널에서 당황했을 때의 생존 루틴을 익힌다. |
| 3 | `tutorial-2-fast-navigation` | 7 | `w/b/e`, `gg/G`, `0/$`로 한 글자씩 걷지 않는 이동 효율을 체감한다. |
| 4 | `tutorial-3-small-edits` | first beat | `x`로 작은 오타를 즉시 수습하며 “이제 편집도 할 수 있다”는 첫 성공을 준다. |

## Include / Exclude

포함:

- 이동: `h`, `j`, `k`, `l`
- 생존: `esc`, `:q!`, `:wq`
- 빠른 이동: `w`, `b`, `e`, `gg`, `G`, `0`, `$`
- 작은 수정 첫 맛보기: `x`

제외:

- `r`, `i/a/A`, `u/ctrl+r` 전체 작은 수정 튜토리얼
- Ex command substitute
- operator/yank/text-object/search/visual
- incident run
- command-choice drill

## Player Arc

1. **나는 방향키 없이 움직일 수 있다.**
2. **터미널에서 꼬였을 때 빠져나올 수 있다.**
3. **한 글자씩 움직이지 않아도 된다.**
4. **작은 오타 하나는 Normal mode에서 바로 수습할 수 있다.**

## Evidence

- 대표 fixture: `test/e2e/playable_ftue_first_five_route.yaml`
- evidence output: `artifacts/e2e/playable_ftue_first_five_route/`
- 기존 long regression: `test/e2e/playable_full_first_five_minute.yaml`

## Follow-Up Inputs For UI-PLAYTEST-001

- FTUE running 화면에서 `MISSION` cue line이 너무 길게 보이는지 확인한다.
- `:q!`/`:wq` command mode 화면에서 floating modal과 command prompt가 충돌하지 않는지 확인한다.
- `tutorial-2-fast-navigation`은 7문항으로 길기 때문에 첫 5분 안에서 피로가 생기는지 점검한다.
- 첫 `x` 성공 modal이 “작은 수정도 가능하다”는 감각을 충분히 주는지 확인한다.
