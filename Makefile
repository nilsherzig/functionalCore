PHONY: test
test:
	@echo "==> Running tests"
	go test -v -cover ./...

PHONY: run
run: test
	@echo "==> Running project"
	go run .
