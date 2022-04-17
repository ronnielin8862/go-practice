package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

func main() {
	dog := Dog{Name: "Kenny"}
	ShowEat(&dog)
	ShowRun(&dog)

	hello1(dog)
	hello1(1234)

	hello2("yayaya")
}

//此種interface 可以傳入 int、string 或者是建構體都行，這令你在設計程式時擁有足夠的彈性來接收任意型態的值。
func hello1(value interface{}) {
	// 透過型態斷言揭露 interface{} 真正的型態。
	switch v := value.(type) {
	// 如果 value 是字串型態。
	case string:
		fmt.Println("value 是字串，內容是 " + v)
	// 如果 value 是 int 型態。
	case int:
		fmt.Println("value 是數值，加上二就是 ", v+2)
	}
}

func hello2(value interface{}) {
	//if you know what is that type, you can do this is quicker
	v := value.(string)
	fmt.Println("value must be string : ", v)
}
