package main

import "fmt"

type paypal struct {
}

func NewPaypal() *paypal {
	return &paypal{}
}

func (p *paypal) MakePayment(amount int) {
	fmt.Println("Paypal payment is done for amount: ", amount)
}

type amazon struct {
}

func NewAmazon() *amazon {
	return &amazon{}
}

func (a *amazon) PayAmazon(amount int) {
	fmt.Println("Amazon payment is done for amount: ", amount)
}

type paymentGateway interface {
	processPayment(amount int)
}

type paypalAdapter struct {
	paypalGateway *paypal
}

func (p *paypalAdapter) processPayment(amount int) {
	p.paypalGateway.MakePayment(amount)
}

type amazonAdapter struct {
	amazonGateway *amazon
}

func (a *amazonAdapter) processPayment(amount int) {
	a.amazonGateway.PayAmazon(amount)
}

func main() {
	paypal := NewPaypal()
	paypalAdapter := &paypalAdapter{
		paypalGateway: paypal,
	}

	amazon := NewAmazon()
	amazonAdapter := &amazonAdapter{
		amazonGateway: amazon,
	}

	paypalAdapter.processPayment(100)

	amazonAdapter.processPayment(200)
}

// type response interface {
// 	getResponse()
// }

// type client struct {
// }

// func (c *client) getResponse(r response) {
// 	r.getResponse()
// }

// type xmlResponse struct {
// }

// func (x *xmlResponse) xmlResponse() {
// 	fmt.Println("XML response")
// }

// type xmlAdapter struct {
// 	xmlGateway *xmlResponse
// }

// func (x *xmlAdapter) getResponse() {
// 	x.xmlGateway.xmlResponse()
// }

// type jsonResponse struct {
// }

// func (j *jsonResponse) jsonResponse() {
// 	fmt.Println("JSON response")
// }

// type jsonAdapter struct {
// 	jsonGateway *jsonResponse
// }

// func (j *jsonAdapter) getResponse() {
// 	j.jsonGateway.jsonResponse()
// }

// func main() {
// 	client := &client{}
// 	xml := &xmlResponse{}
// 	xmlAdapter := &xmlAdapter{
// 		xmlGateway: xml,
// 	}
// 	client.getResponse(xmlAdapter)

// 	json := &jsonResponse{}
// 	jsonAdapter := &jsonAdapter{
// 		jsonGateway: json,
// 	}

// 	client.getResponse(jsonAdapter)
// }
