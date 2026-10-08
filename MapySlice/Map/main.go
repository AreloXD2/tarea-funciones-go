package main

import "fmt"

func masVotada(votos map[string]int) string {
	ganadora := ""
	maximo := -1

	for actividad, cantidad := range votos {
		if cantidad > maximo {
			maximo = cantidad
			ganadora = actividad
		}
	}

	return ganadora
}

func main() {

	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("1. Deportes")
	fmt.Println("2. Videojuegos")
	fmt.Println("3. Cine")
	fmt.Println("4. Música")

	for cont := 1; cont <= 5; cont++ {
		var opcion int
		fmt.Println("Ingrese el voto", cont, "(1-4)")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			votos["deportes"]++
		case 2:
			votos["videojuegos"]++
		case 3:
			votos["cine"]++
		case 4:
			votos["musica"]++
		default:
			fmt.Println("Opción no válida, se pierde el voto")
		}
	}

	fmt.Println("*** RESULTADOS ***")
	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	fmt.Println("Actividad ganadora:", masVotada(votos))
}
