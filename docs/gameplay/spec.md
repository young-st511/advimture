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
- `whole-file-navigation`은 `gg`, `G`, `0`, `$` 각각이 approved + implemented exercise coverage와 replay gate를 통과한다.
- `gg/G`는 NAV-001에서 파일 처음/끝 줄의 첫 column으로 이동하는 pedagogical motion으로 다룬다.
- `vim-ex-command-substitute`는 `:s`, `:%s`, `:2,3s` 각각이 approved + implemented exercise coverage와 replay gate를 통과한다.
- substitute command는 EXCMD-001에서 literal match만 지원하며, scenario success는 buffer target으로 검증한다.
- playable은 approved/implemented tutorial playlist를 ID 순서로 실행한다.
- 현재 playable tutorial 순서는 `tutorial-0-movement`, `tutorial-1-survival`, `tutorial-2-fast-navigation`, `tutorial-3-small-edits`, `tutorial-4-ex-command`다.
- `first-5-minute`는 legacy vertical slice로 retired 상태이며 default playable path에서 실행하지 않는다.
- 화면은 현재 tutorial title과 episode-local exercise count를 표시한다.
- 한 tutorial 마지막 exercise 성공 시 다음 tutorial이 있으면 `Next tutorial: enter`를 표시하고, `enter`로 다음 tutorial에 진입한다.
- exercise 성공 시 기존 progress `Missions` map에 exercise ID를 key로 자동 저장하고, 성공 상태에서 `enter`를 누르면 다음 unlocked exercise로 이동한다.
- 향후 exercise constraint는 최대 입력 수 초과와 금지 입력/금지 우회 전략 사용을 즉시 실패로 처리해야 한다.
- 실패 횟수는 기본 무제한이며, 후반 콘텐츠를 위해 `attempt_limit` 설정 여지는 남긴다.
- 실패 후 재시도는 `r`과 `enter`를 모두 허용한다.
- 초반 튜토리얼 코칭은 개념 힌트 중심이지만, 새 command 첫 소개에서는 command 의미를 명시한다.
- exercise YAML은 `constraints.max_inputs`, `constraints.required_keys`, `constraints.forbidden_keys`, `constraints.attempt_limit`를 선언할 수 있다.
- `constraints.max_inputs` 초과와 `constraints.forbidden_keys` 입력은 Vim state를 추가 진행시키지 않고 즉시 실패한다.
- 목표에 도착했더라도 `constraints.required_keys`가 key trace에 없으면 실패한다.
- 실패 상태는 progress를 저장하지 않으며, 실패 화면은 `Grade: F`, 남은 입력 수, 재시도 안내를 보여준다.
- 실패 화면은 attempt count를 표시하며 `attempt_limit: 0`은 `unlimited`로 표현한다.
- scoring result는 runtime failure reason을 보존하며, `required_keys_missing`은 `IntentSatisfied=false`, `Grade=F`로 평가한다.
- `single-char-edit`, `insert-mode-entry`, `undo-redo-basic`은 approved + implemented tutorial cluster이며 `x`, `r`, `i`, `a`, `A`, `u`, `ctrl+r` coverage와 replay gate를 통과한다.
- `tutorial-3-small-edits`는 7문항짜리 작은 수정 튜토리얼이며 Ex command보다 먼저 실행된다.
- `undo-redo-basic` 문항은 required key 없이 최종 목표에 먼저 도착하지 않도록 target cursor와 optimal trace를 설계한다.

> 재기획이 승인되고 구현된 항목만 여기에 이동한다. 기존 `docs/archived/PLAN.md`, `docs/archived/GAME_DESIGN.md`, `internal/` 구현은 참고 자료일 뿐이다.

## 미확인 사항

- [ ] 첫 5분 플레이 루프를 정의해야 한다.
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
