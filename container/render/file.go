package render

import (
	"fmt"
	"strings"
)

const CODE_GENERATOR_COMMENT = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/breathbath/gotainer_gen/container
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/`

type File struct {
	PackangeName string
	Funcs        []NewFunc
	Imports      map[string]Import
}

func (f *File) AddImport(i Import) {
	f.Imports[i.Alias+i.Path] = i
}

func NewFile() *File {
	return &File{PackangeName: "", Funcs: []NewFunc{}, Imports: make(map[string]Import)}
}

func (nf File) Render() string {
	newFuncStrings := []string{}
	for _, newF := range nf.Funcs {
		newFuncStrings = append(newFuncStrings, newF.Render())
	}

	importPart := ""
	if len(nf.Imports) > 0 {
		importPart += `import (
`
	}
	for _, importItem := range nf.Imports {
		importLine := fmt.Sprintf(`	"%s"`, importItem.Path)
		if importItem.Alias != "" {
			importLine = fmt.Sprintf(`	%s "%s"`,importItem.Alias, importItem.Path)
		}

		importPart += importLine + "\n"
	}

	if len(nf.Imports) > 0 {
		importPart += `)`
	}
	return fmt.Sprintf(`%s
package %s

%s

%s
`,
		CODE_GENERATOR_COMMENT,
		nf.PackangeName,
		importPart,
		strings.Join(newFuncStrings, "\n"),
	)
}
