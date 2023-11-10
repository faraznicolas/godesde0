package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var text string

func Ejercicio2() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero 1")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			break
		}
	}
	for i := 0; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}

	return text
}
