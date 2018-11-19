package main

import (
	"github.com/breathbath/gotainer_gen/container/parser"
	"github.com/breathbath/gotainer_gen/container/render"
)

func main() {
	pathToSave := "/Users/breathbath/Projects/gotainer_gen/src/github.com/breathbath/gotainer_gen/container/mocks/provider/"
	fileToParse := pathToSave + "dataProvide.go"
	reflectedFile, err := parser.ParseFile(fileToParse)
	if err != nil {
		panic(err)
	}

	err = render.WriteFile(reflectedFile, pathToSave + "gotainerInit.go")
	if err != nil {
		panic(err)
	}
}
