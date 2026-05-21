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
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
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
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
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
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-2-small-edits
  title: 실수 되돌리기와 다시 적용
  commands: ["u", "ctrl+r"]
  coverage_required: ["u", "ctrl+r"]
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
    - ["u", "ctrl+r"]
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

### delete-with-motion

```yaml
command_cluster:
  id: delete-with-motion
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-operator-grammar
  title: Operator grammar: 삭제와 motion
  commands: ["d", "dw", "d$", "dd"]
  coverage_required: ["dw", "d$", "dd"]
  oracle: optional
  purpose: Vim의 `operator + motion` 문법으로 단어, 줄 끝, 전체 줄을 한 번에 삭제한다.
  prerequisite: ["word-motion-basic", "whole-file-navigation", "undo-redo-basic"]
  difficulty: intermediate
  useful_when:
    - 위험하거나 불필요한 단어를 한 번에 제거할 때
    - 현재 위치부터 줄 끝까지 노이즈를 삭제할 때
    - 오염된 설정 줄 전체를 제거할 때
  combo_paths:
    - ["dw", "u"]
    - ["d$", "u"]
    - ["dd", "u"]
  common_mistakes:
    - `x`를 반복해 단어 전체를 지우려고 한다.
    - `d`를 누른 뒤 motion을 입력해야 한다는 operator pending 흐름을 놓친다.
    - `dd`가 줄 전체 삭제라는 점을 잊고 `d$`와 혼동한다.
  compatibility_notes:
    - 첫 구현은 count prefix와 text object를 다루지 않는다.
    - `dw`의 공백 포함 범위는 학습 목적의 pedagogical behavior로 고정하고 후속 exact hardening에서 재검토한다.
  design_notes:
    - 첫 playpack은 `dw`, `d$`, `dd`만 다룬다.
    - 모든 문항은 `constraints.required_keys`로 의도 operator sequence를 고정한다.
```

### change-with-motion

```yaml
command_cluster:
  id: change-with-motion
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-operator-grammar
  title: Operator grammar: 변경과 motion
  commands: ["c", "cw", "c$", "cc"]
  coverage_required: ["cw", "c$", "cc"]
  oracle: optional
  purpose: 삭제와 Insert mode 진입을 하나의 operator grammar로 묶어 빠르게 값을 교체한다.
  prerequisite: ["insert-mode-entry", "delete-with-motion"]
  difficulty: intermediate
  useful_when:
    - 잘못된 설정 단어를 새 값으로 교체할 때
    - 현재 위치부터 줄 끝까지 새 문구로 바꿀 때
    - 손상된 한 줄짜리 명령을 통째로 다시 작성할 때
  combo_paths:
    - ["cw", "esc"]
    - ["c$", "esc"]
    - ["cc", "esc"]
  common_mistakes:
    - `d`로 지운 뒤 다시 `i` 또는 `a`를 눌러 두 단계로 처리한다.
    - change operator 후 Insert mode에 들어간다는 점을 잊는다.
    - 수정 후 `esc`로 Normal mode에 돌아오지 않는다.
  compatibility_notes:
    - 첫 구현은 `cw`, `c$`, `cc`만 지원하고 text object change는 후속 루프에서 다룬다.
    - Vim의 세부 whitespace semantics는 pedagogical tier로 단순화한다.
  design_notes:
    - change 문항은 최종 mode가 Normal mode가 되도록 `esc`를 optimal trace에 포함한다.
    - `cw`와 `c$`는 delete 범위와 insert entry가 동시에 학습되도록 짧은 문자열로 설계한다.
```

### yank-put-basic

```yaml
command_cluster:
  id: yank-put-basic
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-operator-grammar
  title: Operator grammar: 복사와 붙여넣기
  commands: ["y", "yw", "y$", "yy", "p", "P"]
  coverage_required: ["yw", "y$", "yy", "p", "P"]
  oracle: optional
  purpose: 안전한 텍스트 조각이나 줄을 register에 보관한 뒤 필요한 위치에 재사용한다.
  prerequisite: ["delete-with-motion", "change-with-motion"]
  difficulty: intermediate
  useful_when:
    - 정상 설정 줄을 아래/위에 복제할 때
    - 구조 코드나 짧은 토큰을 다시 입력하지 않고 재사용할 때
    - 위험한 수동 재입력을 피하고 검증된 값을 붙여넣을 때
  combo_paths:
    - ["yy", "p"]
    - ["yy", "P"]
    - ["yw", "p"]
    - ["y$", "p"]
  common_mistakes:
    - 복사할 수 있는 값을 직접 다시 타이핑한다.
    - yy가 줄 단위 register를 만든다는 점을 p/P 결과와 연결하지 못한다.
    - p와 P의 붙여넣기 방향을 혼동한다.
  compatibility_notes:
    - 첫 구현은 unnamed register만 다룬다.
    - named register, numbered register, system clipboard는 후속 milestone으로 미룬다.
    - 첫 구현은 single-line charwise yank와 linewise yank만 다룬다.
  design_notes:
    - VIM-019는 yank/register를 구현했고, VIM-020에서 put을 구현했다.
    - 첫 playpack은 y/p의 재사용 감각을 먼저 다루고 text object는 다음 playpack으로 분리한다.
```

### text-object-inner-word

```yaml
command_cluster:
  id: text-object-inner-word
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-operator-grammar
  title: Text object: 단어 내부를 대상으로 편집
  commands: ["iw", "diw", "ciw", "yiw"]
  coverage_required: ["diw", "ciw", "yiw"]
  oracle: optional
  purpose: 커서가 단어 중간에 있어도 단어 전체를 대상으로 삭제, 변경, 복사한다.
  prerequisite: ["yank-put-basic", "delete-with-motion", "change-with-motion"]
  difficulty: intermediate
  useful_when:
    - 커서가 단어 중간에 있을 때도 전체 값을 대상으로 잡을 때
    - config value나 command argument를 위치보다 구조 기준으로 다룰 때
    - operator grammar를 motion에서 text object로 확장할 때
  combo_paths:
    - ["diw", "u"]
    - ["ciw", "esc"]
    - ["yiw", "p"]
  common_mistakes:
    - 커서 위치를 단어 시작으로 옮긴 뒤 dw/cw/yw로 처리하려고 한다.
    - i가 Insert mode가 아니라 text object prefix가 되는 문맥을 놓친다.
    - quote/pair object까지 한 번에 외우려고 한다.
  compatibility_notes:
    - 첫 구현은 iw만 다룬다.
    - i\", i', i(, i{와 around object는 후속 루프로 미룬다.
    - count prefix, visual selection, whitespace 세부 semantics는 후속 hardening으로 미룬다.
  design_notes:
    - VIM-021에서 `operator -> i -> w` 3-key pending sequence를 구현했고, VIM-022에서 `diw`, `ciw`, `yiw` semantics를 구현했다.
    - 첫 text object playpack은 단어 내부 object만 다루며, quote/pair는 별도 고급 튜토리얼로 분리한다.
```

### open-line-edit

```yaml
command_cluster:
  id: open-line-edit
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-2-small-edits
  title: 줄 열기와 즉시 입력
  commands: ["o", "O"]
  coverage_required: ["o", "O"]
  oracle: optional
  purpose: 현재 줄 아래나 위에 새 줄을 열고 바로 입력한다.
  prerequisite: ["insert-mode-entry", "undo-redo-basic"]
  difficulty: intermediate
  useful_when:
    - 설정 줄 아래에 fallback이나 새 항목을 추가할 때
    - 현재 줄 위에 guard나 주석을 끼워 넣을 때
    - 줄 단위 구조를 유지하면서 Insert mode 진입 동작을 줄일 때
  combo_paths:
    - ["o", "esc"]
    - ["O", "esc"]
    - ["o", "u"]
  common_mistakes:
    - 새 줄을 만들려고 A나 i로 기존 줄을 망가뜨린다.
    - 아래에 열어야 할지 위에 열어야 할지 o/O 방향을 혼동한다.
    - 입력 후 esc로 Normal mode에 돌아오지 않는다.
  compatibility_notes:
    - 첫 구현은 빈 줄을 삽입한 뒤 Insert mode로 진입한다.
    - indentation, auto-comment, count prefix, insert-mode Enter는 후속 hardening으로 미룬다.
    - 첫 구현은 dot repeat과 연결하지 않는다.
  design_notes:
    - OPEN-LINE-001은 scope를 고정했고, VIM-023에서 engine support를 구현했다.
    - PLAYPACK-006은 o/O 각각을 최소 2문항 이상 다루고 full playlist E2E를 추가했다.
```

### text-object-quote-pair

```yaml
command_cluster:
  id: text-object-quote-pair
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-operator-grammar
  title: Text object: 따옴표 내부 대상 편집
  commands: ["i\"", "di\"", "ci\"", "yi\""]
  coverage_required: ["di\"", "ci\"", "yi\""]
  oracle: optional
  purpose: 커서가 따옴표 내부 어디에 있든 값 내부를 삭제, 변경, 복사한다.
  prerequisite: ["text-object-inner-word", "yank-put-basic", "change-with-motion"]
  difficulty: intermediate
  useful_when:
    - JSON, env, config의 quoted value를 구조 기준으로 바꿀 때
    - 값 시작으로 이동하지 않고 quote 내부 전체를 잡을 때
    - 복사/붙여넣기와 operator grammar를 문자열 값 단위로 확장할 때
  combo_paths:
    - ["ci\"", "esc"]
    - ["di\"", "u"]
    - ["yi\"", "p"]
    - ["/", "ci\""]
  common_mistakes:
    - 커서를 값 시작으로 옮긴 뒤 cw/dw/yw로 처리하려고 한다.
    - operator 뒤의 i를 Insert mode로 착각한다.
    - quote 문자까지 지워야 한다고 착각한다.
  compatibility_notes:
    - 첫 구현은 double quote 내부 object만 다룬다.
    - quote 문자는 대상에 포함하지 않는다.
    - nested pair, escaped quote, single quote, parenthesis, brace, around object, count prefix, visual selection은 후속 hardening으로 미룬다.
  design_notes:
    - TEXT-PAIR-GAP-001은 double quote 내부 object로 첫 scope를 고정했다.
    - VIM-026은 double quote 내부 object의 engine/runtime support를 구현했고, content/E2E는 PLAYPACK-009로 분리한다.
```

### repeat-last-change

```yaml
command_cluster:
  id: repeat-last-change
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-efficiency
  title: 마지막 변경 반복
  commands: ["."]
  coverage_required: ["."]
  oracle: optional
  purpose: 같은 변경을 다시 입력하지 않고 현재 위치에 반복 적용한다.
  prerequisite: ["insert-mode-entry", "change-with-motion", "text-object-inner-word", "open-line-edit"]
  difficulty: intermediate
  useful_when:
    - 여러 줄에 같은 suffix나 flag를 붙일 때
    - 비슷한 설정 값을 연속해서 같은 값으로 교체할 때
    - 방금 만든 줄 추가/값 교체 패턴을 다음 위치에서 재사용할 때
  combo_paths:
    - ["A", "esc", "."]
    - ["cw", "esc", "."]
    - ["ciw", "esc", "."]
    - ["o", "esc", "."]
  common_mistakes:
    - 같은 텍스트를 다시 타이핑해 key count를 늘린다.
    - dot repeat이 motion이 아니라 마지막 change를 반복한다는 점을 놓친다.
    - undo/redo나 yank도 같은 방식으로 반복된다고 기대한다.
  compatibility_notes:
    - 첫 구현은 pedagogical last-change transaction을 사용한다.
    - 반복 대상은 x, r<char>, insert transaction, change transaction, open-line transaction으로 제한한다.
    - delete, yank, put, Ex command, search, macro/register/count prefix는 후속 hardening으로 미룬다.
    - Vim exact undo block semantics는 후속 hardening으로 미룬다.
  design_notes:
    - REPEAT-GAP-001에서 transaction commit/replay 규칙을 고정했다.
    - VIM-024는 엔진 transaction recorder를 구현했고, PLAYPACK-007에서 efficiency tutorial과 E2E를 추가했다.
```

### search-basic

```yaml
command_cluster:
  id: search-basic
  status: approved
  compatibility_tier: pedagogical
  engine_support: implemented
  curriculum_area: chapter-3-navigation
  title: 로그와 설정 literal 검색
  commands: ["/", "n", "N"]
  coverage_required: ["/", "n", "N"]
  oracle: optional
  purpose: 명시된 token을 직접 찾고 다음/이전 match로 빠르게 이동한다.
  prerequisite: ["word-motion-basic", "whole-file-navigation"]
  difficulty: intermediate
  useful_when:
    - 긴 로그에서 ERROR나 timeout 같은 token을 찾을 때
    - 설정 파일에서 같은 key의 다음 발생 위치로 이동할 때
    - 지나친 match에서 이전 match로 되돌아갈 때
  combo_paths:
    - ["/", "enter", "n"]
    - ["/", "enter", "N"]
    - ["/", "enter", "ciw"]
  common_mistakes:
    - 검색어 입력 후 enter를 누르지 않는다.
    - n과 N의 방향 차이를 혼동한다.
    - ?를 backward search로 기대한다.
  compatibility_notes:
    - 첫 구현은 case-sensitive literal search만 지원한다.
    - ? backward search는 TUI hint key와 충돌하므로 첫 구현에서 제외한다.
    - regex, highlight, search history, ignorecase/smartcase는 후속 hardening으로 미룬다.
  design_notes:
    - SEARCH-GAP-001에서 /, n, N만 첫 scope로 고정했다.
    - VIM-025는 search state와 cursor movement를 구현했고, PLAYPACK-008에서 content/E2E를 추가했다.
```

## First 5-Minute Discovery Notes

- `normal-motion-basic`은 현재 엔진과 playable path에서 `h/j/k/l` optimal coverage를 모두 가진 cluster다.
- `survival-save-quit`은 command-line 입력과 app exit semantics를 분리해 구현됐다. 실제 파일 저장/폐기는 아직 수행하지 않는다.
- `word-motion-basic`은 `w/b/e` engine support가 구현됐다. playable 승격 전에는 각 command가 optimal trace에 등장하는 exercise set과 replay gate를 통과해야 한다.
- `whole-file-navigation`은 `gg/G/0/$` engine support가 구현됐다. `gg/G`는 현재 pedagogical tier로 첫 column 이동만 다룬다.
- `vim-ex-command-substitute`는 literal `:s`, `:%s`, `:2,3s` engine support가 구현됐다. Vim regex와 복잡한 flags는 아직 다루지 않는다.
- `delete-with-motion`, `change-with-motion`은 VIM-017/VIM-018에서 engine support가 구현됐고 PLAYPACK-003에서 6문항 tutorial content로 연결됐다. 첫 구현 범위는 `dw`, `d$`, `dd`, `cw`, `c$`, `cc`다.
- `yank-put-basic`은 VIM-019/VIM-020에서 engine support가 구현됐고 PLAYPACK-004에서 5문항 tutorial content로 연결됐다. 첫 구현 범위는 `yw`, `y$`, `yy`, `p`, `P`다.
- `text-object-inner-word`는 VIM-021/VIM-022에서 engine support가 구현됐고 PLAYPACK-005에서 6문항 tutorial content로 연결됐다. 첫 구현 범위는 `iw` 기반 `diw`, `ciw`, `yiw`다.
- `open-line-edit`은 OPEN-LINE-001에서 approved로 승격했고, VIM-023에서 engine support를 구현했다. 첫 구현 범위는 `o`, `O`이며 indentation, auto-comment, count prefix, insert-mode Enter, dot repeat은 제외한다.
- `repeat-last-change`는 REPEAT-GAP-001에서 approved로 승격했고, VIM-024에서 engine support를 구현했다. 첫 구현은 x, r<char>, insert/change/open-line transaction을 대상으로 하며 delete/yank/put/search/macro/register/count prefix는 제외한다.
- `search-basic`은 SEARCH-GAP-001에서 approved로 승격했고, VIM-025에서 engine support를 구현했으며 PLAYPACK-008에서 4문항 tutorial content와 full E2E를 연결했다. 첫 구현은 `/`, `n`, `N` literal search이며 `?`, regex, highlight, search history는 제외한다.
- CONTENT-001 loader는 `engine_support: planned` 콘텐츠를 읽을 수 있되, playable 후보에서는 제외할 수 있어야 한다.

## Approval Packet — VIM-001

사용자 결정으로 첫 playable 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`이다.

추천:

- `normal-motion-basic`: `approved`, `engine_support: implemented`. `h/j/k/l` coverage exercise가 모두 replay gate를 통과한다.
- `survival-save-quit`: `approved`, `engine_support: implemented`. 단 실제 파일 저장/폐기는 후속 앱/파일 작업 루프에서만 다룬다.
- `word-motion-basic`: `approved`, `engine_support: implemented`. 단 exercise/scenario playable 승격은 후속 루프에서 진행한다.

주의: `approved`는 학습 우선순위 승인이고, `implemented` 또는 playable 연결을 의미하지 않는다.
