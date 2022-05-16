package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

//簡単なメール送信用
func main() {

	from := "送信用GMailアドレス"
	to := "送信先アドレス"

	auth := smtp.PlainAuth("", from, "アカウントパスワード", "smtp.gmail.com")

	//メッセージ　※日本語は文字化けする
	msg := []byte("" +
		"From: 送信した人　<" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名　subjectです\r\n" +
		"\r\n" +
		"テスト\r\n" +
		"")
	//メール送信
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		return
	}

	//成功すれば表示
	fmt.Println("success")
}

type Mail struct {
	Auth smtp.Auth
	Addr string
	From mail.Address
}

func NewMail() *Mail {
	return &Mail{
		//SMTP認証情報
		Auth: smtp.PlainAuth("", "test@gmail.com", "password", "stmp.gmail.com"),
		//メール送信先
		Addr: "",
		//
		From: mail.Address{Name: "テスト", Address: "test@gmail.com"},
	}
}

func (m *Mail) writeString(b *bytes.Buffer, s string) *bytes.Buffer {
	_, err := b.WriteString(s)
	if err != nil {
		fmt.Print(err.Error())
	}
	return b
}

func (m *Mail) encodeSubject(subject string) string {

	b := bytes.NewBuffer([]byte(""))
	strs := []string{}
	length := 13
	for k, c := range strings.Split(subject, "") {
		b.WriteString(c)
		if k%length == length-1 {
			strs = append(strs, b.String())
			b.Reset()
		}
	}

	if b.Len() > 0 {
		strs = append(strs, b.String())
	}

	b2 := bytes.NewBuffer([]byte(""))
	b2.WriteString("Subject: ")
	for _, line := range strs {
		b2.WriteString(" =?utf-8?B?")
		b2.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		b2.WriteString("?=\r\n")
	}

	return b2.String()
}
