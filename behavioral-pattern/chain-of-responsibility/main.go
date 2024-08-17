package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID      string
	Product string
	Price   float64
}

type OrderHandler interface {
	SetNext(OrderHandler) OrderHandler
	Handle(*Order) error
}

type OrderProcessor struct {
	handler OrderHandler
}

func (o *OrderProcessor) SetHandler(handler OrderHandler) {
	o.handler = handler
}

func (o *OrderProcessor) Process(order *Order) error {
	return o.handler.Handle(order)
}

type ProductIdHandler struct {
	next OrderHandler
}

func (p *ProductIdHandler) SetNext(next OrderHandler) OrderHandler {
	p.next = next
	return next
}

func (p *ProductIdHandler) Handle(order *Order) error {
	if order.ID != "" {
		if p.next != nil {
			return p.next.Handle(order)
		}
		return nil
	}
	return errors.New("product id not filled")
}

type ProductNameHandler struct {
	next OrderHandler
}

func (p *ProductNameHandler) SetNext(next OrderHandler) OrderHandler {
	p.next = next
	return next
}

func (p *ProductNameHandler) Handle(order *Order) error {
	if order.Product != "" {
		if p.next != nil {
			return p.next.Handle(order)
		}
		return nil
	}
	return errors.New("product name not filled")
}

type ProductPriceHandler struct {
	next OrderHandler
}

func (p *ProductPriceHandler) SetNext(next OrderHandler) OrderHandler {
	p.next = next
	return next
}

func (b *ProductPriceHandler) Handle(order *Order) error {
	if order.Price != 0 {
		if b.next != nil {
			return b.next.Handle(order)
		}
		return nil
	}
	return errors.New("product price not filled")
}

func main() {
	order := &Order{
		ID:      "",
		Product: "Laptop",
		Price:   1000,
	}
	orderProcessor := &OrderProcessor{}

	productIDHandler := &ProductIdHandler{}
	productNameHandler := &ProductNameHandler{}
	productPriceHandler := &ProductPriceHandler{}

	productIDHandler.SetNext(productNameHandler).SetNext(productPriceHandler)

	orderProcessor.SetHandler(productIDHandler)

	if err := orderProcessor.Process(order); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("order is %s\n", order.Product)
	fmt.Println("Order processed successfully")
}
