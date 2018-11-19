/*
* CODE GENERATED AUTOMATICALLY WITH github.com/breathbath/gotainer_gen/container
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package provider

import (
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	"github.com/breathbath/gotainer_gen/container/mocks/file"
	allFileTypes "github.com/breathbath/gotainer_gen/container/mocks/file/fileTypes"
)

var (
	gotainerInitVarGeneralProvider         GeneralProvider
	isGotainerInitVarGeneralProviderInit   bool
	gotainerInitVarDbProviderAdapter       DbProviderAdapter
	isGotainerInitVarDbProviderAdapterInit bool
)

//new func section

func gotainerNewGeneralProvider(
	db db.DbProvider,
	csvfile allFileTypes.CsvFile,
	fileprovider *file.FileProvider,
	name string,
) GeneralProvider {
	return GeneralProvider{
		Db:           db,
		CsvFile:      csvfile,
		FileProvider: fileprovider,
		name:         name,
	}
}

func gotainerNewDbProviderAdapter(
	db db.DbProvider,
) DbProviderAdapter {
	return DbProviderAdapter{
		Db: db,
	}
}

//builds section

func GotainerBuildGeneralProvider() GeneralProvider {
	if !isGotainerInitVarGeneralProviderInit {
		gotainerInitVarGeneralProvider = gotainerNewGeneralProvider(
			db.GotainerBuildDbProvider(),
			allFileTypes.GotainerBuildCsvFile(),
			file.GotainerBuildFileProviderPtr(),
			GotainerBuildstring(),
		)
		isGotainerInitVarGeneralProviderInit = true
	}

	return gotainerInitVarGeneralProvider
}

func GotainerBuildDbProviderAdapter() DbProviderAdapter {
	if !isGotainerInitVarDbProviderAdapterInit {
		gotainerInitVarDbProviderAdapter = gotainerNewDbProviderAdapter(
			db.GotainerBuildDbProvider(),
		)
		isGotainerInitVarDbProviderAdapterInit = true
	}

	return gotainerInitVarDbProviderAdapter
}
