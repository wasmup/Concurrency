package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))         // 32
	fmt.Println(reflect.TypeOf(a).Size()) // 32

	b := [][2]int{{1, 2}, {3, 4}}
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(b))         // 24
	fmt.Println(reflect.TypeOf(b).Size()) // 24
}
