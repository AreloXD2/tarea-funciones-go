package operaciones

func Suma(a, b int) int {
	return a + b
}

func Suma_resta(num1, num2 int) (int, int) {
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

func Sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}
