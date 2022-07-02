package main

import "fmt"

func main() {
	//defer : Con este keyword se logra ejecutar una funcion
	// antes de que todo el programa finalice o se cierre.
	// Se podría decir que es la última función.

	// Ejm: Si abrimos una conexion a una BD con defer podermos
	// declarar que esta se cierre antes de que el programa
	// termina. Esto es una buena practica con GO.
	// Se puede utilizar mas de una sentencia defer dentro de una funcion
	// sin embargo las buenas prácticas dicen que se debe utilizar solo 1
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break --> Se utilizan dentro del bucle For
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
