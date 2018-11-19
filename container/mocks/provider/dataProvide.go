package provider

import (
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	"github.com/breathbath/gotainer_gen/container/mocks/file"
	allFileTypes "github.com/breathbath/gotainer_gen/container/mocks/file/fileTypes"
)

//someComment
var l string = "1"

//papa

//@gotainer=check
type GeneralProvider struct {
	Db           db.DbProvider
	CsvFile      allFileTypes.CsvFile
	FileProvider *file.FileProvider
	name         string
	Tags         []string `@gotainer:"tags"`
}

/**
* @gotainer={autowire=true}
*/
type DbProviderAdapter struct {
	Db db.DbProvider
}
