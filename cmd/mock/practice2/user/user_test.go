package user

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetUsers(t *testing.T) {
	db := &DBI{}
	users, err := db.Get(3)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	fmt.Println(users)
	users2 := GetUsers(db, 14)
	fmt.Println(users2)
}

func TestGetUsersMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	var users1 []User
	users1 = append(users1, User{Id: 3, Name: "userZZZZ"})
	m.EXPECT().Get(gomock.Eq(int64(2))).Return(users1, nil)
	m.EXPECT().Get(gomock.Eq(int64(4))).Return(users1, nil)
	//m.EXPECT().Get(gomock.Not(int64(3))).Return(nil, fmt.Errorf("err"))
	//m.EXPECT().Get(gomock.Any()).Return(users1, nil)

	//var users2 []User
	//users2 = append(users2, User{Id: 33, Name: "user~~~~~"})
	//m.EXPECT().Get(gomock.Eq(33)).Return(users2, nil)
	get, err := m.Get(4)
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	fmt.Println(get)
}
