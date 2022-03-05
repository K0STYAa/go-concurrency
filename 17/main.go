package main

import (
	"errors"
	"fmt"
)

// Алгоритм для отсортированного массива
// log(2) n
func binSearch(data []int, item int) (int, error) {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := data[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("there is no such number in the list")
}

func main() {
	// log (2) 8, те за 3 шага
	arr := []int{2, 3, 5, 7, 11, 13, 17, 19}

	res, err := binSearch(arr, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
