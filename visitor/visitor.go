package visitor

import "fmt"

type Visitor interface {
	visitForSquare(square *Square)
	visitForCircle(circle *Circle)
}

type AreaVisitor struct {
}

func NewAreaVisitor() *AreaVisitor {
	return &AreaVisitor{}
}

func (impl *AreaVisitor) visitForSquare(square *Square) {
	fmt.Printf("AreaVisitor for square %+v\n", *square)
}

func (impl *AreaVisitor) visitForCircle(circle *Circle) {
	fmt.Printf("AreaVisitor for circle %+v\n", *circle)
}

type LengthVisitor struct {
}

func NewLengthVisitor() *LengthVisitor {
	return &LengthVisitor{}
}

func (impl *LengthVisitor) visitForSquare(square *Square) {
	fmt.Printf("LengthVisitor for square %+v\n", *square)
}

func (impl *LengthVisitor) visitForCircle(circle *Circle) {
	fmt.Printf("LengthVisitor for circle %+v\n", *circle)
}

type Square struct {
	side int
}

func NewSquare(side int) *Square {
	return &Square{
		side: side,
	}
}

func (impl *Square) Accept(visitor Visitor) {
	visitor.visitForSquare(impl)
}

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (impl *Circle) Accept(visitor Visitor) {
	visitor.visitForCircle(impl)
}
