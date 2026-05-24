# Gameplay Spec

## 개요

Advimture의 게임플레이, Vim 학습 문항, 내러티브, 미션 구조를 정의한다. Go 기반 TUI adventure game이라는 방향은 유지하되, 기존 구현은 아직 canonical spec으로 승격하지 않는다.

---

## 현재 동작

### 제품 설계 순서

**소스**: `docs/roadmap/PRODUCT.md`, `docs/workflows/vim-learning-loops.md`, `docs/gameplay/domain-contract.md`

- 주된 목적은 Vim 학습이다.
- 플레이어는 Vim을 실제로 유용하게 쓰는 데 필요한 단축어와 명령어를 반복 훈련한다.
- 설계 순서는 `Vim command → Exercise → Scenario`다.
- `Vim command` 단계는 학습할 명령어 묶음, 실무 유용성, 선행 관계, 조합 가능성을 정의한다.
- `Exercise` 단계는 초기 텍스트, 목표 상태, 정답 키 입력, 허용 키, 힌트, 채점 기준을 정의한다.
- `Scenario` 단계는 검증된 exercise를 어드벤처 사건, 캐릭터 대사, 성공/실패 피드백으로 감싼다.
- Command cluster는 `docs/gameplay/command-catalog.md`에 축적한다.
- Command cluster는 `exact`, `pedagogical`, `unsupported` 중 하나의 Vim compatibility tier를 가진다.
- Exercise는 `docs/gameplay/exercise-bank.md`에 축적한다.
- Scenario는 `docs/gameplay/scenario-bank.md`에 축적한다.
- 게임이 읽는 콘텐츠 파일은 repo root `content/` 아래 YAML을 우선한다.
- 초반 콘텐츠는 8문항 이하의 짧은 tutorial playlist 묶음으로 나누는 것을 기본으로 한다.
- 초반 튜토리얼은 “첫 투어”처럼 Vim command를 넓게 맛보게 하고, 중후반부터 생존 어드벤처와 위기 해결 비중을 높인다.
- Ex substitute 계열(`:s`, `:%s`, range substitute)은 초반 튜토리얼에서 분리해 중반 고급 튜토리얼로 다룬다.
- 시나리오 톤은 DevOps/터미널 문제 해결을 기본으로 하되, 과하지 않은 억까 상황은 허용한다.
- 중반 이후 시나리오 톤은 `docs/gameplay/scenario-tone.md`의 “터미널 문제 해결 생존 어드벤처” 기준을 따른다.
- 생존감은 전투/체력/인벤토리가 아니라 로그, 설정, 통신, 임시 패치, 저장/폐기 판단 같은 텍스트 조작 압력에서 온다.
- 시나리오에서 출발해 문항을 끼워 맞추지 않는다.
- 스토리나 세계관은 문항을 더 기억에 남게 해야 하며, command 학습 목표를 방해하면 안 된다.
- Vim runtime은 작은 Go Vim-like engine으로 구현한다. Neovim은 런타임이 아니라 optional oracle test로만 사용한다.
- 엔진은 `State + Key -> State + Events` 순수 전이를 기본 계약으로 삼는다.
- 진행 사항은 자동 저장을 기본으로 하며, 기존 `.advimture` 계열 progress 저장 경계를 유지한다.
- `normal-motion-basic`은 `h`, `j`, `k`, `l` 각각이 optimal trace에 등장하는 approved + implemented exercise set을 가진다.
- `word-motion-basic`은 `w`, `b`, `e` 각각이 optimal trace에 등장하는 approved + implemented exercise set을 가진다.
- word motion exercise set은 replay gate와 E2E assertion gate를 통과해야 한다.
- `survival-save-quit`은 `esc`, `:q!`, `:wq` 각각이 `trained_commands`와 command goal coverage에 등장하는 approved + implemented exercise set을 가진다.
- `:q!`, `:wq`는 실제 파일 작업이 아니라 command-line key trace replay와 runtime command goal로 검증한다.
- 앱 종료 단축키 `q`와 command-line mode의 `q` 입력은 분리한다.
- `ctrl+c`는 앱 종료 단축키가 아니며 runtime key trace로 전달한다. 종료는 playable 일반 모드의 `q`를 기본으로 한다.
- `whole-file-navigation`은 `gg`, `G`, `0`, `$` 각각이 approved + implemented exercise coverage와 replay gate를 통과한다.
- `gg/G`는 NAV-001에서 파일 처음/끝 줄의 첫 column으로 이동하는 pedagogical motion으로 다룬다.
- `vim-ex-command-substitute`는 `:s`, `:%s`, `:2,3s` 각각이 approved + implemented exercise coverage와 replay gate를 통과한다.
- substitute command는 EXCMD-001에서 literal match만 지원하며, scenario success는 buffer target으로 검증한다.
- playable은 approved/implemented playlist를 `category`, `order`, `id` 순서로 실행한다. `tutorial` category는 `incident` category보다 먼저 실행한다.
- 현재 playable tutorial 순서는 `tutorial-0-movement`, `tutorial-1-survival`, `tutorial-2-fast-navigation`, `tutorial-3-small-edits`, `tutorial-4-ex-command`, `tutorial-5-operator-grammar`, `tutorial-6-yank-put`, `tutorial-7-text-object-inner-word`, `tutorial-8-open-line-edit`, `tutorial-9-repeat-last-change`, `tutorial-90-search-basic`, `tutorial-91-text-object-quote-pair`, `tutorial-92-visual-selection`, `tutorial-93-visual-line`다.
- `first-5-minute`는 legacy vertical slice로 retired 상태이며 default playable path에서 실행하지 않는다.
- 첫 5분 canonical route는 `docs/gameplay/first-five-minute-route.md`를 따른다. 현재 기준은 `tutorial-0-movement` 전체, `tutorial-1-survival` 전체, `tutorial-2-fast-navigation` 전체, `tutorial-3-small-edits` 첫 문항까지이며, 새 command/schema/progress 변경 없이 현재 playable route를 제품 첫 루프로 고정한다.
- 화면은 현재 tutorial title과 episode-local exercise count를 표시한다.
- terminal size가 있는 playable 화면은 `MISSION` HUD, `RUNBOOK CONSOLE`, status line 순서로 렌더링한다.
- 진행/재시도/명령 입력 안내는 일반 하단 텍스트가 아니라 structured `FocusPanel` 모델로 관리한다.
- `FocusPanel`은 `kind`, `title`, `lines`를 가진다. tutorial running은 `training`/`TRAINING BRIEF`, incident running은 `incident`/`OPERATOR JUDGMENT`, failed는 `failure`/`RECOVERY REQUIRED`, succeeded는 `success`/`STEP SEALED`, mode-specific 안내는 `mode` kind를 사용한다.
- running/mode-specific `FocusPanel`은 `MISSION` HUD 안의 짧은 cue line으로 접어 현재 목표와 함께 보인다.
- failed/succeeded `FocusPanel`은 `RUNBOOK CONSOLE` 안에서 Zellij floating pane처럼 보이는 modal로 렌더링한다.
- floating modal은 `tea.WindowSizeMsg`로 전달된 terminal width가 있으면 화면 중앙에 배치하고, 좁은 화면에서는 terminal width를 넘지 않도록 폭을 줄인다.
- floating modal은 실패 시 `RECOVERY CHECK`, 성공 시 `RUNBOOK SEALED` 구조로 표시하며 action line(`Retry`, `Next`, `Next tutorial`)이 잘리지 않아야 한다.
- `복구 현황`은 terminal size가 있는 화면에서 별도 큰 pre-console section이 아니라 `MISSION` HUD 내부의 보조 line으로 표시한다.
- running/failed 상태의 `FocusPanel`은 아직 쓰지 않은 `constraints.required_keys`를 tutorial에서는 `Coach: 훈련 키 ...`, incident failure에서는 `복구 힌트: 필요한 키 ...`로 표시한다.
- `?` hint 요청 결과는 `FocusPanel`에 `Hint: ...`로 표시하며, command/search/insert mode 패널에는 실제 입력 처리와 맞지 않는 일반 hint/quit 안내를 섞지 않는다.
- failed/succeeded 상태의 scenario feedback은 briefing 영역이 아니라 `FocusPanel` 안에 표시하며, briefing 영역은 원래 미션 설명을 유지한다.
- 한 tutorial 마지막 exercise 성공 시 다음 tutorial이 있으면 `Next tutorial: enter`를 표시하고, `enter`로 다음 tutorial에 진입한다.
- exercise 성공 시 기존 progress `Missions` map에 exercise ID를 key로 자동 저장하고, 성공 상태에서 `enter`를 누르면 다음 unlocked exercise로 이동한다.
- 성공 FocusPanel은 현재 복구 기록, 기존 progress 기반 최단 복구 기록, 현재 Runbook 복구 완료 수를 표시한다.
- playlist 마지막 exercise 성공 화면도 별도 저장 포맷 변경 없이 같은 debrief와 `Playlist complete` 안내를 표시한다.
- 향후 exercise constraint는 최대 입력 수 초과와 금지 입력/금지 우회 전략 사용을 즉시 실패로 처리해야 한다.
- 실패 횟수는 기본 무제한이며, 후반 콘텐츠를 위해 `attempt_limit` 설정 여지는 남긴다.
- 실패 후 재시도는 `r`과 `enter`를 모두 허용한다.
- 초반 튜토리얼 코칭은 개념 힌트 중심이지만, 새 command 첫 소개에서는 command 의미를 명시한다.
- exercise YAML은 `constraints.max_inputs`, `constraints.required_keys`, `constraints.forbidden_keys`, `constraints.attempt_limit`를 선언할 수 있다.
- `constraints.max_inputs` 초과와 `constraints.forbidden_keys` 입력은 Vim state를 추가 진행시키지 않고 즉시 실패한다.
- `left/right/up/down` 화살표 입력은 `h/j/k/l`로 변환하지 않고 원래 key name으로 runtime에 전달해 `forbidden_keys`가 검출할 수 있어야 한다.
- `ctrl+c` 입력도 quit으로 가로채지 않고 원래 key name으로 runtime에 전달해 content constraint나 unsupported key handling이 처리하게 한다.
- command goal 문항에서 목표와 다른 command를 실행하면 즉시 실패한다. `:q!` 목표에서 `:wq`, `:wq` 목표에서 `:q!`를 실행하면 저장/폐기 의도 차이를 설명하는 회복 피드백을 보여준다.
- 목표에 도착했더라도 `constraints.required_keys`가 key trace에 없으면 성공하지 않는다. 단, `constraints.max_inputs`가 남아 있으면 즉시 실패하지 않고 계속 진행할 수 있다. 이후 목표 상태를 벗어났다가 required key로 다시 목표에 도달하면 성공할 수 있지만, 목표 상태에 머문 채 required key를 나중에 덧붙이는 우회는 `required_keys_missing`으로 실패한다. 입력 제한이 끝났거나 입력 제한이 없으면 `required_keys_missing`으로 실패한다.
- 실패 상태는 progress를 저장하지 않으며, 실패 화면은 `Grade: F`, 남은 입력 수, 재시도 안내를 보여준다.
- 실패 화면은 attempt count를 표시하며 `attempt_limit: 0`은 `unlimited`로 표현한다.
- scoring result는 runtime failure reason을 보존하며, `required_keys_missing`은 `IntentSatisfied=false`, `Grade=F`로 평가한다.
- `single-char-edit`, `insert-mode-entry`, `undo-redo-basic`은 approved + implemented tutorial cluster이며 `x`, `r`, `i`, `a`, `A`, `u`, `ctrl+r` coverage와 replay gate를 통과한다.
- `tutorial-3-small-edits`는 7문항짜리 작은 수정 튜토리얼이며 Ex command보다 먼저 실행된다.
- `undo-redo-basic` 문항은 redo처럼 중간 상태가 목표와 같아질 수 있는 경우에도 required key를 사용할 때까지 성공하지 않아야 하며, `undo-redo-basic-002`는 `x`, `u`, `ctrl+r` 3입력으로 clear된다.
- `tutorial-5-operator-grammar`는 `dw`, `d$`, `dd`, `cw`, `c$`, `cc`를 각각 한 문항씩 다루는 6문항 operator grammar 입문 tutorial이다.
- `delete-with-motion`, `change-with-motion`은 approved + implemented tutorial cluster이며 replay gate와 E2E assertion gate를 통과한다.
- 현재 pedagogical `cw`는 `dw`와 같은 범위로 단어 뒤 공백까지 삭제하므로, 단어 교체 문항은 새 단어 뒤 공백 입력을 optimal trace에 포함한다.
- `tutorial-6-yank-put`은 `yy+p`, `yy+P`, `yw+P`, `y$+p`, `yw+p`를 각각 다루는 5문항 yank/put 재사용 tutorial이다.
- `yank-put-basic`은 approved + implemented tutorial cluster이며 `yw`, `y$`, `yy`, `p`, `P` coverage와 replay gate, E2E assertion gate를 통과한다.
- 첫 yank/put 구현은 unnamed register만 다루며 named register, clipboard, count prefix, visual selection, text object는 후속 루프로 분리한다.
- `tutorial-7-text-object-inner-word`는 `diw`, `ciw`, `yiw`를 각각 두 문항씩 다루는 6문항 text object tutorial이다.
- `text-object-inner-word`는 approved + implemented tutorial cluster이며 `diw`, `ciw`, `yiw` coverage와 replay gate, E2E assertion gate를 통과한다.
- 첫 text object 구현은 `iw`만 다루며 quote/pair object, around object, visual selection, count prefix는 후속 루프로 분리한다.
- `text-object-quote-pair`는 approved + engine implemented command cluster다. 첫 scope는 double quote 내부 object `di"`, `ci"`, `yi"`이며, nested pair, escaped quote, single quote, parenthesis, brace, around object, count prefix, visual selection은 후속 hardening으로 분리한다.
- `tutorial-91-text-object-quote-pair`는 `ci"`, `di"`, `yi"`, `ci"` + `.` 반복을 다루는 4문항 quote text object tutorial이며 replay gate와 full playlist E2E를 통과한다.
- `open-line-edit`은 approved + engine implemented command cluster이며 `o`, `O`는 현재 줄 아래/위에 빈 줄을 삽입하고 Insert mode로 진입한다.
- `tutorial-8-open-line-edit`은 `o` 3문항, `O` 2문항으로 구성된 5문항 tutorial이며 replay gate와 full playlist E2E를 통과한다.
- 첫 `o/O` 구현은 indentation, auto-comment, count prefix, insert-mode Enter, dot repeat 연계를 제외한다.
- `repeat-last-change`는 approved + engine implemented command cluster다. 첫 구현의 last-change transaction은 x, r<char>, insert transaction, change transaction, open-line transaction으로 제한한다.
- `.` repeat 구현에서 delete/yank/put/Ex command/search/macro/register/count prefix와 Vim exact undo block semantics는 후속 hardening으로 미룬다.
- `tutorial-9-repeat-last-change`는 `A`, `ciw`, `o`, `r<char>` 뒤의 `.` 반복을 각각 다루는 4문항 efficiency tutorial이며 replay gate와 full playlist E2E를 통과한다.
- `search-basic`은 approved + engine implemented command cluster다. 첫 구현은 `/query enter`, `n`, `N` literal search이며 `?` backward search, regex, highlight, search history는 제외한다.
- `tutorial-90-search-basic`은 `/`, `n`, `N`, wrap-around literal search를 다루는 4문항 tutorial이며 replay gate와 full playlist E2E를 통과한다.
- 장기 반복 학습 플랫폼은 `docs/roadmap/PLATFORM_RFC_001.md`를 기준으로 검토한다. mastery, spaced review, daily run은 후보이며, progress 저장 포맷 변경은 별도 승인 전까지 구현하지 않는다.
- review queue는 저장 포맷 변경 없이 기존 progress v1 `Missions`와 content library만 읽는다. 메인 첫 화면에서는 `재점검 대상`, 성공 debrief에서는 `잔류 리스크`로 표현한다. candidate reason은 `미복구`, `복구 등급 <grade>`, `복구 입력 <best>/<optimal> keys`로 표시한다.
- 오늘의 복구 루트는 저장하지 않는 daily motivation layer다. 현재 review queue의 primary 대상과 이유를 읽어 `오늘의 복구 루트: 목표 문자까지 이동하기(미복구) 외 2건 대기`처럼 표시한다.
- E2E state summary는 저장 포맷과 별개로 review queue count, primary exercise, primary reason, daily route 문구를 노출한다.
- Incident Run은 tutorial이 아니라 별도 `incident-*` 카테고리로 다룬다. Incident는 새 command를 소개하지 않고 이미 배운 command를 조합해 생존 어드벤처 사건을 해결하는 적용 런이다.
- 중반 이후 incident는 `docs/gameplay/world-frame.md`의 원격 시설 복구국 / Runbook Dispatch 프레임을 따른다.
- `incident-001-hotfix`는 “릴레이 기지 001: 야간 핫픽스 복구”로 표시하며 `/error`, `/timeout` + `n`, `ciw`, `o`, `yy/p`, `:2,3s`를 조합하는 첫 mixed run이다. replay gate와 full playlist E2E를 통과한다.
- `incident-002-structure-recovery`는 “릴레이 기지 002: 구조 설정 재동기화”로 표시하며 `/secret`, `ci"`, `yi"` + `P`, `:%s`, `ci"` + `.`를 조합하는 두 번째 mixed run이다. replay gate와 full playlist E2E를 통과한다.
- `incident-003-visual-recovery`는 “릴레이 기지 003: 오염 구간 격리”로 표시하며 `/contam`, visual `d`, visual `y` + `p`, backward visual `d`, `:%s`를 조합하는 세 번째 mixed run이다. replay gate와 full playlist E2E를 통과한다.
- `incident-004-linewise-block-recovery`는 “릴레이 기지 004: config block 복구”로 표시하며 `/block`, linewise `Vd`, linewise `Vy` + `p`, linewise `VGd`, `:%s`를 조합하는 네 번째 mixed run이다. replay gate와 full playlist E2E를 통과한다.
- command choice drill은 새 command cluster가 아니라 이미 배운 command 중 적절한 도구를 고르는 적용 레이어다. 첫 설계 기준은 `docs/gameplay/command-choice-drills.md`를 따른다.
- command choice playable은 새 schema 승인 전까지 기존 `constraints.required_keys`와 `constraints.forbidden_keys`로 의도한 선택을 고정한다.
- command choice scenario는 정답 key sequence보다 선택 이유를 성공/실패 피드백에서 강화한다.
- `incident-005-command-choice`는 “릴레이 기지 005: 복구 범위 판별”로 표시하며, `choice-001-scope-triage`를 첫 playable command-choice drill로 승격한다. 새 command cluster나 engine 기능을 추가하지 않고 `visual-line-basic`의 `V`, `j`, `d`로 linewise scope choice를 검증한다.
- incident 001/002/003의 exercise는 각 beat마다 2단계 이상의 hint를 제공하며, scenario wording은 target state, optimal keys, constraints를 바꾸지 않는다.
- `visual-char-line`은 approved + implemented command cluster다. 첫 playable tutorial은 같은 줄 charwise `v` selection과 `d/y` 적용을 다룬다.
- VISUAL-GAP-002는 첫 visual 구현 범위를 charwise `v`로 좁혔다. selection은 `kind`, `anchor`, `head`, normalized inclusive `start`/`end`로 표현한다.
- 첫 visual foundation은 `v` 진입, 기존 normal motion으로 head 이동, `esc`/visual mode 중 `v`로 selection 해제를 다룬다.
- visual mode E2E는 screen text만으로 통과하지 않고 app_state `selection` object를 검증한다. E2E-007은 `internal/e2estate`, runner `assert.app_state`, content `e2e_assertions`에 selection shape를 추가했다.
- QA-SEL-001 이후 `replay_status: pass` exercise가 `e2e_assertions.selection`을 선언하면 content replay gate가 optimal key replay 결과의 selection active/kind/anchor/head/start/end를 검증한다.
- VIM-027-TUI-003은 charwise `v` visual mode foundation을 구현했다. TUI는 `Selection: charwise <start> -> <end>`와 selected non-cursor cell `{x}`를 표시하며, cursor cell은 기존 `[x]`를 유지한다.
- VISUAL-OP-001은 첫 visual operator 범위를 같은 줄 charwise selection에 `d`/`y`를 적용하는 것으로 고정했다. `d`는 inclusive range를 삭제하고, `y`는 inclusive range를 unnamed register에 저장하며, 둘 다 normal mode로 돌아가고 selection을 clear한다.
- VIM-028은 같은 줄 charwise visual selection의 `d`/`y` engine을 구현했다. `d`는 삭제 텍스트를 unnamed charwise register에도 저장하며 undo 가능하다. multi-line charwise operator, linewise visual, visual block은 제외한다.
- `tutorial-92-visual-selection`은 `v + motion + d`, `v + motion + y + p`, backward selection normalization을 다루는 3문항 visual selection tutorial이며 replay gate와 full playlist E2E를 통과한다.
- VISUAL-HARDEN-001은 새 visual 기능을 추가하지 않고, boundary motion, empty line, visual yank undo behavior, visual delete undo/register behavior, multi-line unsupported preservation, out-of-range selection clamp를 테스트로 고정했다.
- `V` linewise visual, visual block(`<C-v>`), count prefix, register prefix, indentation command, mouse/terminal selection 연동은 첫 visual slice에서 제외한다.
- VISUAL-LINE-GAP-001은 다음 visual 후보를 linewise `V` + row motion + `d/y`로 좁혔다. multi-line charwise `v` operator와 visual block은 후속으로 미룬다.
- VISUAL-LINE-001은 linewise `V`의 첫 구현 범위를 `V` 진입/해제, row motion, linewise `d/y`, linewise register, TUI/app_state `selection.kind: linewise`로 고정한다.
- `visual-line-basic`은 approved + implemented command cluster다. `tutorial-93-visual-line`은 linewise `V + j + d`, `V + y + j + p`, `V + G + d`를 다루는 3문항 tutorial이며 replay gate와 full playlist E2E를 통과한다.

> 재기획이 승인되고 구현된 항목만 여기에 이동한다. 기존 `docs/archived/PLAN.md`, `docs/archived/GAME_DESIGN.md`, `internal/` 구현은 참고 자료일 뿐이다.

## 미확인 사항

- [x] 첫 5분 플레이 루프를 정의해야 한다. 결정: `docs/gameplay/first-five-minute-route.md`의 FTUE-001 route를 따른다.
- [x] CONTENT-001 loader가 읽을 실제 content file 경로를 정해야 한다. 결정: repo root `content/` 아래 YAML.
- [x] Vim 핵심 영역 coverage rubric을 승인해야 한다. 결정: `docs/gameplay/vim-curriculum-map.md`의 Priority Bands와 Next Playpack Candidate를 따른다.
- [x] 실패/힌트/등급 시스템이 학습 동기를 해치지 않는 기준을 정의해야 한다. 결정: 최대 입력 수/금지 입력은 즉시 실패, 초반 코칭은 개념 중심, 재시도는 무제한 기본.
- [x] 기존 Vim emulator를 유지, 축소, 교체할지 결정해야 한다. 결정: 새 `internal/vimengine`을 만들고 기존 `internal/editor`는 LEGACY-001에서 archive한다.

---

## 수용 기준

> 이 도메인은 **모드 B**로 운영한다. Agent가 의도를 받아 `[draft]` 초안을 작성하고, 사람이 승인하면 `[draft]`를 제거한다.

### 재기획 준비

- 제품 기획을 시작하기 전에 `docs/roadmap/PRODUCT.md`에 목표, 표면, 기둥, 워크스트림이 채워져 있어야 한다.
- [draft] 첫 구현 slice를 시작하기 전에 `docs/roadmap/PROGRAM.md`에 현재 phase와 활성 slice가 있어야 한다.
- [draft] 기존 `docs/archived/GAME_DESIGN.md`에서 재사용할 아이디어는 그대로 구현하지 않고, 이 spec 또는 ExecPlan의 승인된 수용 기준으로 먼저 승격해야 한다.

### Content Gate

- approved + implemented command cluster는 비어 있지 않은 `coverage_required`를 가져야 한다.
- approved + implemented exercise는 `replay_status: pass`가 아니면 load에 실패한다.
- `replay_status: pass`인 approved + implemented exercise는 `optimal_keys` 재생 결과가 목표 상태와 E2E assertion을 만족해야 한다.
- playable exercise 목록은 replay gate를 통과한 exercise만 반환한다.
- playable playlist 목록은 approved/implemented playlist만 반환하며 ID 순서로 정렬된다.
- approved/implemented tutorial playlist는 8문항을 초과하면 load에 실패한다.
- `constraints.max_inputs`와 `constraints.attempt_limit`는 0 이상이어야 한다. 0은 제한 없음이다.
