package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("unit test unary")
}

func TestHelloWorldError(t *testing.T) {
	t.Error("Error unary")
}

func TestNamingReturnValues(t *testing.T) {
	t.Log("1 = ", noNaming())
	no := noNaming()
	t.Log("2 = ", no)
	t.Log("4 = ", naming())
	t.Log("3 = ", no)

	t.Log("5 = ", includNoNaming())
}

func includNoNaming() (t int) {
	t = noNaming()
	println("ttt = ", t)
	return t
}

func noNaming() int {

	x := 5
	defer func() {
		println("1111111")
		x++
		println("22222222")
	}()

	return x
}

func naming() (x int) {
	x = 5

	defer func() {
		println("333333")
		x++
		println("4444444")
	}()
	return
}
