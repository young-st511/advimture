# ExecPlan: FIRST-RUN-POLISH-001 Cue Density and Viewport Evidence

Status: completed
Started: 2026-05-30
Completed: 2026-06-01

## Goal

새 engine/content/schema 없이 첫 실행 UX를 다듬는다. Tutorial running cue의 밀도를 낮추고, success/failure 화면의 review/daily 상단 노출을 조용하게 만들며, 작은 viewport에서도 modal action line이 보이는지 E2E evidence로 확인한다.

## Context

- `PLAYTEST-GATE-001`에서 P0/P1 blocker는 확인되지 않았다.
- P2 후보는 running cue density, review/daily line length, quote 주변 cursor marker 혼동, mid tutorial evidence 편차, viewport smoke 부재다.
- 이번 slice는 출시 전 polish이며 새 Vim command, 새 content, 저장 포맷 변경을 하지 않았다.

## Scope

포함:

- Tutorial running cue에서 `기억할 명령`과 `Coach` 중복을 줄인다.
- Incident 기본 running cue는 정답 key를 노출하지 않는 기존 원칙을 유지한다.
- Success/failure modal의 app_state focus panel 의미는 유지한다.
- Succeeded/failure 화면에서 상단 review/daily line이 modal보다 과하게 시선을 빼앗지 않도록 한다.
- 80x24 success/failure viewport smoke fixture를 추가한다.
- 대표 mid tutorial full route에 final/timeline/app_state evidence 저장을 보강한다.
- 관련 docs/spec/roadmap을 갱신한다.

제외:

- 새 Vim engine 기능
- 새 exercise/content/playpack
- content schema 변경
- progress 저장 포맷 변경
- renderer 대개편
- terminal cell-grid parser 도입
- cursor marker 렌더링 자체 변경

## Acceptance Criteria

- [x] Tutorial running 화면은 command memory와 coach key를 중복 노출하지 않는다.
- [x] Tutorial running `app_state.ui.focus_panel.lines`는 필요한 키/명령 정보를 유지한다.
- [x] Incident running 기본 화면은 command memory/정답 key sequence를 계속 숨긴다.
- [x] Hint/failure 후 incident `참고 명령` 공개는 유지된다.
- [x] Succeeded/failed 화면은 상단 review/daily line을 과하게 노출하지 않고, modal의 `잔류 리스크`/`다음 출격`/`Retry`/`Next` action은 유지한다.
- [x] 80x24 success modal E2E가 action line과 `ui.focus_panel`을 검증한다.
- [x] 80x24 failure modal E2E가 retry action line과 `ui.focus_panel`을 검증한다.
- [x] 대표 mid tutorial route는 `screen_final.txt`, `screen_timeline.txt`, `app_state.json` evidence를 남긴다.
- [x] `go test ./...`, focused E2E, `make release-check`, `git diff --check`가 통과한다.

## Implementation Notes

- Tutorial running 상태에서는 `기억할 명령: ...`이 이미 같은 required key를 설명하면 `Coach: 훈련 키 ...`를 중복 표시하지 않는다.
- Incident running 기본 상태는 command memory를 계속 숨기며, hint/failure 이후 `참고 명령` 공개 규칙을 유지한다.
- Success/failure floating modal이 있을 때는 상단 detailed review/daily line을 숨기고, modal 내부의 residual risk/next action을 primary 안내로 둔다.
- `make e2e-playable`에 80x24 success/failure viewport smoke를 추가했다.
- open-line, repeat, search, visual selection, visual line, char-find full route는 final/timeline/app_state evidence를 남긴다.

## Verification

- Red: `go test ./internal/playable ./internal/playableview`에서 기존 duplicate coach와 floating modal 상단 detailed review line 기대를 실패시켰다.
- Green:
  - `go test ./internal/playable ./internal/playableview`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_hint_affordance.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_success_modal_80x24.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_failure_modal_80x24.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_full.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- Regression:
  - `go test ./...`
  - `make release-check`
  - `git diff --check`

## Next

다음 권장 slice는 `RELEASE-CANDIDATE-001`이다. 새 engine/content를 열기보다 release note, known limitations, 최종 evidence bundle, 태그 후보 기준을 먼저 정리한다.
