package main

import (
	"fmt"
	"time"
)

func Ping(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}
func PrintPing(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}
func Pong(v chan string) {
	for i := 0; ; i++ {
		v <- "Pong"
	}
}
func PrintPong(v chan string) {
	for {
		msg := <-v
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var c chan string = make(chan string)

	go Ping(c)
	go PrintPing(c)
	go Pong(c)
	go PrintPong(c)

	fmt.Scanln()
}
