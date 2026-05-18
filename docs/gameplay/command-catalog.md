# Vim Command Catalog

> Vim 학습 설계의 첫 번째 저장소다. 모든 exercise와 scenario는 여기의 command cluster에서 출발한다.

## 운영 규칙

- 새 항목은 `[draft]` 상태로 추가한다.
- 사람이 승인하면 상태를 `approved`로 바꾼다.
- `approved` command cluster만 exercise 설계에 사용할 수 있다.
- command cluster는 “명령어 나열”이 아니라 하나의 학습 목표로 묶는다.
- 시나리오 아이디어가 먼저 떠오른 경우에도, 먼저 이 catalog의 command cluster로 환원한다.
- 각 cluster는 Vim 호환성 티어를 가진다. 엔진 구현과 oracle 검증 범위는 이 티어를 따른다.

## 상태 값

| 상태 | 의미 |
|------|------|
| `draft` | Agent 또는 사람이 제안했지만 아직 학습 우선순위로 승인되지 않음 |
| `approved` | exercise 설계에 사용할 수 있음 |
| `implemented` | 게임 내 문항과 검증까지 연결됨 |
| `retired` | 더 이상 새 문항에 사용하지 않음 |

## 호환성 티어

| 티어 | 의미 |
|------|------|
| `exact` | Neovim과 같은 최종 buffer, cursor, mode 결과를 목표로 한다. |
| `pedagogical` | 학습 목적상 단순화한 동작을 허용한다. 차이는 `compatibility_notes`에 문서화한다. |
| `unsupported` | 현재 게임에서 다루지 않는다. 입력을 차단하고 이유를 설명한다. |

## Command Cluster 스키마

```yaml
command_cluster:
  id: <kebab-case-id>
  status: draft | approved | implemented | retired
  compatibility_tier: exact | pedagogical | unsupported
  engine_support: implemented | planned | unsupported
  curriculum_area: <chapter-id>
  title: <학습 목표 이름>
  commands: ["<vim-key-or-command>"]
  coverage_required: ["<command that must appear in optimal traces>"]
  oracle: required | optional | none
  purpose: <한 문장 목적>
  prerequisite: ["<command-cluster-id>"]
  difficulty: beginner | intermediate | advanced
  useful_when:
    - <실전에서 유용한 상황>
  combo_paths:
    - ["<다음에 연결될 command>"]
  common_mistakes:
    - <초보자가 자주 하는 실수>
  compatibility_notes:
    - <Vim/Neovim과 다르게 처리하거나 아직 검증하지 않는 부분>
  design_notes:
    - <exercise 설계 시 주의점>
```

## Clusters

### survival-save-quit

```yaml
command_cluster:
  id: survival-save-quit
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-0-survival
  title: Vim 생존: 모드 복귀와 저장/종료
  commands: ["esc", ":q!", ":wq"]
  coverage_required: ["esc", ":q!", ":wq"]
  oracle: none
  purpose: Vim에서 안전하게 빠져나오고, 수정한 내용을 저장 후 종료한다.
  prerequisite: []
  difficulty: beginner
  useful_when:
    - Vim에 처음 진입했을 때 당황하지 않고 Normal mode로 돌아올 때
    - 파일 수정을 저장하고 종료해야 할 때
    - 실수한 변경을 버리고 나가야 할 때
  combo_paths:
    - ["i", "esc", ":wq"]
  common_mistakes:
    - Insert mode에서 바로 `:wq`를 입력하려고 한다.
    - 저장해야 하는 상황과 버려야 하는 상황을 구분하지 못한다.
  compatibility_notes:
    - `esc`는 Vim의 mode 복귀와 동일한 학습 목표로 다룬다.
    - `:q!`, `:wq`는 실제 파일 저장보다 command-line 입력과 exercise command goal을 우선 검증한다.
  design_notes:
    - 첫 학습 문항은 조작 성공 경험을 우선한다.
    - `:q!`와 `:wq`의 차이를 문항 목표 상태로 분리한다.
```

### normal-motion-basic

```yaml
command_cluster:
  id: normal-motion-basic
  status: approved
  compatibility_tier: exact
  engine_support: implemented
  curriculum_area: chapter-1-movement
  title: 기본 커서 이동
  commands: ["h", "j", "k", "l"]
  coverage_required: ["h", "j", "k", "l"]
  oracle: optional
  purpose: Normal mode에서 홈 포지션을 유지하며 커서를 이동한다.
  prerequisite: ["survival-save-quit"]
  difficulty: beginner
  useful_when:
    - 화살표 키 없이 짧은 거리의 목표 위치로 이동할 때
    - 이후 단어 이동과 operator 조합을 배우기 전 공간 감각을 만들 때
  combo_paths:
    - ["w", "b", "e"]
    - ["0", "$"]
  common_mistakes:
    - `h`와 `l`, `j`와 `k` 방향을 헷갈린다.
    - 긴 이동도 계속 `h/j/k/l`로만 해결하려고 한다.
  compatibility_notes:
    - Normal mode의 줄/열 경계 처리와 짧은 줄로 이동할 때의 cursor clamp를 Neovim 결과와 맞추는 것을 목표로 한다.
  design_notes:
    - 긴 거리 이동 문항으로 과도하게 반복시키지 않는다.
    - 이후 빠른 이동 command의 유용성을 체감할 비교군으로 사용한다.
```

### word-motion-basic

```yaml
command_cluster:
  id: word-motion-basic
  status: approved
  compatibility_tier: exact
  engine_support: implemented
  curriculum_area: chapter-1-movement
  title: 단어 단위 이동
  commands: ["w", "b", "e"]
  coverage_required: ["w", "b", "e"]
  oracle: required
  purpose: 단어 경계를 이용해 한 글자씩 이동하지 않고 빠르게 위치를 잡는다.
  prerequisite: ["normal-motion-basic"]
  difficulty: beginner
  useful_when:
    - 긴 설정 줄에서 옵션 이름이나 값으로 빠르게 이동할 때
    - `d`, `c`, `y` operator와 조합하기 전 motion 감각을 익힐 때
  combo_paths:
    - ["dw", "cw", "yw"]
    - ["ciw", "diw", "yiw"]
  common_mistakes:
    - 단어 시작, 단어 끝, 이전 단어의 차이를 구분하지 못한다.
    - 특수문자와 공백에서 커서가 어디로 가는지 예측하지 못한다.
  compatibility_notes:
    - keyword 경계, 공백, 문장부호, 줄 경계는 `internal/vimengine` 단위 테스트와 oracle-style fixture로 고정한다.
  design_notes:
    - `hjkl` 대비 키 입력 수가 줄어드는 문항을 반드시 포함한다.
    - 같은 줄 안 이동과 줄 경계 이동을 별도 exercise로 분리한다.
```

### whole-file-navigation

```yaml
command_cluster:
  id: whole-file-navigation
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-2-navigation
  title: 줄과 파일 단위 이동
  commands: ["gg", "G", "0", "$"]
  coverage_required: ["gg", "G", "0", "$"]
  oracle: optional
  purpose: 긴 파일에서 한 글자씩 이동하지 않고 줄 시작/끝과 파일 처음/끝으로 빠르게 이동한다.
  prerequisite: ["word-motion-basic"]
  difficulty: intermediate
  useful_when:
    - 설정 파일 맨 위의 선언부로 돌아갈 때
    - 로그나 설정 파일의 마지막 줄로 즉시 이동할 때
    - 현재 줄의 키나 값 경계를 빠르게 확인할 때
  combo_paths:
    - ["gg", "0"]
    - ["G", "$"]
    - ["0", "w", "$"]
  common_mistakes:
    - 긴 파일에서도 `j/k`를 반복한다.
    - 현재 줄의 시작과 파일의 시작을 혼동한다.
  compatibility_notes:
    - NAV-001에서는 `gg/G`를 줄의 첫 column으로 이동하는 pedagogical motion으로 다룬다.
    - numeric count prefix와 first non-blank column semantics는 후속 루프에서 다룬다.
  design_notes:
    - 후반 navigation은 모든 이동을 나열하지 않고 `coverage_required`로 좁혀 구현한다.
    - 문항은 `hjkl` 반복보다 키 입력 수가 줄어드는 상황을 사용한다.
```

### vim-ex-command-substitute

```yaml
command_cluster:
  id: vim-ex-command-substitute
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-ex-command
  title: Ex substitute와 range
  commands: [":s", ":%s", ":2,3s"]
  coverage_required: [":s", ":%s", ":2,3s"]
  oracle: optional
  purpose: command-line에서 현재 줄, 전체 파일, 지정 범위의 문자열을 치환한다.
  prerequisite: ["survival-save-quit", "whole-file-navigation"]
  difficulty: intermediate
  useful_when:
    - 현재 줄의 잘못된 토큰을 빠르게 바꿀 때
    - 파일 전체의 반복된 상태값을 한 번에 바꿀 때
    - 특정 줄 범위 안에서만 같은 오타를 고칠 때
  combo_paths:
    - [":s", "n"]
    - [":%s", ":%s///g"]
    - [":2,3s", ":wq"]
  common_mistakes:
    - Normal mode 키 입력과 command-line 입력을 혼동한다.
    - 현재 줄 치환과 전체 파일 치환의 범위를 구분하지 못한다.
  compatibility_notes:
    - EXCMD-001은 Vim regex가 아니라 literal match substitute만 지원한다.
    - 지원 flag는 global 치환을 뜻하는 `g`로 제한한다.
  design_notes:
    - 시나리오 성공은 buffer target으로 검증하고 app quit과 분리한다.
    - range command는 1-based inclusive 줄 번호로 설명한다.
```

### single-char-edit

```yaml
command_cluster:
  id: single-char-edit
  status: draft
  compatibility_tier: pedagogical
  engine_support: planned
  curriculum_area: chapter-2-small-edits
  title: 문자 하나 고치기
  commands: ["x", "r"]
  coverage_required: ["x", "r"]
  oracle: optional
  purpose: Normal mode에서 작은 오타를 Insert mode 진입 없이 빠르게 고친다.
  prerequisite: ["normal-motion-basic", "word-motion-basic"]
  difficulty: beginner
  useful_when:
    - 설정 키나 값의 글자 하나가 틀렸을 때
    - 불필요한 문자 하나를 제거할 때
    - Insert mode로 들어가기 전 작은 수정 감각을 만들 때
  combo_paths:
    - ["x", "u"]
    - ["r", "u"]
    - ["w", "x"]
  common_mistakes:
    - 글자 하나를 고치려고 매번 Insert mode에 들어간다.
    - `x`가 현재 커서 아래 문자를 지운다는 점을 잊는다.
    - `r` 뒤에 교체할 문자 하나만 입력한다는 점을 놓친다.
  compatibility_notes:
    - 첫 구현은 single-width 문자 중심의 pedagogical behavior를 허용한다.
    - count prefix와 multibyte edge case는 후속 exact tier hardening에서 다룬다.
  design_notes:
    - 문항은 목표 문자를 찾는 이동과 실제 수정 입력을 분리해서 설계한다.
    - `constraints.required_keys`로 `x` 또는 `r` 사용을 고정한다.
```

### insert-mode-entry

```yaml
command_cluster:
  id: insert-mode-entry
  status: draft
  compatibility_tier: pedagogical
  engine_support: planned
  curriculum_area: chapter-2-small-edits
  title: Insert mode 진입과 작은 추가
  commands: ["i", "a", "A"]
  coverage_required: ["i", "a", "A", "esc"]
  oracle: optional
  purpose: 커서 주변 또는 줄 끝에 짧은 텍스트를 추가하고 Normal mode로 복귀한다.
  prerequisite: ["survival-save-quit", "normal-motion-basic"]
  difficulty: beginner
  useful_when:
    - 커서 앞/뒤에 값을 끼워 넣을 때
    - 설정 줄 끝에 suffix나 flag를 추가할 때
    - 수정 후 `esc`로 Normal mode를 회복해야 할 때
  combo_paths:
    - ["i", "esc", ":wq"]
    - ["A", "esc", ":wq"]
    - ["w", "a", "esc"]
  common_mistakes:
    - Insert mode에 들어간 뒤 Normal mode로 돌아오지 않는다.
    - `i`와 `a`의 삽입 위치 차이를 혼동한다.
    - 줄 끝 추가에도 `l`을 반복해서 끝까지 이동한다.
  compatibility_notes:
    - 첫 구현은 printable rune 입력과 `esc` 복귀를 우선한다.
    - 복잡한 insert editing, backspace, newline은 후속 루프에서 다룬다.
  design_notes:
    - 첫 문항은 `i`, `a`, `A`의 위치 차이를 짧게 분리한다.
    - 모든 문항은 수정 후 `esc`까지 optimal trace에 포함한다.
```

### undo-redo-basic

```yaml
command_cluster:
  id: undo-redo-basic
  status: draft
  compatibility_tier: pedagogical
  engine_support: planned
  curriculum_area: chapter-2-small-edits
  title: 실수 되돌리기와 다시 적용
  commands: ["u", "<C-r>"]
  coverage_required: ["u", "<C-r>"]
  oracle: optional
  purpose: 작은 수정 실수를 안전하게 되돌리고 필요한 경우 다시 적용한다.
  prerequisite: ["single-char-edit", "insert-mode-entry"]
  difficulty: beginner
  useful_when:
    - 잘못 지운 문자나 변경을 즉시 복구할 때
    - 되돌린 변경이 맞았음을 확인하고 다시 적용할 때
    - 실패를 학습 루프로 회복하는 감각을 만들 때
  combo_paths:
    - ["x", "u"]
    - ["r", "u"]
    - ["u", "<C-r>"]
  common_mistakes:
    - 실수 후 수동으로 다시 편집하려고 한다.
    - undo 후 redo가 가능하다는 점을 모른다.
  compatibility_notes:
    - 첫 구현은 exercise 단위 mutation history를 pedagogical stack으로 다룬다.
    - Vim의 undo block, insert transaction semantics는 후속 hardening에서 다룬다.
  design_notes:
    - 억까 상황은 과하지 않게, 실패 회복의 재미를 주는 정도로 사용한다.
    - undo/redo는 점수 감점보다 학습 회복 메시지를 우선한다.
```

## First 5-Minute Discovery Notes

- `normal-motion-basic`은 현재 엔진과 playable path에서 `h/j/k/l` optimal coverage를 모두 가진 cluster다.
- `survival-save-quit`은 command-line 입력과 app exit semantics를 분리해 구현됐다. 실제 파일 저장/폐기는 아직 수행하지 않는다.
- `word-motion-basic`은 `w/b/e` engine support가 구현됐다. playable 승격 전에는 각 command가 optimal trace에 등장하는 exercise set과 replay gate를 통과해야 한다.
- `whole-file-navigation`은 `gg/G/0/$` engine support가 구현됐다. `gg/G`는 현재 pedagogical tier로 첫 column 이동만 다룬다.
- `vim-ex-command-substitute`는 literal `:s`, `:%s`, `:2,3s` engine support가 구현됐다. Vim regex와 복잡한 flags는 아직 다루지 않는다.
- CONTENT-001 loader는 `engine_support: planned` 콘텐츠를 읽을 수 있되, playable 후보에서는 제외할 수 있어야 한다.

## Approval Packet — VIM-001

사용자 결정으로 첫 playable 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`이다.

추천:

- `normal-motion-basic`: `approved`, `engine_support: implemented`. `h/j/k/l` coverage exercise가 모두 replay gate를 통과한다.
- `survival-save-quit`: `approved`, `engine_support: implemented`. 단 실제 파일 저장/폐기는 후속 앱/파일 작업 루프에서만 다룬다.
- `word-motion-basic`: `approved`, `engine_support: implemented`. 단 exercise/scenario playable 승격은 후속 루프에서 진행한다.

주의: `approved`는 학습 우선순위 승인이고, `implemented` 또는 playable 연결을 의미하지 않는다.
