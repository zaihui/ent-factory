package cmd

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/zaihui/ent-factory/pkg"

	"github.com/spf13/cobra"

	"github.com/zaihui/ent-factory/constants"
)

func init() {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate ent model factory",
		Run:   GenerateFactories,
	}
	rootCmd.AddCommand(cmd)
}

type GenFlags struct {
	SchemaFile      string
	SchemaPath      string
	OutputPath      string
	ProjectPath     string
	Overwrite       bool
	FactoriesPath   string
	AppPath         string
	EntClientName   string
	ModelPath       string
	GenImportFields bool
}

type SetValueParam struct {
	generateForUnexportedFields bool
	fieldContainsImport         bool
	identName                   string
	indentString                string
}

// GenerateFactories generates factory files for a given schema.
func GenerateFactories(cmd *cobra.Command, _ []string) {
	flags := ExtraFlags(cmd)
	f, err := os.Open(flags.SchemaPath)
	if err != nil {
		fail(err.Error())
	}
	defer f.Close()

	createCommonPathIfNeeded(flags)

	if isSingleFile(flags) {
		GenerateFactoryForOneFile(flags)
		return
	}

	processSchemaDirectory(flags, f)
}

func createCommonPathIfNeeded(flags GenFlags) {
	commonPath := fmt.Sprintf("%s/common.go", flags.OutputPath)
	_, err := os.Stat(commonPath)
	if err != nil && os.IsNotExist(err) {
		CreatePathAndCommonFile(commonPath, flags.OutputPath)
	} else if err != nil {
		fail(err.Error())
	}
}

func isSingleFile(flags GenFlags) bool {
	return flags.SchemaFile != "" && flags.SchemaPath == ""
}

func processSchemaDirectory(flags GenFlags, f *os.File) {
	files, err := f.Readdir(0)
	if err != nil {
		fail(err.Error())
	}
	for _, v := range files {
		if !v.IsDir() || shouldBeIgnored(v) {
			continue
		}

		realPath, realOutPutPath := GetRealPathAndFilePath(flags.SchemaPath, v, flags.OutputPath)

		if !flags.Overwrite {
			checkAndCreateFactory(flags, realPath, v.Name(), realOutPutPath)
		} else {
			CreateOneFactory(realPath, v.Name(), realOutPutPath, flags)
		}
	}
}

func shouldBeIgnored(file os.FileInfo) bool {
	for _, n := range constants.IgnoreFolderNames {
		if file.Name() == n {
			return true
		}
	}
	return false
}

func checkAndCreateFactory(flags GenFlags, realPath, dirName, realOutPutPath string) {
	_, err := os.Stat(realOutPutPath)
	switch {
	case err == nil:
		return
	case os.IsNotExist(err):
		CreateOneFactory(realPath, dirName, realOutPutPath, flags)
	default:
		fail(fmt.Sprintf("Error occurred while checking file existence: %s", realOutPutPath))
	}
}

// GenerateFactoryForOneFile only for one model file.
func GenerateFactoryForOneFile(flags GenFlags) {
	schemaName := ExtraNameFromSchemaFilePath(flags.SchemaFile)
	realOutPutPath := fmt.Sprintf("%s/%sfactory/%sfactory.go", flags.OutputPath, schemaName, schemaName)
	CreateOneFactory(flags.SchemaFile, schemaName, realOutPutPath, flags)
}

func GetRealPathAndFilePath(schemaPath string, v os.FileInfo, outputPath string) (string, string) {
	filePath := fmt.Sprintf("%s/%s", schemaPath, v.Name())
	realPath := fmt.Sprintf("%s.go", filePath)
	realOutPutPath := fmt.Sprintf("%s/%sfactory/%sfactory.go", outputPath, v.Name(), v.Name())
	return realPath, realOutPutPath
}

func ExtraNameFromSchemaFilePath(schemaFile string) string {
	endPoints := strings.Split(schemaFile, "/")
	SchemaFileName := endPoints[len(endPoints)-1]
	schemaNames := strings.Split(SchemaFileName, ".")
	schemaName := schemaNames[0]
	return schemaName
}

// CreateOneFactory create one factory.
func CreateOneFactory(realPath, schemaName, realOutPutPath string, flags GenFlags) {
	outReader, err := RunGenerate(realPath, schemaName, realOutPutPath, flags)
	if err != nil {
		fail(err.Error())
	}
	var dest io.Writer
	switch {
	case flags.OutputPath == "":
		dest = os.Stdout
	default:
		_, err := os.Stat(filepath.Dir(realOutPutPath))
		switch {
		case os.IsNotExist(err):
			err2 := os.MkdirAll(filepath.Dir(realOutPutPath), os.FileMode(constants.Perm))
			if err2 != nil {
				fail(err2.Error())
			}
		case !os.IsNotExist(err) && err != nil:
			fail(err.Error())
		}
		dest, err = os.OpenFile(realOutPutPath, os.O_RDWR|os.O_CREATE, os.FileMode(constants.Perm))
		if err != nil {
			fail(err.Error())
		}
	}
	if _, err := io.Copy(dest, outReader); err != nil {
		fail(err.Error())
	}
	cmd := exec.Command("goimports", "-w", realOutPutPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fail(err.Error())
	}
}

func CreatePathAndCommonFile(commonPath string, outputPath string) {
	err := os.MkdirAll(filepath.Dir(commonPath), os.FileMode(constants.Perm))
	if err != nil {
		fail(err.Error())
	}
	endPoints := strings.Split(outputPath, "/")
	packageName := endPoints[len(endPoints)-1]
	code := fmt.Sprintf(`package %s 

                     import "context"

                     type TestSuite interface {
	                      NoError(err error, msgAndArgs ...interface{}) bool
	                      Context() context.Context
                     }`, packageName)
	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		fail(err.Error())
	}
	fc, err := os.OpenFile(fmt.Sprintf("%s/common.go", outputPath), os.O_RDWR|os.O_CREATE, os.FileMode(constants.Perm))
	if err != nil {
		fail(err.Error())
	}
	defer func(fc *os.File) {
		err := fc.Close()
		if err != nil {
			fail(err.Error())
		}
	}(fc)
	_, err = fmt.Fprint(fc, string(formattedCode))
	if err != nil {
		fail(err.Error())
	}
}

func ExtraFlags(cmd *cobra.Command) GenFlags {
	flags := GenFlags{}
	t := reflect.TypeOf(flags)
	flagsValue := reflect.ValueOf(&flags).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonName := pkg.WithFirstCharLower(field.Name)
		//nolint:exhaustive // only have two types
		switch field.Type.Kind() {
		case reflect.String:
			value, err := cmd.Flags().GetString(jsonName)
			if err != nil {
				fail(fmt.Sprintf("get %s failed: %v\n", jsonName, err))
			}
			flagsValue.Field(i).SetString(value)
		case reflect.Bool:
			value, err := cmd.Flags().GetBool(jsonName)
			if err != nil {
				fail(fmt.Sprintf("get %s failed: %v\n", jsonName, err))
			}
			flagsValue.Field(i).SetBool(value)
		}
	}
	if flags.ProjectPath == "" {
		Fatalf("project path cannot be empty")
	}
	if flags.SchemaFile == "" && flags.SchemaPath == "" {
		Fatalf("schema file and schema path must give at lease one")
	}
	if flags.OutputPath == "" {
		Fatalf("output path cannot be empty")
	}
	return flags
}

func fail(msg string) {
	if msg != "" {
		_, err := fmt.Fprintln(os.Stderr, msg)
		if err != nil {
			return
		}
	}
	os.Exit(1)
}

// RunGenerate ===== generate one factory schema =======.
// todo 57 lines of code (exceeds 50 allowed).
func RunGenerate(schemaFile, schemaTypeName, outputPath string, flags GenFlags) (io.Reader, error) {
	// Read input file
	fset := token.NewFileSet()
	astF, err := parser.ParseFile(fset, schemaFile, nil, 0)
	if err != nil {
		return nil, err
	}

	// Look for specified struct type.
	structType, ok := findRequestedStructType(astF, schemaTypeName)
	if !ok {
		return nil, constants.ErrNotDefinition
	}
	path := schemaFile
	endPoints := strings.Split(path, "/")
	packagePrefix := endPoints[len(endPoints)-2]

	fnTypeIdent := funcTypeIdent(structType.Name.Name, true)
	fnParamType := &ast.StarExpr{
		X: ast.NewIdent(packagePrefix + "." + structType.Name.String()),
	}
	paramTypeName := fmt.Sprintf("%s.%s", packagePrefix, structType.Name.String())
	// Initialize output
	outEndPoints := strings.Split(outputPath, "/")
	packageName := outEndPoints[len(outEndPoints)-2]
	astOut := &ast.File{Name: ast.NewIdent(packageName)}

	// Add import
	getImportDef(astOut, structType, true, flags.GenImportFields, flags.ProjectPath, flags.FactoriesPath,
		flags.AppPath, flags.ModelPath, flags.SchemaPath)

	// Add type definition for functional option function signature
	withTypeDef(astOut, fnTypeIdent, fnParamType)

	// Add function for each applicable struct field
	if err := withFunc(astOut, structType, fnTypeIdent, fnParamType, false, true, flags.GenImportFields,
		constants.SkipStructFields); err != nil {
		return nil, err
	}

	if err := NewFunc(astOut, paramTypeName, structType, fnTypeIdent, fnParamType, false, true,
		constants.SkipStructFields, flags.EntClientName, flags.GenImportFields,
	); err != nil {
		return nil, err
	}
	// Generate output file
	out := new(bytes.Buffer)
	if err := printer.Fprint(out, token.NewFileSet(), astOut); err != nil {
		return nil, err
	}
	res, err := pkg.FormatCode(out)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewFunc Create the New instance function. todo 85 lines of code (exceeds 50 allowed).
func NewFunc(astOut *ast.File, paramTypeName string, structType *ast.TypeSpec, fnIdent *ast.Ident,
	fnParamType *ast.StarExpr, generateForUnexportedFields, ignoreUnsupported bool,
	skipStructFields map[string]struct{}, entClient string, genImportFields bool,
) error {
	suiteIndent, suiteNoErrIndent, optsIndent, testCaseIndent, EllipsisIndent, returnIndent, dataIndent, optKeyIndent,
		optValueIndent, fakerIndent, dataResIndent, dataResPosIndent := CreateIndentForNewFunc(fnIdent, entClient,
		structType, paramTypeName)
	newFunc := &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{
				Text: fmt.Sprintf("// %s function for creating one %s instance.",
					constants.FactoryNewFuncName, structType.Name),
			},
		}},
		Name: ast.NewIdent(constants.FactoryNewFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{suiteIndent},
						Type:  testCaseIndent,
					},
					{
						Names: []*ast.Ident{optsIndent},
						Type:  EllipsisIndent,
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{{Type: fnParamType}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						dataResIndent,
					},
					TokPos: 0,
					Tok:    token.DEFINE,
					Rhs: []ast.Expr{
						dataIndent,
					},
				},
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: suiteNoErrIndent,
						Args: []ast.Expr{
							&ast.CallExpr{
								Fun:  fakerIndent,
								Args: []ast.Expr{dataResPosIndent},
							},
						},
					},
				},
				&ast.RangeStmt{
					Key:   optKeyIndent,
					Value: optValueIndent,
					X:     optsIndent,
					Tok:   token.DEFINE,
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ExprStmt{
								X: &ast.CallExpr{
									Fun:  optValueIndent,
									Args: []ast.Expr{dataResPosIndent},
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Return: 0,
					Results: []ast.Expr{
						getFactoryReturn(
							returnIndent,
							structType,
							generateForUnexportedFields,
							ignoreUnsupported,
							genImportFields,
							skipStructFields,
						),
					},
				},
			},
		},
	}
	astOut.Decls = append(astOut.Decls, newFunc)
	return nil
}

func CreateIndentForNewFunc(fnIdent *ast.Ident, entClient string, structType *ast.TypeSpec,
	paramTypeName string) (*ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident,
	*ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident, *ast.Ident,
) {
	// func params
	suiteIndent := ast.NewIdent(constants.SuiteCaseVariable)
	suiteNoErrIndent := ast.NewIdent(fmt.Sprintf("%v.%v", constants.SuiteCaseVariable, constants.SuiteNoErrorFunc))
	optsIndent := ast.NewIdent("opts")
	testCaseIndent := ast.NewIdent(constants.SuiteCaseType)
	EllipsisIndent := ast.NewIdent(fmt.Sprintf("...%s", fnIdent.Name))
	// return params
	returnIndent := ast.NewIdent(fmt.Sprintf("%v.%s", entClient, structType.Name))
	// process params
	dataIndent := ast.NewIdent(fmt.Sprintf("%s{}", paramTypeName))
	optKeyIndent := ast.NewIdent("_")
	optValueIndent := ast.NewIdent("opt")
	// serializer params
	fakerIndent := ast.NewIdent(constants.FakeDataFunc)
	dataResIndent := ast.NewIdent("data")
	dataResPosIndent := ast.NewIdent("&data")
	return suiteIndent, suiteNoErrIndent, optsIndent, testCaseIndent, EllipsisIndent, returnIndent, dataIndent,
		optKeyIndent, optValueIndent, fakerIndent, dataResIndent, dataResPosIndent
}

// withTypeDef makes a type definition declaration for the functional option
// function type and adds it to astOut.
func withTypeDef(astOut *ast.File, fnIdent *ast.Ident, paramType *ast.StarExpr) {
	fnType := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: paramType,
				},
			},
		},
	}

	typeSpec := &ast.TypeSpec{
		Name: fnIdent,
		Type: fnType,
	}

	astOut.Decls = append(astOut.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			typeSpec,
		},
	})
}

// findRequestedStructType searches the input file for a struct type with name
// structName. If found, return the type spec, true; else return nil, false.
func findRequestedStructType(f *ast.File, structName string) (*ast.TypeSpec, bool) {
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		// 如果不是定义的type，就跳过
		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if _, ok := typeSpec.Type.(*ast.StructType); ok &&
				strings.EqualFold(strings.ToLower(typeSpec.Name.Name), strings.ToLower(structName)) {
				return typeSpec, true
			}
		}
	}

	return nil, false
}

// funcTypeIdent returns the identifier for the name of the functional option
// function type.
func funcTypeIdent(structName string, exportFnType bool) *ast.Ident {
	const nameF = "%sFieldSetter"
	var casedStructName string
	if exportFnType {
		casedStructName = pkg.WithFirstCharUpper(structName)
	} else {
		casedStructName = pkg.WithFirstCharLower(structName)
	}
	return ast.NewIdent(fmt.Sprintf(nameF, casedStructName))
}

// withFunc creates a functional option function for each applicable field and
// adds it to astOut.
func withFunc(astOut *ast.File, structType *ast.TypeSpec, fnIdent *ast.Ident, fnParamType *ast.StarExpr,
	generateForUnexportedFields, ignoreEmbedded, genImportFields bool, skipStructFields map[string]struct{},
) error {
	structTypeTyped, ok := structType.Type.(*ast.StructType)
	if !ok {
		panic("bad type for struct type")
	}
	var numFnsAdded int
	// Look at fields. Each entry in list is actually a list: could be embedded
	// field (length 0), "regular" field (length 1), or multiple named fields
	// with same type (length > 1).
	for _, field := range structTypeTyped.Fields.List {
		// No embedded fields
		if len(field.Names) == 0 {
			if ignoreEmbedded {
				continue
			}
			return constants.ErrDisableAllowed
		}
		// No fields whose type is imported from another package
		var fieldContainsImport bool
		ast.Inspect(field, func(n ast.Node) bool {
			_, ok := n.(*ast.SelectorExpr)
			if ok {
				fieldContainsImport = true
				return false
			}
			return true
		})
		if fieldContainsImport {
			if !genImportFields {
				continue
			}
		}
		// Now that we're operating on non-imported types and non-embedded
		// fields, let's look at each actual field name and generate a setter
		// for it.
		numFnsAdded = GenerateWithFunc(astOut, structType, fnIdent, fnParamType, generateForUnexportedFields, field,
			skipStructFields, numFnsAdded)
	}
	if numFnsAdded == 0 {
		return constants.ErrNoFiled
	}
	return nil
}

func GenerateWithFunc(astOut *ast.File, structType *ast.TypeSpec, fnIdent *ast.Ident, fnParamType *ast.StarExpr,
	generateForUnexportedFields bool, field *ast.Field, skipStructFields map[string]struct{}, numFnsAdded int,
) int {
	for _, fieldName := range field.Names {
		if _, ok := skipStructFields[fieldName.Name]; ok {
			continue
		}

		if unicode.IsLower(rune(fieldName.Name[0])) && !generateForUnexportedFields {
			continue
		}
		if pkg.WithFirstCharUpper(fieldName.Name) == constants.IgnoreField {
			continue
		}
		outerParamIdent := ast.NewIdent(pkg.WithFirstCharLower(fieldName.Name) + "Gen")
		functionName := "Set" + pkg.WithFirstCharUpper(fieldName.Name)
		newFunc := &ast.FuncDecl{
			Doc: &ast.CommentGroup{List: []*ast.Comment{
				{
					Text: fmt.Sprintf("// %s Function Optional func for %s.",
						functionName, pkg.WithFirstCharUpper(fieldName.Name)),
				},
			}},
			Name: ast.NewIdent(functionName),
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{outerParamIdent},
							Type:  field.Type,
						},
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{{Type: fnIdent}},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							getInnerFn(
								structType.Name,
								fieldName,
								outerParamIdent,
								fnParamType,
							),
						},
					},
				},
			},
		}
		astOut.Decls = append(astOut.Decls, newFunc)
		numFnsAdded++
	}
	return numFnsAdded
}

// getInnerFn returns a function literal for the inner function - the one that
// does the assignment of the struct field.
func getInnerFn(
	structTypeIdent, fieldIdent, outerParamIdent *ast.Ident,
	innerParamType *ast.StarExpr,
) *ast.FuncLit {
	paramIdent := ast.NewIdent(pkg.WithFirstCharLower(structTypeIdent.Name) + "Gen")
	return &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{paramIdent},
						Type:  innerParamType,
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.SelectorExpr{
							X:   paramIdent,
							Sel: fieldIdent,
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						outerParamIdent,
					},
				},
			},
		},
	}
}

func getFactoryReturn(returnIndent *ast.Ident, structType *ast.TypeSpec, generateForUnexportedFields, ignoreEmbedded,
	genImportFields bool, skipStructFields map[string]struct{},
) *ast.SelectorExpr {
	structTypeTyped, ok := structType.Type.(*ast.StructType)
	if !ok {
		panic("bad type for struct type")
	}
	IndentString := ""

	for _, field := range structTypeTyped.Fields.List {
		// No embedded fields
		if len(field.Names) == 0 {
			if ignoreEmbedded {
				continue
			}
			panic("embedded fields disallowed")
		}

		// No fields whose type is imported from another package
		var fieldContainsImport bool
		identName := ""
		ast.Inspect(field, func(n ast.Node) bool {
			sel, ok := n.(*ast.SelectorExpr)
			if ok {
				fieldContainsImport = true
			}
			if ok && genImportFields {
				ident, ok := sel.X.(*ast.Ident)
				if ok {
					identName = ident.Name
				}
				return false
			}
			return true
		})
		if fieldContainsImport {
			if !genImportFields {
				continue
			}
		}
		setValueParam := SetValueParam{
			generateForUnexportedFields: generateForUnexportedFields,
			fieldContainsImport:         fieldContainsImport,
			identName:                   identName,
			indentString:                IndentString,
		}
		IndentString = GenerateSetValueFunc(field, skipStructFields, setValueParam)
	}
	IndentString = "Create()" + IndentString + ".\n\tSaveX(s.Context())"
	res := ast.SelectorExpr{
		X:   returnIndent,
		Sel: ast.NewIdent(IndentString),
	}
	return &res
}

func GenerateSetValueFunc(field *ast.Field, skipStructFields map[string]struct{}, setValueParam SetValueParam) string {
	identName := setValueParam.identName
	indentString := setValueParam.indentString
	for _, fieldName := range field.Names {
		if _, ok := skipStructFields[fieldName.Name]; ok {
			continue
		}

		if unicode.IsLower(rune(fieldName.Name[0])) && !setValueParam.generateForUnexportedFields {
			continue
		}
		// set value for time do a special default
		setterStr, valueName := getSetStrAndValueName(fieldName, setValueParam.fieldContainsImport, identName)
		if pkg.WithFirstCharUpper(fieldName.Name) == constants.IgnoreField {
			continue
		}
		indentString = indentString + setterStr +
			pkg.WithFirstCharUpper(fieldName.Name) + "(" + valueName + ")"
	}
	return indentString
}

func getSetStrAndValueName(ident *ast.Ident, fieldContainsImport bool, identName string) (string, string) {
	setStr := ".\n\tSet"
	valueName := fmt.Sprintf("data.%s", ident.Name)
	if fieldContainsImport && identName == constants.ImportTime {
		valueName = "time.Now()"
	}
	if ident.Obj == nil {
		return setStr, valueName
	}
	decl, ok := ident.Obj.Decl.(*ast.Field)
	if !ok {
		return setStr, valueName
	}
	_, ok = decl.Type.(*ast.StarExpr)
	if ok {
		setStr = ".\n\tSetNillable"
		if fieldContainsImport && identName == constants.ImportTime {
			valueName = "nil"
		}
	}
	return setStr, valueName
}

func getImportDef(astOut *ast.File, structType *ast.TypeSpec, ignoreEmbedded, genImported bool, projectPath,
	factoriesPath, appPath, modelPath, schemaPath string,
) {
	structTypeTyped, ok := structType.Type.(*ast.StructType)
	if !ok {
		panic("bad type for struct type")
	}
	var importFields []string
	importFields = ImportFieldFunc(structTypeTyped, ignoreEmbedded, genImported, importFields)
	projectImportSpecs := make([]ast.Spec, 0)
	if pkg.SliceContain(importFields, constants.ImportTime) {
		spec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(constants.ImportTime),
			},
		}
		projectImportSpecs = append(projectImportSpecs, spec)
	}
	secondSpecs := secondSpecsFunc(projectPath, schemaPath, appPath, factoriesPath)

	projectImportSpecs = append(projectImportSpecs, secondSpecs...)
	endPoints := strings.Split(modelPath, "/")
	var ModelPathEndPoint string
	if len(endPoints) == 1 {
		ModelPathEndPoint = endPoints[0]
	} else {
		ModelPathEndPoint = endPoints[len(endPoints)-1]
	}
	if pkg.SliceContain(importFields, ModelPathEndPoint) {
		spec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprintf("%s/%s", projectPath, modelPath)),
			},
		}
		projectImportSpecs = append(projectImportSpecs, spec)
	}
	importDecl1 := &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: projectImportSpecs,
	}
	astOut.Decls = append([]ast.Decl{importDecl1}, astOut.Decls...)
}

func ImportFieldFunc(
	structType *ast.StructType,
	ignoreEmbedded bool,
	genImported bool,
	importFields []string,
) []string {
	var numImportFields int

	for _, field := range structType.Fields.List {
		// No embedded fields
		if len(field.Names) == 0 {
			if ignoreEmbedded {
				continue
			}
			fail(constants.ErrDisableAllowed.Error())
		}

		numImportFields += containsImportField(field, genImported, importFields)
	}

	if numImportFields == 0 || !genImported {
		return importFields
	}

	return importFields
}

func containsImportField(field *ast.Field, genImported bool, importFields []string) int {
	var containsImport int

	ast.Inspect(field, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if ok {
			containsImport = 1
		}

		if ok && genImported {
			ident, ok := sel.X.(*ast.Ident)
			if ok && !pkg.SliceContain(importFields, ident.Name) {
				importFields = append(importFields, ident.Name)
			}
			return false
		}

		return true
	})

	return containsImport
}

func secondSpecsFunc(projectPath string, schemaPath string, appPath string, factoriesPath string) []ast.Spec {
	secondSpecs := []ast.Spec{
		&ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote("github.com/bxcodec/faker"),
			},
		},
		&ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprintf("%s/%s", projectPath, schemaPath)),
			},
		},
		&ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprintf("%s/%s", projectPath, appPath)),
			},
		},
		&ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprintf("%s/%s", projectPath, factoriesPath)),
			},
		},
	}
	return secondSpecs
}
