package ext_package

var (
	gotainerStructOnePtrIsInit = false
	gotainerStructOnePtr       *StructOne

	gotainerGetStringParamsMethodResultIsInit = false
	gotainerGetStringParamsMethodResult       map[string]string
)

func GotainerBuildStructOnePtr() *StructOne {
	if !gotainerStructOnePtrIsInit {
		gotainerStructOnePtrIsInit = true
		gotainerStructOnePtr = &StructOne{}
	}
	return gotainerStructOnePtr
}

func GotainerFuncCallGetName() string {
	return name
}

func GotainerFuncCallGetAnotherName() string {
	return getAnotherName()
}

func GotainerFuncCallGetStringParams() (res map[string]string, err error) {
	if !gotainerGetStringParamsMethodResultIsInit {
		gotainerGetStringParamsMethodResultIsInit = true
		gotainerGetStringParamsMethodResult, err = GetStringParams()
	}

	return gotainerGetStringParamsMethodResult, err
}

func GotainerFuncCallGetStringParamsOne(arg1 string) (res string, err error) {
	allParams, err := GotainerFuncCallGetStringParams()
	if err != nil {
		return
	}

	return allParams[arg1], nil
}
