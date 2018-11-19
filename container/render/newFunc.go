package render

import (
	"fmt"
	"strings"
)

type ReturnTypeDecl struct {
	Name      string
	IsPointer bool
}

type File struct {
	PackangeName string
	Funcs        []NewFunc
}

func (nf File) Render() string {
	newFuncStrings := []string{}
	for _, newF := range nf.Funcs {
		newFuncStrings = append(newFuncStrings, newF.Render())
	}
	return fmt.Sprintf(`package %s

%s
`, nf.PackangeName, strings.Join(newFuncStrings, "\n"))
}

type NewFunc struct {
	Ret        ReturnTypeDecl
	StructName string
}

func (nf NewFunc) Render() string {
	pointerPrefix, pointerDeref := "", ""
	if nf.Ret.IsPointer {
		pointerPrefix = "*"
		pointerDeref = "&"
	}
	return fmt.Sprintf(`func GotainerNew%s() %s {
   return %s{}
}
`,
		nf.StructName,
		pointerPrefix+nf.StructName,
		pointerDeref+nf.StructName,
	)
}
