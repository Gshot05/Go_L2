package builder

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import "fmt"

// Product - продукт, который мы строим с помощью паттерна Строитель
type Product struct {
	parts []string
}

// Builder - интерфейс строителя, определяющий методы для пошагового создания продукта
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() *Product
}

// ConcreteBuilder - конкретный строитель, реализующий интерфейс Builder
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

// BuildPart1 - метод для построения первой части продукта
func (b *ConcreteBuilder) BuildPart1() {
	b.product.parts = append(b.product.parts, "Part1")
}

// BuildPart2 - метод для построения второй части продукта
func (b *ConcreteBuilder) BuildPart2() {
	b.product.parts = append(b.product.parts, "Part2")
}

// BuildPart3 - метод для построения третьей части продукта
func (b *ConcreteBuilder) BuildPart3() {
	b.product.parts = append(b.product.parts, "Part3")
}

// GetProduct - метод для получения готового продукта
func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Director - директор, который контролирует процесс построения продукта
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct - метод, который определяет порядок построения продукта
func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

func main() {
	// Создаем конкретного строителя
	builder := NewConcreteBuilder()

	// Создаем директора и передаем ему строителя
	director := NewDirector(builder)

	// Директор контролирует процесс построения продукта
	director.Construct()

	// Получаем готовый продукт от строителя
	product := builder.GetProduct()

	// Выводим информацию о продукте
	fmt.Println("Product parts:", product.parts)
}
