/*
Ejercicio 2 - Impuestos de salario #2

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: salary is less than 10000"
y lanzalo en caso de que “salary” sea menor o igual a  10000.

La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	message string
}

func (s *SalaryError) Error() string {
	return s.message
}

var salaryError error = &SalaryError{
	message: "Error: salary is less than 10000",
}

func checkSalary(salary int) (err error) {
	if salary <= 10000 {
		err = salaryError
	}
	return
}

func main() {
	var salary int = 10000
	if err := checkSalary(salary); err != nil {
		if errors.Is(err, salaryError) {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Must pay tax")
}
