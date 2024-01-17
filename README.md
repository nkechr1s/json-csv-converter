# Clone the repository
git clone https://github.com/nkechr1s/json-csv-converter.git

# Install dependencies
go mod tidy

# this runs the json to csv convert
go run main.go


# this runs the csv to json convert
cd csv_to_json

go run csv_to_json.go
