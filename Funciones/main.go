package main

import "fmt"

/*
func <nombre> (<parametros>) <tipo de retorno> {
	<cuerpo de la función>
	return <valor de retorno>
}
*/

func saludar() {
	fmt.Println("Hola, es mi primera función en Go")
}

func Bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func suma(a, b int) int {
	return a + b
}

func suma_resta(num1, num2 int) (int, int) {
	if num2 > num1 {
		ResSuma := num1 + num2
		ResResta := num2 - num1
		return ResResta, ResSuma
	} else {
		ResSuma := num1 + num2
		ResResta := 0
		return ResResta, ResSuma

	}
}

func mostrarNum(numeros ...int) {
	fmt.Println("Los números ingresados son: ", numeros)
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func main() {

	var usr string

	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&usr)

	saludar()
	Bienvenida(usr)

	r1, r2 := suma_resta(5, 8)
	fmt.Println("La resta es:", r1, "La suma es: ", r2)

	mostrarNum(5, 10, 45, 4, 6)
	fmt.Println("La sumatoria es:", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9))

}
