# ExecPlan: PRE-RC-HARDENING-001 First Release Hardening

Status: completed
Date: 2026-06-02

## Goal

Advimture를 첫 공개 후보로 묶기 전에, 현재 playable loop의 약한 지점을 한 번 더 보강한다. 목표는 새 콘텐츠 확장이 아니라 처음 플레이하는 사람이 덜 헷갈리고, 터미널 UI가 더 안정적으로 읽히며, 검증 evidence가 더 신뢰 가능한 상태를 만드는 것이다.

## Context

- `FIRST-RUN-POLISH-001`은 완료됐고, 다음 후보는 `RELEASE-CANDIDATE-001`이었다.
- 사용자는 release candidate로 바로 닫기보다 pre-RC hardening을 한 번 더 원했다.
- 현재 foundation loop는 새 engine/content 없이도 첫 공개 후보에 가까우나, 첫 5분 route, incident 실패/힌트, success/failure modal, review/daily copy, viewport evidence를 실제 evidence 기준으로 한 번 더 훑었다.

## Scope

포함:

- 첫 5분 route와 대표 incident route의 `screen_final.txt`, `screen_timeline.txt`, `app_state.json` evidence spot review.
- 실패/힌트/재시도, success/failure modal, review/daily copy의 혼동 지점 보강.
- 기존 route의 문구/힌트/검증 보강.
- 80x24 외 대표 viewport smoke로 incident hint wrapping을 80x30에서 검증.
- known limitations, roadmap, verification docs를 현재 상태와 맞춤.
- 다음 goal로 `RELEASE-CANDIDATE-001`을 열 수 있다는 판단 문서화.

제외:

- 새 Vim command 구현
- 신규 대형 tutorial/incident 제작
- 새 content schema
- progress 저장 포맷 변경
- `go.mod`/`go.sum` 변경
- renderer 대개편
- 색상/theme 대공사

## Acceptance Criteria

- [x] 첫 5분 route evidence를 사람이 읽고 P0/P1 UX blocker가 없음을 기록한다.
- [x] 대표 incident route evidence를 사람이 읽고 P0/P1 UX blocker가 없음을 기록한다.
- [x] 실패/힌트/재시도 안내가 초심자에게 다음 행동을 충분히 알려준다.
- [x] Success/failure modal은 다음 행동과 잔류 리스크를 terminal viewport에서 안정적으로 보여준다.
- [x] Review/daily copy가 처음 보는 사람에게 과하게 메타적으로 보이지 않도록 보조 정보로 유지된다.
- [x] 80x24 외 viewport smoke가 incident hint wrapping과 `ui.focus_panel`을 검증한다.
- [x] known limitations와 roadmap 문서가 현재 hardening 결과와 맞는다.
- [x] 새 engine/content/schema/progress 저장 포맷/go.mod 변경이 없다.
- [x] `go test ./...`, `make release-check`, `git diff --check`가 통과한다.

## Implementation Notes

- Running mission cue를 terminal width 기준으로 여러 줄에 감싸도록 변경했다. 긴 incident hint가 한 줄에서 잘리지 않고, `참고 명령`, `Hint`, `?: hint q: quit` 의미가 유지된다.
- `playable_incident_hint_affordance`를 80x30 viewport smoke로 낮추고 final/timeline/app_state evidence를 저장하도록 보강했다.
- 긴 incident full route 001/002/003/004/006/007은 `save_app_state: true`를 추가해 verification spec과 실제 evidence가 맞도록 했다.
- `docs/roadmap/PRE_RC_HARDENING_2026-06-02.md`에 evidence spot review와 P0/P1 blocker 없음, P2 조치 내용을 기록했다.

## Verification

- Red:
  - `go test ./internal/playableview` 실패: long incident hint cue가 한 줄 146 display width로 렌더링됨.
- Focused:
  - `go test ./internal/playable ./internal/playableview`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_hint_affordance.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- Regression:
  - `go test ./...`
  - `make release-check`
  - `git diff --check`

## Result

P0/P1 UX blocker는 없다. 발견된 P2는 incident hint cue truncation과 long incident app_state evidence gap이었고 둘 다 닫았다. 다음 권장 slice는 `RELEASE-CANDIDATE-001`이다.
