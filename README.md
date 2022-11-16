# Captcha

![验证码](./example/file/captcha0.png)
&emsp;&emsp;
![验证码](./example/file/captcha1.png)
&emsp;&emsp;
![验证码](./example/file/captcha2.png)
<br/>

![验证码](./example/file/captcha3.png)
&emsp;&emsp;
![验证码](./example/file/captcha4.png)
&emsp;&emsp;
![验证码](./example/file/captcha5.png)
<br/>

![验证码](./example/file/captcha6.png)
&emsp;&emsp;
![验证码](./example/file/captcha7.png)
&emsp;&emsp;
![验证码](./example/file/captcha8.png)
<br/>

## Usage

```go
package main

import (
	"log"

	"github.com/dushixiang/captcha"
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
```