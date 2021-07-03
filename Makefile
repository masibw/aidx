.PHONY: test
test:
	go test -v ./... -count=1 -cover

.PHONY: lint
lint:
	golangci-lint run --out-format=github-actions --enable=staticcheck,stylecheck,gosimple,gosec,prealloc,gocognit,bodyclose,gofmt
