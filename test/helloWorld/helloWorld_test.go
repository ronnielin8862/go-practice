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
