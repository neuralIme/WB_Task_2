package _3_visitor

import "fmt"

/*
Посетитель — это поведенческий паттерн, который позволяет добавить
новую операцию для целой иерархии классов, не изменяя код этих классов.

Минусы:
Паттерн не оправдан, если иерархия элементов часто меняется;
Может привести к нарушению инкапсуляции элементов.
*/

// Задает поведение для объектов
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

// shape для связи объектов с интерфейсом visitor
type shape interface {
	accept(visitor)
}

// areaCalculator реализуя интерфейс visitor задает поведение объектам
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForRectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

// middleCoordinates реализуя интерфейс visitor задает поведение объектам
type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForRectangle(t *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
