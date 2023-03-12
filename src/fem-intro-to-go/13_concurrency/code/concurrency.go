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

func main() {
	wg.Add(3)
	go printStuff(40)
	go printStuff(3)
	go printStuff(13)
	wg.Wait()
}
