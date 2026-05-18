# Scenario Flow Workbench

> `Scenario Production Harness`를 실제 작업 루프로 실행하기 위한 재사용 가능한 운영 문서다. 목표는 창의적 토론을 많이 하는 것이 아니라, `Vim command -> Exercise -> Scenario` 순서를 지키면서 더 좋은 플레이팩 결정을 반복 생산하는 것이다.

## 언제 사용하나

- 새 콘텐츠 팩을 만들 때
- 기존 playable pack의 scenario 품질을 높일 때
- command/exercise/scenario 중 어느 층이 약한지 판단해야 할 때
- E2E는 통과하지만 플레이 경험이 밋밋하거나 학습 효과가 흐릴 때

## 고정 원칙

1. command-first 원칙을 깨지 않는다.
2. scenario는 exercise의 target state, optimal keys, allowed keys, pass condition을 바꾸지 않는다.
3. engine/schema/progress 변경이 필요하면 이 워크벤치를 멈추고 별도 ExecPlan으로 분리한다.
4. SubAgent 토론은 결론을 내기 위한 장치다. 최종 적용은 Integrator가 must-fix / should-fix / defer로 통합한 뒤 수행한다.
5. `OK`는 작성자가 아니라 Verifier/Skeptic 역할이 준다.

## 입력

- `docs/roadmap/PRODUCT.md`
- `docs/roadmap/PROGRAM.md`
- `docs/gameplay/domain-contract.md`
- `docs/gameplay/spec.md`
- `docs/workflows/vim-learning-loops.md`
- `docs/workflows/scenario-production-harness.md`
- 관련 command/exercise/scenario bank
- 관련 `content/**/*.yaml`
- 관련 `test/e2e/**/*.yaml`
- 최근 E2E evidence 또는 실행 결과

## 역할

### Integrator

책임:

- 작업 범위와 ExecPlan을 고정한다.
- SubAgent 결과를 충돌 없이 합친다.
- 실제 파일 변경을 수행한다.
- E2E와 Go test로 회귀를 검증한다.

판단 기준:

- command/exercise를 바꾸지 않고 scenario만 개선할 수 있으면 scenario layer에 머문다.
- command coverage나 exercise target 자체가 약하면 scenario polish로 덮지 않고 defer 또는 별도 slice로 분리한다.

### Curriculum Reviewer

책임:

- command 순서, 선행 관계, 난이도 상승, coverage gap을 본다.
- 현재 플레이팩이 학습 목표를 과도하게 넓히는지 확인한다.

산출:

```yaml
curriculum_review:
  verdict: OK | REVISE | BLOCKED
  must_fix:
    - <blocking curriculum issue>
  should_fix:
    - <quality issue>
  defer:
    - <needs separate engine/content slice>
```

### Exercise Reviewer

책임:

- exercise goal, target state, optimal trace, allowed/forbidden keys가 학습 목표를 흐리지 않는지 본다.
- replay/E2E를 깨지 않는 범위에서 문항 표현 개선점을 찾는다.

산출:

```yaml
exercise_review:
  verdict: OK | REVISE | BLOCKED
  must_fix:
    - <blocking exercise issue>
  should_fix:
    - <quality issue>
  defer:
    - <requires exercise or engine change>
```

### Scenario Writer

책임:

- briefing, context role, success/failure feedback을 command 학습에 맞게 다듬는다.
- DevOps/터미널 문제 해결 톤과 가벼운 억까를 유지하되, 목표를 흐리지 않는다.

산출:

```yaml
scenario_writer_review:
  verdict: OK | REVISE | BLOCKED
  proposed_copy:
    <scenario_id>:
      briefing: <copy>
      mentor_success: <copy>
      mentor_failure: <copy>
  rationale:
    - <why this improves learning>
```

### Verifier / Skeptic

책임:

- 모든 제안을 `scenario-production-harness.md`의 OK Gate로 다시 본다.
- scenario가 target/keys/pass condition을 건드리려는 순간 reject한다.
- E2E assertion 동기화가 필요한 문구를 표시한다.

산출:

```yaml
verification_report:
  verdict: OK | REVISE | BLOCKED
  must_fix:
    - <blocking issue>
  should_fix:
    - <quality issue>
  e2e_sync_needed:
    - <test file or assertion text>
```

## 실행 순서

1. Integrator가 Work Start Protocol과 ExecPlan을 확인한다.
2. Integrator가 대상 팩의 command/exercise/scenario/E2E 현황을 요약한다.
3. Curriculum Reviewer, Exercise Reviewer, Scenario Writer, Verifier/Skeptic을 독립적으로 실행한다.
4. Integrator가 결과를 다음 표로 통합한다.

```text
issue
  source: curriculum | exercise | scenario | verifier
  classification: must-fix | should-fix | defer
  action: apply | document | split-slice | ignore-with-reason
```

5. `must-fix`가 있으면 먼저 해결한다.
6. `should-fix`는 범위 안에서 해결하되, command/exercise/schema 변경이 필요하면 defer한다.
7. 변경 후 E2E assertion 문구를 동기화한다.
8. 검증 명령을 실행한다.
9. ExecPlan을 completed로 이동하고 `PROGRAM.md`를 갱신한다.

## 적용 기준

적용한다:

- briefing이 command 목적을 더 명확히 하는 변경
- success/failure가 Vim 개념을 더 잘 강화하는 변경
- story tone을 정리하되 target/keys를 바꾸지 않는 변경
- E2E assertion에 쓰이는 화면 문구의 동기화

분리한다:

- 새 command 추가
- exercise target, initial buffer, optimal keys 변경
- content schema 변경
- progress 저장 포맷 변경
- engine behavior 변경

버린다:

- 세계관은 좋아지지만 학습 목표가 흐려지는 변경
- 플레이어가 풀어야 할 조작보다 긴 설명
- 현재 엔진이 지원하지 않는 Vim 기능을 playable처럼 보이게 하는 변경

## OK Checklist

- [ ] command-first 원칙을 지켰다.
- [ ] scenario 변경이 target/keys/pass condition을 바꾸지 않는다.
- [ ] 각 변경은 어떤 Vim 개념을 강화하는지 설명 가능하다.
- [ ] E2E screen assertion 문구가 변경된 copy와 일치한다.
- [ ] 실제 HOME이나 실제 progress 파일을 쓰지 않는다.
- [ ] `go test ./internal/content/...`가 통과한다.
- [ ] `go test ./...`가 통과한다.
- [ ] `make e2e-playable`이 통과한다.

## SubAgent Prompt Pack

### Curriculum Reviewer

```text
Advimture 첫 플레이팩을 curriculum 관점에서 검토하세요.
PRODUCT, PROGRAM, gameplay spec, scenario-production-harness, 대상 content YAML을 읽고 command 순서, 난이도 상승, coverage gap을 평가하세요.
파일을 수정하지 말고 must-fix / should-fix / defer로만 보고하세요.
```

### Exercise Reviewer

```text
Advimture 첫 플레이팩의 exercise 품질을 검토하세요.
각 exercise가 command를 실제로 훈련하는지, target/optimal/allowed keys가 명확한지, scenario polish로 해결 가능한 문제와 별도 slice가 필요한 문제를 분리하세요.
파일을 수정하지 말고 must-fix / should-fix / defer로만 보고하세요.
```

### Scenario Writer

```text
Advimture 첫 플레이팩의 scenario copy를 검토하세요.
DevOps/터미널 문제 해결 톤, 약간의 과하지 않은 억까, 학습 강화 문구를 기준으로 briefing/success/failure 개선안을 제안하세요.
exercise target, optimal keys, allowed keys, pass condition을 바꾸는 제안은 하지 마세요.
파일을 수정하지 말고 제안 copy와 이유만 보고하세요.
```

### Verifier / Skeptic

```text
Advimture 첫 플레이팩과 제안된 변경 방향을 scenario-production-harness OK Gate로 검증하세요.
command-first 위반, scenario-first 역류, E2E assertion 누락, progress/schema/engine 변경 필요성이 있으면 reject하세요.
파일을 수정하지 말고 OK / REVISE / BLOCKED 판정과 must-fix만 보고하세요.
```
