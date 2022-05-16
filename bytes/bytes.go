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
	str := `development/downloads/purchased_ticket/mt_tv_tickets_11501_
		えーっとこの度おぱなちゃん生誕祭おかげさまで2回目を迎えることになりました(パチパチパチパチ) 皆様に愛され甘やかされ(えへへ)成長していく姿をとくとご覧あれ2022
		_aj4O0fmaDwlWmWu.xlsx`
	buf := []byte(str)

	str1 := url.QueryEscape(str)
	println(len(str1), " bytes")

	// println(len(str), " bytes")

	bts := base64.StdEncoding.EncodeToString(buf)

	println(len(bts))
	println(utf8.RuneCountInString(bts), "文字")
}
