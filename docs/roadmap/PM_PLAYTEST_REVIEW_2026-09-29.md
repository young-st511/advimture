# PM 직접 플레이 리뷰와 개선 기획

Date: 2026-09-29
Baseline: `aa685ce`, VERSION `0.4.0`
Status: 감사 완료 / 후속 구현 수용 기준은 `[draft]`
ExecPlan: `docs/exec-plans/active/pm-playtest-audit-001.md`

## 판단

Advimture에는 작동하는 Vim 훈련 기반과 복구 시나리오가 있다. 그러나 입력 실수를 복구하는 기본 동작에 결함이 있고, 실제 복구 작전을 경험하기까지 긴 튜토리얼을 거쳐야 한다. 현재 가장 큰 제품 과제는 콘텐츠 수가 아니라 **입력에 대한 신뢰, 빠른 첫 성취, 풀이를 인정하는 기준, 다시 돌아올 이유**다.

권장 순서는 `입력/실패 복구 → 첫 복구 경험 → 응용 채점 → 힌트/복습 → 외부 사용자 검증`이다. 이번 리뷰는 이 순서를 제안하는 문서이며, PRODUCT/PROGRAM의 승인된 방향을 자동으로 바꾸거나 구현 승인을 대신하지 않는다.

## 직접 사용한 범위와 한계

- 실제 Go 바이너리를 PTY에서 실행했다. 사용자 HOME과 분리한 세션에서 첫 화면, 힌트, 방향키 실패, 재시도, 일반 모드 묶음 입력, 이동 튜토리얼, 생존 첫 문항, 종료/재실행을 직접 조작했다.
- 별도 테스트 진행도 fixture로 incident 001에 진입해 힌트 사용, `/error`, `/timeout`, `n`, 성공 후 다음 단계와 종료를 직접 조작했다. 79개 튜토리얼 전체를 사람이 처음부터 플레이한 것은 아니다.
- 80x24와 110x34에서 대표 화면을 비교했다. 기존 first-five route는 원래의 120x40 설정도 유지했다. TUI 프로젝트 계약에 따라 브라우저 스크린샷 대신 PTY raw ANSI, terminal cell final frame, timeline, app_state, progress를 확인했다.
- 기존 runner를 이용한 기본 경로 5개와 관찰/결함 재현 20개를 실행했다. 25개 모두 정의한 관찰 assertion을 만족했다. **결함 재현의 pass는 버그가 재현됐다는 뜻이며 제품 품질 통과가 아니다.**
- 동일 기준 commit에서 전 턴에 `go test ./...`가 통과했다. 이번에는 앱/runner 빌드와 PTY 검증을 수행했으며 제품 코드는 수정하지 않았다.
- 전문가/Agent 평가다. 실제 초보자의 재미, 이해, 학습 전이, 재방문은 측정하지 않았다. 기본 경로 외 전체 콘텐츠와 모든 터미널 조합이 정상이라고 주장하지 않는다.

## 우선순위 요약

P0는 이번 범위에서 확인하지 못했다. P1은 첫 세션의 조작 신뢰 또는 핵심 학습 약속을 훼손하는 문제, P2는 경험과 반복 동기를 약화하는 문제, P3는 정보 정합성 문제다. 설계 문제의 우선순위는 PM 판단이며 이탈률 측정 결과가 아니다.

| ID | 우선순위 | 분류 | 발견 | 먼저 할 일 |
|---|---|---|---|---|
| PM-01 | P1 | 구현 버그 | Backspace/방향키 이름이 텍스트로 입력됨 | 특수키와 문자 입력을 분리 |
| PM-02 | P1 | 구현 버그 | Insert/Visual 실패 후 Enter 재시도가 무시됨 | 실패 UI action을 Vim mode보다 먼저 처리 |
| PM-03 | P1 | 구현 버그 | Normal 모드의 묶음 입력 `ll`이 조용히 사라짐 | 입력 이벤트 분할과 모드 전환 처리 검증 |
| PM-04 | P1 | 학습 설계 | 올바른 줄 복사 `yyjp`가 실패함 | 튜토리얼 제약과 응용 성공 판정 분리 |
| PM-05 | P1 | 제품 구조 | 첫 incident 전에 튜토리얼 79문항 | 짧은 학습-적용 경로와 진입 선택 설계 |
| PM-06 | P1 | 콘텐츠 UX | 첫 문제의 목표 문자 `c`가 화면에 없음 | 성공 목표를 명시하고 시각적으로 구분 |
| PM-07 | P2 | 렌더링 버그 | 80x24 실패+힌트에서 종료 action이 잘림 | footer 높이를 먼저 확보 |
| PM-08 | P2 | 복습 흐름 결함 | 복습이 끝나도 이미 완료한 다음 튜토리얼로 이동 | 복습 세션의 종료 조건 정의 |
| PM-09 | P2 | 코칭/채점 설계 | 같은 힌트 재열람도 감점, 성공했는데 F | 힌트 단계별 효용/비용과 성공 의미 정렬 |
| PM-10 | P2 | 인터랙션 UX | 1~2키 조작마다 긴 모달, 복구 결과는 가려짐 | 즉시 피드백과 에피소드 보고서 분리 |
| PM-11 | P2 | 정보 설계 | 아직 배우지 않은 문항이 재점검/잔류 리스크로 나옴 | 새 학습과 실제 복습을 구별 |
| PM-12 | P3 | 문서 정합성 | README의 출시 상태/콘텐츠 수가 낡음 | 현재 배포본 안내로 정리 |

## 발견 상세

### PM-01. 수정하려는 키가 오히려 텍스트를 손상시킨다

- 재현 A: `vim-ex-command-substitute-002`에서 `:%s/TODO/DONX` 입력 후 실제 DEL byte(`0x7f`, Backspace)를 보낸다.
- 실제: `:%s/TODO/DONXbackspace`가 표시된다. 별도 재현에서 `:` 다음 실제 Left escape sequence를 보내면 `:left`가 된다.
- 재현 B: `insert-mode-entry-001`에서 `i`, Backspace. 버퍼가 `ac`에서 `abac`로 변하고 입력 제한 초과로 실패한다. `backspace` 문자열의 앞 문자들이 입력 예산을 소모한다.
- 근거: `internal/tuiadapter/adapter.go:72`, `:125`, `:244`의 문자열 분해. `tea.KeyMsg.String()`의 특수키 이름도 일반 입력 문자열로 처리한다.
- evidence: `pm_command_backspace/screen_timeline.txt`, `pm_command_left/screen_timeline.txt`, `pm_insert_backspace/app_state.json`.
- 영향: 배우는 사람이 자연스럽게 오타를 고치려다 앱이 버퍼를 바꾸고 실패시킨다. Vim 전체 지원 여부와 무관하게 키 이름이 데이터로 들어가는 것은 결함이다.
- 방향: 특수키는 명시적 입력 타입으로 분기하고, Backspace의 지원된 수정 동작을 구현한다. 미지원 키는 텍스트를 바꾸지 않고 이유를 알려야 한다.

### PM-02. 모달의 약속이 실패 당시 모드에 따라 달라진다

- 재현: `insert-mode-entry-001`에서 `i`, `z`, `z`, `z`로 입력 제한 실패를 만든 뒤 Enter. 화면은 `다시 시도: r 또는 enter`지만 status는 failed, mode는 insert로 남는다.
- 같은 실패에서 `?`도 힌트를 표시하지 않는다. 비교 실험에서 `r`은 초기 `ac`/normal/running으로 정상 재시작한다.
- Visual 재현: `command-choice-scope-001`에서 `V`, `x`로 실패한 뒤 Enter. visual/failed로 남는다.
- 근거: `internal/playable/model.go:129`에서 UI 상태보다 mode 기반 입력 매핑이 먼저 실행된다. `internal/tuiadapter/adapter.go:80`, `:90`의 Enter/기타 입력 처리가 실패 화면에도 적용된다.
- evidence: `pm_insert_failure_enter`, `pm_insert_failure_hint`, `pm_insert_failure_retry_r`, `pm_visual_failure_enter`.
- 방향: 실패 모달은 독립된 입력 문맥이어야 한다. 표시한 retry/hint/quit은 모든 실패 모드에서 같은 의미로 동작해야 한다.

### PM-03. 같은 문자 입력이 전달 단위에 따라 달라진다

- 재현: 첫 문제에서 `ll`을 하나의 PTY write로 보낸다. running, cursor 0:0에서 변화가 없다. `l`을 각각 보내면 정상적으로 cursor 0:2에 도착한다.
- 근거: `internal/tuiadapter/adapter.go:231` 이후 Normal mode는 단일 rune 또는 `:`/`/`로 시작한 문자열만 받아 나머지는 ignored 처리한다.
- evidence: `pm_normal_batch_80`, 비교 `pm_success_review_80`, 직접 플레이 `manual-first.raw`.
- 한계: 모든 실제 키보드의 빠른 타이핑이 반드시 하나의 이벤트로 합쳐진다고 주장하지 않는다. 묶여 전달된 입력/붙여넣기에 대한 손실은 실제 재현했다.
- 방향: 일반 입력도 순서대로 처리하되 각 키 이후 모드와 성공/실패 상태를 다시 확인한다. 남은 붙여넣기 입력이 다음 문제에 흘러 들어가면 안 된다.

### PM-04. 응용 판단 문제가 사실상 지정 키 재현을 요구한다

- 재현: `command-choice-line-reuse-001`에서 개별 키 `y`, `y`, `j`, `p`.
- 실제: 요구 버퍼 `route=stable / backup= / route=stable / guard=on`에 도착했지만 `required_keys`의 `V`가 없어 failed다. `V`, `y`, `j`, `p`는 같은 4키로 성공한다.
- 근거: `content/exercises/command-choice.yaml:287`, `internal/runtime/session.go:254`. 화면의 참고 명령 목록에는 `yy`도 있다.
- evidence: `pm_line_reuse_alternative`와 `pm_line_reuse_intended`의 buffer/status 비교.
- 분류: 현 명세대로 실행된 동작이므로 런타임 버그로 취급하지 않는다. "줄 전체 재사용 판단"이라는 약속과 채점 규칙이 충돌하는 제품 설계 문제다.
- 방향: 튜토리얼은 훈련할 동작을 사전 고지하고 제한할 수 있다. 응용은 결과/보존 범위/재사용 의도를 만족하는 다른 풀이를 인정한다. 정상적인 대안 풀이를 RedTeam 우회와 구별해야 한다.

### PM-05. 차별점에 도달하기 전 학습 경로가 길다

- 실제 구성: approved tutorial 16개, 총 79 beat가 incident 001보다 먼저 정렬된다. approved incident는 11개, 42 beat다. 현재 playable 전체는 121 beat다.
- 근거: `internal/content/loader.go:321`의 category/order 정렬, `internal/playable/model.go:668`의 first-incomplete 진입. `pm_incident_hints_80/progress.json`의 incident 진입 fixture도 선행 tutorial 79개의 완료를 필요로 한다.
- UI에서 직접 에피소드/응용/복습을 고르는 경로는 확인되지 않았다. E2E의 `complete_before`는 테스트 장치이며 사용자의 진입 기능이 아니다.
- 영향 가설: "복구 게임"을 기대하고 들어온 사람에게 단일 명령 학습과 확인 모달이 오래 지속된다. 신규 콘텐츠를 끝에 붙여도 첫 경험은 개선되지 않는다.
- 방향: 필요한 소수 명령을 배운 직후 첫 복구 작전을 제공한다. 기본 학습 경로와 경험자의 진단/응용 진입을 구별하되 선행 지식이 없으면 해당 개념을 짧게 연결한다.

### PM-06. 가장 첫 문제의 성공 목표가 화면에서 빠졌다

- 재현: fresh start. 화면은 `> [a]bc`와 "목표 문자까지 커서를 이동하세요"를 보여주지만 목표가 `c`라는 말이나 목표 표시는 없다.
- 실제 exercise 목표는 `커서를 c 위로 이동하세요.`다. scenario briefing이 화면에 쓰이면서 구체적 목표가 사라진다.
- 처음 `?`를 눌러도 "오른쪽으로 한 칸 더"라고 나온다. 초기 a에서 c까지는 두 칸이다.
- 근거: `content/exercises/first-five-minutes.yaml:130`, `:153`; `content/scenarios/first-five-minutes.yaml:54`; `internal/playable/model.go:588`의 scenario briefing 전달.
- evidence: `pm_first_screen_80/screen_final.txt`, `manual-first.raw`.
- 방향: 첫 화면만으로 성공 목표를 알아야 한다. 목표 `c`를 문장과 색에 의존하지 않는 표식으로 표시하고, 초기 상태에서도 맞는 힌트를 제공한다.

### PM-07. 좁은 화면에서 보조 행동이 사라진다

- 재현: 80x24 첫 문제에서 Right로 실패 후 `?`.
- 실제: final frame에는 retry와 hint가 있지만 `종료: q`가 없다. app_state의 quit action은 남아 있고 q 자체도 작동한다. 110x34에서는 종료 action이 보인다. 80x24에서 힌트 없이 실패한 비교 사례에서는 종료 action이 보인다.
- 근거: `internal/playableview/render.go:756`의 고정 13행 제한, `:973`의 앞부분 절단과 단일 priority line 보존.
- evidence: `pm_failure_hint_80`, `pm_failure_hint_110`, `pm_failure_no_hint_80`.
- 방향: primary와 보조 action footer 공간을 먼저 예약한다. 설명/힌트 본문을 줄이거나 접더라도 존재하는 행동 안내를 잘라내지 않는다.

### PM-08. 복습 큐가 비어도 복습 흐름이 끝나지 않는다

- fixture: 121문항 중 첫 이동 문항만 완료/A, 마지막 치환 문항만 미완료, 나머지는 완료/S.
- 재현: 마지막 문항을 `:%s/hold/live/g`로 완료 → Enter로 첫 이동 복습 진입 → `l`, `l`로 S 달성 → Enter.
- 실제: review queue는 0인데 이미 완료/S인 `normal-motion-basic-002`가 running으로 열린다. 완료된 전체 캠페인의 순차 진행을 다시 타기 시작한다.
- 근거: `internal/playable/model.go:340`은 복습에서 진입했는지 구별하지 않고 다음 entry로 이동한다.
- evidence: `pm_review_exit_route/app_state.json`, `progress.json`, `screen_timeline.txt`.
- 방향: 복습 세션의 대상과 종료를 명시한다. 큐가 비면 복습 완료, 남으면 다음 복습 대상으로 이동한다. 전체 캠페인 재플레이는 별도 선택이다.

### PM-09. 힌트 비용은 누적되지만 정보는 늘지 않는다

- incident 001 첫 문제에서 첫 힌트는 공통 "바뀌어야 할 범위와 보존할 범위" 문장이다. 이 문제는 텍스트 변경 없이 error로 이동하는 검색 과제다.
- 세 번째 이후는 `CurrentHint()`가 키 입력 수로 고른 문장을 그대로 보여준다. `?`만 반복하면 마지막 상세 힌트까지 진행하지 않는다.
- 재현: 초기 상태에서 `?` 6회 후 `/error`, Enter. 최적 7키로 복구에 성공해 저장되지만 grade는 F다. 세 번째 이후 같은 정보를 다시 읽은 행동도 계속 감점된다.
- 근거: `internal/playable/model.go:802`, `internal/runtime/session.go:207`, `internal/scenario/run.go:100`, `internal/scoring/scoring.go:97`.
- evidence: `pm_incident_hints_80/screen_timeline.txt`, `pm_hint_repeated_success/screen_final.txt`, `app_state.json`.
- 방향: 과제별 첫 힌트, 유한한 단계 표시, 동일 단계 무료 재열람. 개념 힌트는 기본적으로 탐색을 돕고, 정답 공개는 도움받은 완료로 표시한다. 성공 여부와 독립 수행 여부를 혼동시키지 않는다.

### PM-10. 작은 성취보다 보고서를 읽는 비중이 높다

- 직접 플레이한 첫 4문항의 해법은 `ll`, `j`, `hh`, `k`다. 각 성공마다 전체 modal이 나와 배운 점, 현재/최고/목표 기록, Runbook 수, 재점검, 다음 행동을 표시한다.
- 80x24에서는 modal이 편집 결과를 가린다. GOAL이 이중 wrapping되면서 어중간한 줄과 반복 label도 보인다. SIGNAL의 순환 표시는 있으나 관찰한 초반에서는 복구 전후 상태 변화를 충분히 설명하지 않는다.
- evidence: `manual-first.raw`, `manual-resume.raw`, `pm_success_review_80`, `pm_success_review_110`, `pm_baseline_ftue_first_five_route`.
- 분류: 긴 화면과 입력 중단은 관찰 사실이다. 이것이 지루함이나 이탈로 이어지는 정도는 사용자 검증이 필요하다.
- 방향: 단일 조작은 결과를 유지한 짧은 피드백으로, 에피소드 완료는 의미 있는 보고서로 나눈다. 애니메이션은 통신 끊김→연결, 오염→격리 같은 결과를 보여주고 사용자의 입력을 지연시키지 않는다.

### PM-11. 새 학습을 복습 부채처럼 표현한다

- 첫 문제를 완벽하게 끝내도 아직 시도하지 않은 다음 문제를 "재점검 메모", "나중에 다시 풀기"로 보여준다. incident에서는 미학습 문항이 "잔류 리스크"가 된다.
- 근거: `internal/review/review.go:126`은 미완료를 최우선 review candidate로 잡는다. `internal/playable/model.go:408`은 최대 3개만 가져오므로 화면의 3건은 전체 잔여 문항 수가 아니다.
- evidence: `pm_success_review_80/app_state.json`, `pm_success_review_110/screen_final.txt`.
- 방향: "다음 새 미션"과 "다시 연습할 미션"을 구별하고, 추천 3개와 총량을 혼동하지 않게 한다. 오늘의 루트가 실제 날짜 기반 복습 간격을 구현한 것처럼 보이지 않게 한다.

### PM-12. 설치 전 기대를 만드는 문서가 현재 제품과 다르다

- README는 "첫 공개 전 foundation build", 106 exercise/18 playlist 파일이라고 설명한다. 현재 VERSION은 0.4.0이고 확인한 approved playable 구성은 121 beat/27 playlist다. 파일 수와 playlist 객체 수는 별도 지표다.
- 방향: 실제 stable 설치와 플레이 가능한 범위를 일치시키고, 직접 확인한 수치 또는 자동 산출 가능한 정보만 유지한다. 이 감사에서는 README를 수정하지 않았다.

## 후속 제품 기획 [draft]

### 목표와 성공 경험

[draft] 주 대상은 터미널 사용이 가능하고 Vim 기본 조작을 배우거나 막 배운 개발자다. 짧은 복구 과제에서 "무엇을 바꿀지 판단 → 직접 편집 → 복구 결과 확인 → 다른 상황에 적용"을 경험하게 한다. 완전 초보자와 숙련자는 진입 경로를 구별한다.

[draft] 첫 대표 세션은 `작전 목표 → 필요한 조작 맛보기 → 짧은 복구 → 변화 확인 → 다음 작전/복습/종료 선택`으로 구성한다. 시간 목표는 아래 사용자 실험에서 검증하며 자동 E2E 실행 시간을 사람이 플레이하는 시간으로 사용하지 않는다.

### 1. TRUST-RECOVERY-001 — 입력과 실패 복구

관련: PM-01/02/03/06/07. 이 단계가 첫 구현 추천이다.

- [ ] [draft] 특수키와 문자 입력 분리: Backspace가 command/search/insert에서 명시된 수정 동작을 수행하며 키 이름을 삽입하지 않는다. 미지원 특수키도 텍스트를 오염시키지 않는다.
- [ ] [draft] 실패 모달 입력 우선: normal/insert/visual/search/command 실패 상태마다 `Enter`와 `r`이 같은 초기 상태로 재시작하고, `?`는 힌트, `q`는 종료로 동작한다.
- [ ] [draft] `ll`, `gg`, operator 조합의 묶음/개별 전달이 같은 결과를 낸다. 입력 중 mode가 바뀌는 경우와 성공/실패 뒤 남은 입력을 별도로 검증한다.
- [ ] [draft] 첫 목표 `c`와 현재 위치를 색 없이 구별할 수 있고 초기 힌트가 현재 거리와 모순되지 않는다.
- [ ] [draft] 80x24/110x34에서 실패+힌트, insert 실패, visual 실패의 모든 표시 action이 보이고 실제 키 입력과 일치한다.
- 검증 gate: 승인된 기대 동작으로 Red test 작성 → 최소 수정 → focused PTY → `go test ./...` → 관련 전체 playable E2E. 이번 관찰용 "버그 재현 pass"를 회귀 통과 기준으로 복사하지 않는다.

### 2. FIRST-RECOVERY-001 — 짧고 의미 있는 첫 플레이

관련: PM-05/10. 채점 개편의 수용 기준을 함께 확정한 뒤 콘텐츠를 설계한다.

- [ ] [draft] 필수 조작 맛보기 3~5개 뒤 첫 복구 에피소드에 도달하는 대표 경로를 제공한다. 기존 79문항은 학습 라이브러리로 계속 접근 가능하게 한다.
- [ ] [draft] 첫 화면/에피소드 경계에서 이어하기, 학습 경로, 응용 도전의 목적이 구별된다. 자유 진입 시 부족한 선행 개념으로 돌아가는 경로를 제공한다.
- [ ] [draft] 단일 조작 성공은 결과 버퍼와 한 문장 피드백을 함께 유지한다. 상세 기록은 에피소드 보고서에 모으고 다음 행동은 항상 명확히 둔다.
- [ ] [draft] 대표 복구 1개는 동일한 장비/문제의 전후 변화를 보여준다. 애니메이션 없이도 변화의 의미가 읽히고, 애니메이션 중 입력을 잃지 않는다.
- 완료 기준: 첫 복구까지의 경로를 PTY로 재현하고, 대상 사용자 실험에서 첫 복구 중앙값 5분 이내를 목표로 평가한다. 시간 수치는 제품 가설이며 아직 달성하지 않았다.
- 경계: 기존 progress ID와 완료 기록을 보존한다. 순서/재개 정책 변경 영향을 별도 테스트하고 저장 schema가 필요하면 별도 승인한다.

### 3. JUDGMENT-FAIRNESS-001 — 응용 문제의 공정한 성공 판정

관련: PM-04. FIRST-RECOVERY의 응용 문항 출시 gate다.

- [ ] [draft] tutorial의 지정 도구 훈련과 incident의 결과/판단 목표를 문항 및 UI에서 구별한다. 도구 강제는 시작 전에 알려야 한다.
- [ ] [draft] 줄 전체 재사용 목표에서 `yyjp`와 `Vyjp`가 둘 다 성공한다. 복사 범위, 보존 줄, 최종 위치가 의미상 필요한지 각각 정의한다.
- [ ] [draft] 명시적으로 재사용을 요구하는 과제는 실제 yank/put의 의미를 검증한다. 관계없는 필수 키를 한 번 누르는 것으로 의도를 충족했다고 보지 않는다.
- [ ] [draft] 정답 trace와 다르다는 이유만으로 최고 평가를 막지 않는다. 효율은 최적화 피드백으로 사용하고 성공 조건과 구별한다.
- [ ] [draft] delimiter 손상, 정상 영역 변경, 목표 일부만 수정한 풀이와 정상 대안 풀이를 양쪽 모두 테스트한다.
- 완료 기준: 대표 3개 응용 문항마다 복수 유효 풀이와 의미상 잘못된 풀이가 기대한 판정을 받고 이유를 설명한다. 내부 판정 확장으로 충분한지 먼저 검토하며 content schema 변경을 전제하지 않는다.

### 4. COACH-RETURN-001 — 도움받고 다시 연습하는 흐름

관련: PM-08/09/11.

- [ ] [draft] 힌트는 과제별 관찰 → 도구 계열 → 구체적 예시의 유한한 단계이며 현재 단계/최종 단계가 보인다.
- [ ] [draft] 동일 힌트를 재열람해도 비용이 계속 늘지 않는다. 개념 힌트는 무료를 기본안으로 하고 구체적 정답 공개는 "도움받아 완료"로 구별한다.
- [ ] [draft] 성공 완료와 미완료 실패를 혼동하는 F 보고서를 없앤다. 재시도/도움 사용이 평가에 반영되는 방식은 보고서에서 설명한다.
- [ ] [draft] 미학습 다음 과제와 실제 복습 대상을 분리한다. 복습 1개를 끝내 큐가 비면 종료/새 작전 선택을 제공한다.
- [ ] [draft] 복습 큐가 남으면 다음 복습 대상으로 가고 이미 숙련한 순차 튜토리얼로 자동 이탈하지 않는다.
- 완료 기준: 큐 0/1/복수, 완료/S와 완료/도움 사용, 재실행을 fixture로 재현한다. 간격 반복 알고리즘이나 streak는 이번 범위에 필수로 넣지 않는다.

### 5. PLAYER-VALIDATION-001 — 개발량 대신 사용자 가치로 판정

- [ ] [draft] 목표 사용자 6~8명을 모집해 동일 안내로 첫 15~20분을 관찰한다. 경험 수준을 기록하고 진행자가 정답을 알려준 경우를 구별한다.
- [ ] [draft] 플레이 전/후 서로 다른 유사 편집 과제를 실제 Vim에서 풀게 한다. 두 동등 과제 세트의 순서를 나눠 기억/난이도 효과를 줄이고 성공, 도움 요청, 정상 영역 손상을 기록한다.
- [ ] [draft] 학습자는 목표를 자기 말로 설명하고, 마지막에 왜 해당 도구를 골랐는지 말한다. 명령 이름 회상만으로 판단력 향상을 인정하지 않는다.
- [ ] [draft] 48~72시간 뒤 별도 과제로 유지 여부를 확인한다. 요청받아 수행한 후속 테스트와 자발적으로 게임에 돌아온 행동을 따로 기록한다.
- [ ] [draft] 외부 사용자에게 제공하기 전에 실제 stable 설치/실행 안내와 첫 경로를 검증한다. README의 수치와 출시 상태를 동기화한다.

8명 기준의 초기 의사결정 문턱 제안:

| 관찰 | 확장 검토 신호 | 미달 시 우선 조사 |
|---|---|---|
| 설명 없이 첫 복구 완료 | 6명 이상, 중앙값 5분 이내 | 목표 가시성/조작/선행 지식 |
| 실제 Vim의 새로운 유사 과제 | 5명 이상 도움 없이 성공하며 개인별 사전 결과보다 개선 | 정답 암기/판정 제약/문항 전이 |
| 추가 작전 자발적 선택 | 4명 이상이 진행자의 요청 없이 선택 | 피드백/목적/과제 난도 |

위 숫자는 소표본의 제품 의사결정을 위한 가설이다. 통계적 학습 효과나 시장 적합성의 증명이 아니다. 후속 과제는 이미 쉬웠던 사람의 천장 효과를 따로 기록한다. 학습과 자발적 지속이 함께 보이면 확장을 검토하고, 학습만 보이면 훈련 도구 방향을 강화하며, 둘 다 약하면 대표 에피소드를 재설계한다.

## 재현 evidence 안내

모든 아래 경로는 `artifacts/e2e/pm-audit/` 기준이다. raw/progress에는 테스트 fixture만 들어 있으며 실제 사용자 저장 파일을 읽거나 변경하지 않았다. 대용량 raw 및 바이너리는 git에 포함하지 않는다.

- 직접 입력 기록: `manual-first.raw`, `manual-resume.raw`, `manual-incident.raw`.
- 시나리오: `scenarios/pm_*.yaml`. 각 결과 디렉토리에는 `screen_final.txt`, `screen_timeline.txt`, `app_state.json`, `key_trace.txt`, `summary.json` 등이 있다.
- 기본 경로 5개: `pm_baseline_ftue_first_five_route`, `pm_baseline_incident_001_full`, `pm_baseline_command_choice_scope`, `pm_baseline_incident_011_full`, `pm_baseline_resume_first_incomplete`.
- 관찰/재현 20개: `pm_first_screen_80`, `pm_normal_batch_80`, `pm_failure_no_hint_80`, `pm_failure_hint_80`, `pm_failure_hint_110`, `pm_incident_hints_80`, `pm_line_reuse_alternative`, `pm_line_reuse_intended`, `pm_command_backspace`, `pm_command_left`, `pm_insert_backspace`, `pm_substitute_control`, `pm_success_review_80`, `pm_success_review_110`, `pm_insert_failure_enter`, `pm_insert_failure_hint`, `pm_insert_failure_retry_r`, `pm_visual_failure_enter`, `pm_hint_repeated_success`, `pm_review_exit_route`.

로컬 재실행 예:

```sh
artifacts/e2e/pm-audit/bin/e2e-runner --scenario artifacts/e2e/pm-audit/scenarios/pm_command_backspace.yaml --artifacts artifacts/e2e/pm-audit
```

공유용 최소 재현 fixture는 다음과 같다. 기존 runner로 독립 실행할 수 있고, `\u007f`는 문자 "backspace"가 아닌 실제 DEL byte다. 관찰 assertion은 버그가 존재하는 기준 commit을 대상으로 한다.

```yaml
id: pm_command_backspace
command: ["go", "run", "."]
timeout_ms: 30000
terminal: {width: 80, height: 24}
setup: {home: temp, complete_before: vim-ex-command-substitute-002}
steps:
  - wait_screen_contains: "RUNBOOK CONSOLE"
  - send: ":%s/TODO/DONX"
  - wait_ms: 100
  - send: "\u007f"
  - wait_screen_contains: ":%s/TODO/DONXbackspace"
  - send: "esc"
  - wait_ms: 100
  - send: "q"
assert:
  exit_code: 0
  screen_contains: [":%s/TODO/DONXbackspace"]
evidence:
  save_raw_ansi: true
  save_screen_timeline: true
  save_screen_final: true
  save_app_state: true
  save_key_trace: true
  save_summary: true
```

## 작업 기록

- 초기 관찰 fixture 일부에서 아직 생성되지 않은 progress의 저장을 요청해 runner가 evidence 오류를 냈다. 해당 fixture의 `save_progress`를 끄고 재실행했다. 제품 결함으로 집계하지 않았다.
- JSON 형식 YAML에 raw DEL이 들어간 초기 fixture 2개는 YAML parsing이 거부했다. `\u007f` escape로 수정한 뒤 실제 DEL byte 전송을 검증했다.
- 앱/runner 빌드는 성공했다. sandbox의 module stat-cache 쓰기 경고는 있었지만 바이너리 실행과 25개 PTY 검증은 완료했다.
- 제품 코드, content, dependency, progress schema, canonical roadmap는 변경하지 않았다. 후속 개발 승인은 이 문서의 `[draft]` 기준을 대상으로 별도로 받는다.
