### Короткое объявление структуры

Для меня в один момент вдруг стало сюрпризом что структуры можно объявлять не указывая названия полей.

Правда, в коротком объявлении необходимо указывать значения всех полей.

Примеры:

Наша структура
```go
//User ...
type User struct {
	ID         int
	FirstName  string
	LastName   string
	Profession string
	private    bool
}
```
Полное объявление структуры

```go
user1 := User{
	ID:         1,
	FirstName:  "Vasya",
	LastName:   "Pupkin",
	Profession: "senior software engineer",
	private:    true,
}
```

В полном объявлении структуры можно указывать не все поля

```go
user2 := User{
    FirstName: "Vasya",
    LastName:  "Pupkin",
}
```

Короткое объявление структуры. необходимо указывать все поля, даже приватные

```go
user3 := User{2, "Valera", "Ivanov", "junior software engineer", false}
```

