// Siempre va esta primera linea
package main

import (
	//Video8
	//Video11
	"fmt"
	//Video9
	//Video10
	"github.com/franklinacosta/godesde0/variables"
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
	estado, texto := variables.ConviertoATexto(1950)
	fmt.Println(estado)
	fmt.Println(texto)
}
