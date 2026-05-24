# CHOICE-PLAY-001 — Command Choice Drill Playable

Slice-ID: CHOICE-PLAY-001
Created: 2026-05-25
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/choice-play-001-command-choice-playable.md
- docs/exec-plans/completed/ui-playtest-001-hud-usability-pass.md
- content/command_clusters/
- content/exercises/
- content/scenarios/
- content/playlists/
- test/e2e/
- Makefile

## 목표

docs-only로 고정된 command-choice drill을 첫 playable content로 승격한다. 새 Vim engine 기능은 추가하지 않고, 이미 구현된 command 중 상황에 맞는 도구를 고르는 판단을 훈련한다.

## 첫 후보

- Drill: `choice-001-scope-triage`
- Focus: `scope-choice`
- Intended choice: linewise visual `V...d`
- 후보 판단: 단어 하나(`ciw`), quote 값(`ci"`), 줄 묶음(`Vd`) 중 삭제 범위를 고른다.

## 범위

- 포함:
  - command-choice command cluster/content 정의
  - 1~3문항 첫 playable tutorial 또는 applied drill
  - 기존 `constraints.required_keys`/`forbidden_keys`로 의도 선택 고정
  - content replay gate
  - focused E2E
- 제외:
  - 새 schema 필드
  - 새 Vim engine 기능
  - progress 저장 포맷 변경
  - incident pack polish

## 수용 기준

- 새 command를 소개하지 않고 기존 implemented command만 사용한다.
- 각 exercise는 선택 이유가 scenario success/failure copy에 드러난다.
- 첫 playable은 8문항 이하 playlist로 구성된다.
- content load/replay gate를 통과한다.
- focused E2E는 key trace와 app_state target을 검증한다.

## Step 1: Scope Approval

- [ ] 기존 docs candidate를 playable 범위로 구체화
- [ ] command cluster/status 결정
- [ ] playlist 위치와 category 결정

## Step 2: Content + Replay

- [ ] exercise YAML 추가
- [ ] scenario YAML 추가
- [ ] playlist YAML 추가
- [ ] content replay 테스트 통과

## Step 3: E2E + Docs

- [ ] focused E2E 추가
- [ ] Makefile suite 연결
- [ ] spec/docs 갱신
- [ ] `go test ./...`, `make e2e-playable`, `git diff --check`
