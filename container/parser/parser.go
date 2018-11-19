package parser

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
)

const GOTAINER_COMMENT_TAG = "@gotainer"

func ParseFile(filePath string) (ReflectedFile, error) {
	reflFile := NewReflectedFile()

	astFileSet := token.NewFileSet()

	reflImports, err := parseImports(astFileSet, filePath)
	if err != nil {
		panic(err)
	}
	reflFile.Imports = reflImports

	astParsedFile, err := parser.ParseFile(astFileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return reflFile, err
	}

	packageName, err := parsePackageName(filePath, astParsedFile, astFileSet)
	if err != nil {
		return reflFile, err
	}
	reflFile.Package = packageName

	reflFile.Structs = parsStructs(astParsedFile)

	return reflFile, nil
}

func parseImports(astFileSet *token.FileSet, filePath string) ([]ReflectedImport, error) {
	relfImports := []ReflectedImport{}

	astFileImports, err := parser.ParseFile(astFileSet, filePath, nil, parser.ImportsOnly)
	if err != nil {
		return relfImports, err
	}

	for _, astImport := range astFileImports.Imports {
		if astImport.Path == nil {
			continue
		}
		reflectedImport := ReflectedImport{Path: astImport.Path.Value, Alias: "", Namespace: ""}
		if astImport.Name != nil {
			reflectedImport.Alias = astImport.Name.Name
		}

		pathParts := strings.Split(reflectedImport.Path, "/")
		reflectedImport.Namespace = pathParts[len(pathParts)-1]

		relfImports = append(relfImports, reflectedImport)
	}

	return relfImports, nil
}

func parsePackageName(filePath string, astParsedFile *ast.File, astFileSet *token.FileSet) (string, error) {
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check(filePath, astFileSet, []*ast.File{astParsedFile}, nil)
	if err != nil {
		return "", err
	}

	return pkg.Name(), nil
}

func parsStructs(astParsedFile *ast.File) ([]ReflectedStruct) {
	reflectedStructs := []ReflectedStruct{}
	comments := []string{}
	ast.Inspect(astParsedFile, func(node ast.Node) bool {
		return parseNode(node, &comments, &reflectedStructs)
	})

	return reflectedStructs
}

func parseNode(node ast.Node, comments *[]string, reflectedStructs *[]ReflectedStruct) bool {
	switch possibleStructNode := node.(type) {
	case *ast.CommentGroup:
		for _, cc := range possibleStructNode.List {
			comment := cc.Text
			if strings.Contains(comment, GOTAINER_COMMENT_TAG) {
				*comments = append(*comments, cc.Text)
			}
		}
	case *ast.TypeSpec:
		switch possibleStructNode.Type.(type) {
		case *ast.StructType:
			reflectedStruct := NewReflectedStruct()
			reflectedStruct.Name = possibleStructNode.Name.Name
			if len(*comments) > 0 {
				reflectedStruct.Tags = *comments
				*comments = []string{}
			}

			reflectedStruct.Fields = parseStructFields(possibleStructNode.Type.(*ast.StructType))

			*reflectedStructs = append(*reflectedStructs, reflectedStruct)
		}
	}

	return true
}

func parseStructFields(astStruct *ast.StructType) []ReflectedField {
	reflFields := []ReflectedField{}
	rootStructFields := astStruct.Fields.List
	for _, rootStructField := range rootStructFields {
		reflectedField := ReflectedField{}

		if rootStructField.Tag != nil && strings.Contains(rootStructField.Tag.Value, GOTAINER_COMMENT_TAG) {
			reflectedField.Tag = rootStructField.Tag.Value
		}

		if len(rootStructField.Names) > 0 {
			reflectedField.Name = rootStructField.Names[0].Name
		}

		switch rootStructField.Type.(type) {
		case *ast.SelectorExpr:
			selectorExpr := rootStructField.Type.(*ast.SelectorExpr)

			reflectedField.Type = selectorExpr.Sel.Name

			x, ok := selectorExpr.X.(*ast.Ident)
			if ok {
				reflectedField.NameSpace = x.Name
			}
		case *ast.Ident:
			reflectedField.Type = rootStructField.Type.(*ast.Ident).Name
		}

		reflFields = append(reflFields, reflectedField)
	}

	return reflFields
}