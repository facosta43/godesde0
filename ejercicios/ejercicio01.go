// Video13
package ejercicios

import (
	"fmt"
	"strconv"
)

func FuncionEjercicio1(parametro string) (int, string) {
	var convertidoAEntero int
	var error error
	var retornoMensaje string
	convertidoAEntero, error = strconv.Atoi(parametro)
	fmt.Println(error)

	if convertidoAEntero > 100 {
		retornoMensaje = parametro + " Este número es mayor a 100"
	} else {
		retornoMensaje = parametro + " Este número es menor a 100"
	}
	return convertidoAEntero, retornoMensaje
}
