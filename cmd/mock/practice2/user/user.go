package user

import "fmt"

type DB interface {
	Get(id int64) ([]User, error)
}

type DBI struct{}

func (DBI) Get(id int64) ([]User, error) {
	fmt.Println("enter Get")
	//fmt.Println("enter Get === ", d.Get(3))
	user := User{Id: id, Name: "userYOYO"}
	var users []User

	users = append(users, user)
	return users, nil
}

type User struct {
	Id   int64
	Name string
}

func GetUsers(db DB, id int64) []User {
	fmt.Println("GetUsers")
	data, err := db.Get(id)
	if err != nil {
		fmt.Println("err ", err)
		return nil
	}
	return data
}
