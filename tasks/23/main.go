package main

import "fmt"

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[:index+1]...) // Возвращаем срез из элементов массива до индекса и с индекс + 1
}

func main() {
	var index = 4
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = remove(slice, index)
	for i := 0; i < len(slice); i++ { // печатаем массив
		fmt.Println(slice[i])
	}
}
