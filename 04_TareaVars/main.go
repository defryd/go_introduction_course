package main

import (
	"fmt"
)

func main() {

	license := true
	age := 19

	if license == true && age >= 19 {
		fmt.Println("Puede puede seguir avanzando")
	} else if license == false && age > 18 {
		fmt.Println("No puede seguir circulando, no tiene licencia")
	} else if license == true && age == 15 {
		fmt.Println("No puede seguir circulando, es menor de edad")
	}

	fmt.Println("Fin del programa")
}
