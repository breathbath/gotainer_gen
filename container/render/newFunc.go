package render

const NEW_FUNC_PREFIX = "gotainerNew"
const BUILD_FUNC_PREFIX = "GotainerBuild"

type NewFunc struct {
	Ret        ReturnType
	StructName string
	Args       []FuncArgument
}
