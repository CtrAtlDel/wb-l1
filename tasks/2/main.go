package main

import (
	"fmt"
	"sync"
)

func square(x int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик горутин
	c <- x * x
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squareNumbers := make([]int, len(numbers))
	c := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, i := range numbers{ // _ - порядковый номер, i - элемент из numbers
		wg.Add(1)
		go square(i, c, &wg)
	}

	wg.Wait()
	close(c)

	i := 0
	for s := range c {
		squareNumbers[i] = s
		i++
	}

	fmt.Println(squareNumbers)

}
