.PHONY: run e2e-smoke e2e-playable

E2E_GOCACHE ?= $(CURDIR)/artifacts/go-build-cache

run:
	go run .

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
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_selection_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_visual_line_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_focus_panel.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_001_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_002_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_003_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_004_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml
