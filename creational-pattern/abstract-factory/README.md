# Abstract Factory

## Overview

The Abstract Factory Pater provieds a way o create related objects without specifying their concrete classes. This is useful in situation where we want to create objects based on a certain theme or context. For example, if we are building a GUI app, we might want to create a set of related obejects, such as windows, buttons and menus, that all have the same look and feel. By using the abstract factory pattern, we can create these objects without specifying their concrete classes, which makes it easy to change the theme or context of our app.

## Implementation

To Implement the abstract factory pattern in golang, we first define an interface that defines methods for create related objects. This interface is called the abstract factory.
We then define concrete implementations of this interface, which are called concrete factory.. The concrete factory are responsible for creating the actual objects.

```go
type WidgetFactory interface {
    CreateWindow() Window
    CreateButton() Button
    CreateMenu() Menu
}
```

In this example, we define an interface called `WidgetFactory`, which has three methods for creating related object: `CreateWindow`, `CreateButton`, and `CreateMenu`. These methods return objects of type `Window`, `Button`, and `Menu`, respectively.
We can then define concrete implementations of this interface, such as the `WinFactory` and `MacFactory` structs, which implement the `WidgetFactory` interface.

```go
type WinFactory struct{}

func NewWinFactory() *WinFactory {
    return &WinFactory{}
}

func (w *WinFactory) CreateWindow() Window {
    return &WinWindow{}
}

func (w *WinFactory) CreateButton() Button {
    return &WinButton{}
}

func (w *WinFactory) CreateMenu() Menu {
    return &WinMenu{}
}
```

```go
type MacFactory struct{}

func NewMacFactory() *MacFactory {
    return &MacFactory{}
}

func (m *MacFactory) CreateWindow() Window {
    return &MacWindow{}
}

func (m *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (m *MacFactory) CreateMenu() Menu {
    return &MacMenu{}
}
```

In this example, we define 2 concrete factory: `WinFactory` and `MacFactory`, which implement the `WidgetFactory` interface. Each concrete factory is responsible for creating objects of a specific family. For example, the `WinFactory` creates objects with a `Windows` look and feel, while the `MacFactory` creates objects with a `Mac` look and feel.

## Usage

To use the abstract factory pattern, we first create an instance of a concrete factory, such as the `WinFactory` or `MacFactory`. We can then use this factory to create objects of the related family.

```go
func main() {
    var factory WidgetFactory

    if runtime.GOOS == "windows" {
        factory = NewWinFactory()
    } else {
        factory = NewMacFactory()
    }

    window := factory.CreateWindow()
    button := factory.CreateButton()
    menu := factory.CreateMenu()

    fmt.Println(window)
    fmt.Println(button)
    fmt.Println(menu)
}
```
