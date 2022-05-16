package main

import "fmt"

func main() {
	a := 1
	b := 3

	if a == 2 || b == 5 {
		fmt.Println(a, b)
	}
	fmt.Println("success")
}
