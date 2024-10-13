package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

// Point представляет точку в 2D-пространстве с инкапсулированными свойствами x и y
type Point struct {
	x, y float64
}

// NewPoint создает новую точку и возвращает указатель на нее
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance вычисляет расстояние между текущей точкой и другой точкой
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Создание двух точек
	pointA := NewPoint(3, 4)
	pointB := NewPoint(7, 1)

	// Вычисление расстояния между точками
	distance := pointA.Distance(pointB)

	// Вывод результата
	fmt.Printf("Расстояние между точками A(%.2f, %.2f) и B(%.2f, %.2f) равно %.2f\n",
		pointA.x, pointA.y, pointB.x, pointB.y, distance)
}
