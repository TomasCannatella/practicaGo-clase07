/*
Ejercicio 2 - Impuestos de salario #2

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: salary is less than 10000"
y lanzalo en caso de que “salary” sea menor o igual a  10000.

La validación debe ser hecha con la función “Is()” dentro del “main”.

Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código
para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"errors"
	"fmt"
)

var salaryError error = errors.New("Error: salary is less than 10000")

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
