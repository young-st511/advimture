# YANK-TEXT-001 — Yank / Put And Text Object Candidates

Slice-ID: YANK-TEXT-001
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/completed/yank-text-001-yank-put-text-object-candidates.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md

## 목표

Operator grammar 다음 playpack으로 `y/p`와 text object를 어떻게 다룰지 결정한다. 구현은 시작하지 않고, 엔진/콘텐츠 루프가 분리될 수 있도록 범위와 순서를 고정한다.

## 결정

다음 playable milestone은 `yank-put-basic`을 먼저 다룬다. text object는 같은 chapter의 다음 후보로 문서화하되, `y/p` register semantics가 안정된 뒤 별도 루프로 구현한다.

이유:

- `y/p`는 `d/c`와 같은 operator grammar 감각을 복사/재사용으로 확장한다.
- text object는 `operator + i/a + object`의 3단 pending 구조가 필요하다.
- 현재 엔진은 `PendingKey string` 기반이라 text object를 바로 얹으면 register, operator, object range가 한 루프에 섞인다.
- 먼저 unnamed register와 put semantics를 고정해야 `yiw`, `diw`, `ciw`가 자연스럽게 올라간다.

## 다음 루프 순서

1. `VIM-019`: yank/register foundation
   - `KeyY`, unnamed register state, `yy`, `yw`, `y$`
   - linewise/charwise register 구분
   - buffer mutation 없음
   - vimengine/tuiadapter unit
2. `VIM-020`: put basic
   - `p`, `P`
   - charwise put after/before cursor
   - linewise put below/above current line
   - undo/redo, runtime replay smoke
3. `PLAYPACK-004`: yank/put basic tutorial
   - 5-6문항
   - `yy`, `p`, `P`, `yw`, `y$` coverage
   - DevOps 생존 톤에서는 “정상 설정 줄 복제”, “구조 코드 재사용”, “검증된 값 붙여넣기”로 감싼다.
4. `TEXT-OBJECT-001`: text object gap planning
   - `iw` 우선
   - 후보: `diw`, `ciw`, `yiw`
   - 따옴표/괄호 object(`i"`, `i(`)는 `iw` 이후로 미룬다.

## Command Cluster 결정

### yank-put-basic

- status: approved
- engine_support: planned
- compatibility_tier: pedagogical
- commands: `["y", "yw", "y$", "yy", "p", "P"]`
- coverage_required: `["yw", "y$", "yy", "p", "P"]`
- prerequisite: `delete-with-motion`, `change-with-motion`
- 첫 content는 `yw/y$/yy`로 register에 담고 `p/P`로 재사용하는 감각을 목표로 한다.

### text-object-inner-word

- status: draft
- engine_support: planned
- compatibility_tier: pedagogical
- commands: `["iw", "diw", "ciw", "yiw"]`
- coverage_required 후보: `["diw", "ciw", "yiw"]`
- prerequisite: `yank-put-basic`, `delete-with-motion`, `change-with-motion`
- 첫 text object는 `iw`만 다룬다. quote/pair object는 아직 scope 밖이다.

## 제외

- named register
- numbered register
- clipboard/system register
- count prefix
- visual selection
- quote/pair text object
- multi-line text object

## 검증 결과

- `rg "yank-put-basic|text-object-inner-word|VIM-019|VIM-020|PLAYPACK-004|TEXT-OBJECT-001" docs`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 의사결정 로그

- 2026-05-19: text object를 바로 구현하지 않고, `y/p` register 기반을 먼저 만든다.
- 2026-05-19: `yy`와 linewise put은 content에서 실무감이 커서 첫 yank/put playpack에 포함한다.
