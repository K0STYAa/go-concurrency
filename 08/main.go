/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
package main

import (
	"fmt"
)

// добаить обработку ошибок
func setBit(num *int64, pos uint8, bit uint8) int64 {
	if pos > 62 {
		fmt.Println("cannot set because of big pos")
		return *num
	}

	switch bit {
	case 0:
		*num &= 0 << pos
	case 1:
		*num |= 1 << pos

	default:
		fmt.Println("unexpected bit")
	}

	return *num
}

func main() {
	var v int64

	fmt.Printf("%064b\n", setBit(&v, 32, 1))
	fmt.Printf("%064b\n", setBit(&v, 63, 1))
	fmt.Printf("%064b\n", setBit(&v, 32, 1))
}
