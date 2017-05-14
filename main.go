package main

import (
	"fmt"

	"github.com/riccardomc/golab/arraysstrings"
)

func main() {
	fmt.Println(arraysstrings.Unique("abcd"))
	fmt.Println(arraysstrings.Unique("acdc"))
}
