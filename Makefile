.PHONY: gen
gen:
	@buf generate

.PHONY: test
test:
	@go test -race ./...

.PHONY: vet
vet:
	@go vet ./...