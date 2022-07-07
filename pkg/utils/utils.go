package utils

import (
	"encoding/csv"
	"errors"
	"os"
)

func readJSONfile(filename string) ([][]string, error) {

	if csvFile, err := os.Open(filename); err == nil {
		defer csvFile.Close()
		csvReader := csv.NewReader(csvFile)
		data, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("failed to read the file: " + filename)
}

func doMarshall() map[string]interface{} {

}
