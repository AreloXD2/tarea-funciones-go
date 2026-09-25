package contador

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ContarVocales() {
	lector := bufio.NewReader(os.Stdin)

	fmt.Println("*** CONTADOR DE VOCALES ***")
	fmt.Println("Ingrese una frase: ")
	texto, _ := lector.ReadString('\n')
	texto = strings.TrimSpace(texto)

	var cantA, cantE, cantI, cantO, cantU int

	for _, letra := range texto {
		switch letra {
		case 'a', 'á', 'A', 'Á':
			cantA++
		case 'e', 'é', 'E', 'É':
			cantE++
		case 'i', 'í', 'I', 'Í':
			cantI++
		case 'o', 'ó', 'O', 'Ó':
			cantO++
		case 'u', 'ú', 'ü', 'U', 'Ú', 'Ü':
			cantU++
		}
	}

	fmt.Println("Cantidad de 'a':", cantA)
	fmt.Println("Cantidad de 'e':", cantE)
	fmt.Println("Cantidad de 'i':", cantI)
	fmt.Println("Cantidad de 'o':", cantO)
	fmt.Println("Cantidad de 'u':", cantU)
}
