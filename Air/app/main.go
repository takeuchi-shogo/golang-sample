package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "NO",
			"message2": "レスポンスだよ",
		})
	})

	r.POST("/payments", func(c *gin.Context) {
		secretToken := "testtest"
		header := c.GetHeader("X-Komoju-Signature")
		fmt.Println(header)

		buf := make([]byte, 1028)
		n, _ := c.Request.Body.Read(buf)
		body := string(buf[0:n])

		sign, answer := hmacEqual([]byte(body), []byte(header), []byte(secretToken))

		fmt.Println("sign: ", sign)
		fmt.Println("answer: ", answer)

		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"message2": "レスポンスだよ",
		})
	})

	r.Run()
}

func hmacEqual(message, messageMAC, key []byte) (sign string, match bool) {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	//decodedMAC, err := hex.DecodeString(string(messageMAC))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return hex.EncodeToString(expectedMAC), hmac.Equal(messageMAC, expectedMAC)
}
