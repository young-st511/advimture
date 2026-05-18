.PHONY: run e2e-smoke e2e-playable

E2E_GOCACHE ?= $(CURDIR)/artifacts/go-build-cache

run:
	go run .

e2e-smoke:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml

e2e-playable:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_forbidden_retry.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_required_key.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_constraint_max_inputs.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_playlist_next_survival.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_resume_first_incomplete.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_command_quit.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_small_edits_full.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_substitute_current_line.yaml
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_full_first_five_minute.yaml
