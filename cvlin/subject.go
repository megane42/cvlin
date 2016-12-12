package cvlin

import (
	"encoding/csv"
	"os"
)

func Parse (path string) ([][]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comment = '#'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
