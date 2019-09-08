### nil и интерфейсы

nil реализует любой интерфейс

```go
package main

type Sayer interface {
	Say()
}

func Talk(sayer Sayer){
	sayer.Say()
}

func main(){
	Talk(nil)
}
```

В данном случае компилятор не ругнется при компиляции программы, но впоследствии мы поймаем панику.