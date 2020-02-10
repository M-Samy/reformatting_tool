package factories

import (
	ex "reformatting_tool/data/extensions"
	"fmt"
)

type ReaderFactory struct {}

func (rf ReaderFactory) GetReader(readerName string) (error, ex.ReaderExtension) {
	if readerName == "CSV" {
		return nil, ex.CsvReaderFile{}
	}else {
		return fmt.Errorf("reader type not suported"), nil
	}
}