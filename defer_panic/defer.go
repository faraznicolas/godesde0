package defer_panic

import (
	"fmt"
	"log"
)

func VemosError() {
	fmt.Println("Este es un primer Mensaje")
	defer fmt.Println("Este es un tercer Mensaje")
	fmt.Println("Este es un segundo Mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
