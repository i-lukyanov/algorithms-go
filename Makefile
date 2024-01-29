## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## run: start project
.PHONY: run
run:
	go run main.go start

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## check: code syntax checking with linters
.PHONY: check
check:
	golangci-lint run -v --timeout=10m ./...

## test: running unit tests
.PHONY: test
test:
	go test -v -race -count=1 ./...

## test_color: running unit tests (gotestsum)
.PHONY: test_color
test_color:
	go install gotest.tools/gotestsum@latest
	gotestsum \
		--format testname \
		--jsonfile test-output.log

## bench: running benchmark tests
.PHONY: bench
bench:
	go test -bench  . -benchmem ./...

## audit: run quality control checks
.PHONY: audit
audit:
	go run honnef.co/go/tools/cmd/staticcheck@2023.1.3 -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify

## last_tags_list: return 10 latest tags
.PHONY: last_tags_list
last_tags_list:
	git tag --sort=-v:refname | head -n 10

# ==================================================================================== #
# TESTS
# ==================================================================================== #

## test_coverage: analysis of resource usage in the application
.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile coverage.out

## test_coverage_html: analysis of resource usage in the application
.PHONY: test_coverage_html
test_coverage_html:
	go test -coverprofile=coverage.out ./... \
	&& go tool cover -html=coverage.out \
	&& rm coverage.out
