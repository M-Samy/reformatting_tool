package extensions

type ReaderExtension interface {
	ReadFile(filePath string) (error, [][]string)
}
