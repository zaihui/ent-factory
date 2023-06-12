SCHEMA_DIR=./spec/schema

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

migrations:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --target gen/entschema ${SCHEMA_DIR}

self_test:
	go build .
	rm factories/testfactory/testfactory.go
	./ent-factory generate --schemaPath gen/entschema  --outputPath factories --projectPath github.com/zaihui/ent-factory --overwrite=true --genImportFields=true
	rm ent-factory


all_factory:
	@go install github.com/zaihui/ent-factory@latest
	ent-factory generate --schemaPath gen/entschema --outputPath factories  --projectPath github.com/zaihui/ent-factory --factoriesPath factories
