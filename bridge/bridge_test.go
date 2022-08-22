package bridge

import "testing"

func TestCarColor(t *testing.T) {
	b := &BigCar{IColor: &Red{}}
	b.ShowCar()

	s := &SmallCar{IColor: &Black{}}
	s.ShowCar()
}
