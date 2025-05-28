PHONY: test
test:
	go test --race -v --failfast

PHONY: btest
btest:
	go test --race -v --failfast -bench=. -benchmem