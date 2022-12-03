package main

import (
	"fmt"
	"time"
)

func main() {

	str := "2021.04.19"
	layout := "2006.01.02"
	nowTime := time.Now().Unix()
	//t := time.Now().AddDate(0, 1, 0).Unix()
	fmt.Println(nowTime)

	ti := time.Now().Unix()
	fmt.Println(time.Unix(ti, 0))

	ts, _ := time.Parse(layout, str)
	fmt.Println(ts)

	td := time.Now().Unix() + (60 * 60 * 24 * 30)
	tx := time.Unix(td, 0)
	fmt.Println(td)
	fmt.Println(tx)

	//td := t.AddDate(0, 1, 0).Format("2006.01.02")
	//fmt.Println(td)

	/* t1 := time.Now().AddDate(0, 1, 0)
	fmt.Println(t1) */
}
