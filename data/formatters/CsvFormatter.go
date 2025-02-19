package formatters

import (
	"strconv"
	"fmt"
	"reflect"
	resources "reformatting_tool/data"
	validator "reformatting_tool/validations"
)

type CsvFormatter struct {}

func (cs CsvFormatter) FormatData(data interface{}) (error, []resources.Hotel){
	var hotel resources.Hotel
	var hotels []resources.Hotel

	if data == nil || reflect.TypeOf(data).Kind() != reflect.TypeOf([][]string{}).Kind() {
		return fmt.Errorf("data is invalid"), []resources.Hotel{}
	}

	for _, each := range data.([][]string) {

		if !validator.Validate(each[0], each[5], each[2]) {
			continue
		}

		hotel.Name = each[0]
		hotel.Address = each[1]
		hotel.Stars, _ = strconv.ParseFloat(each[2], 64)
		hotel.Contact = each[3]
		hotel.Phone = each[4]
		hotel.URL = each[5]

		hotels = append(hotels, hotel)
	}

	return nil, hotels
}

