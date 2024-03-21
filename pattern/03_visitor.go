package visitor

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import "fmt"

// Element - интерфейс элемента, который должны посещать посетители
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA - конкретный элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB - конкретный элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor: VisitConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor: VisitConcreteElementB")
}

// Client - клиентский код, который использует посетителя для обработки элементов
type Client struct{}

func (c *Client) Operation(elements []Element, visitor Visitor) {
	for _, element := range elements {
		element.Accept(visitor)
	}
}

func main() {
	// Создаем элементы
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	// Создаем посетителя
	visitor := &ConcreteVisitor{}

	// Создаем клиента и выполняем операцию
	client := &Client{}
	client.Operation([]Element{elementA, elementB}, visitor)
}
