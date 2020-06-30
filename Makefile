.PHONY: lint
lint:
	./hack/go-lint.sh

.PHONY: test
test:
	./hack/go-test.sh