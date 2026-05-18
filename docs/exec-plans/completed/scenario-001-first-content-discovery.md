# ExecPlan: First 5-minute scenario content discovery

Slice-ID: SCENARIO-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/README.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/gameplay/content-requirements.md
- docs/exec-plans/completed/scenario-001-first-content-discovery.md

## 목표

첫 5분 플레이 루프를 시나리오 워크플로우로 훑어보며 CONTENT-001 loader가 실제로 받아야 할 콘텐츠 구조를 발견한다.

이번 slice는 스토리를 먼저 확정하는 작업이 아니다. `Vim command -> Exercise -> Scenario` 순서를 유지하면서, scenario layer를 입혀볼 때 추가로 필요한 필드와 검증 요구사항을 찾는다.

## 범위

- 포함: 첫 5분 플레이 루프의 beat draft
- 포함: 각 beat가 요구하는 command/exercise/scenario 데이터 필드
- 포함: 현재 엔진으로 구현 가능한 콘텐츠와 다음 엔진 확장이 필요한 콘텐츠 분리
- 포함: CONTENT-001 loader 입력 요구사항 정리
- 제외: Go loader 구현
- 제외: 새 Vim command 구현
- 제외: 실제 mission unlock/progress schema 변경

## 설계 입력

- 제품 철학: `Vim command -> Exercise -> Scenario`
- 후보 command cluster: `survival-save-quit`, `normal-motion-basic`, `word-motion-basic`
- 현재 엔진 지원: `esc`, `h`, `j`, `k`, `l`
- 다음 후보 엔진 확장: command-line quit/save, `w/b/e`

## 검증 계획

- `rg -n "SCENARIO-001|content-requirements|first-5-minute" docs`
- `rg -n "status: draft|status: approved|status: implemented" docs/gameplay`
- `go test ./...`

## 승인 체크

- [x] 첫 5분 루프가 command-first 순서를 따른다.
- [x] 각 scenario beat가 하나 이상의 exercise를 참조한다.
- [x] 필요한 content loader 필드가 구체적으로 도출된다.
- [x] 현재 구현 가능 콘텐츠와 엔진 확장 필요 콘텐츠가 분리된다.
- [x] 다음 slice CONTENT-001의 입력이 충분하다.

## 의사결정 로그

- 2026-05-18: `normal-motion-basic`은 `engine_support: implemented`로 분리하고, `survival-save-quit`, `word-motion-basic`은 `planned`로 남겼다. CONTENT-001 loader가 planned 콘텐츠를 읽되 playable 후보에서 제외할 수 있어야 하기 때문이다.
- 2026-05-18: 첫 5분 루프는 `docs/gameplay/content-requirements.md`의 `first-5-minute` beat로 정리했다. 이 문서를 CONTENT-001의 입력으로 사용한다.

## 완료 결과

- `docs/gameplay/content-requirements.md` 추가
- command/exercise/scenario bank에 `engine_support` 필드 추가
- 첫 5분 후보 exercise와 scenario draft 보강
- CONTENT-001 acceptance draft 도출
