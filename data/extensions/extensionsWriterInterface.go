package extensions

import resources "reformatting_tool/data"

type WriterExtension interface {
	WriteToFile(filePath string, data []resources.Hotel) error
}
