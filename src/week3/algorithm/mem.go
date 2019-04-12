package main

import (
	"fmt"
	"unsafe"
)
type A1 struct {
	b bool
	a float64
	c int32
}
type A2 struct {
	a float64
	c int32
	b bool
}

func main() {
	a1 := A1{}
	a2 := A2{}
	fmt.Println(unsafe.Sizeof(a1)) //24
	fmt.Println(unsafe.Sizeof(a2)) //16
}
