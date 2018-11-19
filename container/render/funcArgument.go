package render

import (
	"fmt"
	"strings"
)

type FuncArgument struct {
	Name      string
	NameSpace string
	Type      string
	IsPointer bool
	Import    Import
}

func (fa FuncArgument) Render() string {
	typeDeclaration := fa.Type
	if fa.NameSpace != "" {
		typeDeclaration = fa.NameSpace + "." + typeDeclaration
	}
	if fa.IsPointer {
		typeDeclaration = "*" + typeDeclaration
	}

	return fmt.Sprintf("%s %s", strings.ToLower(fa.Name), typeDeclaration)
}
