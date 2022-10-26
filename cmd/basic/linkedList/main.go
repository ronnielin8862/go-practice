package main

import (
	"container/list"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	userList := list.New()

	// 建立3筆 資料 放入linkedList
	a := User{Name: "a", Age: 1}
	_ = userList.PushBack(a)
	b := User{Name: "b", Age: 2}
	bElement := userList.PushBack(b)
	userList.PushBack(User{Name: "c", Age: 3})

	// 直接打印linkedList 出來的是記憶體位置
	fmt.Println("userList: ", userList)

	// 插入資料於b之後
	userList.InsertAfter(User{Name: "d", Age: 4}, bElement)

	// 看到結果是 1 2 4 3
	for e := userList.Front(); e != nil; e = e.Next() {
		user := e.Value.(User)
		fmt.Printf("%+v", user)
	}
}
