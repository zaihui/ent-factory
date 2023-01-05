package schema

import (
	"errors"
	"unicode/utf8"
)

var errLength = errors.New("value is more than the max length")

func MaxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errLength
		}
		return nil
	}
}
