# HELP-AFFORDANCE-001 — Hint, Retry, and Quit Affordance

Slice-ID: HELP-AFFORDANCE-001
Created: 2026-05-26
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/tui-ui-qa-contract.md
- docs/exec-plans/active/help-affordance-001-hint-retry-quit.md
- docs/exec-plans/completed/hud-density-001-mission-hud-priority.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 목표

화면에 노출되는 `?: hint`, `Retry: r or enter`, `q: quit` 안내가 실제 입력 처리와 일치하는지 검증하고, 필요한 경우 문구와 focused E2E를 보강한다.

## 해결할 문제

RedTeam은 대표 루트에서 `?: hint  q: quit`가 보이지만 실제로 `?`를 누르는 경로가 부족하다고 지적했다. 이 slice는 “보이는 affordance가 실제로 동작한다”를 별도 루프로 증명한다.

## 범위

- 포함:
  - tutorial running 상태에서 `?` hint 요청 검증
  - incident running 상태에서 `?` hint 요청 검증
  - failed 상태에서 `r` retry와 `enter` retry 검증
  - running 상태에서 `q` quit 검증
  - mode-specific panel에서 잘못된 quit/hint 문구가 섞이지 않는지 확인
- 제외:
  - search backward `?` 구현
  - hint content schema 변경
  - modal layout 재설계
  - progress 저장 포맷 변경

## 수용 기준

- tutorial에서 `?`를 누르면 FocusPanel/AppState에 `Hint:`가 표시된다.
- incident에서 `?`를 누르면 정답 전체를 과하게 노출하지 않는 hint가 표시된다.
- failed 상태에서 `r`과 `enter` 모두 같은 exercise를 retry한다.
- running 상태에서 `q`는 앱을 종료하며 key trace에 남는다.
- command/search/insert/visual mode panel에는 실제 입력 처리와 맞지 않는 일반 `?: hint q: quit` 문구가 섞이지 않는다.
- focused E2E와 `make e2e-playable`을 통과한다.

## Step 1: Existing Behavior Review

- [ ] existing playable tests와 E2E coverage 확인
- [ ] 누락된 affordance route 식별

## Step 2: Tests and UX Copy

- [ ] tutorial hint E2E 또는 app_state assertion 추가/보강
- [ ] incident hint E2E 추가
- [ ] retry/quit route 보강
- [ ] 필요한 문구 조정

## Step 3: Verification

- [ ] focused E2E
- [ ] `go test ./...`
- [ ] `make e2e-playable`
- [ ] `git diff --check`
