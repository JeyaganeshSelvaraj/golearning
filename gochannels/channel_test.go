package gochannels

import (
	"sync"
	"testing"
)

func TestAsync(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			t.Logf("%d \n", n)
		}(i)
	}
	wg.Wait()
}

func TestChan(t *testing.T) {
	ch := make(chan int, 10)
	exit := make(chan bool, 1)
	var wg sync.WaitGroup
	go func() {
		for n := range ch {
			t.Logf("%d \n", n)
		}
		exit <- true
	}()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n
		}(i)
	}
	wg.Wait()
	close(ch)
	<-exit
}

// func TestHttpServer(t *testing.T) {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, World!")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
