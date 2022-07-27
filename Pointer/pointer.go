package main

import "fmt"

type Person struct {
	Name        string
	Age         int
	Discription *string
}

func main() {
	p := Person{}

	p.Name = "太郎"
	p.Age = 30
	p.Discription = nil

	message := "私は社会人です"

	p.Discription = &message

	fmt.Println(p)
	fmt.Println(*p.Discription)
}
