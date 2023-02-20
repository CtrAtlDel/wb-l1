package main

import (
	"fmt"
)

func input(numbers []int, input chan int) {
	for i := range numbers {
		input <- numbers[i]
	}
	close(input) // больше не будем записывать информацию в канал
}

func output(input, output chan int) {
	for i := range input { // итерируемся по входному каналу
		output <- i * 2 
	}
	close(output) // больше не будем записывать информацию в канал
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8} // массив взят для примера 
	inputChan := make(chan int) // В этот канал записываются числа из массива
	outputChan := make(chan int) // В этот результат операции x^2
	go input(numbers, inputChan)
	go output(inputChan, outputChan)

	for i := range outputChan { // вывод из канала x^2
		fmt.Println(i)
	}
}
