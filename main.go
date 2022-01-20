package main

import "fmt"

//slide
func sld() {
	//Declaración de slides
	//Declaracion simple
	//var dias []int
	//Declaracion e inicializacion
	//dias := []int{1, 2, 3}
	//Declaracion con make
	//dias := make([]int, 5, 10)

	//dias := []string{"domingo", "lunes", "martes", "miercoles", "jueves", "viernes", "sabado"}
	x := []byte{'H', 'O', 'L', 'A'}
	//var diassl []string
	//diassl = dias[5:7]
	//fmt.Println(diassl)
	//fmt.Println(len(diassl))
	//fmt.Println(cap(diassl))
	//dias = append(dias, "dia ocho")
	x[0] = 'B'
	fmt.Printf("%q \n", x)
}

//array
func arr() {
	//Declaracion de arrelgo de 7 posiciones tipo string
	//var dias [7]string
	//Declaración e inicializacion de array
	dias := [7]string{"domingo", "lunes", "martes", "miercoles", "jueves", "viernes", "sabado"}
	//var numeros [10]int
	//Arreglo variable debe crearse y asignarse
	//ids := [...]int{1,2,3,4,5}
	fmt.Println(dias)
	//Asignacion individual
	//dias[0] = "Lunes"
	fmt.Println(dias[1])
}

//Swith
func sw() {
	var dia int
	fmt.Println("Digita un día de la semana en numero:")
	fmt.Scanln(&dia)
	fmt.Println("Tecleaste:", dia)
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	default:
		fmt.Println("¡Dato invalido!")
	}
}

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
	//sw()
	//arr()
	sld()
}
