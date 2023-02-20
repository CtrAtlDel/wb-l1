package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, findNumber int) int {
	index := -1                                   // возвращаем если не нашли
	idx := sort.SearchInts(arr, findNumber)       // сортируем массив и возвращаем индекс элемента, который ищем
	if idx < len(arr) && arr[idx] == findNumber { // проверяем что индекс соответствует элементу
		index = idx
	}
	return index
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Для примера взят такой массив
	findNumber := 6                             // число, которое ищем в массиве
	index := binarySearch(arr, findNumber)      // поиск
	if index == -1 {
		fmt.Printf("Element %d not found in the array\n", findNumber)
	} else {
		fmt.Printf("Element %d found at index %d in the array\n", findNumber, index)
	}
}
