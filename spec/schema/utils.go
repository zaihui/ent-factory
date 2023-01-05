package schema

import (
	"entgo.io/ent/dialect"
)

const (
	LenAesKey      = 43
	LenAppID       = 25
	LenDate        = 10
	LenDesc        = 255
	LenJSON        = 1530
	LenKey         = 4096
	LenMobile      = 11
	LenName        = 100
	LenNormal      = 50
	LenOrder       = 24
	LenPhone       = 20
	LenURL         = 255
	LenUUID        = 32
	LenTwenty      = 20
	LenHundred     = 100
	LenFiveHundred = 500
	DishDesc       = 500

	FALSE = false
)

var (
	DatetimeSchema = map[string]string{dialect.MySQL: "datetime"}
	DateSchema     = map[string]string{dialect.MySQL: "date"}
	DecimalSchema  = map[string]string{dialect.MySQL: "decimal(12,2)"}
	GeoSchema      = map[string]string{dialect.MySQL: "decimal(9,6)"}

	Varchar30  = map[string]string{dialect.MySQL: "varchar(30)"}
	Varchar100 = map[string]string{dialect.MySQL: "varchar(100)"}
	Varchar255 = map[string]string{dialect.MySQL: "varchar(255)"}

	Varchar500 = map[string]string{dialect.MySQL: "varchar(500)"}
)
