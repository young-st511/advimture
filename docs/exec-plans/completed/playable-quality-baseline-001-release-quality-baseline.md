# ExecPlan: PLAYABLE-QUALITY-BASELINE-001 Release-Quality Playable Baseline

Status: completed
Started: 2026-06-02
Completed: 2026-06-02

## Goal

Advimture를 바로 출시하지는 않는다. 대신 현재 playable loop를 "출시 가능한 품질 기준"에 가깝게 만든다.

필요하면 기존 루프의 구조 변경과 리팩터링까지 허용한다. 다만 progress 저장 포맷, 새 의존성, content schema 변경, 대형 신규 시스템처럼 되돌리기 어려운 변경은 별도 checkpoint와 사용자 확인 후 진행한다.

## Why Now

`PRE-RC-HARDENING-001`은 첫 공개 후보 직전 polish처럼 동작했다. 하지만 실제 목표는 release candidate를 포장하는 것이 아니라, 계속 개발하더라도 제품 기준을 "출시 가능한 수준"에 맞추는 것이다.

따라서 다음 작업은 release note나 tag 준비가 아니라, 세계관, UX/UI, 콘텐츠 목표, 엔진/모듈화 기준을 하나의 release-quality baseline으로 세우고 현재 상태를 그 기준에 맞춰 개선하는 것이다.

## Scope

포함:

- 세계관 정합성 검증 및 개선
  - 실제 몰입감 높은 게임 세계관의 기준을 정의한다.
  - 현재 `Runbook Dispatch` 프레임, scenario tone, incident copy가 그 기준을 만족하는지 점검한다.
  - 필요한 경우 문서와 기존 문구를 개선한다.
- UX/UI 개선
  - 출시 가능한 TUI UX 기준을 정의한다.
  - first-run, tutorial, incident, success/failure, review queue, viewport evidence를 기준으로 개선한다.
- 콘텐츠 품질 목표 설정
  - "출시할 만한 콘텐츠 양과 구성" 기준을 정의한다.
  - 현재 콘텐츠를 점검하고 부족분은 기획까지만 한다.
  - 이 goal 안에서는 신규 대형 콘텐츠를 구현하지 않는다.
- 모듈화/엔진 검증
  - 현재 모듈 분리와 Vim engine이 위 목표를 감당하기 충분한지 확인한다.
  - 필요하면 리팩터링 후보와 engine opening 조건을 문서화한다.
- 검증 evidence 갱신
  - representative E2E evidence를 사람이 읽어 UX/UI/world 기준으로 spot review한다.
  - 필요한 viewport smoke나 focused test를 추가할 수 있다.

제외:

- 바로 출시하거나 tag를 찍는 작업
- 신규 대형 tutorial/incident 구현
- progress 저장 포맷 변경
- 사용자 확인 없는 새 의존성 추가
- 사용자 확인 없는 content schema 변경
- 목적 없는 renderer 대개편

## Release-Quality Axes

이 goal은 네 축을 동시에 본다.

| 축 | 기준 질문 | 합격 기준 |
|----|-----------|-----------|
| World | "Vim 문제 풀이"가 아니라 하나의 복구 작전처럼 느껴지는가? | 플레이어 역할, 사건 단위, 명사, 성공/실패 피드백이 `Runbook Dispatch` 프레임 안에서 일관된다. |
| UX/UI | 처음 보는 사람이 무엇을 해야 하는지 즉시 아는가? | 현재 목표, 입력 상태, hint/retry/next action, residual risk가 terminal viewport에서 잘리지 않는다. |
| Content | 출시 가능한 첫 플레이 분량과 반복 동기가 있는가? | tutorial-first route, applied incident route, review loop가 충분한 제품 skeleton을 이루며 부족한 콘텐츠는 기획으로 정리된다. |
| Engine/Modules | 목표 품질을 올리기 위해 무리한 코드 변경이 필요한가? | Vim engine, runtime, content loader, playable model, renderer, E2E가 분리되어 있고 다음 개선 slice를 좁게 열 수 있다. |

## Acceptance Criteria

- [x] `docs/roadmap/PROGRAM.md`, `MIDTERM_TODO.md`, `FORWARD_PLAN.md`가 release-candidate prep이 아니라 release-quality baseline을 다음 active 방향으로 가리킨다.
- [x] 세계관 release-quality 기준과 현재 판정이 문서화되어 있다.
- [x] UX/UI release-quality 기준과 현재 판정이 문서화되어 있다.
- [x] 콘텐츠 출시 품질 기준, 현재 콘텐츠 gap, 구현하지 않을 기획 backlog가 문서화되어 있다.
- [x] 엔진/모듈화가 위 목표를 감당하기 충분한지 판정하고, 부족분은 risk/backlog로 분리한다.
- [x] 대표 evidence(`screen_final.txt`, `screen_timeline.txt`, `app_state.json`) 기준으로 P0/P1 blocker 여부를 기록한다.
- [x] 필요한 최소 개선을 구현하거나, 큰 구조 변경이 필요하면 별도 ExecPlan/checkpoint로 분리한다.
- [x] 변경 후 `go test ./...`, 필요한 focused E2E, `git diff --check`를 통과한다.

## Working Rules

- 새 content implementation은 하지 않는다. 콘텐츠 확장은 목표/구성/beat 기획까지만 한다.
- 큰 구조 변경은 허용하지만, blast radius가 크면 이 ExecPlan 안에 바로 섞지 않고 별도 slice로 쪼갠다.
- "큰 수정" 자체를 막지 않는다. 목표에 더 가까워지는 구조 변경은 열되, evidence와 수용 기준 없이 제품 계약을 흔드는 변경을 막는다.
- progress schema, dependency, content schema는 사용자 확인 전까지 건드리지 않는다.
- UX 개선은 "현재 Vim 조작 목표가 더 잘 보이는가?"를 먼저 통과해야 한다.
- 세계관 개선은 command 학습 목표를 가리면 실패로 본다.

## First Slice

1. Roadmap을 `RELEASE-CANDIDATE-001`에서 `PLAYABLE-QUALITY-BASELINE-001` 중심으로 정정한다.
2. World/UX/Content/Engine 기준과 현재 판정을 하나의 review 문서로 작성한다.
3. 기준선에서 바로 보이는 P0/P1/P2를 나눠 다음 구현 후보를 고른다.

Status: steps 1~3 completed as documentation baseline. Next step is a focused small improvement from the P2 list or a separate checkpoint for any larger UI structure change.

Focused improvement completed:

- Mode-specific `FocusPanel` cues no longer use English `Keys: ...` copy. Insert/search/command/visual mode now use Korean action labels while preserving actual key names such as `esc`, `enter`, and `normal`.
- Mode-specific `FocusPanel` titles no longer use `* CHANNEL` copy. They now read `입력 모드`, `검색 모드`, `명령 모드`, and `선택 모드`.
- Floating modal helper labels no longer use `Mistake`, `Learned`, `Result`, or failure `Next`. They now render as `실수`, `배운 점`, `기록`, and `힌트` while preserving stable action lines such as `Retry: r or enter` and `Next: enter`.
- Header now includes the current track label. Tutorial screens render `Tutorial`, and incident screens render `Runbook Dispatch` before the playlist title so the first viewport communicates the world frame.
- Affected command-mode E2E assertions were updated to verify the Korean command cue.

Content planning completed:

- `docs/roadmap/CONTENT_QUALITY_PLAN_001.md` now defines the minimum release-quality content shape: first control, first efficiency, first edit, core toolbelt, first applied run, judgment drill, and review loop.
- Current content was mapped to that shape without adding new YAML content.
- Incident 001~003 are framed as the first dispatch arc: hotfix recovery -> structure resync -> contamination isolation.
- Incident 005 is framed as the judgment drill set, and future content work is split into planning/review candidates rather than implementation.

Module review completed:

- `docs/roadmap/MODULE_QUALITY_REVIEW_2026-06-02.md` documents current package boundaries, import flow, sufficient module surfaces, risks, and engine-opening conditions.
- Large structural changes are allowed when they move the product toward release-quality, but they are routed through immediate/checkpoint/user-approval paths by blast radius.
- Current verdict: no new Vim engine is required for the release-quality baseline.
- Next structural candidates are `internal/playable` orchestration split or pre-start modal input boundary, both gated by evidence and separate ExecPlan.

Completion audit completed:

- `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md` maps each user-facing requirement to current-state evidence.
- Residual work is classified as follow-up backlog rather than P0/P1 blocker.

## Verification Plan

- Documentation:
  - `git diff --check`
  - roadmap/spec 문서 수동 diff review
- Code-bearing slice 이후:
  - 변경 패키지 `go test`
  - `go test ./...`
  - 필요한 focused E2E
  - `make release-check`는 release gate나 shared behavior 변경 시 실행

## Verification Log

- 2026-06-02 baseline documentation slice:
  - `git diff --check`: pass
  - `go test ./...`: pass
- 2026-06-02 mode cue UX slice:
  - `go test ./internal/playable ./internal/playableview`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_quit.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_substitute_current_line.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_full_first_five_minute.yaml`: pass
- 2026-06-02 mode title UX slice:
  - `go test ./internal/playable ./internal/playableview`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_mismatch_feedback.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_quit.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_substitute_current_line.yaml`: pass
- 2026-06-02 floating modal label UX slice:
  - `go test ./internal/playableview`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_success_modal_80x24.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_failure_modal_80x24.yaml`: pass
- 2026-06-02 content quality planning slice:
  - `git diff --check`: pass
  - `go test ./...`: pass
  - No `content/` YAML files changed.
- 2026-06-02 module quality review slice:
  - `go list ./internal/...`: reviewed
  - `go list -f '{{.ImportPath}}: {{join .Imports ", "}}' ./internal/...`: reviewed
  - `git diff --check`: pass
  - `go test ./...`: pass
  - `git status --short -- go.mod go.sum internal/progress content`: no output
- 2026-06-02 header track/world-frame slice:
  - `go test ./internal/playableview`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`: pass
  - `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`: pass
- 2026-06-02 completion audit slice:
  - `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`: completed
  - `make release-check`: pass
  - `git diff --check`: pass
  - `git status --short -- go.mod go.sum internal/progress content`: no output
