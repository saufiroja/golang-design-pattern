package main

import (
	"errors"
	"fmt"
)

func main() {
	mediator := &ConcreteMediator{
		orderHandler:   OrderHandler{},
		productHandler: ProductHandler{},
		paymentHandler: PaymentHandler{},
	}

	orderCommand := OrderCommand{OrderID: "123"}
	productCommand := ProductCommand{ProductID: "456"}
	paymentCommand := PaymentCommand{PaymentID: "789"}

	// Mengirimkan orderCommand
	result, _ := mediator.Send(orderCommand)
	fmt.Println(result)

	// Mengirimkan productCommand
	result, _ = mediator.Send(productCommand)
	fmt.Println(result)

	// Mengirimkan paymentCommand
	result, _ = mediator.Send(paymentCommand)
	fmt.Println(result)

	// Mengirimkan productQuery
	productQuery := ProductQuery{ProductID: "456"}
	result, _ = mediator.Send(productQuery)
	fmt.Println(result)
}

// Mediator interface
type Mediator[T any] interface {
	Send(command T) (any, error)
}

// ConcreteMediator menggunakan generik
type ConcreteMediator struct {
	orderHandler        OrderHandler
	productHandler      ProductHandler
	paymentHandler      PaymentHandler
	productQueryHandler ProductQueryHandler
}

// Implementasi Send dengan menggunakan generik
func (cm *ConcreteMediator) Send(command interface{}) (interface{}, error) {
	switch cmd := any(command).(type) {
	case OrderCommand:
		return cm.orderHandler.Handle(cmd)
	case ProductCommand:
		return cm.productHandler.Handle(cmd)
	case PaymentCommand:
		return cm.paymentHandler.Handle(cmd)
	case ProductQuery:
		return cm.productQueryHandler.Handle(cmd)
	default:
		return nil, errors.New("Invalid command")
	}
}

// OrderHandler
type OrderCommand struct {
	OrderID string
}

type OrderHandler struct{}

func (oh *OrderHandler) Handle(command OrderCommand) (any, error) {
	// Menangani perintah order
	return fmt.Sprintf("Order %s handled", command.OrderID), nil
}

// ProductHandler
type ProductCommand struct {
	ProductID string
}

type ProductHandler struct{}

func (ph *ProductHandler) Handle(command ProductCommand) (any, error) {
	// Menangani perintah produk
	return fmt.Sprintf("Product %s handled", command.ProductID), nil
}

// PaymentHandler
type PaymentCommand struct {
	PaymentID string
}

type PaymentHandler struct{}

func (ph *PaymentHandler) Handle(command PaymentCommand) (any, error) {
	// Menangani perintah pembayaran
	return fmt.Sprintf("Payment %s handled", command.PaymentID), nil
}

type ProductQuery struct {
	ProductID string
}

type ProductQueryHandler struct{}

func (ph *ProductQueryHandler) Handle(query ProductQuery) (any, error) {
	return fmt.Sprintf("Product %s queried", query.ProductID), nil
}
