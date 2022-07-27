package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()

	if err := f.SaveAs("tmp/sample.xlsx"); err != nil {
		fmt.Println("err: ", err)
	}
}
