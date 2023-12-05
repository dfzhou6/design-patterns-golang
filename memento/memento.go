package memento

import "fmt"

type Memento struct {
	state string
}

func (impl *Memento) GetState() string {
	return impl.state
}

type Originator struct {
	state string
}

func NewOriginator() *Originator {
	return &Originator{}
}

func (impl *Originator) SetState(state string) {
	impl.state = state
}

func (impl *Originator) CreateMemento() *Memento {
	return &Memento{state: impl.state}
}

func (impl *Originator) RestoreMemento(memento *Memento) {
	impl.state = memento.GetState()
}

func (impl *Originator) Show() {
	fmt.Printf("originator's state is %v\n", impl.state)
}

type Caretaker struct {
	mementos []*Memento
}

func NewCareTaker() *Caretaker {
	return &Caretaker{
		mementos: make([]*Memento, 0),
	}
}

func (impl *Caretaker) AddMemento(memento *Memento) {
	impl.mementos = append(impl.mementos, memento)
}

func (impl *Caretaker) GetMemento(index int) *Memento {
	return impl.mementos[index]
}
