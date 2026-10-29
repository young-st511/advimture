# ExecPlan: PM 개선안 순차 구현

Slice-ID: PM-QUALITY-RECOVERY-001
Created: 2026-10-29
Status: completed (implementation; human study pending)

## 승인과 목표

2026-10-29 사용자의 "우선순위를 따라서 하나씩 처리하고 테스트" 요청은 PM_PLAYTEST_REVIEW_2026-09-29.md의 후속 구현안 승인이다. 입력 신뢰성, 빠른 첫 복구, 공정한 응용 판정, 코칭/복습 순서로 구현한다. 외부 사용자 모집/실험은 준비 자료까지 만들고 실제 참여자 결과를 만들어내지 않는다.

## 경계

- progress 저장 포맷, content schema, 기존 exercise ID, 의존성은 유지한다.
- 기존 사용자 완료 기록을 지키고, 경로 선택은 세션 상태로 관리한다.
- PR #1의 감사 branch에서 별도 구현 branch를 만들고 stacked PR로 검토한다.
- 배포/태그/머지는 이번 요청에 포함하지 않는다.

## 승인된 수용 기준과 TODO

- [x] TRUST: Backspace가 insert/search/command에서 실제 삭제를 수행하고 특수키 이름을 문자로 삽입하지 않는다. 실패 모달의 Enter/r/?/q가 실패 당시 모드와 무관하게 동작한다.
- [x] TRUST: 묶음/개별 입력이 같은 결과를 내고 성공/실패 후 남은 입력이 다음 문제로 넘어가지 않는다. 첫 목표 c와 시작 힌트가 명확하며 80x24 modal footer가 보존된다.
- [x] FIRST-RECOVERY: 짧은 대표 경로와 학습/응용/복습 선택을 제공한다. 전체 학습 경로와 기존 재개를 보존하며 단일 동작 결과를 가리지 않는 피드백을 제공한다.
- [x] FAIRNESS: yyjp/Vyjp를 모두 인정하고 실제 복사를 동반하지 않는 위장 입력을 막는다. 대표 3개 문항의 결과/보존 범위를 지키는 대안 풀이를 인정하고 정답 trace 차이만으로 최고 평가를 막지 않는다.
- [x] COACH: 과제별 유한 힌트와 재열람, 도움받은 성공 표시, 학습/복습 구별, 큐 종료/다음 복습을 구현한다.
- [x] 검증: 단계별 Red/Green, 관련 PTY, 전체 Go/E2E, 우회 QA와 80x24/110x34 evidence.
- [x] 문서: spec/contract/roadmap/README와 사용자 검증 프로토콜 동기화. 단계별 커밋 후 PR로 전달한다.
- [x] 외부 실험 준비: 프로토콜 작성. 실제 참여자 모집, 48~72시간 유지/자발적 재방문 측정은 미실시 후속 작업이다.

## 검증

수정 전 테스트가 실패하는지 확인하고 관련 패키지부터 실행한다. E2E는 임시 HOME만 사용하며 artifacts/e2e/에 final/timeline/app_state를 남긴다. 마지막에는 go test ./..., make e2e-playable, git diff --check와 금지 경계 diff를 확인한다.

## 진행 기록

- 2026-10-29: 감사 PR #1은 OPEN이다. 구현은 codex/pm-quality-recovery에서 시작했다.
- 2026-10-29: TRUST/FIRST-RECOVERY 구현과 단위 Red/Green 완료. 전체 Go 테스트, 입력 복구/insert 실패 재시도/첫 복구 경로 PTY 통과. 터미널 OSC 응답을 입력과 분리하는 회귀 테스트를 추가했다. 첫 복구의 5분 목표는 실제 사용자 검증 전까지 미측정이다.
- 2026-10-29: `cb2b7bb`로 첫 단계 커밋. 이후 공정한 응용/코칭/복습/학습 목록 구현. 대표 3개 문항의 6개 정상 풀이, 실제 yank/put event 판정, 반복 힌트 비용과 복습 완료를 단위 검증했다.
- 2026-10-29: `make release-check E2E_GOCACHE=/tmp/advimture-validity-go-cache` 통과: 전체 Go, build, playable E2E 52개(신규 7개 포함). 기존 delimiter/재입력/over-scope RedTeam도 통과. `go vet ./...`, `git diff --check`, 경계 diff 확인 완료.

## 남은 실측

실제 사용자 6~8명의 첫 복구 시간/학습 전이/48~72시간 유지/자발 재방문은 아직 측정하지 않았다. `docs/verification/player-validation-protocol.md`를 실행할 참여자가 필요하다. 자동 구현 완료는 학습 효과 검증 완료를 뜻하지 않는다.
