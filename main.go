package main

import (
	"fmt"

	"github.com/faraznicolas/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	var estado, texto = variables.ConviertoaTexto(10520)
	fmt.Println("estado", estado)
	fmt.Println("texto", texto)
}
