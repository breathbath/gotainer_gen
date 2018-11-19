package file

func GotainerNewFileProvider() FileProvider{
	return FileProvider{}
}

func GotainerNewFileProviderPtr() *FileProvider{
	return &FileProvider{}
}

func GotainerBuildFileProviderPtr() *FileProvider {
	return GotainerNewFileProviderPtr()
}