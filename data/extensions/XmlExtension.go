package extensions

import (
	"encoding/xml"
	"io/ioutil"
	resources "reformatting_tool/data"
)

type XmlFile struct {}

func (xf XmlFile) WriteToFile(filePath string, data []resources.Hotel) (error) {
	filePath = filePath + ".xml"
	// Convert to XML
	file, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	_ = ioutil.WriteFile(filePath, file, 0644)

	return nil
}