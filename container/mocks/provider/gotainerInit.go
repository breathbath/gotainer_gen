/*
* CODE GENERATED AUTOMATICALLY WITH github.com/breathbath/gotainer_gen/container
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package provider

import (
	allFileTypes "github.com/breathbath/gotainer_gen/container/mocks/file/fileTypes"
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	"github.com/breathbath/gotainer_gen/container/mocks/file"
	"time"
)

var (
	gotainerInitVarGeneralProvider GeneralProvider
	isGotainerInitVarGeneralProviderInit bool
	gotainerInitVarDbProviderAdapter DbProviderAdapter
	isGotainerInitVarDbProviderAdapterInit bool
)

//new func section

func gotainerNewGeneralProvider(
	db db.DbProvider,
	csvfilefield allFileTypes.CsvFile,
	dbprovideadapt DbProviderAdapter,
	fileprovider *file.FileProvider,
	name string,
	stringptr *string,
	timefield time.Time,
	duration time.Duration,
	customscalarint MyInt,
	intfield int,
	int64field int64,
	float64field float64,
	boolfield bool,
	uintfield uint,
) GeneralProvider {
	return GeneralProvider{
		Db: db,
		CsvFileField: csvfilefield,
		DbProvideAdapt: dbprovideadapt,
		FileProvider: fileprovider,
		name: name,
		StringPtr: stringptr,
		TimeField: timefield,
		Duration: duration,
		CustomScalarInt: customscalarint,
		IntField: intfield,
		Int64Field: int64field,
		Float64Field: float64field,
		BoolField: boolfield,
		UintField: uintfield,
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
			GotainerBuildDbProviderAdapter(),
			file.GotainerBuildFileProviderPtr(),
			GotainerBuildstring(),
			GotainerBuildstringPtr(),
			time.GotainerBuildTime(),
			time.GotainerBuildDuration(),
			GotainerBuildMyInt(),
			GotainerBuildint(),
			GotainerBuildint64(),
			GotainerBuildfloat64(),
			GotainerBuildbool(),
			GotainerBuilduint(),
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



