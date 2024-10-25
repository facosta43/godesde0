// Video13
package ejercicios

import (
	"strconv"
)

func FuncionEjercicio1(parametro string) (int, string) {
	var convertidoAEntero int
	var error error

	convertidoAEntero, error = strconv.Atoi(parametro)

	if error != nil {
		return 0, "Error controlado: " + error.Error()

	}

	if convertidoAEntero > 100 {
		return convertidoAEntero, parametro + " Este número es mayor a 100"
	} else {
		return convertidoAEntero, parametro + " Este número es menor a 100"
	}
}
