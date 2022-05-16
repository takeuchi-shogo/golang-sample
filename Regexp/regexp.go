package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	name, err := getFileName("mt_tv_images_event_header_89_1_IaO4XuW0HK_2x.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name: ", name)
}

func getFileName(fileName string) (st string, err error) {
	reg := `(?i)(\.jpg|\.jpeg|\.gif|\.png)$`
	str := regexp.MustCompile(reg).FindString(fileName)
	fmt.Println("str: ", str)
	if str == "" {
		return "", errors.New("invalid file")
	}
	extension := str[1:utf8.RuneCountInString(str)]
	fmt.Println("original extension: ", extension)
	fmt.Println("small extension: ", strings.ToLower(extension))
	s := utf8.RuneCountInString(fileName) - utf8.RuneCountInString(str)
	return fileName[0:s], nil
}
