// Video10
package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	Nombre = "Pedro Perez"
	Estado = true
	Sueldo = 1200

	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

// Video11
func ConviertoATexto(numero int) (bool, string) {
	var texto string
	// texto = string(numero)
	texto = strconv.Itoa(numero)
	return true, texto
}
