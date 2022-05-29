package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Age     int
	Name    string
	Surname string
}

func (p Person) Hello() string {
	return "Hello " + p.Name
}

func main() {
	fmt.Printf("%d %d", unsafe.Sizeof(new(int)), unsafe.Sizeof(new(float32)))
}
