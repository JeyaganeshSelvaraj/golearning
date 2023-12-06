package synctest

import (
	"fmt"
	"testing"
)

func fibanocci(c chan<- int, quit <-chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func TestFibanocci(t *testing.T) {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibanocci(c, quit)
}
