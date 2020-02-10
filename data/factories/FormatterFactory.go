package factories

import (
	"reformatting_tool/data/formatters"
	"fmt"
)

type FormatterFactory struct {}

func (ff FormatterFactory) GetFormatter(formatterName string) (error, formatters.Formatter) {
	if formatterName == "CSV" {
		return nil, formatters.CsvFormatter{}
	}else {
		return fmt.Errorf("formatter type not suported"), nil
	}
}