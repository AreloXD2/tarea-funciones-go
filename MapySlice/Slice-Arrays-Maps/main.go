package main

import "fmt"

func main() {

	Nota_Materia := [6][4]float64{
		{7.0, 8.5, 3.0, 5.5},
		{4.5, 9.6, 6.0, 3.2},
		{7.5, 8.7, 9.5, 10.0},
		{9.5, 7.5, 6.4, 4.1},
		{6.3, 6.5, 5.5, 3.4},
		{8.0, 7.0, 6.5, 9.0},
	}

	var promedios []float64

	for cont := 0; cont < len(Nota_Materia); cont++ {
		fila := Nota_Materia[cont][:]

		suma := 0.0
		mayor := fila[0]
		menor := fila[0]

		for _, nota := range fila {
			suma += nota
			if nota > mayor {
				mayor = nota
			}
			if nota < menor {
				menor = nota
			}
		}

		promedio := suma / float64(len(fila))
		promedios = append(promedios, promedio)

		fmt.Println("Estudiante", cont+1)
		fmt.Println("Promedio:", promedio)
		fmt.Println("Nota más alta:", mayor)
		fmt.Println("Nota más baja:", menor)
		fmt.Println("-----------------------------")
	}

	sumaTotal := 0.0
	for _, cont := range promedios {
		sumaTotal += cont
	}
	promedioGeneral := sumaTotal / float64(len(promedios))

	fmt.Println("Promedio general de la clase:", promedioGeneral)
}
