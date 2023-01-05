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
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

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

//nolint:gocognit,cyclop // fix it later
func GenerateFactories(cmd *cobra.Command, _ []string) {
	schemaFile, outputPath, schemaPath, projectPath, factoriesPath, appPath, entClientName, overWrite,
		err := ExtraFlags(cmd)
	if err != nil {
		fail(err.Error())
	}
	if schemaFile == "" && schemaPath == "" {
		Fatalf("schema file and schema path must give at lease one")
	}
	if outputPath == "" {
		Fatalf("output path cannot be empty")
	}
	f, err := os.Open(schemaPath)
	if err != nil {
		fail(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fail(err.Error())
		}
	}(f)
	commonPath := fmt.Sprintf("%s/common.go", outputPath)
	_, err = os.Stat(commonPath)
	if err != nil {
		if os.IsNotExist(err) {
			CreatePathAndCommonFile(commonPath, outputPath)
		} else {
			fail(err.Error())
		}
	}
	if schemaFile != "" && schemaPath == "" {
		schemaName := ExtraNameFromSchemaFilePath(schemaFile)
		realOutPutPath := fmt.Sprintf("%s/%sfactory/%sfactory.go", outputPath, schemaName, schemaName)
		CreateOneFactory(schemaFile, schemaName, realOutPutPath, projectPath, factoriesPath, appPath, entClientName,
			outputPath)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fail(err.Error())
	}
	for _, v := range files {
		if !v.IsDir() {
			continue
		}
		isContinue := false
		for _, n := range constants.IgnoreFolderNames {
			if v.Name() == n {
				isContinue = true
			}
		}
		if isContinue {
			continue
		}
		realPath, realOutPutPath := GetRealPathAndFilePath(schemaPath, v, outputPath)
		if !overWrite {
			//nolint:gocritic // no way to rewrite, because of the os.IsNotExist is not a case
			if _, err := os.Stat(realOutPutPath); err == nil {
				continue
			} else if os.IsNotExist(err) {
				CreateOneFactory(realPath, v.Name(), realOutPutPath, projectPath, factoriesPath, appPath, entClientName,
					outputPath)
			} else {
				fail(fmt.Sprintf("Error occurred while checking file existence: %s", realOutPutPath))
			}
		} else {
			CreateOneFactory(realPath, v.Name(), realOutPutPath, projectPath, factoriesPath, appPath, entClientName, outputPath)
		}
	}
}

func GetRealPathAndFilePath(schemaPath string, v os.FileInfo, outputPath string) (string, string) {
	filePath := fmt.Sprintf("%s/%s", schemaPath, v.Name())
	realPath := fmt.Sprintf("%s.go", filePath)
	realOutPutPath := fmt.Sprintf("%s/%sfactory/%sfactory.go", outputPath, v.Name(), v.Name())
	return realPath, realOutPutPath
}

func ExtraNameFromSchemaFilePath(schemaFile string) string {
	endPoints := strings.Split(schemaFile, "/")
	schemaFileName := endPoints[len(endPoints)-1]
	schemaNames := strings.Split(schemaFileName, ".")
	schemaName := schemaNames[0]
	return schemaName
}

func CreateOneFactory(realPath, schemaName, realOutPutPath, projectPath, factoriesPath, appPath, entClientName,
	outputPath string,
) {
	outReader, err := RunGenerate(realPath, schemaName, realOutPutPath, projectPath, factoriesPath, appPath, entClientName)
	if err != nil {
		fail(err.Error())
	}
	var dest io.Writer
	//nolint:nestif // need it
	if outputPath == "" {
		dest = os.Stdout
	} else {
		_, err := os.Stat(filepath.Dir(realOutPutPath))
		if os.IsNotExist(err) {
			err2 := os.MkdirAll(filepath.Dir(realOutPutPath), os.FileMode(constants.Perm))
			if err2 != nil {
				fail(err2.Error())
			}
		} else if err != nil {
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

func ExtraFlags(cmd *cobra.Command) (string, string, string, string, string, string, string, bool, error) {
	schemaFile, err := cmd.Flags().GetString("schemaFile")
	if err != nil {
		Fatalf("get schema file failed: %v\n", err)
	}
	schemaPath, err := cmd.Flags().GetString("schemaPath")
	if err != nil {
		Fatalf("get schema path failed: %v\n", err)
	}
	outputPath, err := cmd.Flags().GetString("outputPath")
	if err != nil {
		Fatalf("get output path failed: %v\n", err)
	}
	projectPath, err := cmd.Flags().GetString("projectPath")
	if err != nil {
		Fatalf("get project path failed: %v\n", err)
	}
	if projectPath == "" {
		Fatalf("project path cannot be empty")
	}
	overWrite, err := cmd.Flags().GetBool("overwrite")
	if err != nil {
		Fatalf("overwrite setting cannot be empty")
	}

	factoriesPath, err := cmd.Flags().GetString("factoriesPath")
	if err != nil {
		Fatalf("get factories path failed: %v\n", err)
	}
	if factoriesPath == "" {
		factoriesPath = constants.DefaultFactoryPath
	}

	appPath, err := cmd.Flags().GetString("appPath")
	if err != nil {
		Fatalf("get app path failed: %v\n", err)
	}
	if appPath == "" {
		appPath = constants.DefaultAppPath
	}

	entClientName, err := cmd.Flags().GetString("entClientName")
	if err != nil {
		Fatalf("get ent client path failed: %v\n", err)
	}
	if entClientName == "" {
		entClientName = constants.DefaultEntClientName
	}
	entClientName = fmt.Sprintf("app.%s", entClientName)
	return schemaFile, outputPath, schemaPath, projectPath, factoriesPath, appPath, entClientName, overWrite, err
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
func RunGenerate(schemaFile, schemaTypeName, outputPath, projectPath, factoriesPath, appPath, entClientName string) (
	io.Reader, error,
) {
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
	getImportDef(astOut, projectPath, factoriesPath, appPath)

	// Add type definition for functional option function signature
	withTypeDef(astOut, fnTypeIdent, fnParamType)

	// Add function for each applicable struct field
	if err := withFunc(
		astOut,
		structType,
		fnTypeIdent,
		fnParamType,
		false,
		true,
		constants.SkipStructFields); err != nil {
		return nil, err
	}

	if err := NewFunc(
		astOut,
		paramTypeName,
		structType,
		fnTypeIdent,
		fnParamType,
		false,
		true,
		constants.SkipStructFields,
		entClientName,
	); err != nil {
		return nil, err
	}
	// Generate output file
	out := new(bytes.Buffer)
	if err := printer.Fprint(out, token.NewFileSet(), astOut); err != nil {
		return nil, err
	}
	res, err := formatCode(out)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewFunc Create the New instance function.
func NewFunc(astOut *ast.File,
	paramTypeName string,
	structType *ast.TypeSpec,
	fnIdent *ast.Ident,
	fnParamType *ast.StarExpr,
	generateForUnexportedFields, ignoreUnsupported bool,
	skipStructFields map[string]struct{},
	entClient string,
) error {
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
	newFunc := &ast.FuncDecl{
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
		casedStructName = withFirstCharUpper(structName)
	} else {
		casedStructName = withFirstCharLower(structName)
	}
	return ast.NewIdent(fmt.Sprintf(nameF, casedStructName))
}

// withFunc creates a functional option function for each applicable field and
// adds it to astOut.
func withFunc(
	astOut *ast.File,
	structType *ast.TypeSpec,
	fnIdent *ast.Ident,
	fnParamType *ast.StarExpr,
	generateForUnexportedFields, ignoreUnsupported bool,
	skipStructFields map[string]struct{},
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
			if ignoreUnsupported {
				continue
			} else {
				return constants.ErrDisableAllowed
			}
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
			if ignoreUnsupported {
				continue
			} else {
				return constants.ErrCanNotGen
			}
		}

		// Now that we're operating on non-imported types and non-embedded
		// fields, let's look at each actual field name and generate a setter
		// for it.
		for _, fieldName := range field.Names {
			if _, ok := skipStructFields[fieldName.Name]; ok {
				continue
			}

			if unicode.IsLower(rune(fieldName.Name[0])) && !generateForUnexportedFields {
				continue
			}

			outerParamIdent := ast.NewIdent(withFirstCharLower(fieldName.Name) + "Gen")
			newFunc := &ast.FuncDecl{
				Name: ast.NewIdent("Set" + withFirstCharUpper(fieldName.Name)),
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
	}

	if numFnsAdded == 0 {
		return constants.ErrNoFiled
	}

	return nil
}

// getInnerFn returns a function literal for the inner function - the one that
// does the assignment of the struct field.
func getInnerFn(
	structTypeIdent, fieldIdent, outerParamIdent *ast.Ident,
	innerParamType *ast.StarExpr,
) *ast.FuncLit {
	paramIdent := ast.NewIdent(withFirstCharLower(structTypeIdent.Name) + "Gen")
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

func getFactoryReturn(returnIndent *ast.Ident, structType *ast.TypeSpec, generateForUnexportedFields,
	ignoreUnsupported bool, skipStructFields map[string]struct{},
) *ast.SelectorExpr {
	structTypeTyped, ok := structType.Type.(*ast.StructType)
	if !ok {
		panic("bad type for struct type")
	}
	IndentString := ""

	for _, field := range structTypeTyped.Fields.List {
		// No embedded fields
		if len(field.Names) == 0 {
			if ignoreUnsupported {
				continue
			} else {
				panic("embedded fields disallowed")
			}
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
			if ignoreUnsupported {
				continue
			} else {
				panic("cannot generate for fields whose type is imported")
			}
		}

		// Now that we're operating on non-imported types and non-embedded
		// fields, let's look at each actual field name and generate a setter
		// for it.

		for _, fieldName := range field.Names {
			if _, ok := skipStructFields[fieldName.Name]; ok {
				continue
			}

			if unicode.IsLower(rune(fieldName.Name[0])) && !generateForUnexportedFields {
				continue
			}
			IndentString = IndentString + ".\n\tSet" + withFirstCharUpper(fieldName.Name) + "(data." + fieldName.Name + ")"
		}
	}
	IndentString = "Create()" + IndentString + ".\n\tSaveX(s.Context())"
	res := ast.SelectorExpr{
		X:   returnIndent,
		Sel: ast.NewIdent(IndentString),
	}
	return &res
}

func getImportDef(astOut *ast.File, projectPath, factoriesPath, appPath string) {
	importDecl1 := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("github.com/bxcodec/faker"),
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(fmt.Sprintf("%s/gen/entschema", projectPath)),
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
		},
	}
	astOut.Decls = append([]ast.Decl{importDecl1}, astOut.Decls...)
}

func withFirstCharLower(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[0:1]) + s[1:]
}

func withFirstCharUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

func formatCode(buf *bytes.Buffer) (*bytes.Buffer, error) {
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
