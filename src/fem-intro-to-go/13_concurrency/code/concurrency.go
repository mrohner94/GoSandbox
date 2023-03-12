package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func printStuff(num int) {
	defer wg.Done()
	defer handlePanic()

	for i := 0; i < num; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func mathThings(c chan int, someValue int) {
	defer wg.Done()
	// From right to left, calculate the value, send it to the channel
	c <- someValue * 5
}

func main() {
	// wg.Add(3)
	// go printStuff(2)
	// go printStuff(3)
	// go printStuff(1)
	// wg.Wait()

	// resultChan := make(chan int)

	// go mathThings(resultChan, 5)
	// go mathThings(resultChan, 3)

	// num1 := <-resultChan
	// num2 := <-resultChan

	// fmt.Println(num1, num2)


	resultChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mathThings(resultChan, i)
	}

	// Wait for all goroutines to complete before closing the channel
	wg.Wait()

	// Close wont wait for the goroutines to be done without additional synchronization
	close(resultChan)


	for item := range resultChan {
		fmt.Println(item)
	}



}
