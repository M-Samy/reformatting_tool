package factories

import (
	ex "reformatting_tool/data/extensions"
	"fmt"
)

type WriterFactory struct {}

func (wf WriterFactory) GetWriter(writerName string) (error, ex.WriterExtension) {
	if writerName == "JSON" {
		return nil, ex.JsonFile{}
	}else {
		return fmt.Errorf("writter type: %s not supported", writerName), nil
	}
}