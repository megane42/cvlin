package subject

import (
	"encoding/csv"
	"os"
	"github.com/pkg/errors"
)

func LoadSubject(path string) ([][]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load csv file (load error)")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comment = '#'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load csv file (parse error)")
	}

	return data, nil
}
