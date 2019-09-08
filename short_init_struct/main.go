package main

import "fmt"

//User ...
type User struct {
	ID         int
	FirstName  string
	LastName   string
	Profession string
	private    bool
}

func main() {
	//полное объявление структуры
	user1 := User{
		ID:         1,
		FirstName:  "Vasya",
		LastName:   "Pupkin",
		Profession: "senior software engineer",
		private:    true,
	}
	//можно указывать не все поля
	user2 := User{
		FirstName: "Vasya",
		LastName:  "Pupkin",
	}
	//короткое объявление структуры - необходимо заполнять все поля
	user3 := User{2, "Valera", "Ivanov", "junior software engineer", false}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
}
