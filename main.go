package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	variables.Muestroenteros()
	variables.RestoVariables()
	texto, Estado := variables.ConviertoaTexto(42)
	fmt.Println("Texto = ", texto, "Estado = ", Estado)

}
