package main

import "fmt"

func sum(a int, b int) (int, int) { // При таких операциях можно поменять местами две переменные
	a = a - b
	b = a + b
	a = -a + b
	return a, b
}

func xor(a int, b int) (int, int) { // Меняем местами по свойству xor
	a = a ^ b
	b = b ^ a
	a = a ^ b
	return a, b
}

func main() {
	a := 4
	b := 5
	fmt.Println(sum(a, b))
	fmt.Println(xor(a, b))
}
