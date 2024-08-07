package main

import "fmt"

func main() {
	var number1, number2 int

	//Solicita o primeiro número
	fmt.Print("Insira o primeiro número: ")
	fmt.Scanln(&number1)

	//Solicita o segundo número
	fmt.Print("Insira o segundo número: ")
	fmt.Scanln(&number2)

	//Calcula a soma dos dois números
	soma := number1 + number2

	//Imprime o resultado
	fmt.Printf("A soma de %d e %d é %d\n", number1, number2, soma)
}
