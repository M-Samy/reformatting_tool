package main

import (
	"os"
	"fmt"
	"reformatting_tool/data/factories"
	"strings"
	resources "reformatting_tool/data"
)

func main() {

	readerFactory := factories.ReaderFactory{}
	formatterFactory := factories.FormatterFactory{}
	writterFactory := factories.WriterFactory{}

	err, readerFile := readerFactory.GetReader(os.Getenv("READER_EXTENSION"))
	// TODO move if error to common function
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err, formatter := formatterFactory.GetFormatter(os.Getenv("READER_EXTENSION"))
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err, data := readerFile.ReadFile(os.Getenv("SOURCEFILE"))
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err, formattedData := formatter.FormatData(data)
	if err !=  nil{
		fmt.Println(err)
		os.Exit(1)
	}

	MultipleExtensionsWrite(os.Getenv("WRITER_EXTENSION"), formattedData, writterFactory)
}

//TODO move this function outside main
func MultipleExtensionsWrite(validExtensions string, formattedData []resources.Hotel, writterFactory factories.WriterFactory) {
	extensions := strings.Split(validExtensions, ",")
	destinationFile := os.Getenv("DESTFILE")
	for index := range extensions {
		extension := strings.TrimSpace(extensions[index])
		extension = strings.ToUpper(extension)

		err, writerFile := writterFactory.GetWriter(extension)
		if err !=  nil{
			fmt.Println(err)
			continue
		}

		err = writerFile.WriteToFile(destinationFile, formattedData)

		if err !=  nil{
			fmt.Println(err)
		}
	}
}