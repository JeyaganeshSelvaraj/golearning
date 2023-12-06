package synctest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var quit = make(chan bool)

func player(name string, move chan int) {
	turn := 0

	for {

		turn = <-move
		n := rand.Intn(100)
		if n < 5 && turn >= 5 {
			fmt.Printf("Player %s is lost at turn %d\n", name, turn)
			quit <- true
			return
		}

		fmt.Printf("Player %s is moving %d at turn %d\n", name, n, turn)
		turn++
		// Yield the turn back to the opposing player
		time.Sleep(1 * time.Second)
		move <- turn
	}
}

func TestChess(t *testing.T) {
	mv := make(chan int, 1)

	go player("player1", mv)
	go player("player2", mv)
	mv <- 1

	<-quit
}
