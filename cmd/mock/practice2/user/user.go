package user

import "fmt"

type DB interface {
	Get(id int64) ([]User, error)
	GetUsersNameReceiver(id1, id2 int64) string
	AAA(a int)
}

type DBI struct{}

func (DBI) Get(id int64) ([]User, error) {
	fmt.Println("enter Get")
	user := User{Id: id, Name: "userYOYO"}
	var users []User

	users = append(users, user)
	return users, nil
}

type User struct {
	Id   int64
	Name string
}

func GetUsersName(id1, id2 int64, db2 DB) string {
	fmt.Println("GetUsers")
	user1, err := db2.Get(id1)
	user2, err := db2.Get(id2)
	if err != nil {
		fmt.Println("err ", err)
		return err.Error()
	}
	user1Name := user1[0].Name
	user2Name := user2[0].Name

	return user1Name + " " + user2Name
}

func (db DBI) GetUsersNameReceiver(id1, id2 int64) string {
	fmt.Println("GetUsers")
	user1, err := db.Get(id1)
	user2, err := db.Get(id2)
	if err != nil {
		fmt.Println("err ", err)
		return err.Error()
	}
	user1Name := user1[0].Name
	user2Name := user2[0].Name

	return user1Name + " " + user2Name
}
