package main

import (
	"fmt"

	"github.com/leonardoAlonso/godesde0/functions"
	"github.com/leonardoAlonso/godesde0/variables"
)

func main() {
	variables.Get_variables()
	state, text := functions.Convert_to_text(1234)
	fmt.Println(state)
	fmt.Println(text)
}
