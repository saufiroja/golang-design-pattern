package main

import "fmt"

func main() {
	factory := NewTransportationFactory()
	truck := factory.CreateTransportation("Truck")
	pesawat := factory.CreateTransportation("Pesawat")
	kapal := factory.CreateTransportation("Kapal")

	fmt.Println(truck.SetShip())
	fmt.Println(pesawat.SetShip())
	fmt.Println(kapal.SetShip())
}

type Tranportation interface {
	SetShip() string
}

// concrete product
type Truck struct {
	Name string
}

func NewTruck(name string) *Truck {
	return &Truck{
		Name: name,
	}
}

func (t *Truck) SetShip() string {
	return t.Name
}

// concrete product
type Pesawat struct {
	Name string
}

func NewPesawat(name string) *Pesawat {
	return &Pesawat{
		Name: name,
	}
}

func (t *Pesawat) SetShip() string {
	return t.Name
}

// concrete product
type Kapal struct {
	Name string
}

func NewKapal(name string) *Kapal {
	return &Kapal{
		Name: name,
	}
}

func (t *Kapal) SetShip() string {
	return t.Name
}

// factory method
type factory interface {
	CreateTransportation(transportationType string) Tranportation
}

type TransportationFactory struct{}

func NewTransportationFactory() factory {
	return &TransportationFactory{}
}

func (t *TransportationFactory) CreateTransportation(transportationType string) Tranportation {
	switch transportationType {
	case "Truck":
		return NewTruck("Truck")
	case "Pesawat":
		return NewPesawat("Pesawat")
	case "Kapal":
		return NewKapal("Kapal")
	default:
		return nil
	}
}
