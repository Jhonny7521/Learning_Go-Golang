package main

import "fmt"

func main() {

	//Cada variable tiene un valor y una direccion en memoria
	fruta1 := "Manzana"   // valor = Ford , direccion en Memoria= 0xc00004c250
	fruta2 := "Mandarina" // valor = Toyota , direccion en Memoria= 0xc00004c260

	//Con & apuntamos a la direccion de memoria
	fmt.Println("fruta1:", fruta1, &fruta1) // Salida--> fruta1: Manzana 0xc00004c250
	fmt.Println("fruta2:", fruta2, &fruta2) // Salida--> fruta2: Mandarina 0xc00004c260

	// La variable 3 almacena el valor de la variable 1 mas no su direccion de memoria
	fruta3 := fruta1
	fmt.Println("fruta3:", fruta3, &fruta3) // Salida--> fruta3: Manzana 0xc00004c290

	//Sin embargo si actualizamos el valor de fruta1 el valor de fruta3 no se actualiza
	fruta1 = "Platano"
	fmt.Println("fruta1:", fruta1, &fruta1) // Salida--> fruta1: Platano 0xc00004c250
	fmt.Println("fruta3:", fruta3, &fruta3) // Salida--> fruta3: Manzana 0xc00004c290

	//Para ello se utilizan los punteros
	//Primero guardamos el valor de la direccion de memoria de fruta1 en fruta4
	fruta4 := &fruta1
	//Con esto podemos acceder al nuevo valor de fruta4
	fmt.Println("fruta4:", fruta4, &fruta4, *fruta4) // Salida--> fruta4: 0xc00004c250 0xc000006030 Platano

	//Enviado Variables a funciones

	var peso float32

	fmt.Println("Cuál es tu peso en Kg?")
	//Pedimos un ingreso de datos por consola, definimos el dato como flotante y la guardamos en variable altura
	fmt.Scanf("%f", &peso)

	fmt.Println(&peso) //esta variable tiene una direccion en memoria

	//peso = kgALibras(peso)    // enviamos el valor de la variable a la funcion
	kgALibras2(&peso) // enviamos la direccion en memoria de la variable
	fmt.Println("Tu peso en libras es:", peso, "libras")
}

//En esta funcion se genera otra direccion de memoria para pesoData que almacena el valor recibido
func kgALibras(pesoData float32) float32 {
	pesoData = pesoData * 2.2046
	return pesoData
}

//En esta funcion no se genera otra direccion de memoria para pesoData
//sino solo se trabaja directamente con el valor de la direccion recibida
// y el nuevo valor se vuelve a almacenar en la misma direccion. Por ello no es necesario el "return"
func kgALibras2(pesoData2 *float32) {
	*pesoData2 = *pesoData2 * 2.2046
}
