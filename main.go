package main

import (
	"fmt"

	"github.com/leonardoAlonso/godesde0/functions"
)

func main() {
	state, text := functions.Convert_to_text(1234)
	fmt.Println(state)
	fmt.Println(text)
}
