package main

import (
	"fmt"
	"math/big" // Будем использовать библитоку для больших чисел
)

func main() {
	a := big.NewInt(2)
	a.Exp(a, big.NewInt(20), nil) // a = 2^20 возводим в степень
	a.Mul(a, big.NewInt(30))      // a = 2^20 * 30

	b := big.NewInt(2)
	b.Exp(b, big.NewInt(20), nil) // b = 2^20
	b.Mul(b, big.NewInt(15))      // b = 2^20 * 15

	c := big.NewInt(0) // Умножение 
	c.Mul(a, b)
	fmt.Printf("Multiplication: %v\n", c)

	d := big.NewInt(0)
	d.Div(a, b)
	fmt.Printf("Division: %v\n", d) // Деление

	e := big.NewInt(0) // Сложение 
	e.Add(a, b)
	fmt.Printf("Addition: %v\n", e)

	f := big.NewInt(0) // Вычитание
	f.Sub(a, b)
	fmt.Printf("Subtraction: %v\n", f)
}
