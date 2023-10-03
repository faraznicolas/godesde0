package ejercicios

import (
	"strconv"
)

func Conversor(valor string) (int, string) {
	var texto string
	numero, _ := strconv.Atoi(valor)
	if numero > 100 {
		texto = "Es mayor a 100"
	} else if numero < 100 {
		texto = "Es menor a 100"
	} else {
		texto = "es igual a 100"
	}
	return numero, texto
}
