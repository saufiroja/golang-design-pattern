package main

import "fmt"

// abstract product
type Payment interface {
	ProccessPayment(amount float32) float32
	AuthorizePayment(credential string) bool
}

// concrete product
type Gopay struct {
	apiKey string
}

func newGopay() *Gopay {
	return &Gopay{}
}

func (g *Gopay) ProccessPayment(amount float32) float32 {
	fmt.Printf("Paying with Gopay amount %f\n", amount)
	return amount
}

func (g *Gopay) AuthorizePayment(credential string) bool {
	if credential != "123" {
		fmt.Println("AuthorizePayment failed")
		return false
	}
	fmt.Println("AuthorizePayment success")
	return true
}

type Ovo struct {
	token string
}

func newOvo() *Ovo {
	return &Ovo{}
}

func (o *Ovo) ProccessPayment(amount float32) float32 {
	fmt.Printf("Paying with Ovo amount %f\n", amount)
	return amount
}

func (o *Ovo) AuthorizePayment(credential string) bool {
	if credential != "test" {
		fmt.Println("AuthorizePayment failed")
		return false
	}
	fmt.Println("AuthorizePayment success")
	return true
}

// abstract factory
type PaymentMethodFactory interface {
	CreatePayment() Payment
}

// concrete factory

type GopayFactory struct {
}

func newGopayFactory() *GopayFactory {
	return &GopayFactory{}
}

func (g *GopayFactory) CreatePayment() Payment {
	return newGopay()
}

type OvoFactory struct {
}

func newOvoFactory() *OvoFactory {
	return &OvoFactory{}
}

func (o *OvoFactory) CreatePayment() Payment {
	return newOvo()
}

// client

func GetPaymentMethodFactory(paymentMethod string) PaymentMethodFactory {
	if paymentMethod == "gopay" {
		return newGopayFactory()
	}

	if paymentMethod == "ovo" {
		return newOvoFactory()
	}

	return nil
}

func main() {
	var paymentMethod string
	var credential string
	fmt.Scanln(&credential)
	fmt.Scanln(&paymentMethod)

	paymentMethodFactory := GetPaymentMethodFactory(paymentMethod)
	if paymentMethodFactory == nil {
		fmt.Println("Invalid payment method")
		return
	}

	payment := paymentMethodFactory.CreatePayment()
	if !payment.AuthorizePayment(credential) {
		fmt.Println("Payment failed")
		return
	}
	payment.ProccessPayment(100)
}
