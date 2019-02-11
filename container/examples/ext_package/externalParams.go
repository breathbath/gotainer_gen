package ext_package

func GetStringParams() (map[string]string, error) {
	return map[string]string{
		"some_param": "some_value",
	}, nil
}
