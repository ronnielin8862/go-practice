package main

import (
	"fmt"
	"reflect"
)

type APlusB func(int, int) int

type MathStruct struct {
	a      int
	b      int
	aPlusB APlusB
}

func (c *MathStruct) DoMathStruct(a int, b int) {
	fmt.Println("3: 進入DoMathStruct", a-b)
}

func main() {
	mathStruct := MathStruct{
		a: 1,
		b: 2,
		aPlusB: func(a int, b int) int {
			fmt.Println(a + b)
			return a + b
		},
	}

	reflectMathStruct := reflect.ValueOf(&mathStruct)
	fmt.Println("1: ", reflectMathStruct)
	value1 := reflect.ValueOf(10)
	value2 := reflect.ValueOf(2)
	value := []reflect.Value{value1, value2}
	reflectMathStruct.MethodByName("DoMathStruct").Call(value)

	fmt.Println("4: = ", reflectMathStruct.NumMethod()) //1
	fmt.Println("5: = ", reflectMathStruct.Elem().Field(2).Call(value))

	//for i := 0; i < reflectMathStruct.Elem().NumField(); i++ {
	//	fmt.Println(reflectMathStruct.Elem().Field(i))
	//}
}

//func main() {
//	////i := 555
//	//
//	//animal := Animal{
//	//	name:     "cat",
//	//	behavior: "meow",
//	//}
//	//////typeOfI := reflect.TypeOf(i)
//	////fmt.Println("reflect.TypeOf(i) = ", reflect.TypeOf(i))
//	//////fmt.Println("reflect.TypeOf(i).Len() = ", reflect.TypeOf(i).Len())
//	////fmt.Println("reflect.TypeOf(i).Size() = ", reflect.TypeOf(i).Size())
//	////fmt.Println(reflect.Int)
//	////fmt.Println("reflect.ValueOf(i) = ", reflect.ValueOf(i))
//	////fmt.Println("reflect.ValueOf(i).Type() = ", reflect.ValueOf(i).Type())
//	////
//	//fmt.Println("reflect.TypeOf(animal) = ", reflect.TypeOf(animal))
//	//fmt.Println("reflect.TypeOf(animal).Kind() = ", reflect.TypeOf(animal).Kind())
//	////fmt.Println("reflect.TypeOf(animal).Key() = ", reflect.TypeOf(animal).Key())
//	//fmt.Println("reflect.TypeOf(animal).Align() = ", reflect.TypeOf(animal).Align())
//	//
//	//fmt.Println("reflect.ValueOf(animal).Kind() = ", reflect.ValueOf(animal).Kind())
//	//
//	//fmt.Println(reflect.TypeOf(animal).Field(1))
//	//fmt.Println(reflect.TypeOf(animal).Field(1).Index)
//	//fmt.Println(reflect.TypeOf(animal).Field(1).Offset)
//	//fmt.Println(reflect.TypeOf(animal).Field(1).Name)
//
//	i := 1
//	k := 5
//	v := reflect.ValueOf(&i)
//	g := reflect.ValueOf(&k)
//	v.Elem().SetInt(777)
//	g.Elem().SetInt(10)
//	fmt.Println("sdsf", i)
//	fmt.Println("iInterface = ", k)
//}

//func main() {
//	i := 1
//	v := reflect.ValueOf(&i)
//	v.Elem().SetInt(10)
//	fmt.Println(i)
//}
