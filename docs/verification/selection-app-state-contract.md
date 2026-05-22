# Selection App State Contract

## 목적

Visual mode는 “무엇을 선택했는지” 자체가 학습 피드백이다. 따라서 첫 구현 전에 engine, TUI, content assertion, E2E runner가 같은 selection 모델을 사용해야 한다.

## 첫 구현 범위

- 포함: charwise `v`, 기존 normal motion으로 selection head 이동, `esc`로 selection 해제
- 포함: app_state selection assertion
- 포함: TUI 최소 표시
- 제외: `V` linewise visual mode
- 제외: visual block `<C-v>`
- 제외: visual selection에 `d`, `y`, `c`, `>`, `<` 적용
- 제외: count/register prefix
- 제외: mouse/terminal selection 연동

## Selection Model

```yaml
selection:
  active: true
  kind: charwise
  anchor:
    row: 0
    col: 1
  head:
    row: 0
    col: 3
  start:
    row: 0
    col: 1
  end:
    row: 0
    col: 3
```

- `active`: visual selection이 켜져 있으면 true다.
- `kind`: 첫 구현은 `charwise`만 허용한다.
- `anchor`: `v`를 누른 순간의 cursor다.
- `head`: 현재 cursor다.
- `start`/`end`: `anchor`와 `head`를 문서 순서로 정규화한 inclusive range다.
- charwise selection은 cursor cell을 포함한다.
- `esc`는 mode를 `normal`로 되돌리고 selection을 제거한다.
- 첫 slice에서 `v`를 visual mode 중 다시 누르면 selection을 제거하고 normal mode로 돌아간다.

## Engine Contract

`internal/vimengine.State`는 selection을 표현한다.

```go
type SelectionKind string

const SelectionCharwise SelectionKind = "charwise"

type Selection struct {
    Active bool
    Kind   SelectionKind
    Anchor Cursor
    Head   Cursor
    Start  Cursor
    End    Cursor
}
```

- `ModeVisual`은 `"visual"`이다.
- `v` in normal mode: `ModeVisual`로 진입하고 `Anchor`, `Head`, `Start`, `End`를 현재 cursor로 초기화한다.
- normal motion in visual mode: cursor와 `Head`를 함께 이동하고 `Start`/`End`를 다시 계산한다.
- `esc` in visual mode: `ModeNormal`로 돌아가며 selection을 clear한다.
- `v` in visual mode: `ModeNormal`로 돌아가며 selection을 clear한다.
- visual mode 중 unsupported key는 state를 변경하지 않는다.

## TUI Contract

`internal/tuiadapter.ViewModel`은 selection을 그대로 전달한다.

- 상태 줄은 기존 `Mode: visual`을 표시한다.
- selection이 active이면 별도 줄에 `Selection: charwise <start.row>,<start.col> -> <end.row>,<end.col>`를 표시한다.
- buffer rendering은 selected non-cursor cell을 `{x}`로 표시한다.
- cursor cell은 기존 `[x]` 표시를 유지한다.
- empty line selection은 첫 slice에서 만들지 않는다.

## E2E Contract

`internal/e2estate.State`와 runner `assert.app_state`는 같은 shape의 `selection`을 지원한다.

```yaml
assert:
  app_state:
    mode: "visual"
    selection:
      active: true
      kind: "charwise"
      anchor: {row: 0, col: 1}
      head: {row: 0, col: 3}
      start: {row: 0, col: 1}
      end: {row: 0, col: 3}
```

- runner는 `selection.active`, `kind`, `anchor`, `head`, `start`, `end`를 각각 검증할 수 있어야 한다.
- visual mode E2E는 screen text만으로 selection을 통과시키지 않는다.
- content `e2e_assertions.selection`도 같은 의미를 사용한다.

## 첫 구현 Slice 제안

- `E2E-007`: `e2estate`, runner assertion, content assertion schema에 `selection` 추가
- `VIM-027-TUI-003`: `v` charwise selection foundation, motion update, `esc` reset, 최소 TUI 표시
- 후속: visual selection에 `d`/`y` 적용
