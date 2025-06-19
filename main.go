package main

import (
	write "github.com/Thedrogon/GophersCSV/writers"
	// "encoding/csv"
	// "log"
	// "os"
)

func main() {

	write.Create_csv("./db/200.csv" , []string{"first_name","last_name","nick_name"})

	// label_data := [][]string{
	// 	{"first_name", "last_name", "nick_name"},
		
	// }
	// w := csv.NewWriter(os.Stdout) //kind of writer object

	// w.WriteAll(label_data)

	// if err := w.Error(); err != nil {
	// 	log.Fatalln("Error: ", err)
	// }

}