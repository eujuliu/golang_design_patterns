package main

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementos[index]
}
