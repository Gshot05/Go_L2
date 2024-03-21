package fabricMethod

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

// Product - интерфейс продукта
type Product interface {
	Use() string
}

// ConcreteProductA - конкретная реализация продукта A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A used"
}

// ConcreteProductB - конкретная реализация продукта B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B used"
}

// Creator - интерфейс создателя
type Creator interface {
	FactoryMethod() Product
}

// ConcreteCreatorA - конкретная реализация создателя A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB - конкретная реализация создателя B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func main() {
	// Создаем конкретные создатели
	creatorA := &ConcreteCreatorA{}
	creatorB := &ConcreteCreatorB{}

	// Используем фабричный метод для создания продуктов
	productA := creatorA.FactoryMethod()
	productB := creatorB.FactoryMethod()

	// Используем продукты
	fmt.Println(productA.Use()) // Product A used
	fmt.Println(productB.Use()) // Product B used
}
