# Builder Pattern

The builder pattern is a creational pattern that lets you contruct complex objects step by step. The pattern allows you to produce different types and representations of an obect using the same construction code.

## Problem

Lets say we want to create a complex object that has many optional parameters. One way to do this is to create a contructor that takes in all the parameters and provides default values for the optional ones. This approach has a few downsides:

- It can be hard to remember the order of the parameters
- It can be hard to know which parameter are optional and which are required
- The constructor can become very long and hard to read

## Solution

The builder pattern provides a solution to these problem by seperating the contruction of a complex object from its representation. The builder pattern involves the following components:

- A builder interface that defines the steps for constructing the object
- A ConcreteBuilder struct that implements the builder interface and provides methods a way to contruct the object
- A Director struct that uses the builder to construct the object

## Example

```go
type Pizza struct {
    dough     string
    sauce     string
    cheese    string
    toppings  []string
}

type PizzaBuilder interface {
    SetDough(string) PizzaBuilder
    SetSauce(string) PizzaBuilder
    SetCheese(string) PizzaBuilder
    SetToppings([]string) PizzaBuilder
    Build() *Pizza
}

type ConcretePizzaBuilder struct {
    pizza *Pizza
}

func NewConcretePizzaBuilder() *ConcretePizzaBuilder {
    return &ConcretePizzaBuilder{pizza: &Pizza{}}
}

func (cpb *ConcretePizzaBuilder) SetDough(dough string) PizzaBuilder {
    cpb.pizza.dough = dough
    return cpb
}

func (cpb *ConcretePizzaBuilder) SetSauce(sauce string) PizzaBuilder {
    cpb.pizza.sauce = sauce
    return cpb
}

func (cpb *ConcretePizzaBuilder) SetCheese(cheese string) PizzaBuilder {
    cpb.pizza.cheese = cheese
    return cpb
}

func (cpb *ConcretePizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
    cpb.pizza.toppings = toppings
    return cpb
}

func (cpb *ConcretePizzaBuilder) Build() *Pizza {
    return cpb.pizza
}

type Director struct {
    builder PizzaBuilder
}

func NewDirector(builder PizzaBuilder) *Director {
    return &Director{builder: builder}
}

func (d *Director) Construct() *Pizza {
    return d.builder.SetDough("Thin Crust").SetSauce("Tomato").SetCheese("Mozzarella").SetToppings([]string{"Mushrooms", "Olives", "Onions"}).Build()
}
```

In this example, we define the `Pizza` struct and the `PizzaBuilder` interface. The `ConcretePizzaBuilder` struct implements the `ConcretePizzaBuilder` interface and provides a way to contruct the `Pizza` object. The `Director` struct uses the `Pizzabuilder` to construct the `Pizza` object. The `Director` object is not strictly necessary, but it provides a way to simplify the process of constructing the `Pizza` object.

We can use the `Director` and `ConcreteBuilder` to construct a `Pizza` object like this:

```go
builder := NewConcretePizzaBuilder()
director := NewDirector(builder)
pizza := director.Construct()
```
