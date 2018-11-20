package provider

import (
	"github.com/breathbath/gotainer_gen/container/mocks/db"
	"github.com/breathbath/gotainer_gen/container/mocks/file"
	allFileTypes "github.com/breathbath/gotainer_gen/container/mocks/file/fileTypes"
	"time"
)

//someComment
var l string = "1"

//papa

//@gotainer=check
type GeneralProvider struct {
	Db                 db.DbProvider        //Type=SelectorExpr, Type.X(Expr|Ident).Name=db, Type.Sel(Ident).Name=DbProvider, Names(Ident)=Db
	CsvFileField       allFileTypes.CsvFile //Type=SelectorExpr, Type.X(Expr|Ident).Name="CsvFile", Type.Sel(Ident).Name="CsvFile", Names(Ident)=CsvFileField
	DbProvideAdapt     DbProviderAdapter
	FileProvider       *file.FileProvider   //Type=StarExpr, Type.X(Expr|SelectorExpr.X(Expr|Ident).Name=file|SelectorExpr.Sel.Name=FileProvider), Names(Ident)=FileProvider
	name               string               //Type=Expr|Ident,Type(Ident).Name=string, Names=name
	StringPtr          *string              //Type=StarExpr, Type.X(Expr|Ident).Name=string, Names=StringPtr
	TimeField          time.Time
	Duration           time.Duration
	CustomScalarInt    MyInt
	IntField           int
	Int64Field         int64
	Float64Field       float64
	BoolField          bool
	UintField          uint
	MapStringInt       map[string]int
	StructInt          []int
	MapStringCsvFile   map[string]allFileTypes.CsvFile
	StructFileProvider []file.FileProvider
	Tags               []string `@gotainer:"tags"`
}

/**
* @gotainer={autowire=true}
*/
type DbProviderAdapter struct {
	Db db.DbProvider
}
