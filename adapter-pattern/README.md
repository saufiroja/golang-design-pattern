# Adapter Pattern

## Overview

The Adapter Pattern, is used to enable the interaction between two incompatible interfaces by creating a wrapper object that can translate requests between both interfaces. In Golang, this pattern is quite useful when you have two interfaces, but their functionalities do not match, and you need to make them work together.

## Implementation

The Adapter Pattern consists of four components:

1. **Target Interface**:
   The interface that the client is expecting to interact with.
2. **Adapter**:
   The object that implements the Target Interface and wraps the Adaptee.
3. **Adaptee**:
   The interface that needs to be adapted to be used by the client.
4. **Client**:
   The component that uses the Target Interface.

```go
type Target interface {
    Request() string
}

type Adaptee interface {
    SpecificRequest() string
}

type adapteeImpl struct {}

func (a *adapteeImpl) SpecificRequest() string {
    return "Adaptee request"
}

type adapter struct {
    adaptee Adaptee
}

func (a *adapter) Request() string {
    return "Adapter: " + a.adaptee.SpecificRequest()
}

func main() {
    adaptee := &adapteeImpl{}
    target := &adapter{
        adaptee: adaptee,
    }

    fmt.Println(target.Request())
}
```

In this example, we have two interfaces, Target and Adaptee, that are incompatible. We create a concrete implementation of Adaptee, adapteeImpl, and an adapter that implements the Target interface and wraps the Adaptee interface.

# Examples

Let's take an example of a payment gateway that accepts payments from different payment providers like Gopay, Ovo, paypal and Amazon Pay. Each payment provider has its own unique interface to connect with the payment gateway. If we want to add a new payment provider, we need to create a new interface and implement it.

To solve this problem, we can create an adapter for each payment provider that implements the payment gateway's standard interface. Let's take an example of PayPal payment provider:

```go
// gopay.go

type Gopay struct {}

func (p *Gopay) MakePayment(amount float32) bool {
    // connect to Gopay and process payment
    return true
}

// ovo.go

type Ovo struct {}

func (a *Ovo) PayOvo(amount float32) bool {
    // connect to Ovo and process payment
    return true
}

// gateway.go

type PaymentGateway interface {
    ProcessPayment(amount float32) bool
}

type GopayAdapter struct {
    Gopay *Gopay
}

func (p *GopayAdapter) ProcessPayment(amount float32) bool {
    return p.Gopay.MakePayment(amount)
}

type OvoAdapter struct {
    Ovo *Ovo
}

func (a *OvoAdapter) ProcessPayment(amount float32) bool {
    return a.Ovo.PayOvo(amount)
}

// main.go

func main() {
    paymentGateway := &GopayAdapter{
        Gopay: &Gopay{},
    }

    paymentGateway2 := &OvoAdapter{
        Ovo: &Ovo{},
    }

    paymentGateway.ProcessPayment(100)

    paymentGateway2.ProcessPayment(100)
}
```

In this example, we have a PaymentGateway interface that needs to be implemented by all the payment providers. The `Gopay` provider has its own unique interface, so we create an adapter, GopayAdapter, that implements the PaymentGateway interface and wraps the `Gopay` interface. Now, we can use the `Gopay` provider with the payment gateway without changing the payment gateway's implementation. We can also imagine to add an adapter for `Ovo`, which has its own API implemented in `Ovo`.

## Conclusion

In conclusion, the Adapter Pattern is a useful pattern in Golang when you have two incompatible interfaces that need to work together. It creates a wrapper object that can translate requests between both interfaces. The real-world example of a payment gateway shows how we can use this pattern to create adapters for different payment providers without changing the payment gateway's implementation.

## References

https://blog.matthiasbruns.com/golang-adapter-pattern
