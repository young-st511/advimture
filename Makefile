.PHONY: run e2e-smoke e2e-playable

run:
	go run .

e2e-smoke:
	go run ./cmd/e2e-runner --scenario test/e2e/ftue_ctrl_c_quit.yaml

e2e-playable:
	go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml
