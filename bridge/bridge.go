package bridge

import "fmt"

type IColor interface {
	ShowColor()
}

type ICar interface {
	ShowCar()
}

type SmallCar struct {
	IColor IColor
}

func (s *SmallCar) ShowCar() {
	fmt.Println("this is small car")
	s.IColor.ShowColor()
}

type BigCar struct {
	IColor IColor
}

func (b *BigCar) ShowCar() {
	fmt.Println("this is big car")
	b.IColor.ShowColor()
}

type Red struct {
}

func (r *Red) ShowColor() {
	fmt.Println("red color")
}

type Black struct {
}

func (b *Black) ShowColor() {
	fmt.Println("black color")
}
