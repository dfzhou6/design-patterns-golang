package bridge

import "testing"

func TestBridge(t *testing.T) {
	square := NewSquare(&Red{})
	square.ShowShape()
	circle := NewCircle(&Blue{})
	circle.ShowShape()
}
