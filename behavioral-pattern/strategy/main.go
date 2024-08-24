package main

import "fmt"

func main() {
	creditCard := NewCreditCardPayment("1234 5678 9012 3456")
	payPal := NewPayPalPayment("test@mail.com")
	bankTransfer := NewBankTransferPayment("1234567890")

	paymentContext := &PaymentContext{}

	paymentContext.SetPaymentStrategy(creditCard)
	fmt.Println(paymentContext.ProcessPayment(100))

	paymentContext.SetPaymentStrategy(payPal)
	fmt.Println(paymentContext.ProcessPayment(200))

	paymentContext.SetPaymentStrategy(bankTransfer)
	fmt.Println(paymentContext.ProcessPayment(300))
}

// PaymentStrategy is the strategy interface
type PaymentStrategy interface {
	Pay(amount int) string
}

// CreditCardPayment is a concrete strategy
type CreditCardPayment struct {
	CardNumber string
}

func NewCreditCardPayment(cardNumber string) PaymentStrategy {
	return &CreditCardPayment{CardNumber: cardNumber}
}

func (c *CreditCardPayment) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using credit card %s", amount, c.CardNumber)
}

// PayPalPayment is a concrete strategy
type PayPalPayment struct {
	Email string
}

func NewPayPalPayment(email string) PaymentStrategy {
	return &PayPalPayment{Email: email}
}

func (p *PayPalPayment) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using PayPal %s", amount, p.Email)
}

// bankTransferPayment is a concrete strategy
type BankTransferPayment struct {
	AccountNumber string
}

func NewBankTransferPayment(accountNumber string) PaymentStrategy {
	return &BankTransferPayment{AccountNumber: accountNumber}
}

func (b *BankTransferPayment) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using bank transfer %s", amount, b.AccountNumber)
}

// PaymentContext is the context
type PaymentContext struct {
	PaymentStrategy PaymentStrategy
}

func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}

func (p *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	p.PaymentStrategy = strategy
}

func (p *PaymentContext) ProcessPayment(amount int) string {
	if p.PaymentStrategy == nil {
		return "No payment strategy set"
	}

	return p.PaymentStrategy.Pay(amount)
}
