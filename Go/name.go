package main

import (
	"fmt"
)

func main() {
	names := []string{"Morticia", "Gomez", "Wednesday", "Pugsley"}

	for _, n := range names {
		fmt.Println(n)
	}
}
