package main

import (
	"fmt"
	"math"
)

type Point struct { // Структура точки имеет 2 координаты
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) DistanceTo(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy) // по теореме Пифагора рассчитаем расстояние между точками
}

func main() {
	p := NewPoint(1, 2) // 
	q := NewPoint(4, 6) // 
	distance := p.DistanceTo(q)
	fmt.Println(distance)
}
