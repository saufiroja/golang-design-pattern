package main

import (
	"fmt"
)

func main() {
	// create the receiver
	product := newProduct(1, "Product 1", 100.0)

	// create the commands
	createCommand := newCreateProductCommand(product)
	updateCommand := newUpdateProductCommand(product, "Product 1 Updated", 150.0)

	// create the invoker
	commandInvoker := &CommandInvoker{}
	commandInvoker.AddCommand(createCommand)
	commandInvoker.AddCommand(updateCommand)

	// execute the commands
	commandInvoker.ExecuteCommand()
}

// create the receiver
type Product struct {
	Id    int
	Name  string
	Price float64
}

func newProduct(id int, name string, price float64) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func (p *Product) createProduct() string {
	return fmt.Sprintf("Product created: %s", p.Name)
}

func (p *Product) updateProduct(name string, price float64) string {
	p.Name = name
	p.Price = price
	return fmt.Sprintf("Product updated: %s, New price: %.2f", p.Name, p.Price)
}

// command interface
type Command interface {
	Execute() error
}

// concrete command type for creating a product
type CreateProductCommand struct {
	Product *Product
}

func newCreateProductCommand(product *Product) *CreateProductCommand {
	return &CreateProductCommand{Product: product}
}

func (cpc *CreateProductCommand) Execute() error {
	fmt.Println(cpc.Product.createProduct())
	return nil
}

// concrete command type for updating a product
type UpdateProductCommand struct {
	Product *Product
	Name    string
	Price   float64
}

func newUpdateProductCommand(product *Product, name string, price float64) *UpdateProductCommand {
	return &UpdateProductCommand{
		Product: product,
		Name:    name,
		Price:   price,
	}
}

func (upc *UpdateProductCommand) Execute() error {
	fmt.Println(upc.Product.updateProduct(upc.Name, upc.Price))
	return nil
}

// invoker
type CommandInvoker struct {
	Command []Command
}

func (ci *CommandInvoker) AddCommand(c Command) {
	ci.Command = append(ci.Command, c)
}

func (ci *CommandInvoker) ExecuteCommand() error {
	for _, c := range ci.Command {
		if err := c.Execute(); err != nil {
			return err
		}
	}
	return nil
}
