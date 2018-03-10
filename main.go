package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	x       = 0
	y       = 0
	width   = 400
	height  = 300
	quality = 100
)

func main() {
	img := image.NewRGBA(image.Rect(x, y, width, height))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 255, 0})
		}
	}

	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(200 * 64),
		Y: fixed.Int26_6(150 * 64),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString("僕は西川和希です")

	file, _ := os.Create("output.jpg")
	defer file.Close()

	err := jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
	if err != nil {
		panic(err)
	}

	log.Println("generated!")
}
