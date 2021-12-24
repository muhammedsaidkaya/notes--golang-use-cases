package main

import (
	"fmt"
	"strconv"
)

type Shape interface {
	area() float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Square struct {
	value float32
}


func (rectangle Rectangle) area() float32 {
	return rectangle.height * rectangle.width
}

func (square Square) area() float32 {
	return square.value * square.value
}


func (rectangle *Rectangle) setWidth(width float32) {
	rectangle.width = width
}

func (rectangle *Rectangle) setHeight(height float32) {
	rectangle.height = height
}

func (square *Square) setValue(value float32)  {
	square.value = value
}

func add(x, y int) int {
	return x + y
}

func main() {
	var rectangle = &Rectangle{height: 2}
	var square = &Square{value: 5}

	rectangle.setWidth(10)
	square.setValue(20)

	liste := []Shape{rectangle, square}

	for _, shape := range liste {
		fmt.Println(shape.area())
	}


	out := fmt.Sprintf("%s", strconv.Itoa(add(2, 5)))
	fmt.Println(out)

	test := func() string {
		return fmt.Sprintf("%s", "hello world")
	}
	fmt.Printf(test())
}
