package main

import (
	"fmt"
	"sync"
)

func sqr(x int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1
	c <- x * x // записываем результат в канал
}

func summ(sum *int, x int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock() // Блокирум доступ, чтобы в данный момент доступ был только у одной горутине к sum
	defer mu.Unlock() // После return разблокируем доступ
	*sum += x
	wg.Done()
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	c := make(chan int, len(numbers))
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // Будем использовать mutex для последовательного доступа горутин

	for _, j := range numbers {
		wg.Add(1) // увеличиваем счетчик на 1
		 go sqr(j, c, &wg)
	}
	wg.Wait() // ожидаем записи квадратов в канал
	close(c) // Закрываем канал, поскольку дальше будем использовать его только для чтения
	for x := range c { // при подсчете суммы важен порядок выполнения, поэтому это либо очередь из горутин, либо простой подсчет данных
		wg.Add(1)
		go summ(&sum, x, &mu, &wg) 
	}
	wg.Wait()

	fmt.Printf("Sum: %d \n", sum)
}