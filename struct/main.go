package main

import "fmt"

func main() {

	type Sample struct {
		A string
		B int
		C bool
		D interface{}
	}

	fmt.Println(Sample{})
}
