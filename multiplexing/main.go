package main

import (
	"fmt"
	"time"
)

// Simulates a fast-producing channel
func fastProducer(ch chan<- string) {
	for {
		time.Sleep(500 * time.Millisecond)
		ch <- "Fast message"
	}
}

// Simulates a slow-producing channel
func slowProducer(ch chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		ch <- "Slow message"
	}
}

// Worker that can be cancelled using a done channel
func cancellableWorker(done <-chan struct{}, id int) {
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d: Received cancellation signal, stopping", id)
			return
		default:
			fmt.Printf("Worker %d: Working...", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("=== Basic Multiplexing with Select ===")
	fast := make(chan string)
	slow := make(chan string)
	go fastProducer(fast)
	go slowProducer(slow)
	for i := 0; i < 5; i++ {
		select {
		case msg := <-fast:
			fmt.Print("Received from fast: ", msg)
		case msg := <-slow:
			fmt.Print("Received from slow: ", msg)
		}
	}

	fmt.Println("=== Select with Timeout ===")
	timeoutChan := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		timeoutChan <- "Delayed response"
	}()
	select {
	case msg := <-timeoutChan:
		fmt.Print("Received: ", msg)
	case <-time.After(1 * time.Second):
		fmt.Print("Timeout! No message received.")
	}

	fmt.Println("=== Non-Blocking Select with Default ===")
	nonBlockingChan := make(chan string)
	select {
	case msg := <-nonBlockingChan:
		fmt.Print("Got: ", msg)
	default:
		fmt.Print("No message available. Skipping...")
	}

	fmt.Println("=== Using Select for Cancellation ===")
	done := make(chan struct{})
	go cancellableWorker(done, 1)
	time.Sleep(3 * time.Second)
	fmt.Print("Sending cancel signal...")
	close(done)
	time.Sleep(1 * time.Second)

	fmt.Println("=== Fan-In Pattern (Multiple Sources to One) ===")
	source1 := make(chan string)
	source2 := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			source1 <- fmt.Sprintf("Source 1 - %d", i)
			time.Sleep(400 * time.Millisecond)
		}
		close(source1)
	}()
	go func() {
		for i := 0; i < 3; i++ {
			source2 <- fmt.Sprintf("Source 2 - %d", i)
			time.Sleep(700 * time.Millisecond)
		}
		close(source2)
	}()
loop:
	for {
		select {
		case msg, ok := <-source1:
			if !ok {
				source1 = nil
				continue
			}
			fmt.Print("Fan-In Received from source1: ", msg)
		case msg, ok := <-source2:
			if !ok {
				source2 = nil
				continue
			}
			fmt.Print("Fan-In Received from source2: ", msg)
		}
		if source1 == nil && source2 == nil {
			break loop
		}
	}
	fmt.Println("=== Done ===")
}
