package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Cloneable {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) Clone() Cloneable {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
}

func main() {
	person1 := NewPerson("John", 30)
	person2 := person1.Clone().(*Person)

	fmt.Println(person1)
	fmt.Println(person2)
}
