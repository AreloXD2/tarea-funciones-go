package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {

	fmt.Println("😁Bienvenido a la CLase de Paquetes 😘")
	mensaje := saludo.Saludar("Arelo")
	fmt.Println(mensaje)

	Operacion1 := operaciones.Suma(5, 6)
	fmt.Println("Llamando la opercion 1 Sumar: ", Operacion1)

	Op1, Op2 := operaciones.Suma_resta(6, 4)
	fmt.Println("Llamando la operación 2 Suma_resta: ", Op1, Op2)

	Operacion3 := operaciones.Sumatoria(6, 5, 4, 7, 8, 9)
	fmt.Println("Llamando a la operación 3 Sumatoria: ", Operacion3)

}
