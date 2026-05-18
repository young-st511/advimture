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
  engine_support: planned
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
    - `:q!`, `:wq`는 실제 파일 저장보다 exercise 종료/성공 상태 전이를 우선 검증한다.
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

## First 5-Minute Discovery Notes

- `normal-motion-basic`은 현재 엔진과 playable path에서 바로 파일 기반 콘텐츠로 승격 가능한 cluster다. 다만 현재 draft exercise는 `h`, `k` optimal coverage가 부족하므로 playable 확장 전 보강이 필요하다.
- `survival-save-quit`은 첫 경험의 불안감을 줄이는 데 중요하지만 command-line 입력과 app exit semantics가 필요하다.
- `word-motion-basic`은 `w/b/e` engine support가 구현됐다. playable 승격 전에는 각 command가 optimal trace에 등장하는 exercise set과 replay gate를 통과해야 한다.
- CONTENT-001 loader는 `engine_support: planned` 콘텐츠를 읽을 수 있되, playable 후보에서는 제외할 수 있어야 한다.

## Approval Packet — VIM-001

사용자 결정으로 첫 playable 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`이다.

추천:

- `normal-motion-basic`: `approved`. 단 `h`, `k` coverage exercise를 후속 루프에서 반드시 보강한다.
- `survival-save-quit`: `approved`, `engine_support: planned`.
- `word-motion-basic`: `approved`, `engine_support: implemented`. 단 exercise/scenario playable 승격은 후속 루프에서 진행한다.

주의: `approved`는 학습 우선순위 승인이고, `implemented` 또는 playable 연결을 의미하지 않는다.
