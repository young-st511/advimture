# Release Workflow

Advimture는 `major.minor.patch` 버전을 사용한다. 버전은 `VERSION` 파일과 `vX.Y.Z` git tag를 함께 맞춘다.

## 배포 방식

- 기본 배포 채널은 `young-st511/homebrew-tap`의 source build formula다.
- `main` push와 PR은 CI에서 `go test ./...`, `go vet ./...`, `make e2e-smoke`를 실행한다.
- `vX.Y.Z` tag push는 release workflow를 실행한다.
- release workflow는 `make release-check`를 통과한 뒤 GitHub Release를 만들고, source archive의 `sha256`을 계산해 Homebrew tap formula를 갱신한다.
- tap 갱신은 `HOMEBREW_TAP_SSH_KEY` secret으로 수행한다. 이 secret은 `young-st511/homebrew-tap`에만 write 가능한 deploy key의 private key여야 한다.

## Homebrew Tap Secret 설정

자동 tap 갱신을 처음 켤 때는 사람이 로컬에서 아래 절차를 실행한다. 개인 PAT 대신 tap repo 전용 deploy key를 사용한다.

```sh
ssh-keygen -t ed25519 -C "advimture-release-homebrew-tap" -N "" -f ~/.ssh/advimture_homebrew_tap
gh repo deploy-key add ~/.ssh/advimture_homebrew_tap.pub --repo young-st511/homebrew-tap --title "advimture release workflow" --allow-write
gh secret set HOMEBREW_TAP_SSH_KEY --repo young-st511/advimture < ~/.ssh/advimture_homebrew_tap
```

secret 저장 후 private key 파일을 로컬에 계속 둘 필요가 없으면 안전하게 삭제한다.

```sh
rm ~/.ssh/advimture_homebrew_tap ~/.ssh/advimture_homebrew_tap.pub
```

## ExecPlan 완료 후 배포 순서

1. active ExecPlan의 수용 기준과 검증 결과를 채운다.
2. 완료된 ExecPlan은 `docs/exec-plans/completed/`로 이동한다.
3. 관련 변경만 stage하고 커밋한다.
4. 변경 성격에 맞춰 `make release-patch`, `make release-minor`, `make release-major` 중 하나를 실행한다.
5. release script가 `VERSION`을 올리고 `make release-check`를 통과하면 release commit과 `vX.Y.Z` tag를 push한다.
6. GitHub Actions release workflow가 GitHub Release와 Homebrew stable formula를 갱신한다.

## 게임 버전 기준

### Patch

플레이어 입장에서는 같은 버전대의 안정화로 느껴지는 변경이다.

- crash, rendering, input, scoring, replay, E2E flake 수정
- 오탈자, 짧은 카피, 힌트, briefing, success/failure 문구 조정
- 기존 exercise의 목표/정답/학습 의도를 바꾸지 않는 밸런스 조정
- Homebrew formula, CI, release script, 문서 보강
- 저장 포맷과 기존 progress 의미를 바꾸지 않는 내부 리팩터링

### Minor

새로운 플레이 경험이나 학습 콘텐츠가 추가되지만 기존 플레이어의 저장 데이터와 기본 조작 기대를 깨지 않는 변경이다.

- 새 tutorial, incident, exercise pack, command-choice drill 추가
- 새 Vim command cluster 또는 이미 학습한 command의 새 적용 루프 추가
- 새 UI affordance, review loop, hint/debrief 기능 추가
- backward-compatible progress 필드 추가 또는 저장하지 않는 동기부여/추천 레이어 추가
- 첫 공개 가능한 Brew 배포처럼 새 배포 채널을 정식으로 여는 변경

### Major

기존 플레이어가 “다른 게임 버전”으로 인식할 수준이거나 저장/학습 계약이 깨지는 변경이다.

- progress 저장 포맷 reset, 비호환 migration, 기존 completion 의미 변경
- 기본 조작 체계, command semantics, scoring model의 큰 변경
- 기존 playlist/exercise ID를 대규모로 바꾸거나 삭제해 저장된 진행과 충돌하는 변경
- 제품 정체성, 세계관, 핵심 게임 루프의 재정의
- 지원 플랫폼/터미널 요구사항을 크게 올리는 변경

## 수동 확인 명령

```sh
make release-check
./advimture --version
brew install --HEAD young-st511/tap/advimture
brew test young-st511/tap/advimture
```
