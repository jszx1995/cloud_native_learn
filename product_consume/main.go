package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i + i
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func consumer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	consumer(ch)
}
