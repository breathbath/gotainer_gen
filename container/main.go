package main

import (
	"fmt"
	"github.com/breathbath/gotainer_gen/container/parser"
)

func main() {
	curFile := "/Users/breathbath/Projects/gotainer_gen/src/github.com/breathbath/gotainer_gen/container/mocks/mockstr.go"
	reflectedFile, err := parser.ParseFile(curFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflectedFile)
}
