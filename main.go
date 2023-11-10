package main

import (
	"fmt"
	"runtime"

	"github.com/faraznicolas/godesde0/ejercicios"
	"github.com/faraznicolas/godesde0/files"
	"github.com/faraznicolas/godesde0/funciones"
	"github.com/faraznicolas/godesde0/iteraciones"
	"github.com/faraznicolas/godesde0/teclado"
	"github.com/faraznicolas/godesde0/variables"
)

func main() {
	funciones.Exponencia(2)
}

func mainOld() {
	funciones.LlamarClosure()
	files.LeoArchivo()
	numero, text := ejercicios.Conversor("as")
	fmt.Println(numero)
	fmt.Println(text)
	variables.MuestroEnteros()
	variables.RestoVariables()
	var estado, texto = variables.ConviertoaTexto(10520)
	fmt.Println("estado", estado)
	fmt.Println("texto", texto)
	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Sistema operativa", os)
	} else {
		fmt.Println("Sistema operativa", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s", os)
	}

	iteraciones.Iterar()
	teclado.IngresoNumeros()
}
