package main

import (
	"fmt"
	"math/rand"
	"time"
)

// gen возвращает канал,
// который производит числа 1, 2, 3,…
// Чтобы остановить подлежащую goroutine, закройте канал.
func gen() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func read(data <-chan int) {
	for i := range data {
		fmt.Println(i)
	}
}

// 1) просто закрыть канал
// 2) выставить таймаут
// 3) прерывать на sigint (из задания 4)
func main() {
	number := gen()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)

	data := make(chan int)
	timeout := time.After(1 * time.Second)

	go read(data)

	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			close(data)
			return

		case data <- rand.Intn(10):
		default:
			data <- rand.Intn(10)
		}
	}
}
