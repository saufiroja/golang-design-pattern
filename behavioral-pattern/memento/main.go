package main

import "fmt"

func main() {
	doc := &Document{content: "Hello, World!"}

	saveContent := &SaveContent{}

	// save the initial state of the document
	saveContent.Save(doc.Save())

	// modify the document
	doc.content = "Hello, Go!"

	// save the modified state of the document
	saveContent.Save(doc.Save())

	// // restore the initial state of the document
	// doc.Restore(saveContent.Pop())
	// fmt.Println(doc.content) // Output: Hello, World!

	// modify the document
	doc.content = "Hello, Golang!"

	// save the modified state of the document
	saveContent.Save(doc.Save())

	// restore the initial state of the document
	doc.Restore(saveContent.Pop())
	fmt.Println(doc.content) // Output: Hello, Go!
}

// Memento
type Memento struct {
	state string
}

// Originator
type Originator interface {
	Save() Memento
	Restore(m Memento)
}
type Document struct {
	content string
}

func (d *Document) Save() Memento {
	return Memento{state: d.content}
}

func (d *Document) Restore(m Memento) {
	d.content = m.state
}

// Caretaker
type SaveContent struct {
	mementos []Memento
}

func (s *SaveContent) Save(m Memento) {
	s.mementos = append(s.mementos, m)
}

func (s *SaveContent) Pop() Memento {
	if len(s.mementos) == 0 {
		return Memento{}
	}

	mementos := s.mementos[:len(s.mementos)-1]
	memento := mementos[len(mementos)-1]
	return memento
}
