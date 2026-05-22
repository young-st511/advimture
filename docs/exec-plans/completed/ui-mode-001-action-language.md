# UI-MODE-001 — Tutorial and Incident Action Language

Slice-ID: UI-MODE-001
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: playable-ui-copy
Allowed-Paths:
- docs/exec-plans/active/ui-mode-001-action-language.md
- docs/exec-plans/completed/ui-mode-001-action-language.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/playable/model.go
- internal/playable/model_test.go
- test/e2e/playable_incident_001_full.yaml

## 목표

Tutorial과 incident의 action panel 언어를 분리한다. Tutorial은 새 command 학습을 위해 훈련 키를 직접 노출하고, incident는 처음부터 정답 key sequence를 과하게 노출하지 않고 판단 cue와 점진 힌트를 우선한다.

## 완료 내용

- `gameEntry`가 playlist category를 보존하도록 했다.
- running tutorial은 기존 `Coach: 훈련 키 ...`를 유지한다.
- running incident는 `판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요.`를 표시하고 `Coach: 훈련 키 ...`를 숨긴다.
- failed incident는 retry 맥락에서 `복구 힌트: 필요한 키 ...`를 표시한다.
- command/search/visual mode action panel은 기존 안내를 유지했다.
- progress 저장 포맷, content schema, E2E state shape는 변경하지 않았다.

## 검증 결과

- `go test ./internal/playable/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml`
- `make e2e-smoke`
- `git diff --check`

참고: sandbox의 기본 Go build cache 접근이 막혀 repo 내부 `artifacts/go-build-cache`를 `GOCACHE`로 지정해 검증했다.

## 후속 작업

다음 slice는 `UI-EVIDENCE-001`이다. 화면 변경을 반복할 때 LLM이 최종 화면과 누적 timeline을 더 쉽게 확인할 수 있도록 E2E evidence를 보강한다.
