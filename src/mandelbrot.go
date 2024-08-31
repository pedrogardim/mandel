package mandelbrot

import (
	"fmt"
	"math/cmplx"
	"sync"
	"time"
)

const (
	SIZE           = 2000
	MAX_ITERATIONS = 50
	THRESHOLD      = 2.0
	WORKER_COUNT   = 8
)

func Generate(center [2]float64, zoom float64) [SIZE][SIZE]int {
	arr := [SIZE][SIZE]int{}
	var wg sync.WaitGroup
	rows := make(chan int, SIZE)

	// Start workers
	for i := 0; i < WORKER_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range rows {
				for x := 0; x < SIZE; x++ {
					xval := float64(x)/float64(SIZE)*2.0 - 1.0
					yval := float64(y)/float64(SIZE)*2.0 - 1.0
					it := manderbrotFormula(complex(xval, yval), complex(0, 0), 0)
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

func manderbrotFormula(input complex128, current complex128, iteration int) *int {
	if cmplx.Abs(current) > THRESHOLD {
		return &iteration
	}
	if iteration >= MAX_ITERATIONS {
		return nil
	}

	value := current*current + input
	return manderbrotFormula(input, value, iteration+1)
	// return nil
}
