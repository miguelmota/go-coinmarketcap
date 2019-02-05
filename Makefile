all:
	@echo "no default"

.PHONY: test/v1
test/v1:
	@go test -v v1/*.go

.PHONY: test/v2
test/v2:
	@go test -v v2/*.go

.PHONY: test/pro/v1
test/pro/v1:
	@go test -v pro/v1/*.go

.PHONY: test
test: test/v1 test/v2 test/pro/v1

.PHONY: release
release:
	@rm -rf dist
	@goreleaser
