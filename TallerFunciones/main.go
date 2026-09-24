package main

import "fmt"

// Taller Funciones

func main() {
	var entrada string
	for {
		fmt.Println(".... MENÚ: ....")
		fmt.Println("Elija las siguientes opciones porfavor:")
		fmt.Println("1. Primera opción: Promedio Curso")
		fmt.Println("2. Segunda opción: Suma de números")
		fmt.Println("3. Tercera opción: Celsius a Fahrenheit ")
		fmt.Println("4. Cuarta opción: Fahrenheit a Celsius")
		fmt.Println("Para acabar el programa escriba -salir- o 0")

		fmt.Println("Ingrese su opción elejida:")
		fmt.Scan(&entrada)

		if entrada == "0" || entrada == "salir" {
			break
		} else if entrada == "1" {
			Opcion1()
		} else if entrada == "2" {
			Opcion2()
		} else if entrada == "3" {
			Opcion3()
		} else if entrada == "4" {
			Opcion4()
		} else {
			fmt.Println("Opción NO disponible.")
		}

	}

}

func Opcion1() {
	var estudiantes int
	var nota float64
	var sumaNotas float64
	fmt.Println(".... PRIMERA OPCIÓN .... ")
	fmt.Println("Ingrese la cantidad de estudiantes del curso: ")
	fmt.Scan(&estudiantes)
	for contador := 1; contador <= estudiantes; contador++ {
		fmt.Println("Ingrese la nota del estudiante", contador, "(0 a 100): ")
		fmt.Scan(&nota)
		sumaNotas = sumaNotas + nota
	}

	promedio := averageGrade(sumaNotas, estudiantes)
	fmt.Println("El promedio es: ", promedio)

	if promedio >= 70 {
		fmt.Println("El promedio del curso es APROBADO")
	} else if promedio < 70 {
		fmt.Println("EL promedio del curso es REPROBADO")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Excellent Performance")

	case promedio >= 80 && promedio < 90:
		fmt.Println("Good performance")

	case promedio >= 70 && promedio < 80:
		fmt.Println("Satisfactory performance")

	case promedio < 70:
		fmt.Println("Needs improvement")
	}

}

func averageGrade(suma float64, cantidad int) float64 {
	promedio := suma / float64(cantidad)
	return promedio
}

func Opcion2() {
	fmt.Println(".... SEGUNDA OPCIÓN .... ")
	fmt.Println("Ingrese un número para sumar del 1 hasta ese número ingresado: ")

	var num int
	fmt.Scan(&num)

	var acumulador int
	for contador := 1; contador <= num; contador++ {
		acumulador = acumulador + contador
	}

	fmt.Println("La suma de 1 hasta el numero ingresado:", num, "es: ", acumulador)
}

func Opcion3() {
	fmt.Println(".... TERCERA OPCIÓN .... ")
	fmt.Println("Ingrese la temperatura en Celsius: ")

	var celsius float64
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Println("La temperatura en Fahrenheit es: ", fahrenheit)
}

func Opcion4() {
	fmt.Println(".... CUARTA OPCIÓN .... ")
	fmt.Println("Ingrese la temperatura en Fahrenheit: ")

	var fahrenheit float64
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Println("La temperatura en Celsius es: ", celsius)
}
