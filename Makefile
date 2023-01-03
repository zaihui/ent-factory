##################################
#######       Setup       ########
##################################
.PHONY: setup

setup:
	@go mod download

##################################
#######        Tool       ########
##################################
.PHONY: fmt lint clean

fmt:
	@gofumpt -w .
	@gofmt -d -w -e .

lint:
	@golangci-lint run ./...

clean:
	@git clean -fdx ${COVERAGE_FILE}
