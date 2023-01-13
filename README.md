# Ent Factory
[![CI](https://github.com/zaihui/ent-factory/workflows/CI/badge.svg)](https://github.com/zaihui/ent-factory)
![Golang](https://img.shields.io/badge/Golang-1.18%7C1.19-blue)

A Golang Factory Code Generation Command Tool For Ent.

## Function
Simplify the process of creating mock model objects for ent schema. 
- Auto reading the struct of ent model schemas to create factory functions.
  - Support build-in type and imported type field
  - PS: it will generate the imported type fields in default
- Auto fake data for each field as default.
- Allow self define the value of any field independently.

## Todo
- [ ] Better Format
  - [ ]Add A Blank Space Between With Functions. It's tricky, because printer rules against gofmt rules.
  - [ ]Import Sort by fmt rules
- [ ] Unit Test Coverage
- [ ] Version Compatibility Test
- [X] ~~Generate imported fields have a switch to control~~
- [X] ~~Auto Format Code~~
- [X] ~~Support Time Optional Function~~
- [X] ~~Support Json Field Optional Function~~
- [X] ~~Code Lint~~
- [X] ~~CI Add Auto Lint~~
- [X] ~~Option param to control overwrite exist factory or not~~

## Struct
```
-- cmd
    -- root.go
    -- generate.go  // main logic
-- constants
    -- constants.go
    -- errors.go
-- factories        // the test generate code in here
-- gen              // ent generate code in here
-- service          // simulate app and ent client for test
-- spec             // ent tables definition dir
-- main.go
-- README.md
-- Makefile
```

## Requirement
- Golang 1.18～1.19
- [Ent](https://entgo.io) 0.10.0 - 0.11.4
- [Faker](https://github.com/bxcodec/faker) 2.0
## How to install
### Install
```bash
go install github.com/zaihui/ent-factory@latest
```
### Setup
#### Flags
- schemaPath
  - path of your entschema, normally is "gen/entschema"
  - **required**
- outputPath
  - path of your factories
  - the endpoint must be "factories"
  - **required**
- projectPath
  - the relative path of this project
  - eg. `github.com/zaihui/ent-factory`
  - **required**
- overwrite
  - whether overwrite these exist files
  - the default value is `false`
- factoriesPath
  - the relative path of these factories located in this project
  - eg. `factories`, means `projectPath/factories`
  - the default value is `factories`
- appPath
  - the relative path of app client
  - eg. `service/app` means `github.com/zaihui/ent-factory/service/app`
  - the default value is `service/app`
- entClientName
  - the name of ent client, means `appPath.entClientName`
  - the default value is `EntClient`
- genImportFields
  - the bool value to set whether generate those imported fields of the models
  - the default value is `true`
#### Makefile Command
> CAUTION!: This Command will **NOT** regenerate every factory based on the current struct of table.
> 
> If you want to **Regenerate**, please use the `overwrite` flag. Or delete what you want to regenerate then use this 
> command.
```bash
all_factory:
  go run ent-factory generate --schemaPath {your ent schema path} --outputPath {path of your factories} ----projectPath {your project module path}
```
Sample 
```bash
all_factory:
  go run ent-factory generate --schemaPath gen/entschema --outputPath  /Users/lvxinyan/zaihui/ent-factory/factories --projectPath github.com/zaihui/ent-factory
## for one factory
ent-factory generate --schemaFile gen/entschema/{one ent model file}.go --outputPath  /Users/lvxinyan/zaihui/ent-factory/factories --projectPath github.com/zaihui/ent-factory
```


## Sample
### The sample of a factory
```go
package booktableorderfactory

import (
	"github.com/bxcodec/faker"
	"github/zaihui/ent-factory/factories"
	"github/zaihui/ent-factory/gen/entschema"
	"github/zaihui/ent-factory/service/app"
)

type BookTableOrderFieldSetter func(*entschema.BookTableOrder)


func SetUID(uIDGen string) BookTableOrderFieldSetter {
	return func(bookTableOrderGen *entschema.BookTableOrder) {
		bookTableOrderGen.UID = uIDGen
	}
}

func SetUserUID(userUIDGen string) BookTableOrderFieldSetter {
	return func(bookTableOrderGen *entschema.BookTableOrder) {
		bookTableOrderGen.UserUID = userUIDGen
	}
}
func SetOrderNo(orderNoGen string) BookTableOrderFieldSetter {
	return func(bookTableOrderGen *entschema.BookTableOrder) {
		bookTableOrderGen.OrderNo = orderNoGen
	}
}
func SetType(typeGen string) BookTableOrderFieldSetter {
	return func(bookTableOrderGen *entschema.BookTableOrder) {
		bookTableOrderGen.Type = typeGen
	}
}

func New(s factories.TestSuite, opts ...BookTableOrderFieldSetter) *entschema.BookTableOrder {
	data := entschema.BookTableOrder{}
	s.NoError(faker.FakeData(&data))
	for _, opt := range opts {
		opt(&data)
	}
	return app.EntClient.BookTableOrder.Create().
		SetUID(data.UID).
		SetUserUID(data.UserUID).
		SetOrderNo(data.OrderNo).
		SetType(data.Type).
		SaveX(s.Context())
}
```
### How to use this factory
```go
package main
import (
	your_location_of_factories/booktableorderfactory
)
// s is an instance of test suite
booktableorderfactory.new(s) 
// if you want to customize the value of a field, eg. UID
booktableorderfactory.new(s, booktableorderfactory.SetUID("your uid"))
```

## Special Thank
Special Thanks splunk. The part of codegen about function optional pattern is based on this [project](https://github.com/splunk/go-generate-builder-opts) from splunk.

I edit some logic of it to implement my need, so I haven't imported it as a package to use it.

## How to contribute
### Process
1. Fork this repo
2. Clone your repo to local
3. Create a branch based on master
4. Add your code
5. You can use the `make self_test` to test your code. It will build this package, then generate code for the table in `spec/schema`.
6. Use `make fmt` format your code, and use `make lint` to check the code styles.
7. Follow the base rule to create a pull request, happy coding!
### Base Rule
1. PR only can have one commit, and it needs to rebase the master branch
2. It must have UTs, if it is possible.
3. Commit Message must fit the format. 
```
EF-{Number}(label): title
label:
- feat, for feature
- fix, for bug fix
- ut, for ut
- doc, for document
- refactor, for refactor code
```
Happy Coding, Happy Sharing!