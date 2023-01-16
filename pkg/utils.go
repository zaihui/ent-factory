package pkg

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
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

func WithFirstCharLower(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[0:1]) + s[1:]
}

func WithFirstCharUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

func FormatCode(buf *bytes.Buffer) (*bytes.Buffer, error) {
	code, err := io.ReadAll(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read code from buffer: %w", err)
	}

	formattedCode, err := format.Source(code)
	if err != nil {
		return nil, fmt.Errorf("failed to format code: %w", err)
	}
	return bytes.NewBuffer(formattedCode), nil
}
