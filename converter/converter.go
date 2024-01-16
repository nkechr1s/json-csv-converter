package converter

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

type Translation struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// json to csv
func ConvertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var translations map[string]string
	if err := json.NewDecoder(sourceFile).Decode(&translations); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"key", "value"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for key, value := range translations {
		var csvRow []string
		csvRow = append(csvRow, key, value)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

// csv to json
func ConvertCSVToJSON(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	reader := csv.NewReader(sourceFile)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	translations := make(map[string]string)

	// Skip the first record assuming it's the header "key": "value"
	for _, record := range records[1:] {
		if len(record) >= 2 {
			translations[record[0]] = record[1]
		}
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(translations); err != nil {
		return err
	}

	return nil
}
