package main

import "fmt"

func main() {

	//Utilizado cuando queremos iterar sobre una variable en epecífica
	// modulo := 5 % 2
	// switch modulo {
	switch modulo := 4 % 2; modulo { // switch "definicion de variable"; "parseo de la variable"
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Switch sin condicion su uso es similar al de tener varios "if"
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
