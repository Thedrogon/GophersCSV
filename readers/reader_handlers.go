package readers

import (
	"os"
)

func Read_csv(path string) ([]byte , error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil,err
	}

	return data, nil
}
