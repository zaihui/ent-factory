# Ent Factory
![Golang](https://img.shields.io/badge/Golang-1.18%7C1.19-blue)
[![CI](https://github.com/zaihui/ent-factory/workflows/CI/badge.svg)](https://github.com/zaihui/ent-factory)
[![Maintainability](https://api.codeclimate.com/v1/badges/0f7322740bcbc6d7c2dd/maintainability)](https://codeclimate.com/github/zaihui/ent-factory/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaihui/ent-factory)](https://goreportcard.com/report/github.com/zaihui/ent-factory)


一个为Ent Go生成模型实例工厂代码的工具。

## 功能
简化为ent模式创建模拟模型对象的过程。
- 自动读取ent模型模式的结构以创建工厂函数。
  - 支持内置类型和导入类型字段。
  - 注意：默认情况下，它将生成导入类型字段。
- 默认为每个字段自动生成假数据。
- 允许使用函数可选模式自定义任何字段的值。
- 默认情况下，表的ID字段不能被编辑，因此不会为ID字段生成with函数和setter。

## 待办事项
- [ ] 单元测试覆盖率
- [ ] 版本兼容性测试
- [ ] 提高代码可维护性，从B提升到A
- [ ] 读取外键，支持创建相关实例后
- [X] 为CI/CD自动生成的CI配置
- [X] 代码Lint修复，从C提升到B
- [X] 为非只读的导入字段提供Setter
- [X] 更好的格式
  - [X] 在With函数之间添加一个空格。通过为每个函数添加文档来修复。
  - [X] 按fmt规则排序导入。 (通过exec.Command修复)
- [X] 为每个函数添加文档
- [X] 生成导入字段有一个开关进行控制
- [X] 自动格式化代码
- [X] 支持时间可选函数
- [X] 支持Json字段可选函数
- [X] 添加代码Lint
- [X] CI添加自动Lint
- [X] 选项参数控制是否覆盖现有工厂


## 结构
```
-- cmd
    -- root.go
    -- generate.go  // 主要逻辑
-- constants
    -- constants.go
    -- errors.go
-- factories        // 测试生成的代码在这里
-- gen              // ent生成的代码在这里
-- service          // 模拟app和ent客户端进行测试
-- spec             // ent表定义目录
-- main.go
-- README.md
-- Makefile
```

## 需求
- Golang 1.18～1.19
- [Ent](https://entgo.io) 0.10.0 - 0.11.4
- [Faker](https://github.com

/bxcodec/faker) 2.0
## 如何安装
### 安装
```bash
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/zaihui/ent-factory@latest
```
### 设置
#### 标志
- schemaPath
  - 你的entschema的路径，通常是 "gen/entschema"
  - **必需**
- outputPath
  - 你的工厂的路径
  - 终点必须是 "factories"
  - **必需**
- projectPath
  - 这个项目的相对路径
  - 例如 `github.com/zaihui/ent-factory`
  - **必需**
- overwrite
  - 是否覆盖这些存在的文件
  - 默认值是 `false`
- factoriesPath
  - 这些工厂在这个项目中的相对路径
  - 例如 `factories`，意味着 `projectPath/factories`
  - 默认值是 `factories`
- appPath
  - app客户端的相对路径
  - 例如 `service/app` 意味着 `github.com/zaihui/ent-factory/service/app`
  - 默认值是 `service/app`
- entClientName
  - ent客户端的名字，意味着 `appPath.entClientName`
  - 默认值是 `EntClient`
- genImportFields
  - 设置是否生成模型的导入字段的布尔值
  - 默认值是 `true`
#### Makefile命令
> 注意！：默认情况下，此命令将**不会**根据表的当前结构重新生成每个工厂。
> 
> 如果你想**重新生成**，请使用 `overwrite` 标志。或者删除你想重新生成的，然后使用这个命令。
```bash
all_factory:
  go run ent-factory generate --schemaPath {你的ent schema路径} --outputPath {你的工厂的路径} ----projectPath {你的项目模块路径}
```
示例 
```bash
all_factory:
  go run ent-factory generate --schemaPath gen/entschema --outputPath  /Users/lvxinyan/zaihui/ent-factory/factories --projectPath github.com/zaihui/ent-factory
## for one factory
ent-factory generate --schemaFile gen/entschema/{one ent model file}.go --outputPath  /Users/lvxinyan/zaihui/ent-factory/factories --projectPath github.com/zaihui/ent-factory
```


## 示例
### 工厂的示例
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
		bookTable

OrderGen.UserUID = userUIDGen
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
### 如何使用这个工厂
```go
package main
import (
	your_location_of_factories/booktableorderfactory
)
// s 是测试套件的一个实例
order := booktableorderfactory.new(s) 
// 如果你想自定义一个字段的值，例如 UID
order2: = booktableorderfactory.new(s, booktableorderfactory.SetUID("你的uid"))
```

## 特别感谢
特别感谢splunk。关于函数可选模式的代码生成部分基于splunk的这个[项目](https://github.com/splunk/go-generate-builder-opts)。

我修改了一些逻辑来实现我的需求，所以我没有将它作为一个包导入使用。

## 如何贡献
### 流程
1. Fork这个仓库
2. 将你的仓库克隆到本地
3. 基于master创建一个分支
4. 添加你的代码
5. 你可以使用 `make self_test` 来测试你的代码。它将构建这个包，然后为 `spec/schema` 中的表生成代码。
6. 使用 `make fmt` 格式化你的代码，使用 `make lint` 检查代码风格。
7. 遵循基本规则创建一个pull request，快乐编程！
### 基本规则
1. PR只能有一个提交，并且需要rebase master分支
2. 必须有UTs，如果可能的话。
3. 提交消息必须符合格式。 
```
EF-{Number}(label): title
label:
- feat, for feature
- fix, for bug fix
- ut, for ut
- doc, for document
- refactor, for refactor code
```
快乐编程，快乐分享！