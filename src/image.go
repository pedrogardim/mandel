package mandelbrot

import (
	"image"
	"image/color"
)

func SaveImage(array [WIDTH][HEIGHT]int) *image.Gray {
	// Dimensions of the array
	height := len(array)
	width := len(array[0])

	// Create a new grayscale image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Set pixels based on the array values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := uint8(array[x][y])
			img.Set(x, y, color.Gray{Y: v * 3})
		}
	}

	return img
}
