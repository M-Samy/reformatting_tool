package formatters

import resources "reformatting_tool/data"

type Formatter interface {
	FormatData(data interface{}) (error, []resources.Hotel)
}
