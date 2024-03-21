Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод: 2 1. defer'ы добавляются в стек обратным порядком(последний defer исоплнится первый). Каждый вызов defer связан с текущим контекстом выполнения функции, то есть с кадром стека, где он был определен.

```
