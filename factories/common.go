package factories

import "context"

type TestSuite interface {
	NoError(err error, msgAndArgs ...interface{}) bool
	Context() context.Context
}
