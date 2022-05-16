package main

import (
	"fmt"
	"time"
)

func main() {

	// utc := time.Now().UTC()
	// localTime := utc.Truncate(time.Hour).Local() // 切り捨ててUTC->Localに戻す

	// fmt.Println(localTime)

	// date := localTime.AddDate(0, 0, -6)
	// fmt.Println(date)

	// d := localTime.Sub(date)
	// fmt.Printf("%T %v", int(d.Hours()), int(d.Hours()))
	t := time.Now()

	h := 1000
	ts := time.Date(t.Year(), t.Month(), t.Day(), t.Hour()-(h), 0, 0, 0, time.Local)

	fmt.Println(ts)

}
