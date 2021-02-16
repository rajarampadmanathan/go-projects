package main

import "fmt"

type shape interface {
	getArea() float64
}
type triagle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

func (t triagle) getArea() float64 {
	return t.base * t.height / 2
}

func (t square) getArea() float64 {
	return 2 * t.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{side: 4}
	printArea(s)
	t := triagle{height: 4, base: 10}
	printArea(t)
}
