package render

const REUSABLE_BLOCKS = `
{{ define "InitFlagName"}}is{{fup .}}Init{{end}}
`
const TEMPLATE = `{{.Comment}}
package {{.PackageName}}

import (
{{- range $importItem := .Imports}}
	{{if $importItem.Alias}}{{$importItem.Alias}} {{end}}"{{$importItem.Path}}"
{{- end}}
)

var (
{{- $moduleVarsPrefix := .ModuleVarsPrefix}}
{{- range $initFunc := .Funcs}}
	{{$moduleVarsPrefix}}{{$initFunc.StructName}} {{if $initFunc.Ret.IsPointer}}*{{end}}{{$initFunc.StructName}}
	{{template "InitFlagName" print $moduleVarsPrefix $initFunc.StructName}} bool
{{- end}}
)

//new func section
{{- $newFuncPrefix := .NewFuncPrefix}}
{{range $func := .Funcs}}
func {{$newFuncPrefix}}{{$func.StructName}}(
{{- range $arg := $func.Args}}
	{{lower $arg.Name}} {{if $arg.IsPointer}}*{{end}}{{if $arg.NameSpace}}{{$arg.NameSpace}}.{{end}}{{$arg.Type}},
{{- end}}
) {{if $func.Ret.IsPointer}}*{{end}}{{$func.StructName}} {
	return {{$func.StructName}}{
{{- range $arg := $func.Args}}
		{{$arg.Name}}: {{lower $arg.Name}},
{{- end}}
	}
}
{{end}}
//builds section
{{ $buildFuncPrefix := .BuildFuncPrefix}}
{{- range $func := .Funcs}}
func {{$buildFuncPrefix}}{{$func.StructName}}() {{$func.StructName}} {
	if !{{template "InitFlagName" print $moduleVarsPrefix $func.StructName}} {
		{{$moduleVarsPrefix}}{{$func.StructName}} = {{$newFuncPrefix}}{{$func.StructName}}(
         {{- range $arg := $func.Args}}
			{{ if $arg.NameSpace}}{{$arg.NameSpace}}.{{end}}
			 {{- $buildFuncPrefix}}{{$arg.Type}}
			 {{- if $arg.IsPointer}}Ptr{{- end}}(),
         {{- end}}
		)
		{{template "InitFlagName" print $moduleVarsPrefix $func.StructName}} = true
    }

	return {{$moduleVarsPrefix}}{{$func.StructName}}
}
{{end}}
`
