package extensions

import (
	"encoding/json"
	"os"
	resources "reformatting_tool/data"
)

type JsonFile struct {}

func (f JsonFile) WriteToFile(filePath string, data []resources.Hotel) (error) {
	filePath = filePath + ".json"
	// Convert to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	jsonFile.Write(jsonData)
	jsonFile.Close()

	return nil
}