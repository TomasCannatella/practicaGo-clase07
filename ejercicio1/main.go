/*
Ejercicio 1 - Impuestos de salario #1

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: the salary entered does not reach the taxable minimum"
y lanzalo en caso de que “salary” sea menor a 150.000.
De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.
*/
package main

import (
	"fmt"
)

type SalaryError struct {
	Message string
}

func (sE *SalaryError) Error() string {
	return sE.Message
}

func main() {
	var salary int = 149000
	var err error = &SalaryError{"Error: the salary entered does not reach the taxable minimum"}

	if salary < 150000 {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Must pay tax")
}
