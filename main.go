package main

import (
	"json-csv-converter/converter"
	"log"
)

func main() {
	if err := converter.ConvertJSONToCSV("translations.json", "exports/translations.csv"); err != nil {
		log.Fatal(err)
	}
	if err := converter.ConvertCSVToJSON("exports/translations.csv", "exports/new_translations.json"); err != nil {
		log.Fatal(err)
	}
}
