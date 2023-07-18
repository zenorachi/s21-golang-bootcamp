package main

import "image/color"

type Pixel struct {
	x, y int
	clr  color.Color
}

var pixels []Pixel

func init() {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if (y >= 150 && y <= 160) && ((x >= 100 && x <= 140) || (x >= 170 && x <= 210)) {
				pixels = append(pixels, Pixel{x: x, y: y, clr: color.White})
			} else if (((x >= 70 && x <= 80) || (x >= 230 && x <= 240)) && (y >= 80 && y <= 250)) || (((y >= 80 && y <= 90) || (y >= 240 && y <= 250)) && x >= 70 && x <= 240) ||
				(((x >= 90 && x <= 150) || (x >= 160 && x <= 220)) && y >= 100 && y <= 230) {
				pixels = append(pixels, Pixel{x, y, color.Black})
			} else if (x >= 150 && x <= 160 && y >= 90 && y <= 250) ||
				(((x >= 80 && x <= 90) || (x >= 220 && x <= 230)) && (y >= 90 && y <= 240)) || (((y >= 90 && y <= 100) || (y >= 230 && y <= 240)) && x >= 80 && x <= 230) {
				pixels = append(pixels, Pixel{x: x, y: y, clr: color.RGBA{R: 255, A: 255}})
			} else if x >= 110 && x <= 140 && y >= 120 && y <= 130 {
				pixels = append(pixels, Pixel{x: x, y: y, clr: color.White})
			} else {
				pixels = append(pixels, Pixel{x: x, y: y, clr: color.White})
			}
		}
	}
}
