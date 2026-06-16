# ExecPlan: Review Findings Quality Hardening

Slice-ID: QUALITY-HARDENING-001
Created: 2026-06-12
Status: completed
Scope-Mode: normal
Allowed-Paths:
- `docs/exec-plans/active/quality-hardening-001-review-findings.md`
- `internal/progress/**`
- `internal/app/**`
- `internal/playable/**`
- `internal/content/**`
- `cmd/e2e-runner/**`

## 목표

전체 코드 품질 리뷰에서 확인한 데이터 신뢰성, 콘텐츠 검증, QA 시나리오 검증의 조용한 실패 지점을 닫는다. 진행도 저장 포맷과 콘텐츠 ID/파일명은 유지하고, 기존 Homebrew release plan의 embedded content fallback 동작을 보존한다.

## 범위

- 포함: progress load/save 오류 표면화, playable 저장 실패 UX, YAML strict decode, approved playlist beat 검증, tutorial playlist 길이 제한 범위 정리
- 제외: progress JSON 필드 변경, content schema/ID/파일명 변경, 새 Go 의존성 추가, Homebrew tap/formula 변경, 신규 게임 콘텐츠 추가

## 수용 기준

- progress 파일이 없을 때는 기존처럼 fresh progress로 시작한다.
- progress main 파일이 손상됐고 backup이 유효하면 backup을 로드한다.
- progress main 파일과 backup을 모두 신뢰할 수 없으면 fresh progress로 조용히 대체하지 않고 오류를 반환한다.
- 앱 시작 시 progress 로드 오류가 발생하면 플레이 가능한 화면에서 오류가 사용자에게 보인다.
- 미션 성공 후 progress 저장이 실패하면 완료 상태를 저장 완료로 표시하지 않고, 다음 미션으로 넘어가지 않으며, 사용자가 오류를 볼 수 있다.
- 콘텐츠 YAML 로더는 알 수 없는 필드를 거부하며 디스크 로딩과 embedded FS 로딩 모두 같은 규칙을 사용한다.
- E2E scenario YAML 로더는 알 수 없는 필드를 거부한다.
- approved/implemented playable playlist의 beat는 playable exercise와 playable scenario를 참조해야 하며, 누락 또는 비플레이어블 참조가 조용히 skip되지 않는다.
- 8-beat 제한은 tutorial category의 approved playlist에만 적용되고, incident category의 approved playlist에는 적용하지 않는다.
- 새 의존성을 추가하지 않는다.

## 검증 계획

- `gofmt` 대상 Go 파일
- `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/progress ./internal/app ./internal/playable ./internal/content ./cmd/e2e-runner`
- `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go vet ./...`
- `make e2e-smoke`
- 가능하면 `make e2e-playable`
- `git diff` 확인

## 의사결정 로그

- 2026-06-12: 사용자가 코드 품질 리뷰 결과를 하나씩 성공 기준을 설정해 진행하도록 승인했다.
- 2026-06-12: 진행도 저장 포맷과 콘텐츠 schema는 변경하지 않고, 오류 처리와 검증 규칙만 강화한다.
- 2026-06-12: 활성 Homebrew release plan의 embedded content fallback 변경은 보존한다.
- 2026-06-12: progress/app/playable/content/e2e-runner 대상 패키지 테스트, 전체 `go test ./...`, `go vet ./...`, `make e2e-smoke`, `make e2e-playable`이 통과했다.

## 미해결 질문

- 없음.
