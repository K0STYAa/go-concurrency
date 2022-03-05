/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
// https://go.dev/blog/pipelines
package main

import "fmt"

/*
	Генерируем данные из слайса
	отправляем и в канал
*/
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

/*
	Читаем данных из канала
	обрабатываем их
	отправляем в канал
*/
func work(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	for res := range work(gen(arr)) {
		fmt.Println(res)
	}
}
