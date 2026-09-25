package convertor

import "fmt"

func Euros() {
	var dolar float64
	fmt.Println("Ingrese su valor en Dolares: ")
	fmt.Scanln(&dolar)
	resultado := dolar * 0.8785
	fmt.Println("La conversión de ", dolar, " $ a Euros son: ", resultado, " euros")
}

func Libras() {
	var dolar float64
	fmt.Println("Ingrese su valor en Dolares: ")
	fmt.Scanln(&dolar)
	resultado := dolar * 0.7558
	fmt.Println("La conversión de ", dolar, " $ a Libras son: ", resultado, " LB")
}

func Won() {
	var dolar float64
	fmt.Println("Ingrese su valor en Dolares: ")
	fmt.Scanln(&dolar)
	resultado := dolar * 1366.8
	fmt.Println("La conversión de ", dolar, " $ a Wones son: ", resultado, " wones")
}

func BTC() {
	var dolar float64
	fmt.Println("Ingrese su valor en Dolares: ")
	fmt.Scanln(&dolar)
	resultado := dolar / 84200.0
	fmt.Println("La conversión de ", dolar, " $ a Bitcoins son: ", resultado, " BTC")
}
