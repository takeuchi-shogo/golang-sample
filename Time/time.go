package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	ts := time.Date(t.Year(), t.Month(), 1, 24, 0, 0, 0, time.Local)
	fmt.Println(ts)

	td := time.Unix(1630317600, 0)
	fmt.Println(td)

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T", tokyo)

	fmt.Println(time.Local)

	// for i := 1; i >= 0; i-- {
	// 	t1 := t.AddDate(0, -(i), 0)
	// 	ts := time.Date(t1.Year(), t1.Month(), 1, 0, 0, 0, 0, time.Local)

	// 	tl := time.Date(t1.Year(), t1.Month()+1, 1, 0, 0, 0, 0, time.Local)
	// 	fmt.Println("count:", i)

	// 	for d := 1; d <= tl.AddDate(0, 0, -1).Day(); d++ {

	// 		ds := time.Date(ts.Year(), ts.Month(), d, 0, 0, 0, 0, time.Local)

	// 		fmt.Println(ds)
	// 	}

	// 	// for h := 0; h <= 23; h++ {
	// 	// 	t3 := time.Date(ts.Year(), ts.Month()+1, 1, h, 0, 0, 0, time.Local)

	// 	// 	fmt.Println(t3)
	// 	// }
	// }
}
