# ExecPlan: Progress adapter

Slice-ID: VIM-009
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-009-progress-adapter.md
- internal/progressadapter/**

## 목표

Scenario Runtime의 완료 결과를 기존 progress 저장 모델에 반영 가능한 completion record로 변환하는 Progress Adapter를 만든다. 이번 slice는 파일 저장이 아니라 순수 변환과 in-memory apply만 다룬다.

## 범위

- 포함: `internal/progressadapter` 패키지 생성
- 포함: scenario state + elapsed time을 mission completion으로 변환
- 포함: mission completion을 progress model copy에 적용
- 제외: `~/.advimture/progress.json` 파일 쓰기
- 제외: save schema 변경
- 제외: Bubble Tea 연결

## 검증 계획

- `go test ./internal/progressadapter`
- `go test ./...`

## 승인 체크

- [x] completed scenario state가 mission completion으로 변환된다.
- [x] incomplete state는 completion으로 변환되지 않는다.
- [x] progress apply는 입력 progress를 직접 mutate하지 않는다.
- [x] filesystem에 의존하지 않는다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: 기존 `internal/progress` save schema는 변경하지 않았다. 이번 slice는 scenario 결과를 기존 mission progress에 반영 가능한 형태로 바꾸는 adapter만 만든다.
- 2026-05-18: 파일 저장은 제외했다. 실제 `~/.advimture/progress.json` 쓰기는 사용자 데이터에 영향을 주므로 별도 승인/검증 루프로 다룬다.
