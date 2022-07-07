package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{Name: "John", Age: 30}
	user2 := User{Name: "Mary", Age: 25}
	user3 := User{Name: "Mary", Age: 25}

	fmt.Println(user1 == user2)
	fmt.Println(user2 == user3)
}
