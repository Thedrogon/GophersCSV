package writers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//"encoding/csv"

func Create_csv(filename string, fields []string) { //path of file , fields slice 
	f, err := os.OpenFile(fmt.Sprintf("./db/%v",filename), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(fields)
	w.Flush()

	if err:= w.Error() ; err != nil{
		log.Fatalln("Error message : ",err)
	}
}



func Add_data(filename string, fields [][]string) {
	f, err := os.OpenFile(fmt.Sprintf("./db/%v",filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(fields)

	if err:= w.Error() ; err != nil{
		log.Fatalln("Error message : ",err)
	}
}
