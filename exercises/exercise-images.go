package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{
	w, h int
	color uint8
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
  return color.RGBA{img.color + uint8(x), img.color + uint8(y), 255, 255}
}

func main() {
	m := Image{200, 150, 12}
	pic.ShowImage(m)
}

