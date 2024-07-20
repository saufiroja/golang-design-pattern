package main

import "fmt"

// Product
type Pizza struct {
	Dough    string
	Sauce    string
	Cheese   string
	Toppings []string
}

// Builder interface
type PizzaBuilder interface {
	SetDough(dough string) *ConcretePizzaBuilder
	SetSauce(sauce string) *ConcretePizzaBuilder
	SetCheese(cheese string) *ConcretePizzaBuilder
	SetToppings(toppings []string) *ConcretePizzaBuilder
	Build() *Pizza
}

// Concrete Builder
type ConcretePizzaBuilder struct {
	pizza *Pizza
}

func NewConcretePizzaBuilder() PizzaBuilder {
	return &ConcretePizzaBuilder{pizza: &Pizza{}}
}

func (pb *ConcretePizzaBuilder) SetDough(dough string) *ConcretePizzaBuilder {
	pb.pizza.Dough = dough
	return pb
}

func (pb *ConcretePizzaBuilder) SetSauce(sauce string) *ConcretePizzaBuilder {
	pb.pizza.Sauce = sauce
	return pb
}

func (pb *ConcretePizzaBuilder) SetCheese(cheese string) *ConcretePizzaBuilder {
	pb.pizza.Cheese = cheese
	return pb
}

func (pb *ConcretePizzaBuilder) SetToppings(toppings []string) *ConcretePizzaBuilder {
	pb.pizza.Toppings = toppings
	return pb
}

func (pb *ConcretePizzaBuilder) Build() *Pizza {
	return pb.pizza
}

// Director
type Director struct {
	builder PizzaBuilder
}

func NewDirector(builder PizzaBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Pizza {
	return d.builder.SetDough("Thin Crust").SetSauce("Tomato").SetCheese("Mozzarella").SetToppings([]string{"Mushrooms", "Olives", "Onions"}).Build()
}

func main() {
	builder := NewConcretePizzaBuilder()
	director := NewDirector(builder)
	pizza := director.Construct()

	fmt.Println(pizza)
}
