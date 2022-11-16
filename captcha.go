package captcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ToBase64(img image.Image) (base64Encoding string, err error) {

	emptyBuff := bytes.NewBuffer(nil)
	err = png.Encode(emptyBuff, img)
	if err != nil {
		return "", nil
	}

	imgBytes := emptyBuff.Bytes()[0:emptyBuff.Len()]
	base64Encoding = "data:image/jpeg;base64," + toBase64(imgBytes)
	return base64Encoding, nil
}

func New() (*Captcha, error) {
	// 读取字体
	fontBytes, err := os.ReadFile("./actionj.ttf")
	if err != nil {
		return nil, err
	}
	// 解析字体
	_font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(_font, &truetype.Options{Size: 32})

	option := Option{
		WordLength:  5,
		LineCount:   1,
		CircleCount: 1,
		Width:       130,
		Height:      48,
		CodeType:    Default,
		Mode:        Normal,
		FontFace:    face,
		Colors:      DefaultColors,
	}
	return NewWithOption(option), nil
}

func NewWithOption(option Option) *Captcha {
	if len(option.Colors) == 0 {
		option.Colors = DefaultColors
	}
	dc := gg.NewContext(option.Width, option.Height)
	dc.SetFontFace(option.FontFace)
	return &Captcha{
		option: option,
		dc:     dc,
	}
}

type Captcha struct {
	option Option
	dc     *gg.Context
}

func (c *Captcha) code() string {
	code := c.drawCode()
	c.drawBesselLine()
	c.drawCircle()
	return code
}

func (c *Captcha) Create() (image.Image, string) {
	return c.dc.Image(), c.code()
}

func (c *Captcha) randomStr(chars string) (randStr string) {
	charsLen := len(chars)
	n := c.option.WordLength
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

func (c *Captcha) drawCode() string {
	var word, result string
	if c.option.Mode == Normal {
		switch c.option.CodeType {
		case Default:
			word = c.randomStr(DefaultStr)
		case OnlyNumber:
			word = c.randomStr(OnlyNumberStr)
		case OnlyLetter:
			word = c.randomStr(OnlyLetterStr)
		}
		result = word
	} else {
		word, result = c.randomFormula()
	}

	for i, w := range strings.Split(word, "") {
		c.dc.SetColor(c.randomColor())
		c.dc.DrawString(w, float64(i*20+10), 40)
	}

	return result
}

func (c *Captcha) randomFormula() (string, string) {
	n1 := randomNumber(11, 20)
	n2 := randomNumber(1, 10)
	ops := []rune("+-*/")
	op := ops[randomNumber(0, 4)]

	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)

	var result int
	switch op {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '×':
	case '*':
		result = n1 * n2
	case '÷':
	case '/':
		result = n1 / n2
	}
	return s1 + string(op) + s2 + "=?", strconv.Itoa(result)
}

func (c *Captcha) drawBesselLine() {
	ctx := c.dc
	width := c.option.Width
	height := c.option.Height

	for i := 0; i < c.option.LineCount; i++ {
		var x0 = 5
		var y0 = randomNumber(5, height/2)
		var x1 = randomNumber(width/4, width/4*3)
		var y1 = randomNumber(5, height-5)
		var x2 = width - 5
		var y2 = randomNumber(height/2, height-5)

		if rand.Intn(2) == 0 {
			// 二阶贝塞尔曲线
			ctx.MoveTo(float64(x0), float64(y0))
			ctx.QuadraticTo(float64(x1), float64(y1), float64(x2), float64(y2))

			c.drawCurve()
		} else {
			// 三阶贝塞尔曲线
			var x3 = randomNumber(width/4, width/4*3)
			var y3 = randomNumber(5, height-5)
			ctx.MoveTo(float64(x0), float64(y0))
			ctx.CubicTo(float64(x1), float64(y1), float64(x3), float64(y3), float64(x2), float64(y2))

			c.drawCurve()
		}
	}
}

func (c *Captcha) drawCurve() {
	c.dc.SetRGBA(0, 0, 0, 0)
	c.dc.FillPreserve()
	c.dc.SetColor(c.randomColor())
	c.dc.SetLineWidth(2)
	c.dc.Stroke()
}

func (c *Captcha) drawCircle() {
	width := c.option.Width
	height := c.option.Height

	for i := 0; i < c.option.CircleCount; i++ {
		var r = randomNumber(5, 10)
		x := rand.Intn(width - 25)
		y := rand.Intn(height - 15)
		c.dc.DrawCircle(float64(x), float64(y), float64(r))
		c.dc.SetColor(c.randomColor())
		c.dc.SetLineWidth(2)
		c.dc.Stroke()
	}
}

func (c *Captcha) randomColor() color.RGBA {
	colorLen := len(c.option.Colors)
	return c.option.Colors[rand.Intn(colorLen)]
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
