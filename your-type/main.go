package main

import "fmt"

//MySlice ...
type MySlice []int

//Add ...
func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

//Count ...
func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	sl := MySlice([]int{1, 2, 3})
	sl.Add(8)
	fmt.Println(sl.Count(), sl)
}
