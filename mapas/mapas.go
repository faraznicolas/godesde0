package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenas Aires"
	fmt.Println(paises)

	campeonato := map[string]int{

		"Barcelona":    39,
		"Real Madrid":  38,
		"Boca Juniors": 40,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, puntaje %d \n", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("Equipo existe %t, puntaje %d \n", existe, puntaje)

}
