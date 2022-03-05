//Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

// для полноценно реализованного сета
// Сложность: O(m·n), где m и n — количество элементов переданного и текущего множеств соответственно.
// пересечения сетов
func intersection1(x []int, y []int) []int {
	// set - коллекция, которая реализует основные математические операции над множествами
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// отмечаем элементы тру (игнорируем повторы)
	for _, val := range x {
		set1[val] = true
	}

	for _, val := range y {
		set2[val] = true
	}

	// берем большее множесто и итерируемся по нему
	// а в меньшем проверяем на вхождение
	// если да, добавляем в слайс

	out := make([]int, 0, len(x))
	for i := range set2 {
		if set1[i] {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	x := []int{2, 1, 4, 3}
	y := []int{5, 6, 7, 4, 3}

	fmt.Println(intersection1(x, y))
}
