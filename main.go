// Siempre va esta primera linea
package main

import (
	//Video8
	//Video11
	"fmt"
	//Video9
	//Video10

	//Video12

	"github.com/franklinacosta/godesde0/ejercicios"
)

// Variables con icono azul, Funciones con el icono morado
// Funciones
func main() {
	//Video8
	// fmt.Println("Hola mundo")

	//Video9
	// variables.MuestroEnteros()

	//Video10
	// variables.RestoVariables()

	//Video11
	// estado, texto := variables.ConviertoATexto(1950)
	// fmt.Println(estado)
	// fmt.Println(texto)

	//Video12
	// os := runtime.GOOS
	// if os := runtime.GOOS; os == "Linux." || os == "OS X." || os == "darwin" {
	// 	fmt.Println("Esto no es windows, es", os)
	// } else {
	// 	fmt.Println("Esto es windows")
	// 	fmt.Println(os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es linux")
	// case "windows":
	// 	fmt.Println("Esto es windows")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	//Video13
	numero, mensaje := ejercicios.FuncionEjercicio1("101")
	fmt.Println(numero, mensaje)
}
