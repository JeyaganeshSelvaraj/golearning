package synctest

import (
	"fmt"
	"testing"
	"time"
)

var limit = 10_000_000

func plainSum(elems []float64) float64 {
	sum := 0.0
	for _, elem := range elems {
		sum += elem
	}
	return sum
}
func sum(elems []float64, res chan<- float64) {
	sum := 0.0
	for _, elem := range elems {
		sum += elem
	}
	res <- sum
}

func TestConcurSum(t *testing.T) {
	elems := []float64{}
	for i := 1; i <= limit; i++ {
		elems = append(elems, 1.0)
	}
	c := make(chan float64)
	start := time.Now()
	totalElems := len(elems)
	go sum(elems[:totalElems/2], c)
	go sum(elems[totalElems/2:], c)
	first, second := <-c, <-c
	elapsed := time.Since(start)
	fmt.Printf("first: %f  second: %f elapsed time: %v", first, second,
		elapsed)
	start = time.Now()
	answer := plainSum(elems)
	elapsed = time.Since(start)
	fmt.Printf("\nplain sum: %f elapsed time: %v", answer, elapsed)
}
