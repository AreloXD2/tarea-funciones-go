package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PalabraRepetida(conteo map[string]int) string {
	ganadora := ""
	maximo := 0

	for palabra, veces := range conteo {
		if veces > maximo {
			maximo = veces
			ganadora = palabra
		}
	}

	return ganadora
}

func main() {

	fmt.Println("Ingrese una frase (sin signos, en minúscula):")
	lector := bufio.NewReader(os.Stdin)
	texto, _ := lector.ReadString('\n')

	palabras := strings.Fields(texto)

	conteo := map[string]int{}

	for cont := 0; cont < len(palabras); cont++ {
		conteo[palabras[cont]]++
	}

	fmt.Println("----- REPETICIONES -----")
	for palabra, veces := range conteo {
		fmt.Println(palabra, ":", veces)
	}

	fmt.Println("Palabra que más se repite:", PalabraRepetida(conteo))
}
