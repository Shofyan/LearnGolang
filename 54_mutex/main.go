package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	sync.Mutex
	val int
}

func (c *Counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *Counter) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var meter Counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
