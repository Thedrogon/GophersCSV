package readers

import (
	"log"
	"os"
)

func Read_csv() ([]byte , error) {

	data, err := os.ReadFile("./db/200.csv")
	if err != nil {
		log.Fatalln(err)
	}

	return data, err
}
