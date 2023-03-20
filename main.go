package main

import "github.com/leonardoAlonso/godesde0/iteraciones"

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
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Esto no es windows")
	// } else {
	// 	fmt.Println("Esto es windows")
	// 	fmt.Println(os)
	// }

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Esto es linux")
	// case "darwin":
	// 	fmt.Println("Esto es darwin")
	// default:
	// 	fmt.Printf("Esto es: %s\n", os)
	// }

	// numero, mensaje := ejercicio01.EvaluateString("109")
	// fmt.Println(numero)
	// fmt.Println(mensaje)

	/*
		for {
			Esto es un ciclo infinito
		}

		break -> es una forma para truncar un cliclo

	*/

	iteraciones.Iterar()

}
