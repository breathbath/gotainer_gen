package render

import (
	"fmt"
	"strings"
)

const NEW_FUNC_PREFIX = "GotainerNew"

type NewFunc struct {
	Ret        ReturnType
	StructName string
	Args       []FuncArgument
}

func (nf NewFunc) Render() string {
	pointerPrefix, pointerDeref := "", ""
	if nf.Ret.IsPointer {
		pointerPrefix = "*"
		pointerDeref = "&"
	}
	argsStr := ""
	for _, arg := range nf.Args {
		argsStr += arg.Render() + ", "
	}
	argsStr = strings.TrimRight(argsStr, ", ")

	return fmt.Sprintf(`func %s%s(%s) %s {
   return %s{}
}
`,
		NEW_FUNC_PREFIX,
		nf.StructName,
		argsStr,
		pointerPrefix+nf.StructName,
		pointerDeref+nf.StructName,
	)
}
