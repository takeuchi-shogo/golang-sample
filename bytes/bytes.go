package main

import (
	"encoding/base64"
	"net/url"
	"unicode/utf8"
)

func main() {
	abyte := ""
	b := "上手"
	println(len(b))
	for i := 0; len(abyte) < 1024; i++ {
		abyte += "あ"
	}

	println(len(abyte), "バイト")
	println(utf8.RuneCountInString(abyte), "文字")
	str := `アリスは川辺でおねえさんのよこにすわって、なんにもすることがないのでとても退屈（たいくつ）しはじめていました。
		一、二回はおねえさんの読んでいる本をのぞいてみたけれど、そこには絵も会話もないのです。`
	buf := []byte(str)

	str1 := url.QueryEscape(str)
	println(len(str1), " bytes")

	// println(len(str), " bytes")

	bts := base64.StdEncoding.EncodeToString(buf)

	println(len(bts))
	println(utf8.RuneCountInString(bts), "文字")
}
