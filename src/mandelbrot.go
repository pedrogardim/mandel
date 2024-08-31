package mandelbrot

import (
	"fmt"
	"math"
	"math/cmplx"
	"sync"
	"time"
)

const (
	SIZE         = 1000
	THRESHOLD    = 2.0
	WORKER_COUNT = 8
)

func Generate(center [2]float64, zoom float64) [SIZE][SIZE]int {
	arr := [SIZE][SIZE]int{}
	var wg sync.WaitGroup
	rows := make(chan int, SIZE)
	z := 1 / zoom
	maxIterations := int(math.Round(zoom * 1.5))
	if maxIterations < 20 {
		maxIterations = 20
	}

	// Start workers
	for i := 0; i < WORKER_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range rows {
				for x := 0; x < SIZE; x++ {
					xval := float64(x)/float64(SIZE)*z*2 - z + center[0]
					yval := float64(y)/float64(SIZE)*z*2 - z + center[1]
					it := manderbrotFormula(complex(xval, yval), complex(0, 0), 0, maxIterations)
					if it != nil {
						arr[x][y] = *it
					} else {
						arr[x][y] = 0
					}
				}
			}
		}()
	}

	t := time.Now()

	// Distribute work to workers
	for y := 0; y < SIZE; y++ {
		rows <- y
	}
	close(rows)

	wg.Wait()
	fmt.Println("Time taken: ", time.Since(t))

	return arr
}

func manderbrotFormula(input complex128, current complex128, iteration int, maxIterations int) *int {
	if cmplx.Abs(current) > THRESHOLD {
		return &iteration
	}
	if iteration >= maxIterations {
		return nil
	}

	value := current*current + input
	return manderbrotFormula(input, value, iteration+1, maxIterations)
	// return nil
}
