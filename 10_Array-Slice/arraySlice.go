package main

import (
	"fmt"
	"strings"
)

func main() {

	//Array
	//los arrays son inmutables no pueden cambiar su tamaño, los arreglos en Go sólo pueden contener
	//elementos de un mismo tipo.
	//Si se alcanza y excede la capacidad del arreglo, ocurre un error

	// Formas de declarar un Array
	var arrayName [4]int // --> Forma normal de crear un arreglo vacío
	arrayName[0] = 1
	arrayName[1] = 2
	arrayName[2] = 3
	arrayName[3] = 4
	// arrayName := [4]int{}  --> Forma corta de crear un arreglo vacío
	// arrayName := [4]int{"1", "2", "3", "4"}  --> Forma corta de crear un arreglo con datos
	// arrayName := [...]int{"1", "2", "3", "4"}  --> Forma corta de crear un arreglo sin necesidad de establecer su longitud

	//len --> cuantos elements hay en array, cap--> capacidad maxima de datos de array
	fmt.Println(arrayName, len(arrayName), cap(arrayName))

	//Recorriendo un arreglo
	for i := 0; i < len(arrayName); i++ {
		fmt.Println(arrayName[i])
	}

	//range devuelve un indice "i" y un valor "v" por iteración
	for i, v := range arrayName {
		fmt.Println(i, v)
	}

	marcasAutos := [...]string{"Toyota", "Kia", "Hyundai", "Jeep"}

	for i := 0; i < len(marcasAutos); i++ {
		fmt.Println(marcasAutos[i])
	}

	for _, v := range marcasAutos {
		fmt.Println(v)
	}

	//Slice
	//Un Slice se podria decir que es un segmento de un arreglo.
	//En comparacion con los arreglos los Slices son mas flexibles, ya que podemos agregar,
	//remover y copiar elementos de un slice a otro.

	//Para crear un slice se debe definir su nombre, su tipo y su capacidad inicial

	//declaración de un Slice - similar a declarar un arreglo pero sin especificar su longitud
	var nums []int //Slice sin tamaño definido
	fmt.Println(nums)

	//Declaración de un Slice con la función make
	//Si la capacidad no se define el valor por defecto es igual a la longitud
	var nums1 []int           //Declaración de slice
	nums1 = make([]int, 3, 3) //Inicialización con 3 elementos y capacidad de 3

	nums2 := make([]int, 3, 3) // Declaración en 1 sola línea

	fmt.Println(nums1)
	fmt.Println(nums2)

	// --- Operaciones y Acciones sobre Slices ---

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice se utiliza al manejar arrays, slices o listas para interactuar
	// con sus elementos
	fmt.Println(slice[0])   // Imprime el elemento en la posición --> 0
	fmt.Println(slice[:3])  // Imprime los elementos hasta uno antes de la posición 3
	fmt.Println(slice[2:4]) // Imprime los elementos desde la posición 2 hasta una antes del 4
	fmt.Println(slice[4:])  // Imprime los elementos desde la posición 4

	// slice[a:b] //Acceder a los alementos desde a-b
	// slice[:b] //Acceder a los alementos desde 0-b
	// slice[a:] //Acceder a los alementos desde a-len(a)
	// slice[:] //Acceder a todos los elementos

	// Donde a es el límite inferior (inclusivo) y b el límite superior (exclusivo)

	//Puedes inicializar un slice con la información total o parcial de otro gracias a los límites antes mencionados:
	slice2 := slice[1:3] //Crea e inicializa slice2 con los índices 1 y 2 de slice1
	slice3 := slice[:]   //Crea e inicializa slice3 con el contenido total de slice1

	fmt.Println(slice2)
	fmt.Println(slice3)
	//¡NOTA IMPORTANTE!: Si se inicializa un slice como en los ejemplos anteriores, es importante tener en cuenta que
	//lo que se crea realmente es un apuntador al espacio de memoria del slice que se usa como información base;
	//por tanto, si se hace un cambio sobre alguno de ellos, el otro también se verá afectado.

	// -- Copiar un Slice --
	//Se puede realizar una copia de un slice con la función copy()
	//copy(slice_destino, slice_original)
	var slice4 []int
	copy(slice4, slice)

	// -- Agregar datos a un Slice --
	//la función append() añade uno o varios elementos nuevos desde el último índice del slice.
	//si el último índice ya está utilizado, se aumenta dinámicamente la capacidad del slice al
	//doble y se añade dicho elemento en la última posición
	slice = append(slice, 7, 8, 9) //Se añaden 3 elementos nuevos
	slice = append(slice, 10)      //Se pueden añadir N o 1 elemento nuevo

	// -- Eliminar un valor de un Slice --
	razasDePerros := []string{"labrador", "poodle", "doberman", "shitzu", "beagle"}
	fmt.Println(razasDePerros) // Salida --> [labrador poodle doberman shitzu beagle]
	razasDePerros = append(razasDePerros[:2], razasDePerros[2+1:]...)
	fmt.Println(razasDePerros) // Salida --> [labrador poodle shitzu beagle]

	// -- Ejercicio Palindromos --
	palindromo("Amor a Roma")

}

func palindromo(cadena string) {

	var cadenaInversa string
	cadena = strings.ToLower(cadena)

	for i := len(cadena) - 1; i >= 0; i-- {
		cadenaInversa += string(cadena[i])
	}

	if cadena == cadenaInversa {
		fmt.Println("La cadena de texto es palindromo")
	} else {
		fmt.Println("La cadena de texto NO es palindromo")
	}
}
