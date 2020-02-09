package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"encoding/json"

	resources "reformatting_tool/data"
)

func main() {
	csvFile, err := os.Open("./hotels.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	headersArr := make([]string, 0)
	for _, headE := range csvData[0] {
		headersArr = append(headersArr, headE)
	}

	//Skip the header row
	csvData = csvData[1:]

	var hotel resources.Hotel
	var hotels []resources.Hotel

	for _, each := range csvData {
		hotel.Name = each[0]
		hotel.Address = each[1]
		hotel.Stars, _ = strconv.Atoi(each[2])
		hotel.Contact = each[3]
		hotel.Phone = each[4]
		hotel.URL = each[5]
		hotels = append(hotels, hotel)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(hotels)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
