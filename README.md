# Clone the repository
git clone https://github.com/nkechr1s/json-csv-converter.git

# Install dependencies
go mod tidy

# this runs the json to csv convert
go run main.go

cd csv_to_json
# this runs the csv to json convert
go fun csv_to_json.go
