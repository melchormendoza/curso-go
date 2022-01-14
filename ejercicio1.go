package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	fmt.Println("Contador de numeros impares")
	fmt.Println("Digita el primer numero:")
	fmt.Scanln(&numero1)
	fmt.Println("Digite el segundo numero:")
	fmt.Scanln(&numero2)
	for {
		if numero1 >= numero2 {
			fmt.Println("Digite nuevamente el segundo numero:")
			fmt.Scanln(&numero2)
		} else {
			break
		}
	}
	fmt.Println("De los numeros comprendidos entre: ", numero1, "y", numero2)
	for i := numero1; i <= numero2; i++ {
		if i%2 == 0 {
			fmt.Printf("%d es un numero par\n", i)
		} else {
			fmt.Printf("%d es un numero impar\n", i)
		}

	}
}
