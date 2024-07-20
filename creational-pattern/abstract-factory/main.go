package main

import (
	"fmt"
)

// abstract factory interface
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func NewSportsFactory(brand string) (ISportsFactory, error) {
	return &Adidas{}, nil
}

// abstract product
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// abstract product
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

// concrete product
type AdidasShoe struct {
	Shoe
}

// concrete product
type AdidasShirt struct {
	Shirt
}

// concrete factory
type Adidas struct{}

func NewAdidas() ISportsFactory {
	return &Adidas{}
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

func main() {
	factory, err := NewSportsFactory("adidas")
	if err != nil {
		fmt.Println(err)
		return
	}

	shirt := factory.makeShirt()
	shoe := factory.makeShoe()

	fmt.Printf("Logo: %s\n", shirt.getLogo())
	fmt.Printf("Size: %d\n", shirt.getSize())

	fmt.Printf("Logo: %s\n", shoe.getLogo())
	fmt.Printf("Size: %d\n", shoe.getSize())
}
