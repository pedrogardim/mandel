package mandelbrot

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ParseParams(r *http.Request) ([2]float64, float64, error) {
	c := [2]float64{0, 0}

	center := r.URL.Query().Get("c")
	splited := strings.Split(center, ",")
	if len(splited) != 2 {
		return c, 0, errors.New("invalid center")
	}

	zoom, err := strconv.ParseFloat(r.URL.Query().Get("z"), 64)
	if err != nil {
		return c, 0, errors.New("invalid zoom")
	}

	x, err := strconv.ParseFloat(splited[0], 64)
	if err != nil {
		return c, 0, errors.New("invalid x center")
	}
	c[0] = x

	y, err := strconv.ParseFloat(splited[1], 64)
	if err != nil {
		return c, 0, errors.New("invalid y center")
	}
	c[1] = y

	return c, zoom, nil
}
