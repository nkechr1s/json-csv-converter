package main

import (
	"json-csv-converter/converter"
	"log"
)

func main() {
	if err := converter.ConvertJSONToCSV("translations.json", "csv_to_json/exports/translations.csv"); err != nil {
		log.Fatal(err)
	}
}
