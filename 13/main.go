/*
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Printf("a = %d, b = %d\n", a, b)

	//множественного присвоения (tuple assignment).
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)

	a += b  //a+b ; b
	b -= a  //a+b ; -a
	a += b  //b ; -a
	b *= -1 //b ; a

	fmt.Printf("a = %d, b = %d\n", a, b)

}
