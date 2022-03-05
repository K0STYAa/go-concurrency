/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// worker than make squares
func pow(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		// если убрать, то только одна горутина будет выполняться, так как другие горутины не будут запланированы
		//time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
	// done with worker
	wg.Done()
}

func main() {
	// предотвратить много (ненужных) переключений контекста
	var wg sync.WaitGroup

	// буферизированный не лочиться, если полностью не заполнен
	data := []int{2, 4, 6, 8, 10}
	tasks := make(chan int, len(data)+1)
	results := make(chan int, len(data)+1)

	// launching 5 worker goroutines
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go pow(&wg, tasks, results, i)
	}

	// вкидываем значения в канал таск
	for _, i2 := range data {
		tasks <- i2
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// wait until all workers done their job

	// receving results from all workers
	wg.Wait()

	for i := 0; i < len(data); i++ {
		//result := <-results // non-blocking because buffer is non-empty
		fmt.Println("[main] Result", i, ":", <-results)
	}
}
