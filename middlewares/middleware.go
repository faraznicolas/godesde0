package middlewares

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")
	resut := operacionesMidd(sumar)(2, 3)
	fmt.Println(resut)
	resut = operacionesMidd(restar)(2, 3)
	fmt.Println(resut)
	resut = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(resut)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio operacion")
		return f(a, b)
	}
}
