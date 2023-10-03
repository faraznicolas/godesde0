package ejercicios

import (
	"strconv"
)

func Conversor(valor string) (int, string) {
	var texto string
	numero, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error"
	}
	if numero > 100 {
		texto = "Es mayor a 100"
	} else if numero < 100 {
		texto = "Es menor a 100"
	} else {
		texto = "es igual a 100"
	}
	return numero, texto
}
