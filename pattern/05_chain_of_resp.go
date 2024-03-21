package chainofresp

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import "fmt"

// Handler - интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler - базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// ConcreteHandlerA - конкретная реализация обработчика A
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA: обработан запрос", request)
	} else if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerA: невозможно обработать запрос", request)
	}
}

// ConcreteHandlerB - конкретная реализация обработчика B
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB: обработан запрос", request)
	} else if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerB: невозможно обработать запрос", request)
	}
}

func main() {
	// Создаем обработчики
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Устанавливаем следующий обработчик для обработчика A
	handlerA.SetNext(handlerB)

	// Отправляем запросы на обработку
	handlerA.Handle("A") // обработчик A
	handlerA.Handle("B") // обработчик B
	handlerA.Handle("C") // невозможно обработать
}
