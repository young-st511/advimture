# ExecPlan: 수정 후 같은 기준 PM 재실험

Slice-ID: PM-RETEST-001
Created: 2026-10-29
Status: completed (expert playtest; no external participant experiment)

## 목표와 범위

사용자의 "다시 같은 기준으로 실험" 요청에 따라 ad56412 버전을 실제 PTY에서 다시 사용한다. 이전 PM-01~12 해결 여부, 첫 복구/학습 목록/힌트/복습 신규 경로의 회귀를 점검한다. 전문가 직접 사용과 자동 재현을 실제 외부 참가자 학습/재방문 실험으로 보고하지 않는다.

이번 범위는 감사/재현/우선순위 재평가다. 제품 수정은 수행하지 않으며 독립 HOME과 artifacts/e2e/pm-retest/만 사용한다. PR #2 기반 별도 감사 branch에서 보고서를 커밋하고 stacked PR로 전달한다.

## 완료 기준

- [x] 실제 바이너리 직접 사용: 첫 복구/실패/오타 수정/경로 선택/복습.
- [x] 80x24와 110x34 비교, final/timeline/app_state/key trace evidence.
- [x] 이전 12개 이슈: 6개 해결 확인, 5개 부분 해결, 1개 경로 구현 확인/사용자 효과 미측정.
- [x] 기존 52개 회귀와 신규 탐색 28개를 구분했다. 관찰 재현 pass와 별도로 품질 기대 8개의 미충족을 기록했다.
- [x] 재현 단계/영향/근거/우선순위와 다음 권고를 PM_RETEST_REVIEW_2026-10-29.md에 기록했다.
- [x] 변경은 감사 문서/검사 항목뿐이다. 문서 commit 후 PR로 전달한다.

## 기준

- 이전 감사: docs/roadmap/PM_PLAYTEST_REVIEW_2026-09-29.md.
- 승인 동작: docs/gameplay/spec.md 및 completed/pm-quality-recovery-001.md.
- 같은 최소 viewport 및 PTY QA 계약을 사용한다.

## 결과

- 기준 commit ad56412의 실제 직접 PTY 3세션 및 80x24/110x34 비교 완료.
- 기존 회귀 52/52 통과, 추가 탐색 28개 관찰 assertion 충족. 별도 품질 기대 8개는 미충족이며 보고서에서 구현/정책/QA 문제를 구별했다.
- `go test -count=1 ./...` 통과. 제품 코드, content, 의존성, progress 저장 포맷은 변경하지 않았다.
- 독립 HOME/evidence: artifacts/e2e/pm-retest/. 실제 사용자 학습/재방문과 플레이 시간은 미측정이다.
