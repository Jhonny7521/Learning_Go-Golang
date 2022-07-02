package main

import "fmt"

func main() {

	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	for j := 20; j >= 0; j-- {
		fmt.Println(j)
	}

	fmt.Printf("\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	sharks := []string{"tiburón martillo", "gran tiburón blanco", "cazón", "con volantes", "cabeza de toro", "réquiem"}

	for i, shark := range sharks {
		fmt.Println(i, shark) //Se imprime el indice y valor de cada elemento en sharks
	}

	//For forever el bucle se ejecuta por siempre hasta que lo detengamos con ctrl + c
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
