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

	fmt.Printf("After update1 %v\n", zoo)

	for idx, _ := range zoo {
		zoo[idx].legs = 999
	}

	fmt.Printf("After update2 %v\n", zoo)

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
}
