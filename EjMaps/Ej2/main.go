package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func mostrarProducto(prod Producto) {
	fmt.Println("Nombre:", prod.Nombre)
	fmt.Println("Precio:", prod.Precio)
	fmt.Println("Cantidad:", prod.Cantidad)
}

func valorTotal(prod Producto) float64 {
	return prod.Precio * float64(prod.Cantidad)
}

func hayStock(prod Producto) bool {
	return prod.Cantidad > 0
}

func main() {

	productos := []Producto{
		{Nombre: "Laptop", Precio: 800.50, Cantidad: 5},
		{Nombre: "Mouse", Precio: 15.50, Cantidad: 20},
		{Nombre: "Teclado", Precio: 45.00, Cantidad: 0},
	}

	fmt.Println("===== INVENTARIO =====")

	for cont := 0; cont < len(productos); cont++ {
		prod := productos[cont]

		mostrarProducto(prod)
		fmt.Println("Valor total:", valorTotal(prod))

		if hayStock(prod) {
			fmt.Println("Estado: disponible")
		} else {
			fmt.Println("Estado: sin stock")
		}

		fmt.Println("---------------------")
	}
}
