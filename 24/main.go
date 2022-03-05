/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func create(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func distance(p1, p2 *point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := create(10, 10)
	p2 := create(12, 54)

	fmt.Println(distance(p1, p2))

}
