package main

import "fmt"

// Estructura de una función

// func functionName(a int) int { // funcion que retorna un valor
//   return a * 2
// }
// func: Indica que se esta definiendo una función
// functionName: nombre de la función
// (a int): los paréntesis enmarcan los parametros que recibirá la función.
//					Se tiene que definir el tipo de dato de los parámetros.
// int: Se especifica el tipo de dato que retornará la función (en caso sea necesario un retorno de valor).
// { ... }: Las llaves delimitan el alcanse del scope de la función.
// return: Palabra clave para retornar un valor. Este valor debe coincidir con el tipo de dato que retorna la función.

func normalFunction(message string) {
	fmt.Println("Hola mundo")
}

func tripleArgumento(a, b int, c string) { // "a" y "b" son del mismo tipo de dato int
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // funcion que retorna un valor
	return a * 2
}

func singleReturn(a int) (b int) { //función que retorna una variable definida --> (b int)
	b = a * 2
	return // como b ya fue definida no es necesario declararla en el return
}

func doubleReturn(a int) (c, d int) { // funcion que retorna dos valores
	return a, a * 2 // imprimimoe en el mismo orden primero "c" -->a y luego "d" --> a*2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgumento(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	//llamamos a la funcion con doble retorno y guardamos datos en 2 variables
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	//Si solo utilizaremos 1  dato el otro lo declaramos con "_"
	dato1, _ := doubleReturn(5)
	fmt.Println("dato1", dato1)
}
