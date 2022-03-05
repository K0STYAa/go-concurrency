package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// костыль
var wg sync.WaitGroup

type counter struct {
	number uint64
}

func create() *counter {
	return &counter{
		number: 0,
	}
}

func (c *counter) add(num uint64) {
	atomic.AddUint64(&c.number, num)
	wg.Done()
}

func (c *counter) read() uint64 {
	return atomic.LoadUint64(&c.number)
}

func main() {
	c := create()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.add(1)
	}

	wg.Wait()
	fmt.Println(c.read())
}
