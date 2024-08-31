package main

import (
	"bytes"
	"image/png"
	mandelbrot "lightstand/mandel/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		points, zoom, err := mandelbrot.ParseParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		arr := mandelbrot.Generate(points, zoom)
		img := mandelbrot.SaveImage(arr)

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			http.Error(w, "Unable to encode image.", http.StatusInternalServerError)
			return
		}

		// Set the content type to image/png
		w.Header().Set("Content-Type", "image/png")

		// Write the image data to the response
		w.Write(buf.Bytes())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
