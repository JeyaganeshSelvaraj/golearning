package synctest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var output1 float64
var output2 float64
var output3 float64
var output4 float64
var wg sync.WaitGroup
var max = 1_000_000

func worker1() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < max; index++ {
		output = append(output, 89.6)
		sum += 89.6
	}
	output1 = sum
	fmt.Printf("Total elements in result output of worker1 %d\n", len(output))
}
func worker2() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < max; index++ {
		output = append(output, 899.6)
		sum += 899.6
	}
	output2 = sum
	fmt.Printf("Total elements in result output of worker2 %d\n", len(output))
}
func worker3() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < max; index++ {
		output = append(output, 8999.6)
		sum += 8999.6
	}
	output3 = sum
	fmt.Printf("Total elements in result output of worker3 %d\n", len(output))
}

func worker4() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < max; index++ {
		output = append(output, 8999.6)
		sum += 8999.6
	}
	output4 = sum
	fmt.Printf("Total elements in result output of worker4 %d\n", len(output))
}
func TestConcurrent(t *testing.T) {

	wg.Add(8)
	start := time.Now()
	worker1()
	worker2()
	worker3()
	worker4()
	elapsed := time.Since(start)
	fmt.Println("\nTime for 4 workers in series: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f  \nOutput3: %f  \nOutput4: %f\n",
		output1, output2, output3, output4)
	start = time.Now()
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	elapsed = time.Since(start)
	fmt.Println("\nTime for 4 workers in concurrent: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f  \nOutput3: %f  \nOutput4: %f\n",
		output1, output2, output3, output4)
	wg.Wait()
}
