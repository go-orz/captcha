package main

import (
	"captcha"
	"log"
)

func main() {
	c, err := captcha.New()
	if err != nil {
		log.Fatal(err)
	}
	img, word := c.Create()
	println("captcha:", word)
	base64Encoding, err := captcha.ToBase64(img)
	if err != nil {
		log.Fatal(err)
	}
	println("base64:", base64Encoding)
}
