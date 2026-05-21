# TEXT-PAIR-GAP-001 — Quote/Pair Text Object Scope

Slice-ID: TEXT-PAIR-GAP-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-and-content-catalog
Allowed-Paths:
- docs/exec-plans/completed/text-pair-gap-001-quote-pair-scope.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/vim-curriculum-map.md
- content/command_clusters/text-object-quote-pair.yaml
- internal/content/loader_test.go

## 목표

quote/pair text object의 첫 구현 범위를 `ci"`, `di"`, `yi"` 중심으로 고정하고, engine 구현과 playpack 제작을 별도 slice로 분리한다.

## 범위

- 포함:
  - `text-object-quote-pair` command cluster approval packet
  - content command cluster YAML 추가
  - VIM-026/PLAYPACK-009 분리 계획
  - 제외 항목 명시
- 제외:
  - engine 구현
  - exercise/scenario/playlist YAML
  - nested pair, escaped quote, around object, count prefix, visual selection

## 결정

- 첫 engine slice는 double quote 내부 text object만 구현한다.
- 포함 command는 `di"`, `ci"`, `yi"`다.
- selection 범위는 같은 줄에서 커서를 포함하는 가장 가까운 double quote pair 내부다.
- quote 문자는 대상에 포함하지 않는다.
- 대상 pair를 찾지 못하면 unsupported/no-op 성격으로 상태를 보존한다.
- single quote, parenthesis, brace는 PLAYPACK-009 이후 별도 hardening 후보로 남긴다.

## 수용 기준

- completed: command catalog에 `text-object-quote-pair`가 approved/planned로 추가된다.
- completed: content command cluster YAML에 같은 cluster가 추가된다.
- completed: docs는 VIM-026과 PLAYPACK-009를 별도 루프로 분리한다.
- completed: 첫 구현 제외 항목이 command catalog, curriculum, roadmap 중 하나 이상에 명시된다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/content/...`
- passed: `git diff --check`
