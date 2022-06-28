package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Marco",
		idade: 56,
	}

	fmt.Println(x)
}
