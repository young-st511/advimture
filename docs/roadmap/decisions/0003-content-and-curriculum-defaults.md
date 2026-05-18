# 0003. Content and curriculum defaults

Date: 2026-05-18
Status: accepted

## Decision

Advimture의 중기 콘텐츠/커리큘럼 기본값을 아래처럼 고정한다.

- 첫 playable 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`로 간다.
- 콘텐츠 파일 포맷은 YAML을 우선한다. JSON은 필요해질 때 export/import 대상으로 검토한다.
- 게임이 읽는 콘텐츠는 repo root의 `content/` 아래에 둔다.
- draft 콘텐츠도 실제 파일로 둘 수 있다. 단 `status: draft` 또는 `engine_support: planned`는 playable 후보에서 제외한다.
- 다음 엔진 확장은 `w/b/e` word motion을 우선한다.
- 후반부에는 `gg`, `G`, line motions, `:` command, substitute, range command까지 폭넓게 다룬다.
- 시나리오 톤은 DevOps/터미널 문제 해결을 기본으로 한다. 약간의 과하지 않은 억까 상황은 허용하지만 학습 목표를 흐리면 안 된다.
- command cluster 승격은 사람이 명시 승인한다. exercise/scenario draft는 Agent가 만들고 Verifier OK 후 다음 단계로 진행할 수 있다.
- playable 연결은 approved 이상만 허용한다.
- coverage는 cluster 목적에 맞춘다. 초반 기본 이동처럼 방향 감각 자체가 목표인 cluster는 모든 command가 optimal trace에 등장해야 한다. 후반 범용 이동/조합 cluster는 모든 command를 억지로 같은 초반 루프에 넣지 않고, cluster별 `coverage_required`로 판단한다.
- 첫 버전은 `optimal_keys` 중심으로 학습한다. 대체 정답 허용은 후속 루프에서 추가한다.
- 진행 사항은 자동 저장한다. 기본 저장 위치는 기존 `.advimture` 계열 progress 저장을 유지한다.

## Consequences

- CONTENT-001은 YAML loader와 root `content/` 구조를 우선한다.
- loader는 draft/planned 콘텐츠를 읽을 수 있지만 playable candidate로 승격하지 않는다.
- GAMELOOP-001은 자동 저장을 기본 요구사항으로 포함한다.
- VIM-012 이후에도 `gg/G`, Ex command, substitute, range command를 curriculum backlog에서 잊지 않는다.
