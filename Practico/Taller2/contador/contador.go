package contador

import "fmt"

func ContarVocales() {
	var texto string

	fmt.Println("*** CONTADOR DE VOCALES ***")
	fmt.Println("Ingrese una palabra o texto (sin espacios): ")
	fmt.Scan(&texto)

	var cantA, cantE, cantI, cantO, cantU int

	for i := 0; i < len(texto); i++ {
		letra := texto[i]

		switch letra {
		case 'a', 'A':
			cantA++
		case 'e', 'E':
			cantE++
		case 'i', 'I':
			cantI++
		case 'o', 'O':
			cantO++
		case 'u', 'U':
			cantU++
		}
	}

	fmt.Println("Cantidad de 'a':", cantA)
	fmt.Println("Cantidad de 'e':", cantE)
	fmt.Println("Cantidad de 'i':", cantI)
	fmt.Println("Cantidad de 'o':", cantO)
	fmt.Println("Cantidad de 'u':", cantU)
}
