package main

import (
	"fmt"
	"time"
)

func main() {
	// --- BASIC SELECT: waiting on multiple channels with delay ---
	fmt.Println("=== Example 1: select between two channels with delay ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	// --- TIMEOUT using select and time.After ---
	fmt.Println("=== Example 2: Timeout with select ===")

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Done processing"
	}()

	select {
	case result := <-ch:
		fmt.Println("Result:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: operation took too long")
	}

	// --- DEFAULT CASE: non-blocking select ---
	fmt.Println("=== Example 3: Default (non-blocking select) ===")

	ch3 := make(chan int)

	select {
	case val := <-ch3:
		fmt.Println("Received:", val)
	default:
		fmt.Println("No value available, continuing...")
	}

	// --- LOOP WITH SELECT and MULTIPLE CHANNELS ---
	fmt.Println("=== Example 4: Loop with select ===")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			c1 <- fmt.Sprintf("C1 value %d", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 2; i++ {
			c2 <- fmt.Sprintf("C2 value %d", i)
			time.Sleep(700 * time.Millisecond)
		}
	}()

	timeout := time.After(3 * time.Second)
loop:
	for {
		select {
		case msg := <-c1:
			fmt.Println("From C1:", msg)
		case msg := <-c2:
			fmt.Println("From C2:", msg)
		case <-timeout:
			fmt.Println("Timed out. Exiting loop.")
			break loop
		}
	}

	// --- CLOSED CHANNEL HANDLING ---
	fmt.Println("=== Example 5: Handling closed channel ===")

	closedChan := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			closedChan <- i
		}
		close(closedChan)
	}()

	for {
		select {
		case val, ok := <-closedChan:
			if !ok {
				fmt.Println("Channel closed. Exiting.")
				goto end
			}
			fmt.Println("Received:", val)
		}
	}
end:
	fmt.Println("Program finished.")
}
