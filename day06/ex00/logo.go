package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

const (
	size = 300
	name = "amazing_logo.png"
)

func main() {
	img := generateLogo()
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateLogo() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for _, pixel := range pixels {
		img.Set(pixel.x, pixel.y, pixel.clr)
	}

	return img
}
