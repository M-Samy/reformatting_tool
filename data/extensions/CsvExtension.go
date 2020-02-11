package extensions

import (
	"os"
	"encoding/csv"
)

type CsvReaderFile struct {}

func (f CsvReaderFile) ReadFile(filePath string) (error, [][]string){
	csvFile, err := os.Open(filePath)
	if err != nil {
		return err, [][]string{}
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		return err, [][]string{}
	}

	//Skip the header row
	csvData = csvData[1:]

	return nil, csvData
}
