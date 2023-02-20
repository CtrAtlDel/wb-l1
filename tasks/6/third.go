package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(stopFlag *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stopFlag {
			fmt.Println("Worker stopped")
			return
		}
		fmt.Println("Working...")
	}
}

// Идея такая: можем передать флаг по указателю в бесконечный цикл, затем через 5 секунд поменять его значение
// тогда программма завершиться через 5 секунд
func main() {
	var stopFlag bool
	var wg sync.WaitGroup
	wg.Add(1) 
	go worker(&stopFlag, &wg)
	// Wait for some time
	time.Sleep(5 * time.Second)
	// Signal worker to stop
	stopFlag = true
	// Wait for worker to stop
	wg.Wait()
	fmt.Println("Program terminated")
}
