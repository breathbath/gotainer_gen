package render

import (
	"bytes"
	"github.com/breathbath/gotainer_gen/container/parser"
	"os"
	"strings"
	"text/template"
)

func RenderFile(reflFile parser.ReflectedFile) string {
	fileToRender := NewFile()
	fileToRender.PackageName = reflFile.Package

	imports := map[string]parser.ReflectedImport{}
	for _, imp := range reflFile.Imports {
		nameSpace := imp.Namespace
		if imp.Alias != "" {
			nameSpace = imp.Alias
		}
		imports[nameSpace] = imp
	}

	for _, reflStruct := range reflFile.Structs {
		newFunc := NewFunc{StructName: reflStruct.Name, Ret: ReturnType{IsPointer: false}}
		for _, field := range reflStruct.Fields {
			if field.Type == "" {
				continue
			}

			arg := FuncArgument{
				Name:      field.Name,
				NameSpace: field.NameSpace,
				Type:      field.Type,
				IsPointer: field.IsPointer,
			}
			reflImport, ok := imports[field.NameSpace]
			if ok {
				arg.Import = Import{Alias: reflImport.Alias, Path: reflImport.Path}
				fileToRender.AddImport(arg.Import)
			}
			newFunc.Args = append(newFunc.Args, arg)
		}

		fileToRender.Funcs = append(fileToRender.Funcs, newFunc)
	}

	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"fup": func(input string) string {
			return strings.ToUpper(string(input[0])) + input[1:]
		},
	}
	tmpl := template.Must(template.New("constructorFile").Funcs(funcMap).Parse(TEMPLATE + REUSABLE_BLOCKS))
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, fileToRender)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func WriteFile(reflFile parser.ReflectedFile, path string) error {
	contents := RenderFile(reflFile)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(contents))
	return err
}
