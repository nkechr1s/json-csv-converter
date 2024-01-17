# Clone the repository
git clone https://github.com/nkechr1s/json-csv-converter.git

# Install dependencies
go mod tidy
go run main.go //this runs the json to csv convert

cd csv_to_json
go fun csv_to_json.go //this runs the csv to json convert
