package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64) // делаем мапу где ключ число с шагом и значение - группа элементов

	for _, temp := range temps {
		lowerBound := int(math.Floor(temp/10.0)) * 10         // получаем нижнее значение и будем использовать его как ключ
		groups[lowerBound] = append(groups[lowerBound], temp) // добавляем значение по ключу
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temps)

	for k, v := range groups { // Печатаем группы
		fmt.Printf("%d: %v\n", k, v)
	}
}
