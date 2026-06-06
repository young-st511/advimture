# CONTENT-ARC-POLISH-001 — First Dispatch Arc Polish

Status: completed
Created: 2026-06-06

## 영향 도메인

- Gameplay content: 기존 scenario copy만 조정한다. `content` schema, exercise target state, optimal keys, constraints는 변경하지 않는다.
- Roadmap/docs: first dispatch arc의 완료 상태와 evidence 기준을 문서에 반영한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`: FocusPanel과 scenario feedback은 Vim 조작/판단 이유를 먼저 설명해야 한다.
- `docs/gameplay/domain-contract.md`: content schema ID/파일명 변경 금지.
- 사용자 제공 중기 Plan: incident 001~003을 "원인 신호 추적 -> 구조 재동기화 -> 오염 구간 격리" 첫 복구 arc로 읽히게 한다.

## Step 1: Incident 001~003 Copy Polish

- 목표: incident 001~003이 첫 대표 복구 아크로 이어지도록 title/briefing/success/failure copy를 다듬는다.
- 변경 파일:
  - `content/scenarios/incident-run.yaml`
  - `content/scenarios/incident-structure.yaml`
  - `content/scenarios/incident-visual.yaml`
- 충족 기준:
  - 각 briefing은 상황 1문장 + Vim 조작/판단 목표 1문장을 넘지 않는다.
  - success/failure copy는 세계관 감탄보다 Vim 조작/판단 이유를 먼저 말한다.
  - 새 YAML exercise를 추가하지 않는다.
- Boundaries 주의:
  - `target_state`, `optimal_keys`, `constraints`, schema 필드, scenario/exercise ID를 변경하지 않는다.
- 상세 작업:
  - [x] 001 copy를 원인 신호 추적 arc로 정리한다.
  - [x] 002 copy를 구조 재동기화 arc로 정리한다.
  - [x] 003 copy를 오염 구간 격리 arc로 정리한다.
  - [x] focused E2E의 exact copy assertion을 최신화한다.

## Step 2: Docs/Evidence Sync

- 목표: first dispatch arc의 현재 품질 기준과 evidence 묶음을 문서에 반영한다.
- 변경 파일:
  - `docs/roadmap/CONTENT_QUALITY_PLAN_001.md`
  - `docs/gameplay/world-frame.md`
- 충족 기준:
  - incident 001~003 final/timeline evidence를 사람이 읽었을 때 첫 복구 작전 arc로 설명 가능해야 한다.
  - 신규 콘텐츠 추가보다 arc framing 보강이 이번 scope임을 명시한다.
- 상세 작업:
  - [x] `CONTENT_QUALITY_PLAN_001.md`에 완료 상태와 evidence 기준을 반영한다.
  - [x] `world-frame.md`에 First Dispatch arc 언어를 반영한다.

## 검증 계획

- `go test ./internal/content ./internal/playable ./internal/playableview`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_003_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
- `git diff --check`

## 완료 기록

- Completed on 2026-06-06.
- `content/scenarios/incident-run.yaml`, `incident-structure.yaml`, `incident-visual.yaml` copy를 Vim 조작/판단 이유 중심으로 보강했다.
- `playable_incident_001_full`, `playable_incident_002_full`, `playable_incident_003_full`, `playable_ftue_first_five_route` 통과.
- `go test ./internal/content ./internal/playable ./internal/playableview` 통과.
