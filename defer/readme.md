### Аргументы отложенных функций


Что выведет этот код
```go
package main

import "fmt"

func main() {
	task1()
}

func getSomeVars(text string) string {
	fmt.Println("exec getSomeVars")
	return text
}

func task1() {
	fmt.Println("TASK 1")
	text := "init text"
	defer fmt.Println("after work")
	defer fmt.Println(getSomeVars(text))
	text = "change text"
	fmt.Println("work")
}
```

Может показаться что раз функция `getSomeVars()` находится внутри блока `defer` то в переменной `text` должен оказаться 
измененный текст `change text`. Но вывод программы вас немного удивит.

```
TASK 1
exec getSomeVars
work
init text
after work
```

Так произошло потому что аргументы отложенных функций вычисляются при объявлении defer!

Как мы можем изменить наш код, чтобы получить в выводе измененный текст?
Использовать анонимную функцию например.

```go
func task2() {
	fmt.Println("TASK 2")
	text := "init text"
	defer fmt.Println("after work")
	defer func() {
		fmt.Println(getSomeVars(text))
	}()
	text = "change text"
	fmt.Println("work")
}
```

Тогда мы получим следующий вывод

```
TASK 2
work
exec getSomeVars
change text
after work
```