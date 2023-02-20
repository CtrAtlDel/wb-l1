package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)                            // разобьем строку на слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // будем менять местами i, j элементы пока не дойдем до середины
		words[i], words[j] = words[j], words[i] // поменяем местами i и j элемент, где i - начальный элемент строки, j- конечный
	}
	return strings.Join(words, " ") // Добавим пробел к каждому слову
}

func main() {
	input := "snow dog sun — sun dog snow"
	output := reverseWords(input)
	fmt.Println(output)
}
