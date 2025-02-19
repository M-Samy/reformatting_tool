package factories

import (
	"fmt"
	ex "reformatting_tool/data/extensions"
)

type WriterFactory struct {}

func (wf WriterFactory) GetWriter(writerName string) (error, ex.WriterExtension) {
	if writerName == "JSON" {
		return nil, ex.JsonFile{}
	}else if writerName == "XML" {
		return nil, ex.XmlFile{}
	} else {
		return fmt.Errorf("writter type: %s not supported", writerName), nil
	}
}