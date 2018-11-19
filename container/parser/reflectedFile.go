package parser

import "encoding/json"

type ReflectedFile struct {
	Package string
	Imports []ReflectedImport
	Structs []ReflectedStruct
}

func NewReflectedFile() ReflectedFile {
	return ReflectedFile{Imports: []ReflectedImport{}, Structs: []ReflectedStruct{}, Package: ""}
}

func (rf ReflectedFile) String() string {
	jsonStr, _ := json.Marshal(rf)
	return string(jsonStr)
}
