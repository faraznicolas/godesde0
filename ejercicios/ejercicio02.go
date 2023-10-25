package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Ejercicio2() {
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
		fmt.Println(i * numero)
	}
}
