package main

import (
	"fmt"
	"log"

	read "github.com/Thedrogon/GophersCSV/readers"
	// "encoding/csv"
	// "log"
	// "os"
)

func main() {


	data,err := read.Read_csv()
	
	data_new := string(data)
	if err!= nil{
		log.Fatalln(err)
	}

	fmt.Println(data_new)
	// if err := write.Create_csv("200.csv", []string{"first_name", "last_name", "nick_name"}); err != nil {
	// 	fmt.Println(err)
	// }

	// if err := write.Add_data("200.csv", [][]string{
	// 	{"cloud", "river", "stone"},
	// 	{"forest", "sky", "breeze"},
	// 	{"mountain", "wave", "tree"},
	// 	{"shadow", "flame", "dust"},
	// 	{"valley", "star", "wind"},
	// 	{"ocean", "hill", "spark"},
	// 	{"meadow", "moon", "rock"},
	// 	{"desert", "lake", "fog"},
	// 	{"canyon", "sun", "leaf"},
	// 	{"glacier", "field", "mist"},
	// }); err != nil {
	// 	fmt.Println(err)
	// }

	// label_data := [][]string{
	// 	{"first_name", "last_name", "nick_name"},

	// }
	// w := csv.NewWriter(os.Stdout) //kind of writer object

	// w.WriteAll(label_data)

	// if err := w.Error(); err != nil {
	// 	log.Fatalln("Error: ", err)
	// }

}
