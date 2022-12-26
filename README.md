# Ent Factory
a golang factory library for ent.

## Function
Simplify the process of creating mock model objects for ent schema. 
- Auto reading the struct of ent model schemas to create factory functions.
- Auto fake data for each field as default.
- Allow self define the value of any field independently.

## Struct
-- cmd
    -- root.go
    -- cmd.go
-- errors.go
-- main.go

## How to use
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

## Todo
- [ ] Support Time Optional Function
- [ ] Support Json Field Optional Function
- [ ] Lint
- [ ] Unit Test Coverage
- [ ] Version Compatibility Test
