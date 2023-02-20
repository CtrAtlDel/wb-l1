package main

import (
	"fmt"
	"sync"
)

func sqr(x int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // декрементируем счетчик
	c <- x * x // записываем значение квадратов (порядок не важен)
}


func main() {
	numbers := []int{2, 4, 6, 8, 10} // взят из условия
	c := make(chan int, len(numbers))
	sum := 0
	var wg sync.WaitGroup

	for _, j := range numbers {
		wg.Add(1) // увеличиваем счетчик горутин
		go sqr(j, c, &wg) // ps используем только один экземпляр wg, для корректной работы
	}
	wg.Wait() // ждем чтобы main поток не завершился раньше горутин
	close(c) // закрываем канал для записи
	for j := range c { // итерируемся по каналу и суммируем данные
		sum += j
	}
	fmt.Printf("Sum: %d \n", sum)
}
