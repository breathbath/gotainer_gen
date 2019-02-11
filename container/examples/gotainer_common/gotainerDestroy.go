package gotainer_common

import "github.com/breathbath/gotainer_gen/container/examples/definition"

func Destroy() error {
	err := definition.GotainerDestroyConn()
	//all other garbage collectors come here
	return err
}