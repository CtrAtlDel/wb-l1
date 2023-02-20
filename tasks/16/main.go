package main

import "fmt"

func quickSort(array []int, start int, end int) []int {
	if start < end {
		var p int
		array, p = partition(array, start, end) 
		array = quickSort(array, start, p-1)
		array = quickSort(array, p+1, end)
	}

	return array
}

func quickSortStart(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

func partition(arr []int, low, high int) ([]int, int) { // разбиваем на части, согласно алгоритму сортировки
	pivot := arr[high] //
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot { // меняем местами элементы, если числа меньше крайнего правого числа в секции
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] //
	return arr, i // возвращаем отсортированный кусок и индекс
}

// Реализация алгоритма быстрой сортировки
func main() {
	fmt.Println(quickSortStart([]int{5, 6, 7, 2, 1, 0}))
}
