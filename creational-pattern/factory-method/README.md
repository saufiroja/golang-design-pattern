# Factory Method

The Factory Patern is a Creational Pattern that provides an interface for creating objects in a superclass, but allows
subclasses to alter the type of objects that will be created. In simpler terms, it is a way of creating an object without
specifying the exact class of object that will be created.

The Factory Method Pattern is useful when we want to create objects without knowing their exact class. Instead of creating objects
directly, we delegate the creation to a factory that can decide which class of object to create based on certain conditions.
This makes our code more flexible and easier to maintain.

# Implementation

In Golang, we can implement the Factory Method pattern using an interface and a struct. The interface will define a method that will be implemented by the struct to create objects. Here is an example:

```go
// Define an interface for creating objects
type Product interface {
    GetName() string
}

// Define a struct that implements the interface
type ConcreteProduct struct {
    name string
}

func NewConcreteProduct(name string) *ConcreteProduct {
    return &ConcreteProduct{name: name}
}

func (p *ConcreteProduct) GetName() string {
    return p.name
}

// factory method
type Factory interface {
    CreateProduct() Product
}

// Define a struct that implements the factory interface
type ConcreteFactory struct {
}

func NewConcreteFactory() *ConcreteFactory {
    return &ConcreteFactory{}
}

func (f *ConcreteFactory) CreateProduct() Product {
    return &ConcreteProduct{name: "ConcreteProduct"}
}
```

In this example, we have defined an interface `Product` with a method `GetName()` which will be implemented by the struct `ConcreteProduct`. We also define the `factory` interface with a method `CreateProduct()` which will be implemented by the struct `ConcreteFactory`.

The `CreateProduct()` method in the `ConcreteFactory` struct creates an instance of the `ConcreteProduct` struct and returns it as a `Product` interface. This way, the client code can create objects without knowing the exact class of the object.

# Usage

Here is an example of how we can use the Factory Method pattern to create objects:

```go
func main() {
    factory := NewConcreteFactory()
    product := factory.CreateProduct()
    fmt.Println(product.GetName())
}
```

In this example, we create an instance of the `ConcreteFactory` struct and use it to create an instance of the `ConcreteProduct` struct. We then call the `GetName()` method on the product to get its name.

# Advantages of Using the Factory Method Pattern

- Flexibility: The Factory Method pattern allows us to create objects without knowing their exact class. This makes our code more flexible and easier to maintain.
- Scalability: The Factory Method pattern allows us to add new classes of objects without modifying the existing code. This makes our code more scalable and easier to extend.
- Decoupling: The Factory Method pattern decouples the creation of objects from their usage. This makes our code more modular and easier to test.

# Conclusion

The Factory Pattern is a useful pattern for creating objects without knowing their exact class. In Golang, we can implement this pattern using an interface and a struct.
By using this pattern, we can make our code more flexible. So, if you want to improve the scalability and maintainability of your code, consider using the Factory Method pattern.

# References

https://blog.matthiasbruns.com/golang-factory-method-pattern
