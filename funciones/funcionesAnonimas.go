package funciones

import "fmt"

func Calculos() {
	calculo := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println(calculo(1, 15))
	calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Println(calculo(1, 15))
}
