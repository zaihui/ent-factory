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
	DefaultEntClientName = "EntClient"
	IngoreFolderNames    = []string{"enttest", "hook", "migrate", "partnerregister", "predicate", "runtime"}
)
