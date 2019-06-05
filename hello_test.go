package gotest

import (
    "testing"
    "fmt"
)

func TestAdd(t *testing.T) {
    fmt.Println("TestAdd: 1+2=", add(1,2))
}