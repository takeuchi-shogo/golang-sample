package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}

	fmt.Println(slice)

	slice = append(slice, 100)

	fmt.Println(slice)

	slice = append([]int{100}, slice...)

	fmt.Println(slice)
}
