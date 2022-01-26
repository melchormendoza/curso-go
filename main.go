package main

import (
	"fmt"
	"os"
)

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

//Variables
func vari() {
	//Declaracion de variables
	var numero int
	var nombre, apellidos string
	//asignacion de valores a variables
	numero = 10
	nombre = "Melchor"
	apellidos = "Mendoza Ramirez"
	//Variable de asignacion dinamica
	ciudad := "Oaxaca"
	fmt.Println("Hello world!", numero, nombre, apellidos, ciudad)
}

//Operadores
func op() {
	//Operadores
	a := 8
	b := 3
	fmt.Println("Operacion a/b:", a/b)
	var resultado bool
	resultado = 5 < 7
	fmt.Println("Resultado de evaluar 5<7 y tirarlo a una varibale de tipo bool:", resultado)
}
func ciclo() {
	//Estructuras de control
	//Contador declarado fuera del for
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//Variable inicializada dentro del for
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}
	//If con variable inicializada y comparada en la misma sentencia
	if a := 5; a <= 4 {
		fmt.Println("a es menor o igual")
	}
}
func sld2() {
	//Declaracón de slide
	x := make([]byte, 4, 100)
	fmt.Println(x)
	//inicializacion
	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)
	fmt.Printf("Slide X: %q \n", x)
	for i := 0; i < len(x); i++ {
		fmt.Printf("Slide: [%d]%q \n", i, x[i])
	}
	//x[5] = 'Y'
	//fmt.Printf("%q \n", x)
	x = append(x, 'M', 'U', 'N', 'D', 'O')
	fmt.Printf("%q \n", x)
}
func copysld() {
	origen := []int{1, 2, 3}
	destino := []int{3, 4, 5}
	copy(destino, origen)
	fmt.Println(origen, destino)
	//-------------------
	origen2 := []int{1, 2, 3}
	destino2 := make([]int, 2)
	copy(destino2, origen2)
	fmt.Println(origen2, destino2)
	//--------------------
	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}
	copy(destino3, origen3)
	fmt.Println(origen3, destino3)
}
func rango() {
	nombres := []string{
		"Melchor",
		"Alejandro",
		"Alisson",
	}
	for i, n := range nombres {
		fmt.Printf("el nombre %q esta en el indice %d \n", n, i)
	}
	for _, n := range nombres {
		fmt.Println(n)
	}
}
func mapa() {
	//mapa 1
	x := make(map[string]string)
	x["nombre"] = "Melchor"
	x["edad"] = "39"
	fmt.Println(x)
	fmt.Println(x["edad"])
	//Mapa 2 con longitud
	y := make(map[string]string, 2)
	fmt.Println(y)
	//mapa 3 declaracion diferente
	z := map[string]int{
		"Melchor":   39,
		"Alisson":   15,
		"Alejandro": 8,
	}
	fmt.Println(z)
	//eliminacion de elementos de un Mapa
	delete(z, "Alejandro")
	fmt.Println(z)
	a, b := z["Alejandro"]
	fmt.Println(a, b)
	//Recorrer un mapa con range
	for nombre, edad := range z {
		fmt.Printf("La edad de %s es %d \n", nombre, edad)
	}
}

//Implementacion de una función
func devuelveNombre(nombre string) {
	fmt.Println("Tu nombres es:", nombre)
}
func sumatoria(n1 int, n2 int) int {
	return n1 + n2
}
func resto(numero1 int, numero2 int) (r int) {
	r = numero1 - numero2
	return
}

//Variadic function|
func suma(numeros ...int) (r int) {
	for _, numero := range numeros {
		r += numero
	}
	return r
}
func multi(numero int) (r1, r2, r3 int) {

	r1 = numero * 4
	r2 = numero * 5
	r3 = numero * 7

	return
}
func texto() (c int, data []byte, err error) {
	f, err := os.Open("texto.txt")
	if err != nil {
		panic(err)
	}
	data = make([]byte, 23)
	c, err = f.Read(data)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%q \n %d \n", data, c)
	return c, data, err
}
func main() {
	cantidad, data, error := texto()
	fmt.Println("Caracteres: \n", cantidad)
	fmt.Printf("%q \n", data)
	fmt.Println(error)
	/*
		a, b, c := multi(5)
		fmt.Println(a, b, c)
	*/
	//fmt.Println(multi(2))
	/*
		cad := "Hola"
		imp := func() {
			fmt.Println(cad)
		}
		imp()
	*/
	//sum := suma
	//fmt.Println(sum(2, 3, 4))
	//fmt.Println(suma(6, 1, 2, 4))
	//fmt.Println(suma(67, 14, 12, 49))
	//fmt.Println(suma())
	//fmt.Println(resto(19, 7))
	//fmt.Println(sumatoria(3, 7))
	//devuelveNombre("Melchor MR")
	//mapa()
	//rango()
	//copysld()
	//sld2()
	//ciclo()
	//|op()
	//vari()
	//sw()
	//arr()
	//sld()
}
