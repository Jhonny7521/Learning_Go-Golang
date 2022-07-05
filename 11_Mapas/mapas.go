package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int) //make(map["tipoLLave"]"tipoDato")

	//Agregar valores al Map
	m["Jose"] = 14
	m["Luis"] = 20
	m["Maria"] = 18

	fmt.Println(m) //En la salida maps no muestra comas sino espacios en blanco para separar cada "llave: valor"

	//Eliminar valores al Map --> delete(mapaName, "llaveName")
	delete(m, "Maria")
	// Debenos tener cuidado al momento de escribir el nombre de la llave, ya que
	// Si la escribimos mal, el compilador no nos arrojará error.

	//Para saber el tamaño de un Map
	lenMap := len(m)
	fmt.Println(lenMap)

	//Recorrer un Map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value := m["Jose"]
	fmt.Println(value) // Salida --> 14

	//Encontrar valor no declarado en un map
	value1 := m["Joselyn"]
	fmt.Println(value1) // Salida --> 0
	//Esto es porque al no encontrar encontar "llave" en el map Go la crea y le asigna un valor por defecto de 0

	//Verificar si un valor existe en un map
	value2, ok := m["Jose"] //La variable ok almacena true si el valor existe, de lo contrario es false
	fmt.Println(value2, ok) // Salida --> 14 true

	value3, ok := m["Joselyn"]
	fmt.Println(value3, ok) // Salida --> 0 false

	//Mapa que almacena arrays u objetos
	estudiantes := make(map[string][]int)

	estudiantes["Mario"] = []int{13, 14, 15}
	estudiantes["Alberto"] = []int{16, 17, 18}

	fmt.Println(estudiantes)

	//Mostrar un elemento del obj almacenado
	fmt.Println(estudiantes["Mario"][1])
}
