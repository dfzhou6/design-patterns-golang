package memento

import "testing"

func TestMemento(t *testing.T) {
	careTaker := NewCareTaker()
	originator := NewOriginator()
	originator.SetState("state 1")
	careTaker.AddMemento(originator.CreateMemento())
	originator.SetState("state 2")
	careTaker.AddMemento(originator.CreateMemento())
	originator.Show()
	originator.RestoreMemento(careTaker.GetMemento(0))
	originator.Show()
	originator.RestoreMemento(careTaker.GetMemento(1))
	originator.Show()
}
