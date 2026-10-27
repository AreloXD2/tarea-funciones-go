package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (prod *Producto) ActualizarPrecio(nuevoPrecio float64) {
	prod.Precio = nuevoPrecio
}

func (prod *Producto) ActualizarCantidad(nuevaCantidad int) {
	prod.Cantidad = nuevaCantidad
}

func (prod Producto) MostrarInfo() {
	fmt.Println("ID:", prod.ID)
	fmt.Println("Nombre:", prod.Nombre)
	fmt.Println("Precio: $", prod.Precio)
	fmt.Println("Cantidad:", prod.Cantidad)
	fmt.Println("-----------------")
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.Productos[p.ID] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.Productos, id)
}

func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.Productos[id]
	return p, existe
}

func (inv Inventario) ListarProductos() {
	for _, p := range inv.Productos {
		p.MostrarInfo()
	}
}

func main() {
	inv := NuevoInventario()

	mouse := NuevoProducto(1, "Mouse", 15.5, 10)
	teclado := NuevoProducto(2, "Teclado", 30, 5)
	monitor := NuevoProducto(3, "Monitor", 100, 4)

	fmt.Println(" Monitor ANTES de actualizar")
	monitor.MostrarInfo()
	monitor.ActualizarPrecio(90)
	monitor.ActualizarCantidad(6)

	fmt.Println(" Monitor DESPUÉS de actualizar ")
	monitor.MostrarInfo()

	inv.AgregarProducto(mouse)
	inv.AgregarProducto(teclado)
	inv.AgregarProducto(monitor)

	fmt.Println("Lista Inicial de Productos")
	inv.ListarProductos()

	inv.EliminarProducto(1)
	fmt.Println("Inventario tras eliminar el ID 1 ")
	inv.ListarProductos()

	fmt.Println("Búsqueda de Producto")
	producto, encontrado := inv.BuscarProducto(2)

	if encontrado == true {
		producto.MostrarInfo()
	} else {
		fmt.Println("Error: No hay ese producto en el sistema")
	}

	fmt.Println("Búsqueda de Producto que NO existe")
	productoFalso, encontradoFalso := inv.BuscarProducto(67)

	if encontradoFalso == true {
		productoFalso.MostrarInfo()
	} else {
		fmt.Println("Error: No hay ese producto en el sistema")
	}
}
