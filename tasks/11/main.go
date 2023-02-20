package main

import "fmt"

func intersect(a, b []int) []int {
	set := make(map[int]bool) // создаем мапу в которой ключ будет элемент массива, а значение по этому ключу есть ли он уже в мапе или нет
	var result []int

	for _, val := range a {
		set[val] = true // заполняем первым множеством
	}

	for _, val := range b {
		if set[val] { // проверяем есть ли уже этот элемент в множестве set, если есть добавляем его в пересечение
			result = append(result, val)
		}
	}

	return result
}

func main() {
	var set1 = []int{1, 2, 3, 4}
	var set2 = []int{3, 4, 6, 5}
	set3 := intersect(set1, set2)
	for _, v := range set3 {
		fmt.Println(v)
	}
}
