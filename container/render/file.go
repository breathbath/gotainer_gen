package render

const CODE_GENERATOR_COMMENT = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/breathbath/gotainer_gen/container
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/`

const MODULE_VARS_PREFIX = "gotainerInitVar"

type File struct {
	Comment         string
	PackageName     string
	Funcs           []NewFunc
	Imports         map[string]Import
	NewFuncPrefix   string
	BuildFuncPrefix string
	ModuleVarsPrefix string
}

func (f *File) AddImport(i Import) {
	f.Imports[i.Alias+i.Path] = i
}

func NewFile() *File {
	return &File{
		Comment:       CODE_GENERATOR_COMMENT,
		PackageName:   "",
		Funcs:         []NewFunc{},
		Imports:       make(map[string]Import),
		NewFuncPrefix: NEW_FUNC_PREFIX,
		BuildFuncPrefix: BUILD_FUNC_PREFIX,
		ModuleVarsPrefix: MODULE_VARS_PREFIX,
	}
}
