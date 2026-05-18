.PHONY: run e2e-smoke e2e-playable

E2E_GOCACHE ?= $(CURDIR)/artifacts/go-build-cache

run:
	go run .

e2e-smoke:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml

e2e-playable:
	GOCACHE=$(E2E_GOCACHE) go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml
