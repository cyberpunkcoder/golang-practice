/*
author: cyberpunkprogrammer
date: 11-18-2020
tutorial: https://www.golang-book.com/books/intro/9
*/

package main

import (
	"fmt"
	"math"
)

var c Circle

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type MultiShape struct {
	shapes []Shape
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
