/*
Ejercicio 4 - Impuestos de salario #4

Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary”
indicando que no alcanza el mínimo imponible (el mensaje mostrado por
consola deberá decir: “Error: the minimum taxable amount is 150,000
and the salary entered is: [salary]”, siendo [salary] el valor de tipo
int pasado por parámetro).
*/
package main

import (
	"errors"
	"fmt"
)

var salaryError error = errors.New("Error: salary is less than 10000")

func checkSalary(salary int) (err error) {
	if salary <= 150000 {
		err = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is:%d . %w", salary, salaryError)
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
