// Video15
package iteraciones

import (
	"fmt"
)

func Iterar() {
	//i := 0
	for i := 0; i < 10; {
		if i == 0 {
			println("Primer for")
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		if i == 0 {
			println("Segundo for")
		}
		fmt.Println(i)
	}

	for i := 0; i < 100; i += 5 {
		if i == 0 {
			println("Tercero for")
		}
		fmt.Println(i)
	}

	for i := 100; i > 10; i -= 5 {
		if i == 100 {
			println("Cuarto for")
		}
		fmt.Println(i)
	}

	for i := 10; i > 1; i-- {
		if i == 10 {
			println("Quinto for")
		}
		if i == 6 {
			continue // pasa de este punto nuevamente a la instrucción for, sin hacer nada de lo que esta más abajo
		}
		if i == 6 {
			break // se detiene
		}
		fmt.Println(i)
	}
}
