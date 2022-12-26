package constants

import (
	"errors"
)

var (
	ErrDisableAllowed = errors.New("embedded fields disallowed")
	ErrNotDefinition  = errors.New("could not find struct type in definition file")
	ErrCanNotGen      = errors.New("cannot generate for fields whose type is imported")
	ErrNoFiled        = errors.New("no fields in struct (aside from ignored errors)")
)
