package main

type IPizza interface {
	getPrice() int
}

type VeggieMania struct {}

func (v*VeggieMania) getPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (c * TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()

	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()

	return pizzaPrice + 10
}
