package parser

type ReflectedStruct struct {
	Name   string
	Tags   []string
	Fields []ReflectedField
}

func NewReflectedStruct() ReflectedStruct {
	return ReflectedStruct{Name: "", Tags: []string{}, Fields: []ReflectedField{}}
}
