package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := &ErMy{"E1"}
	fmt.Println(e1) // E1

	e2 := fmt.Errorf("E2:%w", e1) // note: %w only
	fmt.Println(e2)               // E2:E1

	var v2 *ErMy
	fmt.Println(v2) // <nil>

	fmt.Println(errors.As(e2, &v2)) // true
	fmt.Println(v2)                 // E1
}

type ErMy struct {
	Msg string
}

func (p *ErMy) Error() string {
	return p.Msg
}
