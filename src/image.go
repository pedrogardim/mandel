package mandelbrot

import (
	"image"
	"image/color"
)

func SaveImage(array [SIZE][SIZE]int) *image.Gray {
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

	// for s := 0; s < SIZE; s++ {
	// 	img.Set(s, SIZE/2, color.Gray{Y: 128})
	// 	img.Set(SIZE/2, s, color.Gray{Y: 128})
	// }

	return img
}
