package state

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// State - интерфейс состояния
type State interface {
	Handle()
}

// ConcreteStateA - конкретная реализация состояния A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Обработка состояния A")
}

// ConcreteStateB - конкретная реализация состояния B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Обработка состояния B")
}

// Context - контекст, управляющий состоянием
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем начальное состояние
	context.SetState(&ConcreteStateA{})

	// Выполняем запрос
	context.Request() // Обработка состояния A

	// Изменяем состояние
	context.SetState(&ConcreteStateB{})

	// Выполняем запрос еще раз
	context.Request() // Обработка состояния B
}
