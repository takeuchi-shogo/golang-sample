package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	a := "3.95"
	b, _ := strconv.ParseFloat(a, 64)
	fmt.Println(b - 4)

	var ja string = "Go言語"
	var en string = "golang"

	title := "えーっとこの度"

	fmt.Printf("Title: (%s)\n", string([]rune(title)[:20]))

	//文字数を出力
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))

	fmt.Println(en, "len:", utf8.RuneCountInString(en))
}
