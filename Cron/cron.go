package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		call()
	})

	c.AddFunc("30 8 * * *", func() {})

	c.AddFunc("@every 3s", call2)

	c.Start()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {})

	// ここではポート番号 8080 で起動。適当に番号を当てていい
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf(err.Error())
	}

}

func call() {
	fmt.Printf("1 second!")
}

func call2() {
	fmt.Printf("4 second!")
}
