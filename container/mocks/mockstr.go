package mocks

import (
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	sss "github.com/breathbath/gotainer_gen/container/mocks/db/subdb"
)

//someComment
var l string = "1"

//papa

//@gotainer=lala
type MockStr struct {
	ADbField    db.Db
	aSubDbField sss.SubDb
	name        string
	Ints        []int `gotainer:"intes"`
}

/**
* @gotainer={autowire=true}
*/
type AnotherStruct struct {
	lala MockStr
}
