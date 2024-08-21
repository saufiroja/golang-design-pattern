package main

import "fmt"

func main() {
	eventManager := &EventManager{}
	emailNotifier := &EmailNotification{}
	smsNotifier := &SMSNotification{}

	eventManager.Register(emailNotifier)
	eventManager.Register(smsNotifier)

	eventManager.Notify("User logged in")
}

// observer interface
type Observer interface {
	Update(event string)
}

// concrete observer A
type EmailNotification struct{}

func (e *EmailNotification) Update(event string) {
	fmt.Println("EmailNotifier received event:", event)
}

// concrete observer B
type SMSNotification struct{}

func (s *SMSNotification) Update(event string) {
	fmt.Println("SMSNotifier received event:", event)
}

// subject interface
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(event string)
}

// concrete subject
type EventManager struct {
	observers []Observer
}

func (e *EventManager) Register(observer Observer) {
	e.observers = append(e.observers, observer)
}

func (e *EventManager) Unregister(observer Observer) {
	for i, o := range e.observers {
		if o == observer {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
			break
		}
	}
}

func (e *EventManager) Notify(event string) {
	for _, observer := range e.observers {
		observer.Update(event)
	}
}
