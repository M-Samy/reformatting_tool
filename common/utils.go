package common

import (
	"strings"
	"os"
	"fmt"
	"reformatting_tool/data/factories"
	resources "reformatting_tool/data"
)

func MultipleExtensionsWrite(validExtensions string, formattedData []resources.Hotel, writerFactory factories.WriterFactory) {
	extensions := strings.Split(validExtensions, ",")
	destinationFile := os.Getenv("DESTFILE")
	for index := range extensions {
		extension := strings.TrimSpace(extensions[index])
		extension = strings.ToUpper(extension)

		err, writerFile := writerFactory.GetWriter(extension)
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
