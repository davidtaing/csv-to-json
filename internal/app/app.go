package app

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func ConvertCSVToJSON(records [][]string) error {
	jsonBytes, err := MarshalToJSON(records)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal csv data")
		return err
	}

	_, err = os.Stdout.Write(jsonBytes)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing data to json", err)
		return err
	}

	return nil
}

func ReadCSVFromFile(path string) ([][]string, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "Successfully opened the CSV file")

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func ReadCSVFromStdin() ([][]string, error) {
	r := csv.NewReader(os.Stdin)
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func MarshalToJSON(records [][]string) ([]byte, error) {
	var data []map[string]string

	if len(records) > 0 {
		columnNames := records[0]

		for _, row := range records[1:] {
			record := make(map[string]string)
			for i, value := range row {
				if i < len(columnNames) {
					record[columnNames[i]] = value
				}
			}
			data = append(data, record)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling data:", err)
		return nil, err
	}

	return jsonBytes, nil
}
