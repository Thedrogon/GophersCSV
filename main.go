package main

import (
	"fmt"

	write "github.com/Thedrogon/GophersCSV/writers"
	// "encoding/csv"
	// "log"
	// "os"
)

func main() {

	if err := write.Create_csv("200.csv" , []string{"first_name","last_name","nick_name"}); err != nil{
		fmt.Println(err)
	}

	if err := write.Add_data("201.csv" , [][]string{{"Sayan","Mukherjee","Sayan"}}); err!=nil{
		fmt.Println(err)
	} 



	// label_data := [][]string{
	// 	{"first_name", "last_name", "nick_name"},
		
	// }
	// w := csv.NewWriter(os.Stdout) //kind of writer object

	// w.WriteAll(label_data)

	// if err := w.Error(); err != nil {
	// 	log.Fatalln("Error: ", err)
	// }

}