package main

import (
	"fmt"
	"reflect"
)

type myType struct {
	someField int
}

func getSomething(i int) interface{} {
	if i == 1 {
		return 5
	} else if i == 2 {
		return "hello"
	} else if i == 3 {
		c := make(chan int)
		return c
	} else if i == 4 { // можно вернуть и собственный тип данных
		var m myType
		return m
	}
	return true
}

func otherCheck(o interface{}) { // Также можно получить через reflection информацию о типе
	t := reflect.TypeOf(o)  // Получаем тип
	v := reflect.ValueOf(o) // Получаем значение
	fmt.Println("type - ", t)
	fmt.Println("value -", v)
	fmt.Println("==========")
}

func main() {

	// будем использовать "переключатель типов"
	if r, ok := getSomething(1).(int); ok {
		fmt.Println(r)
	}
	if r, ok := getSomething(2).(string); ok {
		fmt.Println(r)
	}
	if r, ok := getSomething(3).(chan int); ok {
		fmt.Println(r)
	}
	if r, ok := getSomething(6).(bool); ok {
		fmt.Println(r)
	}
	if r, ok := getSomething(4).(myType); ok {
		fmt.Println(r)
	}
	otherCheck(getSomething(4))
	otherCheck(getSomething(1))
}
