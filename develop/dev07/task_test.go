package main

import (
	"testing"
	"time"
)

// Тест для функции or.
func TestOr(t *testing.T) {
	// Функция, создающая канал, который будет закрыт после задержки.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Создаем каналы с разными задержками.
	c1 := sig(2 * time.Hour)
	c2 := sig(5 * time.Minute)
	c3 := sig(1 * time.Second)

	// Вызываем функцию or, передавая каналы.
	out := or(c1, c2, c3)

	// Создаем канал для синхронизации завершения теста.
	done := make(chan struct{})

	// Запускаем горутину, которая будет проверять закрытие канала out.
	go func() {
		// Ожидаем закрытия канала out.
		<-out
		// Посылаем сообщение о завершении работы внутренней горутины.
		close(done)
	}()

	// Ожидаем завершения работы внутренней горутины.
	<-done
}
