package arreglos_slices

import "fmt"

var tablaS []int = []int{1, 1, 23, 1}
var arreglo [10]int = [10]int{4, 45, 88, 97, 456, 13, 548, 312, 18, 4}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]
	porcion2 := arreglo[:5]
	porcion3 := arreglo[6:7]
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
