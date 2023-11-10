package files

import (
	"fmt"
	"os"

	"github.com/faraznicolas/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var text string = ejercicios.Ejercicio2()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, text)
	archivo.Close()
}

func SumaTabla() {
	var text string = ejercicios.Ejercicio2()
	if !Append(filename, text) {
		fmt.Println("Error al concatenar el archivo")
		return
	}
}

func Append(filen string, text string) bool {
	arch, err := os.OpenFile(filen, os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Err al abrir el archivo")
		return false
	}
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Err al escribir el archivo")
		return false
	}
	arch.Close()
	return true

}

func LeoArchivo() {
	archivo, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return
	}
	fmt.Println(string(archivo))
}
