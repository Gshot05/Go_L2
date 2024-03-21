package streategy

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

// Strategy - интерфейс стратегии
type Strategy interface {
	DoOperation(int, int) int
}

// ConcreteStrategyAdd - конкретная реализация стратегии сложения
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// ConcreteStrategySubtract - конкретная реализация стратегии вычитания
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// Context - контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем стратегию сложения
	context.SetStrategy(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("10 + 5 =", result) // 10 + 5 = 15

	// Устанавливаем стратегию вычитания
	context.SetStrategy(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("10 - 5 =", result) // 10 - 5 = 5
}
