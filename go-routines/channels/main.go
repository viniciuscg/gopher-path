package main

import (
	"fmt"
	"time"
)

func main() {
	// --- BASIC CHANNEL CREATION ---
	ch := make(chan string) // unbuffered channel

	// Start a goroutine that sends a value
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	// Receive the value (this blocks until data is sent)
	msg := <-ch
	fmt.Println(msg)

	// --- BUFFERED CHANNEL ---
	buffered := make(chan int, 3) // can hold up to 3 items without blocking
	buffered <- 1
	buffered <- 2
	buffered <- 3

	fmt.Println("Buffered values:", <-buffered, <-buffered, <-buffered)

	// --- CHANNEL DIRECTION (SEND-ONLY / RECEIVE-ONLY) ---
	sendOnly := make(chan<- string)
	recvOnly := make(<-chan string)

	// --- RANGE OVER CHANNEL ---
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // Must close to range over
	}()

	for num := range numbers {
		fmt.Println("Received:", num)
	}

	// --- CHECKING IF CHANNEL IS CLOSED ---
	data := make(chan int)
	go func() {
		data <- 10
		close(data)
	}()

	val, ok := <-data
	if ok {
		fmt.Println("Got value:", val)
	} else {
		fmt.Println("Channel closed")
	}

	// --- SELECT STATEMENT WITH MULTIPLE CHANNELS ---
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
