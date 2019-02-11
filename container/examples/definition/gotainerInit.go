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

	gotainerConnIsInit = false
	gotainerConn       Conn
)

func GotainerBuildMyStructWithDifferentReferences(nonCached bool) (res *MyStructWithDifferentReferences, err error) {
	if !gotainerMyStructWithDifferentReferencesPtrIsInit {
		gotainerMyStructWithDifferentReferencesPtrIsInit = true
		someNameFromParams, err := ext_package.GotainerFuncCallGetStringParamsOne("name", nonCached)
		if err != nil {
			return
		}
		gotainerMyStructWithDifferentReferencesPtr = &MyStructWithDifferentReferences{
			SomeName:           ext_package.GotainerFuncCallGetName(nonCached),        //GotainerFuncCallGetName is generated and made the protected variable exposed to the outside
			SomeNameFromMethod: ext_package.GotainerFuncCallGetAnotherName(nonCached), //getAnotherName is protected so GotainerFuncCallGetAnotherName is generated to expose it
			SomeNameFromParams: someNameFromParams,
			SomeNameFromConst:  ext_package.CONST_NAME,
			StructOnePtr:       ext_package.GotainerBuildStructOnePtr(nonCached),
			AllInterfaceImpls: []SomeInterface{
				GotainerBuildInterfaceImplOne(nonCached),
				GotainerBuildInterfaceImplTwo(nonCached),
				GotainerBuildInterfaceImplTwoPtr(nonCached),
			},
			SomeInts:                       getInts(), //we don't cache here so no indirect methods are generated
			SomeStaticInterfaceVariantHere: GotainerBuildInterfaceImplOne(nonCached),
			SomeDynamicInterfaceVariantHere: SetDynamicInterfaceImplementation(
				ext_package.GotainerFuncCallGetName(nonCached),
				GotainerBuildInterfaceImplOne(nonCached),
				GotainerBuildInterfaceImplTwo(nonCached),
			),
			AllInterfaceImplsImplicit: []SomeInterface{
				GotainerBuildInterfaceImplOne(nonCached),
				GotainerBuildInterfaceImplTwo(nonCached),
			},
			Conn: GotainerBuildConn(nonCached),
		}
	}

	return gotainerMyStructWithDifferentReferencesPtr, nil
}

func GotainerBuildInterfaceImplOne(nonCached bool) InterfaceImplOne {
	if nonCached {
		return InterfaceImplOne{}
	}

	if !gotainerInterfaceImplOneIsInit {
		gotainerInterfaceImplOneIsInit = true
		gotainerInterfaceImplOne = InterfaceImplOne{}
	}
	return gotainerInterfaceImplOne
}

func GotainerBuildInterfaceImplTwo(nonCached bool) InterfaceImplTwo {
	if nonCached {
		return InterfaceImplTwo{}
	}

	if !gotainerInterfaceImplTwoIsInit {
		gotainerInterfaceImplTwoIsInit = true
		gotainerInterfaceImplTwo = InterfaceImplTwo{}
	}
	return gotainerInterfaceImplTwo
}

func GotainerBuildInterfaceImplTwoPtr(nonCached bool) *InterfaceImplTwo {
	if nonCached {
		return &InterfaceImplTwo{}
	}

	if !gotainerInterfaceImplTwoPtrIsInit {
		gotainerInterfaceImplTwoPtrIsInit = true
		gotainerInterfaceImplTwoPtr = &InterfaceImplTwo{}
	}
	return gotainerInterfaceImplTwoPtr
}
func GotainerBuildConn(nonCached bool) Conn {
	if nonCached {
		return Conn{}
	}

	if !gotainerConnIsInit {
		gotainerConnIsInit = true
		gotainerConn = Conn{}
	}
	return gotainerConn
}

func GotainerDestroyConn() error {
	if gotainerConnIsInit {
		return gotainerConn.Destroy()
	}
	return nil
}