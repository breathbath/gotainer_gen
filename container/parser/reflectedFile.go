package parser


type ReflectedFile struct {
	Package string
	Imports []ReflectedImport
	Structs []ReflectedStruct
}

func NewReflectedFile() ReflectedFile {
	return ReflectedFile{Imports: []ReflectedImport{}, Structs: []ReflectedStruct{}, Package: ""}
}
