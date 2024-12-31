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
	Nombre = "Juan"
	Estado = true
	Sueldo = 1000.50
	Fecha = time.Now()
	fmt.Println("Nombre = ", Nombre, "Estado = ", Estado, "Sueldo = ", Sueldo, "Fecha = ", Fecha)
}

func ConviertoaTexto(numero int) (string, bool) {

	texto := strconv.Itoa(numero)
	return texto, true

}
