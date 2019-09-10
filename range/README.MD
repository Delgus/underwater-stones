### Range

Начнем программирование на Go с основ. 
Функция range является одной из самых используемых в Go. Вот пример использования:
```go
package main

import "fmt"

type Animal struct {
	name string
	legs int
}

func main() {
	zoo := []Animal{
		Animal{"Dog", 4},
		Animal{"Chicken", 2},
		Animal{"Snail", 0},
	}

	fmt.Printf("Before update %v\n", zoo)

	for _, animal := range zoo {
		animal.legs = 999
	}

	fmt.Printf("After update %v\n", zoo)
}
```
Вышеприведённый код выглядит довольно невинно. Однако вы можете удивиться, узнав, что два fmt.Printf() 
выражения дают одинаковые результаты
```
Before update [{Dog 4} {Chicken 2} {Snail 0}]
After update [{Dog 4} {Chicken 2} {Snail 0}]
```
Значения (хранятся как animal), по которым мы проходимся с помощью range, являются не указателями на значения, 
а копиями значений из zoo.

Чтобы изменить элемент слайса, мы должны изменить этот элемент через указатель:
```go
for idx, _ := range zoo {
		zoo[idx].legs = 999
}
```

Но! range по слайсу указателей возвращает указатель!
```go
zoo2 := []*Animal{
		&Animal{"Dog", 4},
		&Animal{"Chicken", 2},
		&Animal{"Snail", 0},
	}

	for _, animal := range zoo2 {
		animal.legs = 999
	}

	for _, animal := range zoo2 {
		fmt.Printf("%v\n",animal)
	}
```

```
&{Dog 999}
&{Chicken 999}
&{Snail 999}
```