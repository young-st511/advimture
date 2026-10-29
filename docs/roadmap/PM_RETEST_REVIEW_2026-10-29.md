# PM 재실험: 수정 후 같은 기준 직접 플레이

Date: 2026-10-29
Baseline: `ad56412` (PR #2의 구현 버전, VERSION 0.4.0)
Status: 감사 완료 / 후속 변경 제안은 `[draft]`
Previous: `PM_PLAYTEST_REVIEW_2026-09-29.md`

## 결론

기존 특수키 오염, 실패 모드의 Enter/힌트, 게임 본문의 묶음 입력, 정상 대안 풀이, 실패 모달 footer는 실제 사용에서도 개선됐다. 첫 복구 경로는 4개 연습 후 원인 추적/상태 복구까지 진행할 수 있다.

그러나 기존 12개 이슈를 전부 해결했다고 판정할 수는 없다. 이번 재판정은 **6개 해결 확인, 5개 부분 해결, 1개는 경로 구현 확인 후 사용자 효과 미측정**이다. 특히 checkpoint와 복수 복습 경로의 검증이 이전 완료 판정에서 부족했다.

추가로 잔여/회귀 개선점 8개를 정리했다. P0는 관찰하지 못했다. P1은 메뉴 입력 손실 1개와 첫 학습의 오타 복구 정책 1개이며, 나머지는 P2다. P1 정책 문제는 현재 max_inputs 명세 위반을 뜻하지 않는다.

## 방법과 검증 범위

- 같은 기준: 실제 바이너리/독립 HOME/PTY 키 입력, final frame와 timeline, app_state, 이전 12개 문제의 재현, 첫 실행/실패/힌트/재시도/응용/복습을 확인했다.
- 직접 PTY 세션 3개: 첫 복구 전체(오타/재시도/상세 힌트 포함), 저등급 2문항 복습, incident 008의 긴 힌트. 실제 로그는 `manual-first.raw`, `manual-review.raw`, `manual-long-hint.raw`다.
- 80x24를 기본으로 실패 footer와 긴 힌트는 110x34와 비교했다. 기존 full route의 원래 viewport도 유지했다.
- 기존 회귀 52개를 새로 빌드한 바이너리로 다시 실행해 52개 모두 통과했다. 추가 탐색 28개는 관찰한 동작을 재현했다. 새 품질 기대와의 비교 8개는 모두 미충족으로 기록했다. **관찰 assertion 통과와 제품 개선 기대 충족은 별개다.**
- `go test -count=1 ./...`를 실행했다. 제품 코드/콘텐츠/의존성/저장 포맷은 이번 감사에서 변경하지 않았다.
- 실제 외부 참가자 6~8명의 학습/재방문 실험은 미실시다. Agent가 이미 풀이를 아는 상태의 조작 시간으로 5분 목표나 학습 효과를 평가하지 않는다. 완료 진행도 fixture도 실제 사용자 플레이 이력으로 취급하지 않는다.

## 이전 12개 이슈 재판정

| 이전 ID | 재판정 | 이번 확인 |
|---|---|---|
| PM-01 특수키 오염 | 해결 확인 | 실제 DEL로 명령 오타 수정 후 성공, Left가 문자로 들어가지 않음, literal left 붙여넣기도 구별 |
| PM-02 실패 모드 action | 해결 확인 | insert/visual/search/command 실패에서 힌트/Enter 재시도 확인 |
| PM-03 묶음 입력 | 부분 해결 | 게임의 ll은 성공. 학습 목록 jj는 여전히 무시됨(R1) |
| PM-04 정상 대안 거부 | 해결 확인(대표 범위) | yyjp/Vyjp 모두 S. 기본 회귀의 재입력/구분자/범위 RedTeam도 통과. 전체 Vim 풀이 공간을 검증했다는 뜻은 아님 |
| PM-05 첫 incident까지 79문항 | 경로 구현 확인 / 효과 미측정 | m→첫 복구로 4개 연습 뒤 작전 진입. 자발적 발견/첫 복구 시간은 외부 사용자 검증 필요 |
| PM-06 첫 목표 누락 | 해결 확인 | 처음 화면에서 목표 c를 읽을 수 있음 |
| PM-07 실패+힌트 footer 잘림 | 해결 확인(원래 재현) | 80x24/110x34 모두 retry/hint/quit 표시. 별도 running 긴 힌트에서 header 손실 발견(R6) |
| PM-08 복습 종료 | 부분 해결 | 1개를 S로 개선하면 종료. 복수 문항은 counter가 고정되고 도움받은 완료가 재순환(R4) |
| PM-09 힌트 비용/효용 | 부분 해결 | 무료 첫 힌트 후 S, 상세 반복 후 A로 bounded. 무료/등급 영향 표시는 충돌(R5) |
| PM-10 긴 보고서/행동 집중 | 부분 해결 | 짧은 경로에서 결과 buffer와 checkpoint 유지. 경로/episode counter 혼재, 성공의 종료 안내 부족(R3/R7) |
| PM-11 새 학습을 복습 부채로 표시 | 부분 해결 | 일반 성공 보고서는 다음 새 미션으로 바뀜. checkpoint에서는 재점검으로 다시 표시(R3) |
| PM-12 README 신선도 | 해결 확인 | stable/현재 소스 개선을 구분하고 낡은 콘텐츠 수치 제거 |

## 우선순위

| ID | 우선순위 | 분류 | 발견 | 다음 변경의 완료 기준 초안 |
|---|---|---|---|---|
| R1 | P1 | 입력 버그 | 메뉴 jj는 무시되고 j,j만 이동 | 묶음/개별 메뉴 입력이 같은 선택에 도달 |
| R2 | P1 | 학습 정책 | 첫 복구 연습에서 오타를 수정할 입력 여유가 없음 | 첫 연습의 타이핑 오타를 Backspace로 고쳐 완료 가능 |
| R3 | P2 | 정보 설계/표시 누락 | checkpoint가 미학습을 재점검으로 표시하고 다른 episode counter가 혼재 | 현재 경로의 다음 단계/진행 수만 primary로 표시 |
| R4 | P2 | counter 버그/복습 정책 | 다음 문항에서도 복습 1/2, 도움받은 완료 2개가 재순환 | session 진행 수 갱신, 이번 연습 완료와 숙련 재추천 구별 |
| R5 | P2 | 비용 affordance | 무료와 등급에 영향이 같은 힌트에 동시 표시 | 실제 무료/유료 단계와 화면 안내 일치 |
| R6 | P2 | 높이 처리 버그 | 80x24 긴 running 힌트에서 header 소실 | header/목표/console/action에 높이 예산 확보 |
| R7 | P2 | 성공 action 누락 | 성공에서 Enter만 보이고 작동하는 q/m은 숨어 있음 | 성공 decision surface에 종료/경로 선택도 표시 |
| R8 | P2 | QA 관찰 계약 | 화면은 메뉴/checkpoint인데 ui.focus_panel은 training/success | 실제 활성 화면 identity를 기계 evidence에서도 식별 |

우선순위는 이전 감사와 동일한 PM 판단 기준이다. 실제 이탈률이나 이용자 불만 빈도를 측정해 정한 수치는 아니다.

## 상세 재현

### R1. 게임에서 해결한 묶음 입력 문제가 메뉴에 남았다

1. fresh start에서 m, 6으로 학습 목록을 연다.
2. 하나의 PTY write로 jj를 보내고 Enter.
3. 첫 에피소드 normal-motion-basic-001이 열린다. 같은 j를 나눠 두 번 보내면 word-motion-basic-001이 열린다.

메뉴 분기가 묶음 분해보다 먼저 실행되고, selectRoute는 문자열 j만 처리한다. 근거: `internal/playable/model.go:155`, `internal/playable/routes.go:18`.

이는 묶여 전달된 키/붙여넣기에 대한 실제 입력 손실이다. 모든 키보드의 빠른 타이핑이 반드시 하나의 이벤트로 묶인다는 주장은 하지 않는다. 조작이 반응하지 않았다는 안내도 없다.

Evidence: `retest_menu_batch`, `retest_menu_separate`의 key_trace/app_state 비교.

### R2. 삭제는 되지만 첫 연습의 오타 복구는 실패한다

1. 첫 복구 4번째 연습(text-object-inner-word-003)의 목표는 ciwup + Esc다.
2. ciwop로 오타를 내고 Backspace 두 번, up, Esc로 고치려 한다.
3. 7번째 입력에서 예산이 소진돼 실패한다. 버퍼는 status=o에 남고 수정 완료까지 갈 수 없다.

Backspace 기능은 정상이다. `max_inputs: 6`이 최적 풀이 6키와 같아 첫 오타부터 수정 여유가 없다. 직접 플레이에서도 같은 실패/재시도를 경험했다.

근거: `content/exercises/text-object-inner-word.yaml:98`, `:126`; `internal/runtime/session.go:141`. 현재 입력 제한 명세대로 실행된 것이므로 구현 버그로 보고하지 않는다. 첫 학습의 실수 복구와 최적화 도전의 기준을 나눌 제품 정책이 필요하다.

[draft] 첫 복구의 연습은 hard budget을 완화하고 최적 입력 수를 기록/도전 목표로 둔다. 고정 키 수 도전은 시작 전 고지한다. 기존 목표/보존 조건과 필수 개념은 유지한다.

Evidence: `retest_typo_budget`, `manual-first.raw`.

### R3. checkpoint에 이전 정보 설계가 다시 노출된다

첫 복구 첫 문제를 ll로 성공하면 결과 buffer는 유지된다. 하지만 상단에 다음과 같이 나온다.

```text
복구 메모: 재점검 3건 · 다음: 경고 지점으로 이동하기
Exercise: 1/4
경로 ... 첫 복구 1/6
```

아직 시도하지 않은 문항이 재점검으로 표시된다. 실제 이 경로의 다음은 경고 줄이 아니라 단어 이동이다. 다음 작전까지 완료한 마지막 화면은 "첫 출격을 마쳤습니다"와 Runbook 2/6을 함께 보여준다. 저장된 원래 episode 진행률은 맞아도 현재 짧은 경로의 완료 여부를 읽기 어렵다.

원인: `recoveryStatusLine`이 candidate reason과 현재 route를 보지 않고 count/title을 재점검으로 표시한다(`internal/playableview/render.go:303`). 원래 playlist-local Exercise/Runbook count도 새 경로에 그대로 노출된다.

[draft] 새 학습과 복습을 모든 렌더링 경로에서 구별하고, 현재 route 진행을 primary로 둔다. 원래 episode 기록은 상세 보고서로 내린다.

Evidence: `retest_checkpoint_debt/screen_final.txt`, `baseline/baseline_playable_first_recovery_route/screen_final.txt`, `manual-first.raw`.

### R4. 복습 완료/진행/숙련 추천이 섞여 있다

fixture는 normal-motion-basic-001/002가 완료/A인 두 문항이다.

- m→4는 두 번째 문항을 연다. j로 S를 받고 Enter로 다음 복습에 가도 footer는 복습 1/2다.
- 둘 다 힌트를 사용해 A로 완료하면 두 번째→첫 번째→두 번째로 돌아온다. 직접 플레이와 PTY에서 한 바퀴 재진입을 확인했다.
- 단일 후보를 S로 개선하면 경로 선택으로 돌아가는 기존 수정은 정상이다.

review 전환은 routePos를 증가시키지 않는다(`internal/playable/model.go:418`). 후보는 최고 등급 S가 아니면 계속 남으며, 이번 세션에서 이미 완료한 문항의 집합은 없다(`internal/review/review.go:126`).

counter 고정은 구현 결함이다. 도움받아 A로 완료한 문항을 즉시 다시 요구하는 것은 세션 완료 정책 문제다. S로 풀거나 q로 종료하는 경로가 있으므로 진행 불가/P0로 분류하지 않는다.

[draft] 이번 연습 완료와 향후 숙련 복습을 구별한다. 선택한 두 문항을 한 번씩 연습하면 세션을 닫고, 더 연습하기를 별도 선택으로 제공한다. counter는 실제 시도 순서로 갱신한다.

Evidence: `retest_review_counter`, `retest_review_assisted_cycle`, 정상 비교 `retest_review_finish`, `manual-review.raw`.

### R5. 개념 도움은 무료지만 화면은 감점을 경고한다

incident 첫 힌트는 [1/3 · 무료]이고 이를 사용한 최적 풀이가 S를 받는 것까지 확인했다. 같은 화면에는 "· 등급에 영향", action bar에는 "힌트 - grade 영향"이 남는다. 이 문구는 상세 단계의 잠재 비용을 개념 단계에서도 구별 없이 경고한다.

실제 비용 로직은 개선됐다. 남은 문제는 사용자가 도움을 누를지 결정할 때 읽는 계약이다. 근거: `internal/playable/model.go:935`, `internal/playableview/render.go:380`.

[draft] 개념 무료/상세 도움 사용을 단계별로 표시하고, 이미 공개한 힌트 재열람 비용도 일관되게 표시한다.

Evidence: `retest_free_hint_frame`, `retest_free_hint`, `retest_detail_hint`.

### R6. 긴 running 힌트에서는 80x24 높이 처리가 부족하다

incident-search-scope-001(5줄 buffer)에서 ?를 세 번 누른다. 80x24는 ADVIMTURE header가 사라지고 MISSION부터 보인다. 110x34에서는 header가 유지된다. 직접 PTY에서도 세 번째 힌트 후 header가 사라지는 것을 확인했다.

실패 모달의 footer 문제는 해결됐지만, running HUD 전체 높이를 예약하지 않아 새 hint/command memory/route footer가 화면을 넘는다(`internal/playableview/render.go:137`). final_screen_max_lines만 확인하면 terminal grid는 항상 24행 안에 들어가므로 이 손실을 잡지 못한다.

[draft] running hint 1/2/3단계와 최장 buffer 조합에서 header/현재 목표/console/필수 action의 위치를 검증한다. 보조 정보는 접거나 상세 보기로 이동한다.

Evidence: `retest_hint_long`, `retest_hint_long_wide`, `manual-long-hint.raw`.

### R7. 일반 성공에서 종료/경로 선택이 감춰진다

fresh start에서 ll로 성공하면 모달은 "다음 단계: enter"만 표시한다. q와 m은 동작하지만 성공 decision surface에 표시되지 않는다. 계속 진행을 중단하고 싶은 사용자는 앞 화면에서 키를 기억해야 한다.

근거: `internal/playable/model.go:624`의 같은 playlist/다음 playlist action 반환. 실패 모달에는 quit이 표시되는 것과 대비된다.

[draft] 성공도 다음 행동과 보조 경로 선택/종료를 노출하고 80x24에서 함께 보존한다.

Evidence: `retest_success_hidden_quit/screen_final.txt`와 q 입력의 정상 종료.

### R8. 새 화면을 app_state가 정확히 구분하지 못한다

- m 메뉴가 실제로 떠 있는데 ui.focus_panel.kind는 training이다.
- 첫 복구의 성공 checkpoint가 inline으로 보이는데 app_state는 success다.

View는 메뉴를 직접 반환하거나 화면 DTO만 checkpoint로 교체한다(`internal/playable/model.go:292`, `:322`). State는 run에서 focus panel을 다시 만들어 기록한다(`:365`). 따라서 active UI identity의 불일치다. 이로 인해 메뉴 상태/inline 전환이 machine assertion에서 빠진다.

[draft] 실제 활성 screen/panel identity를 공유한 DTO로 만들고 menu/checkpoint를 관찰 가능하게 한다. 저장 포맷에 반영할 필요는 없다. 종료 후 마지막 행이 지워질 수 있는 evidence도 active viewport와 구별한다.

Evidence: `retest_menu_state`의 screen_final/app_state, `retest_checkpoint_debt`의 screen_final/app_state.

## 다음 순서 제안

1. [draft] R1 메뉴 입력과 R2 첫 학습 오타 허용을 짧은 수정 묶음으로 처리한다. 정상 풀이/보존 조건은 그대로 검증한다.
2. [draft] R3/R4의 현재 경로와 복습 세션 정보를 정리하고, R5/R7의 도움 비용/종료 안내를 같은 화면 계약으로 맞춘다.
3. [draft] R6/R8을 고정 회귀에 추가한다. header와 실제 활성 UI 의미를 assert하고, 정상/실패 상태 text만으로 UI 성공을 판정하지 않는다.
4. 위 묶음을 닫은 뒤 준비된 실제 사용자 실험을 수행한다. 새로운 콘텐츠 확장보다 첫 세션/실제 Vim 전이/자발적 추가 플레이를 판단 근거로 삼는다.

이 제안은 이번 재실험의 결과이며 제품 수정 승인을 자동으로 의미하지 않는다. 이번 작업의 변경은 감사 문서와 검사 항목뿐이다.

## 재현 자료와 검사 결과

root: `artifacts/e2e/pm-retest/` (git ignored)

- `baseline-index.json`: 기존 52개 경로를 새 바이너리로 다시 실행한 결과, 52/52.
- `exploration/retest_*/`: 추가 탐색 28개. 각 디렉토리의 summary/key_trace/app_state/final/timeline/raw를 확인한다.
- `findings-check.json`: 위 8개 품질 기대에 대한 met=false 기록. R2 및 R4의 정책 항목은 기존 명세 위반과 구별한다.
- `check-observations.rb`: machine state/화면 비교를 재생성하는 관찰 스크립트.
- `manual-*.raw`: 직접 사용한 3개 세션.

로컬 재실행:

```sh
artifacts/e2e/pm-retest/bin/e2e-runner --scenario artifacts/e2e/pm-retest/scenarios/retest_menu_batch.yaml --artifacts artifacts/e2e/pm-retest/exploration
ruby artifacts/e2e/pm-retest/check-observations.rb
```

공유용 핵심 재현은 본문의 초기 fixture/키 입력만으로 가능하다. 예: R1은 fresh start → m → 6 → 단일 write jj → Enter, 비교는 j와 j를 분리한다. R4의 진행도 fixture는 다음과 같다.

```json
{"missions":{"normal-motion-basic-001":{"completed":true,"best_grade":"A"},"normal-motion-basic-002":{"completed":true,"best_grade":"A"}}}
```

감사 fixture의 수정 이력도 제품 실패와 구별했다. literal left 검사에서 runner의 send:left는 실제 방향키로 변환되므로 bracketed paste로 바꿔 재검증했다. 빈 복습 안내는 실행 중 timeline에 표시되지만 q 종료가 마지막 행을 지우므로 final-only assertion을 active 안내 판정으로 사용하지 않았다. 제품 버그로 추가 집계하지 않았다.
