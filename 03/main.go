/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
)

func sumSquares(data []int) int {
	if len(data) < 1 {
		return 0
	}

	ch := make(chan int, len(data))
	defer close(ch)

	go func() {
		for i := 0; i < len(data); i++ {
			go func(val int) {
				ch <- val * val
			}(data[i])
		}
	}()

	var sum int
	for i := 0; i < len(data); i++ {
		sum += <-ch
	}

	return sum
}

func main() {
	data := []int{2, 4, 6, 8, 10}
	sum := sumSquares(data)

	fmt.Printf("res = %d\n", sum)
}
