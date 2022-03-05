/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
package main

import (
	"fmt"
)

// главная сложность - выбор опорного элемента
// худший случай n*n
// средний n log2 n
func qs(data []int) []int {
	// базовый случай
	if len(data) < 2 {
		return data
	}

	less := make([]int, 0)
	greater := make([]int, 0)

	// получаем худший случай
	pivot := data[0]
	for _, i := range data[1:] {
		if i > pivot {
			// больше опорного
			greater = append(greater, i)
		} else {
			//меньших опорного
			less = append(less, i)
		}
	}
	// сортируем меньшие и большие
	data = append(qs(less), pivot)
	data = append(data, qs(greater)...)
	return data
}

func main() {
	data := []int{2, 4, 3, 8, 5, 4, 6, 1, 7}

	//sort.Ints(data)
	fmt.Println(qs(data))
}
