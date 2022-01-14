package main

import "fmt"

func main() {
	//Declaracion de variables
	var numero int
	var nombre, apellidos string
	//asignacion de valores a variables
	numero = 10
	nombre = "Melchor"
	apellidos = "Mendoza Ramirez"
	//Variable de asignacion dinamica
	ciudad := "Oaxaca"
	//Operadores
	a := 8
	b := 3
	fmt.Println("Operacion", a/b)
	var resultado bool
	resultado = 5 < 7
	fmt.Println("Resultado", resultado)
	//Estructuras de control
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
	if a := 2; a <= 4 {
		fmt.Println("a es menor")
	}
	fmt.Println("Hello world!", numero, nombre, apellidos, ciudad)
}
