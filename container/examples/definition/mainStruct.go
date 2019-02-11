package definition

import (
	"github.com/breathbath/gotainer_gen/container/examples/ext_package"
)

func getInts() []int {
	return []int{1, 2, 3}
}

type MyStructWithDifferentReferences struct {
	//disables initialisation or if InitOnlyTaggedInput == true will be default behaviour
	DbNo ext_package.StructOne `gotainer:"none"`
	//find global var name in package ext_package
	SomeName string `gotainer:"ext_package.name"`
	//call global function getName() of ext_package package
	SomeNameFromMethod string `gotainer:"ext_package.getAnotherName()"`
	//call global function getParams of ext_package package and extract name from map key "name"
	SomeNameFromParams string `gotainer:"ext_package.GetStringParams()['name']"`
	//the same as global var will also work for constants
	SomeNameFromConst string `gotainer:"ext_package.CONST_NAME"`
	//initialises the pointer constructor in the ext_package package
	StructOnePtr *ext_package.StructOne
	//build slice of 3 interface implementations
	AllInterfaceImpls []SomeInterface `gotainer:"[]InterfaceImplOne,[]InterfaceImplTwo,[]*InterfaceImplTwo"`
	//build slice of all new SomeInterface implementations
	AllInterfaceImplsImplicit []SomeInterface `gotainer:"SomeInterface"`
	//call getInts to get integers, don't use cached version of dependency
	SomeInts []int `gotainer:"getInts(),no-cache"`
	//helps to find the correct implementation by fully qualified name, if none is specified first found is returned
	SomeStaticInterfaceVariantHere  SomeInterface `gotainer:"InterfaceImplOne"`
	//call some custom setter to let it decide which dependency to set, provide all needed dependencies for it
	SomeDynamicInterfaceVariantHere SomeInterface `gotainer:"SetDynamicInterfaceImplementation(ext_package.name, InterfaceImplOne, InterfaceImplTwo)"`
	//Conn implements Destroyable interface therefore it's destructor will appear in the common gargabe collection func gotainerDestroy
	Conn Conn
}

func SetDynamicInterfaceImplementation(
	externalScalarParam string,
	one InterfaceImplOne,
	two InterfaceImplTwo,
) SomeInterface {
	if externalScalarParam == "" {
		return one
	}

	return two
}
