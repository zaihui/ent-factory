# Ent Factory
a golang factory library for ent.

## Function
Simplify the process of creating mock model objects for ent schema. 
- Auto reading the struct of ent model schemas to create factory functions.
- Auto fake data for each field as default.
- Allow self define the value of any field independently.

## Struct
```
-- cmd
    -- root.go
    -- cmd.go
-- constants
    -- constants.go
    -- errors.go
-- main.go
-- README.md
```


## How to install
### Install
```bash
go install github.com/zaihui/ent-factory 
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
#### Makefile Command
```bash
all_factory:
  go run ent-factory generate --schemaPath {your ent schema path} --outputPath {path of your factories}
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
### how to use this factory
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

## Todo
- [ ] Support Time Optional Function
- [ ] Support Json Field Optional Function
- [ ] Lint
- [ ] Unit Test Coverage
- [ ] Version Compatibility Test
