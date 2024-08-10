package main

import "fmt"

// builder interface
type IPerson interface {
	SetName(name string) string
	SetAge(age int) int
	SetPhone(phone string) string
	SetPerson() *Person
}

// product
type Person struct {
	name  string
	age   int
	phone string
}

// concrete builder
type PersonBuilder struct {
	name  string
	age   int
	phone string
}

func newPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (p *PersonBuilder) SetName(name string) string {
	p.name = name
	return p.name
}

func (p *PersonBuilder) SetAge(age int) int {
	p.age = age
	return p.age
}

func (p *PersonBuilder) SetPhone(phone string) string {
	p.phone = phone
	return p.phone
}

func (p *PersonBuilder) SetPerson() *Person {
	return &Person{
		name:  p.name,
		age:   p.age,
		phone: p.phone,
	}
}

// director
type PersonDirector struct {
	builder IPerson
}

func newPersonDirector(builder IPerson) *PersonDirector {
	return &PersonDirector{builder: builder}
}

func (pd *PersonDirector) setBuilder(builder IPerson) {
	pd.builder = builder
}

func (pd *PersonDirector) build(name, phone string, age int) *Person {
	pd.builder.SetName(name)
	pd.builder.SetAge(age)
	pd.builder.SetPhone(phone)
	return pd.builder.SetPerson()
}

// client
func main() {
	builder := newPersonBuilder()
	director := newPersonDirector(builder)

	person := director.build("John Doe", "123456789", 30)

	fmt.Println(person)
}
