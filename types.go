package captcha

import (
	_ "embed"
	"image/color"

	"golang.org/x/image/font"
)

type CodeType int

const (
	Default    CodeType = iota // 数字和字母
	OnlyNumber                 // 纯数字
	OnlyLetter                 // 纯字母
)

const (
	DefaultStr    = "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	OnlyNumberStr = "23456789"
	OnlyLetterStr = "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"
)

type Mode int

const (
	Normal Mode = iota
	Formula
)

type Option struct {
	WordLength  int
	LineCount   int
	CircleCount int
	Width       int
	Height      int
	CodeType    CodeType
	Mode        Mode
	FontFace    font.Face
	Colors      []color.RGBA
}

var defaultColors = []color.RGBA{
	{
		R: 0,
		G: 135,
		B: 255,
		A: 255,
	},
	{
		R: 51,
		G: 153,
		B: 51,
		A: 255,
	},
	{
		R: 255,
		G: 102,
		B: 102,
		A: 255,
	},
	{
		R: 255,
		G: 153,
		B: 0,
		A: 255,
	},
	{
		R: 153,
		G: 102,
		B: 0,
		A: 255,
	},
	{
		R: 153,
		G: 102,
		B: 153,
		A: 255,
	},
	{
		R: 51,
		G: 153,
		B: 153,
		A: 255,
	},
	{
		R: 102,
		G: 102,
		B: 255,
		A: 255,
	},
	{
		R: 0,
		G: 102,
		B: 204,
		A: 255,
	},
	{
		R: 204,
		G: 51,
		B: 51,
		A: 255,
	},
	{
		R: 0,
		G: 153,
		B: 204,
		A: 255,
	},
	{
		R: 0,
		G: 51,
		B: 102,
		A: 255,
	},
}

//go:embed font/actionj.ttf
var defaultFont []byte
