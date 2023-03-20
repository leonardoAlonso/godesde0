package main

import (
	"fmt"
	"runtime"
)

func main() {
	// variables.Get_variables()
	// userinput.UserInput()
	// state, text := functions.Convert_to_text(1234)
	// fmt.Println(state)
	// fmt.Println(text)
	// convertion.Convertion()
	// mytime.UseTime()
	// mypointers.UsePointer()

	/*
		La estructura if nos permite asignar y comparar en una misma linea
		similar al operador warlus de python

		==
		>
		<
		>=
		<=
		&& and
		|| or
	*/
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("Esto es: %s\n", os)
	}

}
