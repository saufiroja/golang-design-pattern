# Prototype Pattern

The prototype pattern is a creational design pattern that allows obeject to be created from existing objects of the same class. This pattern is useful when creating objects is expensive or time consuming, and you want to avoid unnecessary duplication of objects.

To Implement the prototype pattern in go, you can define an interface that includes a method to clone the objects. Then, you can define a struct that impelements this interface and provides a clone method that returns a copy of itself.

## Example

```go
type Cloneable interface {
    Clone() Cloneable
}

type Person struct {
    Name string
    Age int
}

func NewPerson(name string, age int) Cloneable {
    return &Person{
        Name: name,
        Age: age,
    }
}

func (p *Person) Clone() Cloneable {
    return &Person{
        Name: p.Name,
        Age: p.Age,
    }
}
```

In this example, we define a `Clonable` interface that includes a `Clone` method. Then, we define a `Person` struct that implement this interface and provides a `Clone` that return a copy of itself.
In the `main` function, we create a `person1` object and then use the `clone` method to create a copy of it `person2`. We can then print both objects to confirm that `person2` is a copy of `person1`.

```go
func main() {
    person1 := NewPerson("John", 30)
    person2 := person1.Clone().(*Person)

    fmt.Printf("person1: %+v\n", person1)
    fmt.Printf("person2: %+v\n", person2)
}
```
