# 0005 — Remote Recovery Bureau World Frame

Date: 2026-05-22

## Status

Accepted

## Context

`Structure Editing and Applied Survival` 완료 후 점검에서 incident run이 기능적으로는 통과했지만, 아직 하나의 복구 작전보다 종합시험처럼 느껴질 위험이 확인됐다.

사용자는 이제 세계관을 본격적으로 얹고 싶다고 요청했고, 세 명의 시나리오 작가 SubAgent가 세계관 후보를 제안했다. 메타 리뷰어는 `원격 시설 복구국 / Runbook Dispatch`를 기본 골격으로 두고, `침묵한 릴레이 기지` 감각만 얇게 섞는 혼합안을 추천했다.

## Decision

Advimture의 중반 이후 세계관 프레임은 **원격 시설 복구국 / Runbook Dispatch**로 한다.

플레이어는 낡은 원격 시설의 장애 runbook을 Vim으로 복구하는 **원격 복구 오퍼레이터**다. 각 incident는 시설 하나, 콘솔 하나, 장애 하나를 닫는 짧은 복구 작전으로 다룬다.

개별 사건에는 `릴레이`, `신호`, `폭풍권`, `잔류 리스크` 같은 명사를 얇게 사용할 수 있다. 단, 세계관 설명은 command 학습을 가리지 않아야 하며, 브리핑은 “상황 1문장 + Vim 조작 목표 1문장” 원칙을 따른다.

## Consequences

- `incident-001-hotfix`는 “릴레이 기지 001: 야간 핫픽스 복구”로 리프레이밍한다.
- `incident-002-structure-recovery`는 “릴레이 기지 002: 구조 설정 재동기화”로 리프레이밍한다.
- review queue, best record, key count는 “잔류 리스크”, “재점검”, “최단 복구 기록” 같은 언어로 포장할 수 있다.
- visual mode는 “오염 구역 지정”으로 확장할 수 있다.
- 장대한 SF lore, 전투/체력/인벤토리, 관료적 티켓 처리 반복은 피한다.

