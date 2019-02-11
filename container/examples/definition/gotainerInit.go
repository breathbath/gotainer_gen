package definition

import "github.com/breathbath/gotainer_gen/container/examples/ext_package"

var (
	gotainerInterfaceImplOneIsInit = false
	gotainerInterfaceImplOne       InterfaceImplOne

	gotainerInterfaceImplTwoIsInit = false
	gotainerInterfaceImplTwo       InterfaceImplTwo

	gotainerInterfaceImplTwoPtrIsInit = false
	gotainerInterfaceImplTwoPtr       *InterfaceImplTwo

	gotainerMyStructWithDifferentReferencesPtrIsInit = false
	gotainerMyStructWithDifferentReferencesPtr       *MyStructWithDifferentReferences
)

func GotainerBuildMyStructWithDifferentReferences(nonCached bool) (res *MyStructWithDifferentReferences, err error) {
	if !gotainerMyStructWithDifferentReferencesPtrIsInit {
		gotainerMyStructWithDifferentReferencesPtrIsInit = true
		someNameFromParams, err := ext_package.GotainerFuncCallGetStringParamsOne("name")
		if err != nil {
			return
		}
		gotainerMyStructWithDifferentReferencesPtr = &MyStructWithDifferentReferences{
			SomeName:           ext_package.GotainerFuncCallGetName(),        //GotainerFuncCallGetName is generated and made the protected variable exposed to the outside
			SomeNameFromMethod: ext_package.GotainerFuncCallGetAnotherName(), //getAnotherName is protected so GotainerFuncCallGetAnotherName is generated to expose it
			SomeNameFromParams: someNameFromParams,
			SomeNameFromConst:  ext_package.CONST_NAME,
			StructOnePtr:       ext_package.GotainerBuildStructOnePtr(),
			AllInterfaceImpls: []SomeInterface{
				GotainerBuildInterfaceImplOne(),
				GotainerBuildInterfaceImplTwo(),
				GotainerBuildInterfaceImplTwoPtr(),
			},
			SomeInts:                       getInts(), //we don't cache here so no indirect methods are generated
			SomeStaticInterfaceVariantHere: GotainerBuildInterfaceImplOne(),
			SomeDynamicInterfaceVariantHere: SetDynamicInterfaceImplementation(
				ext_package.GotainerFuncCallGetName(),
				GotainerBuildInterfaceImplOne(),
				GotainerBuildInterfaceImplTwo(),
			),
		}
	}

	return gotainerMyStructWithDifferentReferencesPtr, nil
}

func GotainerBuildInterfaceImplOne() InterfaceImplOne {
	if !gotainerInterfaceImplOneIsInit {
		gotainerInterfaceImplOneIsInit = true
		gotainerInterfaceImplOne = InterfaceImplOne{}
	}
	return gotainerInterfaceImplOne
}

func GotainerBuildInterfaceImplTwo() InterfaceImplTwo {
	if !gotainerInterfaceImplTwoIsInit {
		gotainerInterfaceImplTwoIsInit = true
		gotainerInterfaceImplTwo = InterfaceImplTwo{}
	}
	return gotainerInterfaceImplTwo
}

func GotainerBuildInterfaceImplTwoPtr() *InterfaceImplTwo {
	if !gotainerInterfaceImplTwoPtrIsInit {
		gotainerInterfaceImplTwoPtrIsInit = true
		gotainerInterfaceImplTwoPtr = &InterfaceImplTwo{}
	}
	return gotainerInterfaceImplTwoPtr
}
