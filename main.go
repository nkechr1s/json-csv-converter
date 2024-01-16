package main

import (
	"json-csv-converter/converter"
	"log"
)

func main() {
	if err := converter.ConvertJSONToCSV("translations.json", "translations.csv"); err != nil {
		log.Fatal(err)
	}
	if err := converter.ConvertCSVToJSON("translations.csv", "new_translations.json"); err != nil {
		log.Fatal(err)
	}
}
