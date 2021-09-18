package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("unit test helloWorld")
}

func TestHelloWorldError(t *testing.T) {
	t.Error("Error helloWorld")
}
