package ext_package

var (
	gotainerStructOnePtrIsInit = false
	gotainerStructOnePtr       *StructOne

	gotainerGetStringParamsMethodResultIsInit = false
	gotainerGetStringParamsMethodResult       map[string]string

	gotainerGetAnotherNameFuncResultIsInit = false
	gotainerGetAnotherNameFuncResult string
)

func GotainerBuildStructOnePtr(nonCached bool) *StructOne {
	if nonCached {
		return &StructOne{}
	}

	if !gotainerStructOnePtrIsInit {
		gotainerStructOnePtrIsInit = true
		gotainerStructOnePtr = &StructOne{}
	}
	return gotainerStructOnePtr
}

func GotainerFuncCallGetName(nonCached bool) string {
	return name
}

func GotainerFuncCallGetAnotherName(nonCached bool) string {
	if nonCached {
		return getAnotherName()
	}
	if !gotainerGetAnotherNameFuncResultIsInit {
		gotainerGetAnotherNameFuncResultIsInit = true
		gotainerGetAnotherNameFuncResult = getAnotherName()
	}
	return gotainerGetAnotherNameFuncResult
}

func gotainerFuncCallGetStringParams(nonCached bool) (res map[string]string, err error) {
	if nonCached {
		return GetStringParams()
	}

	if !gotainerGetStringParamsMethodResultIsInit {
		gotainerGetStringParamsMethodResultIsInit = true
		gotainerGetStringParamsMethodResult, err = GetStringParams()
	}

	return gotainerGetStringParamsMethodResult, err
}

func GotainerFuncCallGetStringParamsOne(arg1 string, nonCached bool) (res string, err error) {
	allParams, err := gotainerFuncCallGetStringParams(nonCached)
	if err != nil {
		return
	}

	return allParams[arg1], nil
}
