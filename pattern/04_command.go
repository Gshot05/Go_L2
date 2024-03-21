package command

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

// Command - интерфейс команды
type Command interface {
	Execute()
}

// Receiver - получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// ConcreteCommand - конкретная реализация команды
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Invoker - вызывающий объект
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// Создаем получателя
	receiver := &Receiver{}

	// Создаем команду и передаем получателя
	command := NewConcreteCommand(receiver)

	// Создаем вызывающий объект
	invoker := &Invoker{}

	// Назначаем команду вызывающему объекту
	invoker.SetCommand(command)

	// Вызываем команду через вызывающий объект
	invoker.ExecuteCommand()
}
