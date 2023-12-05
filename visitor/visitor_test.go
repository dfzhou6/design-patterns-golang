package visitor

import "testing"

func TestVisitor(t *testing.T) {
	areaVisitor := NewAreaVisitor()
	lenVisitor := NewLengthVisitor()

	square := NewSquare(10)
	circle := NewCircle(20)

	square.Accept(areaVisitor)
	circle.Accept(areaVisitor)

	square.Accept(lenVisitor)
	circle.Accept(lenVisitor)
}
