// Реализовать собственную функцию sleep.
package main

// https://go.dev/src/time/sleep.go#L9

import (
	"fmt"
	"time"
)

func customSleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	fmt.Println("Fell asleep")
	customSleep(2)
	fmt.Println("Woke up")

}
