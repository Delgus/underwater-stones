package main

import "fmt"

func main() {
	task1()
	task2()
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
