package groutine

import (
	"sync"
	"testing"
)

func TestGroutine(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	//time.Sleep(time.Millisecond * 50)
	t.Logf("counter = %d", counter)
}
