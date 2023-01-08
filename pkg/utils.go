package pkg

import (
	"fmt"
	"reflect"
	"strings"
)

func SliceContain(iterableType interface{}, value interface{}) bool {
	v := reflect.ValueOf(iterableType)

	switch kind := reflect.TypeOf(iterableType).Kind(); kind {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == value {
				return true
			}
		}

	case reflect.Map:
		if v.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true
		}
	case reflect.String:
		s := v.String()
		ss, ok := value.(string)
		if !ok {
			panic("类型不匹配")
		}

		return strings.Contains(s, ss)
	default:
		panic(fmt.Sprintf("类型：%s 不受支持", iterableType))
	}

	return false
}
