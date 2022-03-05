package main

import (
	"fmt"
	"sync"
)

// gen - преобразует слайс чисел в канал, который выдает целые числа из списка
// gen - запускает го-лямбду, которая отправляет числа по каналу, из закрывает канал после отправки всех значений
func gen(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range data {
			out <- i
		}
		close(out)
	}()

	return out
}

// pw - получает целые числа из канала и возвращает канал, который выдает квадраты
func pw(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

// Отправка по закрытому каналу вызывает panic, поэтому важно убедиться,
// что все goroutine выполнены до вызова close.
func merge(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Запускаем output goroutine
	// для каждого входного канала в chs.
	// output копирует значения из c в out
	// до тех пор пока c не закрыт, затем вызывает wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chs))
	for _, c := range chs {
		go output(c)
	}

	// Запускаем goroutine чтобы закрыть out
	// когда все output goroutine заверешены.
	// Это должно начнаться после вызова wg.Add.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	data := []int{2, 4, 6, 8, 10}
	in := gen(data)

	// распеделяем работу по 2м горутинам
	work1 := pw(in)
	work2 := pw(in)

	for n := range merge(work1, work2) {
		fmt.Println(n) // 4 затем 9, или 9 затем 4
	}
}
