package synctest

import (
	"sync"
	"sync/atomic"
	"testing"
)

var counter int
var atomicCounter atomic.Int32

func TestGlobalValueUpdate(t *testing.T) {
	var wg sync.WaitGroup
	number := 1000
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	t.Logf("%d", counter)
}

func TestGlobalValueUsingAtomic(t *testing.T) {
	var wg sync.WaitGroup
	number := 1000
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			atomicCounter.Add(1)
		}()
	}
	wg.Wait()
	t.Logf("%d", atomicCounter.Load())
}

func TestGlobalValueUsingLock(t *testing.T) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	number := 1000
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter++
		}()
	}
	wg.Wait()
	t.Logf("%d", counter)
}
