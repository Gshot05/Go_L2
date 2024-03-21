package facade

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import "fmt"

// Subsystem1 - компонент подсистемы
type Subsystem1 struct{}

// Operation1 - метод первого компонента подсистемы
func (s *Subsystem1) Operation1() string {
	return "Subsystem1: Operation1"
}

// Operation2 - метод второго компонента подсистемы
func (s *Subsystem1) Operation2() string {
	return "Subsystem1: Operation2"
}

// Subsystem2 - еще один компонент подсистемы
type Subsystem2 struct{}

// Operation1 - метод первого компонента второй подсистемы
func (s *Subsystem2) Operation1() string {
	return "Subsystem2: Operation1"
}

// Operation2 - метод второго компонента второй подсистемы
func (s *Subsystem2) Operation2() string {
	return "Subsystem2: Operation2"
}

// Facade - фасад, предоставляющий унифицированный интерфейс к подсистеме
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

// NewFacade - конструктор фасада, инициализирующий подсистемы
func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

// Operation - упрощенный метод, скрывающий сложность подсистемы от клиента
func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	// Фасад делегирует вызовы методов подсистемы через унифицированный интерфейс
	result += f.subsystem1.Operation1()
	result += f.subsystem1.Operation2()
	result += f.subsystem2.Operation1()
	result += f.subsystem2.Operation2()
	return result
}

func main() {
	// Использование фасада
	facade := NewFacade()
	result := facade.Operation()
	fmt.Println(result)
}
