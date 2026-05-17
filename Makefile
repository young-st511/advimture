.PHONY: run e2e-smoke

run:
	go run .

e2e-smoke:
	go run ./cmd/e2e-runner --scenario test/e2e/ftue_ctrl_c_quit.yaml
