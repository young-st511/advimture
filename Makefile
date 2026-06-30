.PHONY: run build test release-check release-major release-minor release-patch e2e-smoke e2e-playable

E2E_GOCACHE ?= $(CURDIR)/artifacts/go-build-cache
VERSION ?= $(shell cat VERSION 2>/dev/null || printf dev)
LDFLAGS ?= -s -w -X main.version=$(VERSION)

run:
	go run .

build:
	GOCACHE=$(E2E_GOCACHE) go build -ldflags "$(LDFLAGS)" .

test:
	GOCACHE=$(E2E_GOCACHE) go test ./...

release-check: test build e2e-playable

release-major:
	scripts/release.sh major

release-minor:
	scripts/release.sh minor

release-patch:
	scripts/release.sh patch

e2e-smoke:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml

e2e-playable:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_forbidden_retry.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_arrow_forbidden.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_ctrl_c_does_not_quit.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_required_key.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_max_inputs.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_success_modal_80x24.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_viewport_failure_modal_80x24.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_playlist_next_survival.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_resume_first_incomplete.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_command_quit.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_command_mismatch_feedback.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_small_edits_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_substitute_current_line.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_full_first_five_minute.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_operator_grammar_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_yank_put_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_inner_word_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_repeat_last_change_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_search_basic_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_quote_pair_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_bracket_pair_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_selection_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_line_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_char_find_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_focus_panel.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_hint_affordance.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_003_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_004_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_006_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_007_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_008_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_009_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_010_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_011_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_011_redteam_scope_guard.yaml
