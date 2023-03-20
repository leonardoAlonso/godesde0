package functions

import "strconv"

/*
	Una funcion puede recibir tantos parametros como necesite
	cada parametro debe especificar el tipo
	si una funcion va a retornar un valor entonces lo definiremos entre parentesis
*/

func Convert_to_text(numero int) (bool, string) {
	text := strconv.Itoa(numero)
	return true, text
}
