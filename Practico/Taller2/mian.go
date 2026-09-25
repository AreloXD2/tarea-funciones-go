package main

import (
	"Taller2/contador"
	"Taller2/convertor"
	"fmt"
)

func main() {
	var entrada int

	for {
		fmt.Println("*** MENÚ PRINCIPAL ****")
		fmt.Println("1. Dólares a Euros")
		fmt.Println("2. Dólares a Libras Esterlinas")
		fmt.Println("3. Dólares a Won (Surcoreano)")
		fmt.Println("4. Dólares a Bitcoins (BTC)")
		fmt.Println("5. Contador de Vocales")
		fmt.Println("0. Salir")

		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&entrada)

		if entrada == 0 {
			fmt.Println("Fin del programa.")
			break
		}

		switch entrada {
		case 1:
			convertor.Euros()
		case 2:
			convertor.Libras()
		case 3:
			convertor.Won()
		case 4:
			convertor.BTC()
		case 5:
			contador.ContarVocales()
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
