package main

import "fmt"

func reverseString(str string) (result string) {
	for _, v := range str { // итерируемся по строке и добавляем в начало result значение итератора
		result = string(v) + result
	}
	return
}

func main() {
	str := "главрыба"
	fmt.Println(str)
	fmt.Println(reverseString(str))
}
