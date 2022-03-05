/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func setParam() (string, error) {
	switch len(os.Args) {
	case 1:
		return "", errors.New("add second argument")

	case 2:
		return os.Args[1], nil

	default:
		return "", errors.New("too many arguments")
	}
}

func worker(data <-chan interface{}, instance int) {
	for i := range data {
		fmt.Println(i, "worker", instance)
	}
}

func main() {
	param, err := setParam()
	if err != nil {
		log.Fatalln(err)
	}

	sigch := make(chan os.Signal, 1)
	data := make(chan interface{})
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)

	workerNum, err := strconv.Atoi(param)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < workerNum; i++ {
		go worker(data, i)
	}

	// loop
	for {
		// select блокируется до тех пор, пока один из блоков case не будет выполнен,
		select {
		// ловим прерывание
		case <-sigch:
			close(sigch)
			close(data)
			return
		// ловим данные
		default:
			data <- "hello world"
			data <- 1
			time.Sleep(time.Millisecond)
		}
	}
}
