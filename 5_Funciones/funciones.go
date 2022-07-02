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

// FUNCIONES VARIÁDICAS
// Son funciones que reciben varios argumentos
// al declarar "..." le indico a la funcion que estare recibiendo varios parámetros
// Esto solo se puede hacer como ultimo parámetro. No se acepta --> (numeros ...int, valorInicial int)
func sumar(valorInicial int, numeros ...int) (total int) {
	// el total inicial es 0
	// recorrer todos los numeros
	for _, numero := range numeros {
		// en cada iteración sumar al total el valor del numero
		total = numero + total
	}
	// retornar el valor total
	return
}

// FUNCIONES RECURSIVAS
// Son funciones que se invocan y ejecutan a sí mismas.
// Es necesario que estas funciones cuenten con una condicion que permita que la recursión termine.
func factorial(n int) (res int) {
	if n == 1 { // Condicion que permite terminar la recursividad
		res = 1
		return
	} else {
		res = n * factorial(n-1)
		return
	}
}

// Funcion que recive una función anónima como argumento
func frase(i func(p, q string) string) {
	fmt.Println(i("Ojo", "por"))
}

// Retorno de una funcióne anónimae
func frase2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myf
}

func datosCirculo(radio float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return 3.1416 * radio * radio
	}

	perimetro = func() float64 {
		return 2 * 3.1416 * radio
	}

	return
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

	// FUNCIONES ANONIMAS COMO ARGUMENTOS O LAMBDAS

	//Función Anónima: Son funciones que ejecutan alguna accion pero no tienen un nombre definido.
	//Estas funciones funciones debes estar declaradas dentro del scope main o alguna otra función NO anónima.
	//Esta funcion se ejecutará en cuanto la ejecución del código llegue a ella.

	func() { fmt.Println("Hola mundo!") }()

	//Función anónima con argumentos
	//Se envia parámetros dentro de los ultimos paréntesis --> ("GeeksforGeeks")
	func(ele string) { fmt.Println(ele) }("GeeksforGeeks")

	// ---------------
	nombre := "Maria"
	// Se puede asignar una función anónima a una variable
	saludo := func() {
		fmt.Println("Hola mundo!, mi nombre es", nombre)
	}
	// invocamos a la funcion de esta manera
	saludo()

	valor := func(a, b string) string {
		return a + b + "ojo"
	}
	// Enviamos la función anómima como argumento a otra función
	frase(valor)

	valor2 := frase2()
	fmt.Println(valor2("Completa la", "frase: "))

	area, perimetro := datosCirculo(10)
	fmt.Println("El area del circulo es", area())
	fmt.Println("El perimetro del circulo es", perimetro())
}
