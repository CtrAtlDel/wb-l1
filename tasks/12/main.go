package main

import "fmt"

func createSet(strs []string) map[string]bool {
	set := make(map[string]bool) 
	for _, str := range strs {
		set[str] = true
	}
	return set
}


// Идея как в прошлой задаче:
// Создадим мапу, где ключ это слово, а значение - true, тогда можно узнать
// содержиться ли слово в мапе. Дубликатов не будет так как по ключу попадаем на то же слово
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := createSet(strs)
	fmt.Println(set)
}