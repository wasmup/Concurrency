package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `["a", "b"]`
	var v []string
	fmt.Println(v == nil)

	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}
