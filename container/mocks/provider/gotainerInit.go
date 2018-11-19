/*
* CODE GENERATED AUTOMATICALLY WITH github.com/breathbath/gotainer_gen/container
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package provider

import (
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	allFileTypes "github.com/breathbath/gotainer_gen/container/mocks/file/fileTypes"
	"github.com/breathbath/gotainer_gen/container/mocks/file"
)

func GotainerNewGeneralProvider(db db.DbProvider, csvfile allFileTypes.CsvFile, fileprovider *file.FileProvider, name string) GeneralProvider {
   return GeneralProvider{}
}

func GotainerNewDbProviderAdapter(db db.DbProvider) DbProviderAdapter {
   return DbProviderAdapter{}
}

