package main

import (
	"fmt"
)

func do(value interface{}) {
	fmt.Printf("value %v is type %T\n", value, value)
}

func main() {
	c := make(chan int)
	do(1)
	do(true)
	do("abc")
	do(c)
	do(1.17)
}
