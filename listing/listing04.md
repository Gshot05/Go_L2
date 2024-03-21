Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Deadlock. Так как мы открываем канал, но не закрываем его, и горутина читатель ожидает ещё данные для чтение.

```