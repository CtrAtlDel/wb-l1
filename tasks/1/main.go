package main

import "fmt"

type Human struct { // Базовая структура Human с двумя рандомными полями
	name string
	age  int
}

func (h *Human) getHumanAge() int {
	return h.age
}

func (h *Human) getHumanName() string {
	return h.name
}

type Action struct { // Структура Action 
	Human 
}

func main() {
	var action = &Action{Human: Human{name: "Ivan", age: 21}}
	fmt.Println(action.getHumanAge())
	fmt.Println(action.getHumanName())
}

// https://golangify.com/composition-and-forwarding
// Можно добавить просто к типу, необязательно к структуре
// type library struct
// func (l *library) printLibrary() {
// 	for i,lb := range *l {
// 		fmt.Println(i, lb)
// 	}
// }
