# ExecPlan: Homebrew Tap Release

Slice-ID: RELEASE-BREW-001
Created: 2026-06-11
Status: paused
Scope-Mode: normal
Allowed-Paths:
- `docs/exec-plans/active/release-brew-001-homebrew-tap.md`
- `README.md`
- `LICENSE`
- `VERSION`
- `main.go`
- `Makefile`
- `AGENTS.md`
- `internal/app/**`
- `internal/content/**`
- `.github/**`
- `scripts/**`
- `docs/release.md`
- `docs/guardrails.md`
- Homebrew tap repository: `young-st511/homebrew-tap`

## 목표

Advimture를 Homebrew 개인 tap으로 설치할 수 있게 만든다. 첫 배포 방식은 GoReleaser 없이 Homebrew source build formula를 사용하며, Brew 설치 후에도 콘텐츠 파일 누락 없이 실행되는 단일 바이너리 경험을 우선한다. ExecPlan 완료 후에는 변경을 커밋하고 major/minor/patch 중 하나로 버전을 올려 tag push 기반 자동 배포를 실행할 수 있게 한다.

## 범위

- 포함: 앱 버전 출력, embedded content fallback, 공개 배포용 license, README 설치 안내, Homebrew tap repo 생성, source build formula 작성, release workflow, version bump script, 하네스 release 안내
- 제외: `homebrew/core` 제출, GoReleaser/prebuilt cask 자동화, 새 Go 의존성 추가, progress 저장 포맷 변경, content schema 변경

## 수용 기준

- 사용자가 `brew install --HEAD young-st511/tap/advimture`로 설치할 수 있는 tap repo와 formula 초안이 있다.
- formula는 GitHub source checkout에서 Go source build로 `advimture` 바이너리를 설치한다.
- 첫 stable tag 이후 `brew install young-st511/tap/advimture`로 전환할 수 있는 후속 단계가 README/tap README에 명확히 남아 있다.
- Brew로 설치된 바이너리는 repo root의 `content/` 디렉터리 없이도 기본 playable content를 로드할 수 있다.
- `advimture --version` 또는 `advimture version`은 TUI를 띄우지 않고 버전 문자열을 출력한다.
- GitHub tag `vX.Y.Z` push 시 release workflow가 release gate를 실행하고, GitHub Release를 만들며, `young-st511/homebrew-tap` formula의 stable `url`/`sha256`을 갱신한다.
- release workflow는 개인 PAT 대신 `young-st511/homebrew-tap`에만 쓰기 가능한 deploy key secret을 사용한다.
- `scripts/release.sh major|minor|patch`는 깨끗한 작업트리에서 `VERSION`을 bump하고 release gate 통과 후 release commit/tag/push를 수행한다.
- major/minor/patch 기준은 보편적인 게임 버전 기준에 맞춰 하네스 문서에 명시한다.
- 새 Go 의존성 없이 구현한다.
- progress 저장 포맷은 변경하지 않는다.

## 검증 계획

- `gofmt` 대상 Go 파일
- `go test ./internal/content/... ./internal/app/...` 또는 해당 패키지 테스트
- `go test ./...`
- 가능하면 `go build .`
- 가능하면 Homebrew formula syntax/audit 또는 로컬 install 검증
- `go vet ./...`
- release workflow shell syntax 검토
- `git diff` 확인

## 의사결정 로그

- 2026-06-11: 사용자가 Brew 배포 진행과 `young-st511/*` 아래 homebrew repo 생성을 요청했다.
- 2026-06-11: GoReleaser는 사용하지 않고 `source build formula + 개인 tap`으로 시작한다. GoReleaser의 prebuilt cask 자동화는 반복 release가 필요해질 때 재검토한다.
- 2026-06-11: Brew 설치 경험을 단순하게 만들기 위해 `content/`는 바이너리에 embed하고, 파일 시스템 content root가 있으면 기존 개발/테스트 경로를 유지한다.
- 2026-06-11: 첫 tag가 아직 없고 로컬 `main`이 원격보다 앞서 있으므로, tap의 첫 formula는 stable URL 없이 `--HEAD` 설치를 지원한다.
- 2026-06-11: `young-st511/homebrew-tap` public repo를 만들고 `Formula/advimture.rb`를 push했다. `brew install --HEAD young-st511/tap/advimture`와 `brew test young-st511/tap/advimture`가 통과했다.
- 2026-06-12: 사용자가 자동 배포와 하네스 release 기준 추가를 요청했다. tag push 기반 GitHub Actions release와 tap deploy key 방식을 선택한다.
- 2026-06-12: 자동 tap 갱신에는 `HOMEBREW_TAP_SSH_KEY` secret이 필요하다. 보안 리뷰가 agent의 private key secret 업로드를 차단했으므로, `docs/release.md`에 사람이 실행할 secret 설정 절차를 남긴다.
- 2026-06-17: 사용자가 실패/성공 modal과 action affordance UX blocker를 먼저 닫기로 승인했다. `UI-MODAL-ACTION-HIERARCHY-001` 완료 전까지 release/tag/push 작업은 보류한다.
- 2026-06-17: `UI-MODAL-ACTION-HIERARCHY-001`은 완료됐다. 다만 사용자는 바로 출시가 아니라 출시 가능한 수준의 개발을 원하므로, release/tag/push 작업은 명시적 재개 요청 전까지 계속 보류한다.

## 미해결 질문

- 첫 실제 release tag는 첫 공개 가능한 Brew 배포이므로 `minor` bump로 `v0.1.0`을 후보로 둔다.
