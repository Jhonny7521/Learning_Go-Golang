package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 2
	valor2 := 1

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// Else if
	var edad uint8
	fmt.Scanln("Cual es tu edad?")
	fmt.Scanf("%d", &edad)

	if edad < 13 {
		fmt.Println("Puedes ver solo películas clasificación A")
	} else if edad < 15 {
		fmt.Println("Puedes ver solo películas clasificación A y B")
	} else if edad < 18 {
		fmt.Println("Puedes ver solo películas clasificación A, B y B15")
	} else {
		fmt.Println("Puedes ver cualquier clasificación")
	}

	//Con AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	//Con OR
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("asdf") //value --> valor convertiro, err --> si ocurre un error
	//si existe un error "err" no sera nil, si no existe un error "err" sera nil.
	if err != nil { // Nil es la forma en la que go determian si una funcion tiene un error
		log.Fatal(err) //Si existe un error imprime en consola lo que sucedio y termina el codigo
	}
	fmt.Println("Value:", value)

	//Operadores de comparación --> Todas regresa true o false segun la comparación
	//	Operador		Descripción
	//	x == y			Es igual x igual a y ?
	//	x != y			Es x diferente de y ?
	//	x < y				Es x menor que y?
	//	x <= y			Es x menor o igual que y?
	//	x > y				Es x mayor que y?
	//	x >= y			Es x mayor o igual que y?
}
