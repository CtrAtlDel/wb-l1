package main

import (
	"fmt"
	"sync"
)

func sqr(x int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- x * x
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	c := make(chan int, len(numbers))
	sum := 0
	var wg sync.WaitGroup

	for _, j := range numbers {
		wg.Add(1)
		sqr(j, c, &wg)
	}
	close(c)
	for j := range c { // при подсчете суммы важен порядок выполнения, поэтому это либо очередь из горутин, либо простой подсчет данных
		sum += j
	}
	wg.Wait()
	fmt.Printf("Sum: %d \n", sum)
}
