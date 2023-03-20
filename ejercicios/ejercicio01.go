package ejercicios

import "strconv"

func EvaluateString(str string) (int32, string) {

	numero, err := strconv.Atoi(str)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	if numero >= 100 {
		return int32(numero), "Mayor a 100"
	}
	return int32(numero), "Menor a 100"

}
