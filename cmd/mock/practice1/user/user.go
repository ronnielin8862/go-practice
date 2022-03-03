package user

import "github.com/ronnielin8862/go-practice/cmd/mock/practice1/person"

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) string {
	return u.Person.Get(id)
}
