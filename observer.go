package main

import "fmt"

type Observer interface {
	update(string)
	getID() string
}

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type Item struct {
	observers []Observer
	name      string
	inStock   bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)

	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) deregister(o Observer) {
	i.observers = removeFromSlice(i.observers, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}

func removeFromSlice(observers []Observer, observerToRemove Observer) []Observer {
	observersLength := len(observers)
	for i, observer := range observers {
		if observerToRemove.getID() == observer.getID() {
			observers[observersLength-1], observers[i] = observers[i], observers[observersLength-1]

			return observers[:observersLength-1]
		}
	}

	return observers
}

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}
