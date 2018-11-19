package db

func GotainerNewDbProvider() DbProvider{
	return DbProvider{}
}

func GotainerBuildDbProvider() DbProvider {
	return GotainerNewDbProvider()
}
