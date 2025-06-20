package writers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

// "encoding/csv"
func Exist_check(path string) (error) {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return errors.New("file does not exist")
	} else if err != nil {
		return fmt.Errorf("error checking file: %v", err)
	} 

	return nil
}

func Create_csv(filename string, fields []string) error { //path of file , fields slice
	f, err := os.OpenFile(fmt.Sprintf("./db/%v", filename), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error message : %v", err)
	}

	defer f.Close()

	w := csv.NewWriter(f) //accepts the io.writer but in this case I have passed *os.file

	w.Write(fields)
	w.Flush()

	if err := w.Error(); err != nil {
		return fmt.Errorf("error message : %v", err)
	}

	return nil
}

func Add_data(filename string, fields [][]string) error{

	err  := Exist_check(fmt.Sprintf("./db/%v",filename))

	if err != nil {
		return fmt.Errorf("the error is: %v",err)
	}

	f, err := os.OpenFile(fmt.Sprintf("./db/%v", filename), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error message : %v", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(fields)

	if err := w.Error(); err != nil {
		return fmt.Errorf("error message : %v", err)
	}

	return nil
}
