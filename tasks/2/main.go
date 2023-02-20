package main

import (
	"fmt"
	"sync"
)

func square(x int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик горутин
	c <- x * x      // записываем результат в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squareNumbers := make([]int, len(numbers)) // делаем слайс из количества элементов в массиве
	c := make(chan int, len(numbers))          // делаем канал из количества элементов в массиве

	var wg sync.WaitGroup

	for _, i := range numbers { // _ - порядковый номер, i - элемент из numbers
		wg.Add(1) 				// увеличиваем считчик горутин
		go square(i, c, &wg)
	}

	wg.Wait() // ждем пока горутины завершаться (чтобы main поток раньше не завершился)
	close(c) // закрываем канал для записи

	i := 0
	for s := range c {
		squareNumbers[i] = s // копируем в слайс
		i++
	}
	fmt.Println(squareNumbers) // выводим слайс
}
