# Post Visual Review — 2026-05-22

## 목적

`PLAYPACK-010` 완료 직후, 다음 중기 플랜을 새 command 확장으로 바로 열지 않고 Architecture, Learning/Game UX, Scenario/World 관점에서 재점검했다.

## 입력 상태

- `VISUAL-GAP-002`, `E2E-007`, `VIM-027-TUI-003` 완료: charwise visual selection state/render/app_state assertion 기반이 있다.
- `VISUAL-OP-001`, `VIM-028` 완료: 같은 줄 charwise visual `d/y` engine이 있다.
- `PLAYPACK-010` 완료: `tutorial-92-visual-selection` 3문항과 full E2E가 있다.
- progress 저장 포맷은 여전히 v1 `Missions` map 중심이며, schema 변경은 승인 대상이다.

## 종합 결론

다음 2~3주는 새 Vim 기능을 곧장 늘리기보다 **post-visual hardening + applied mastery**로 간다.

구조 관점에서는 engine/runtime/content/playable/e2e/progress 경계가 대체로 건강하다. 다만 selection이 들어오면서 `internal/vimengine` 단일 파일의 책임이 커졌고, content replay가 selection assertion mismatch를 아직 충분히 검증하지 못한다.

학습 관점에서는 커리큘럼이 실무 Vim 능력으로 이어지는 축을 갖췄다. 하지만 tutorial 수가 넓어졌으므로 이제는 “새 command를 아는가”보다 “상황에 맞는 도구를 고르는가”를 훈련해야 한다.

시나리오 관점에서는 `원격 시설 복구국 / Runbook Dispatch` 프레임이 충분하다. 세계관을 더 키우기보다 incident beat가 하나의 runbook 조치처럼 이어지게 만들고, visual selection을 실제 복구 작전에서 쓰게 하는 편이 좋다.

## 리뷰별 핵심 신호

### Architecture

- 좋은 점: engine/runtime/content/TUI/E2E/progress adapter 경계가 유지되고 있다.
- 리스크: visual selection, operator range, undo/register 로직이 `internal/vimengine`에 계속 쌓인다.
- 리스크: `e2e_assertions.selection`은 존재하지만 content replay 단계의 mismatch 검증은 보강 여지가 있다.
- 권장: selection replay hardening과 vimengine 내부 파일 분리를 먼저 닫는다.

### Learning / Game UX

- 좋은 점: movement, edit, operator, yank/put, text object, search/substitute, quote object, visual까지 실무 조합이 넓어졌다.
- 리스크: 바로 새 command를 추가하면 넓은 맛보기로 흩어질 수 있다.
- 권장: command choice drill, visual applied incident, runbook continuity pass를 통해 적용력을 강화한다.

### Scenario / World

- 좋은 점: Runbook Dispatch 프레임은 visual selection에도 잘 맞는다.
- 리스크: incident가 아직 독립 exercise 묶음처럼 보일 수 있다.
- 권장: `릴레이 기지 003: 오염 구간 격리`처럼 visual selection을 실제 복구 동사로 쓰는 incident를 만든다.

## 결정

새 중기 플랜 이름은 **Post-Visual Applied Mastery and Hardening**으로 한다.

우선순위:

1. 검증 구멍을 먼저 닫는다.
2. 엔진 내부 분리로 다음 visual 확장 비용을 낮춘다.
3. 같은 줄 charwise visual의 invariant를 굳힌다.
4. 새 command가 아니라 visual을 적용하는 incident를 만든다.
5. tutorial과 incident 사이의 runbook 흐름을 강화한다.
6. linewise visual은 구현이 아니라 gap planning으로만 연다.

## 승인 게이트

- progress schema v2, mastery, spaced review, daily run, attempt history 저장은 사용자 승인 전까지 구현하지 않는다.
- `V`, multi-line visual, visual block, count/register prefix는 최소 ExecPlan 대상이다.
- E2E는 계속 temp HOME 또는 progress fixture를 사용하며 실제 `~/.advimture`를 쓰지 않는다.
- 새 dependency는 ExecPlan과 사용자 승인 없이는 추가하지 않는다.
