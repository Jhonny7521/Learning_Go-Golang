package main

import (
	"fmt"
	"math"
)

func main() {

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras
	base := 12          // ":=" Declaramos que esa variable no ha sido creada antes y la creamos y le asignamos un valor
	var altura int = 14 // declaramos el tipo de dato y le asignamos un valor
	var area int        // declaramos el tipo de dato pero no le asignamos un valor

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d) // imprimimos para ver valores por defecto --> 0 0 " " False

	// Area de cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado es: ", areaCuadrado)

	//Operaciones Aritméticas
	x := 10
	y := 50

	//Suma
	result := y + x
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = y * x
	fmt.Println("Multiplicacion:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Reto: Escribir un procedimiento para hallar el área de un rectangulo, triángulo y circulo

	//Área de rectángulo
	baseR := 30
	alturaR := 20

	areaR := baseR * alturaR
	fmt.Printf("El área del rectángulo es %d \n", areaR)

	//Área de Triángulo
	baseT := 30
	alturaT := 20

	areaT := (baseT * alturaT) / 2
	fmt.Printf("El área del triángulo es %d \n", areaT)

	//Área de Círculo
	radioC := 30.0
	areaC := math.Pi * (math.Pow(radioC, 2))
	fmt.Printf("El área del círculo es %f \n", areaC)
}
