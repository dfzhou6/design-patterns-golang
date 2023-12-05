package bridge

import "fmt"

type Shape interface {
	ShowShape()
}

type Color interface {
	ShowColor()
}

type BaseShape struct {
	c Color
}

func (impl *BaseShape) ShowShape() {
	panic("pls implement me")
}

type Square struct {
	*BaseShape
}

func NewSquare(color Color) *Square {
	return &Square{
		&BaseShape{c: color},
	}
}

func (impl *Square) ShowShape() {
	fmt.Println("this is shape Square")
	impl.c.ShowColor()
}

type Circle struct {
	*BaseShape
}

func NewCircle(color Color) *Circle {
	return &Circle{
		&BaseShape{c: color},
	}
}

func (impl *Circle) ShowShape() {
	fmt.Println("this is shape circle")
	impl.c.ShowColor()
}

type Red struct {
}

func (impl *Red) ShowColor() {
	fmt.Println("this is color red")
}

type Blue struct {
}

func (impl *Blue) ShowColor() {
	fmt.Println("this is color blue")
}
