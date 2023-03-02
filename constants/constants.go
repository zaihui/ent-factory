package constants

var (
	FakeDataFunc         = "faker.FakeData" // faker library faker data func
	SuiteCaseType        = "factories.TestSuite"
	SuiteCaseVariable    = "s"       // instance of suite case
	SuiteNoErrorFunc     = "NoError" // error check func for suite case
	FactoryNewFuncName   = "New"
	SkipStructFields     = map[string]struct{}{"Edges": {}}
	Perm                 = 0o777
	DefaultAppPath       = "service/app"
	DefaultFactoryPath   = "factories"
	DefaultEntClientName = "app.EntClient"
	DefaultModelPath     = "spec/schema"
	ImportTime           = "time"
	IgnoreFolderNames    = []string{
		"enttest", "hook", "migrate", "partnerregister", "predicate", "runtime",
		"schema",
	} // schema is the old version(0.10.X) ent generated
	IgnoreField = "ID"
)
