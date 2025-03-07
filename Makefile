.PHONY: debug_run run
debug_run:
	go run ./cmd/int2 --verbose --scanner baseline tests/main.int2

run:
	go run ./cmd/int2 tests/main.int2
