package main

import (
	"os"
	"reformatting_tool/data/factories"
	validator "reformatting_tool/validations"
	utils "reformatting_tool/common"
)

func main() {

	readerFactory := factories.ReaderFactory{}
	formatterFactory := factories.FormatterFactory{}
	writerFactory := factories.WriterFactory{}

	err, readerFile := readerFactory.GetReader(os.Getenv("READER_EXTENSION"))
	validator.CheckError(err)

	err, formatter := formatterFactory.GetFormatter(os.Getenv("READER_EXTENSION"))
	validator.CheckError(err)

	err, data := readerFile.ReadFile(os.Getenv("SOURCEFILE"))
	validator.CheckError(err)

	err, formattedData := formatter.FormatData(data)
	validator.CheckError(err)

	utils.MultipleExtensionsWrite(os.Getenv("WRITER_EXTENSION"), formattedData, writerFactory)
}
