package main

import "fmt"

func main() {

	i := 0

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	println("i:", i)
	// 	for i := 0; i < 2; i++ {
	// 		test := testing()
	// 		fmt.Println(test)
	// 	}

	// 	m := map[string]bool{}
	// 	strArray := [6]string{"tanjiro", "nezuko", "zenitu", "inosuke", "kanao", "nezuko"}

	// 	for _, s := range strArray {
	// 		if !m[s] {
	// 			m[s] = true
	// 			fmt.Println(s)
	// 		}
	// 	}

	// 	count := 0
	// 	for i := 0; i <= 10; i++ {
	// 		count = count + i
	// 		fmt.Println(count)
	// 	}
	// }

	// func testing() int {
	// 	x := 1 + 1
	// 	return x
	// }

	// func SliceUnique(target []string) (unique []string) {
	// 	m := map[string]bool{}

	// 	for _, v := range target {
	// 		if !m[v] {
	// 			m[v] = true
	// 			unique = append(unique, v)
	// 		}
	// 	}

	// 	return unique
}
