# Advimture

Vim을 실제로 유용하게 쓰기 위한 Go + Bubble Tea 기반 터미널 어드벤처 게임입니다.

플레이어는 `Runbook Dispatch`의 원격 복구 오퍼레이터가 되어, Vim 명령으로 고장 난 설정과 로그를 복구합니다. 설계 우선순위는 항상 `Vim command -> Exercise -> Scenario`입니다. 세계관은 재미를 돕는 얇은 레이어이고, 핵심은 반복 가능한 Vim 학습입니다.

## 현재 플레이 가능 범위

현재 앱은 첫 공개 전 foundation build입니다.

- 생존 조작: `esc`, `:q!`, `:wq`
- 이동: `h/j/k/l`, `w/b/e`, `gg/G/0/$`
- 작은 수정: `x`, `r`, `i/a/A`, `u`, `ctrl+r`, `o/O`
- operator grammar: `dw`, `d$`, `dd`, `cw`, `c$`, `cc`, `yw`, `y$`, `yy`, `p/P`
- text object: `diw`, `ciw`, `yiw`, `di"`, `ci"`, `yi"`, `di'`, `ci'`, `yi'`
- 효율/탐색: `.`, `/`, `n`, `N`, `:s`, `:%s`, `:2,3s`
- visual/inline target: `v`, `V`, visual `d/y`, `f/t`, `df/dt/cf/ct`
- 게임 루프: tutorial playlist, incident run, 실패 후 재시도, 힌트, 복습 queue, 성공 debrief, 자동 저장

콘텐츠는 `content/`의 YAML로 관리하며, 현재 106개 exercise와 18개 exercise 파일, 18개 playlist 파일이 연결되어 있습니다.

## 요구 사항

- Go `1.26.1` 이상
- ANSI terminal
- 권장 터미널 크기: 100x30 이상

좁은 터미널에서도 동작하지만, 긴 한국어 안내와 modal이 많이 줄바꿈될 수 있습니다.

## 실행

Homebrew tap으로 개발 중인 최신 `main`을 source build 설치할 수 있습니다.

```sh
brew install --HEAD young-st511/tap/advimture
```

stable tag 배포 후에는 아래 명령으로 설치할 수 있습니다.

```sh
brew install young-st511/tap/advimture
```

소스 체크아웃에서 바로 실행할 수도 있습니다.

```sh
go run .
```

또는 Makefile을 사용할 수 있습니다.

```sh
make run
```

## 진행 저장

진행 상황은 자동으로 저장됩니다.

```text
~/.advimture/progress.json
~/.advimture/progress.json.bak
```

처음부터 다시 플레이하려면 progress 파일을 백업하거나 삭제합니다.

```sh
mv ~/.advimture/progress.json ~/.advimture/progress.json.manual-backup
```

## 검증

단위 테스트:

```sh
make test
```

첫 문제 smoke E2E:

```sh
make e2e-smoke
```

현재 playable route 전체 E2E:

```sh
make e2e-playable
```

공개 전 로컬 release gate:

```sh
make release-check
```

E2E runner는 테스트 전용 임시 HOME을 사용하므로 실제 `~/.advimture/progress.json`을 건드리지 않습니다. 실행 evidence는 `artifacts/e2e/` 아래에 생성되며 git에는 포함하지 않습니다.

## 빌드

```sh
make build
```

생성되는 `advimture` 바이너리는 `.gitignore` 대상입니다.

## 알려진 제한

- Vim 전체 구현이 아니라 학습용 Vim-like engine입니다.
- search는 case-sensitive literal `/`, `n`, `N`만 지원합니다. `?`는 backward search가 아니라 게임 힌트 키입니다.
- Ex substitute는 literal `:s`, `:%s`, `:2,3s` 중심이며 Vim regex 전체를 지원하지 않습니다.
- named register, macro, count prefix, visual block, 복잡한 indentation/auto-comment는 아직 범위 밖입니다.
- progress v1은 완료/최고 기록 중심입니다. spaced review, mastery, streak 저장은 후속 설계 대상입니다.
- CI와 tag 기반 release workflow가 있습니다. 로컬 공개 전 release gate는 `make release-check`입니다.

## 문서와 개발 워크플로우

이 프로젝트는 plan-first + spec-first 방식으로 개발합니다.

- 제품 방향: `docs/roadmap/PRODUCT.md`
- 현재 phase와 다음 slice: `docs/roadmap/PROGRAM.md`
- rolling plan: `docs/roadmap/FORWARD_PLAN.md`
- 실행 계획: `docs/exec-plans/`
- Gameplay spec: `docs/gameplay/spec.md`
- Verification spec: `docs/verification/spec.md`
- Release workflow: `docs/release.md`

비사소한 변경은 `docs/exec-plans/active/`에 ExecPlan을 만든 뒤 진행합니다.
