package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"captcha"
)

func main() {
	for i := 0; i < 10; i++ {
		c, err := captcha.New()
		if err != nil {
			log.Fatal(err)
		}
		img, word := c.Create()
		println("captcha:", word)
		captchaPng, err := os.Create(fmt.Sprintf("example/file/captcha%d.png", i))
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(captchaPng, img); err != nil {
			log.Fatal(err)
		}
		captchaPng.Close()
	}
}
