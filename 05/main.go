/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//var start time.Time
//
//func init() {
//	start = time.Now()
//}

func setParam() (string, error) {
	switch len(os.Args) {
	case 1:
		return "", errors.New("too few arguments... plz add second argument")

	case 2:
		return os.Args[1], nil

	default:
		return "", errors.New("too many arguments")
	}
}

func read(data <-chan int) {
	for i := range data {
		fmt.Println(i)
	}

}

func main() {
	// тут ещё бы сделать проверку на -time :_)
	param, err := setParam()
	if err != nil {
		log.Fatalln(err)
	}

	t, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
	}

	data := make(chan int)
	timeout := time.After(time.Duration(t) * time.Second)

	go read(data)

	for {
		// select блокируется до тех пор, пока один из блоков case не будет выполнен
		select {
		case <-timeout:
			fmt.Println("timeout")
			close(data)
			return

		default:
			data <- rand.Intn(10)
		}
	}
}
