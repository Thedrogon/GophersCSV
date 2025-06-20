package readers

import (
	"os"
)

func Read_csv(path string) (string , error) {

	data, err := os.ReadFile(path)
	var data_string string = string(data)
	if err != nil {
		return "",err
	}

	return data_string, nil
}
