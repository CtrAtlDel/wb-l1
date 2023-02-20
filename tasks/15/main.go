package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // нужно объявить функцию createHugeString
	justString = v[:100]
}

func createHugeString(size int) string { // например можно создать строку через repeat
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
}
