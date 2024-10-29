package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaNumerica() {
	var result int
	var error error
	var numeroIngresado int
	var scaner = bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número: ")

	if scaner.Scan() {
		numeroIngresado, error = strconv.Atoi(scaner.Text())

		if error != nil {
			fmt.Println("El valor ingresado no corresponde a un número")
			TablaNumerica()
		} else {
			for i := 1; i < 11; i++ {
				result = numeroIngresado * i
				fmt.Printf("%d X %d = %d \n", numeroIngresado, i, result)
			}
		}
	}
}
