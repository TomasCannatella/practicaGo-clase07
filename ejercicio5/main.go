/*
Ejercicio 5 -  Impuestos de salario #5

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora
como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000,
se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80
o un número negativo, la función debe devolver un error.
El mismo tendrá que indicar
“Error: the worker cannot have worked less than 80 hours per month”.
*/

package main

import (
	"errors"
	"fmt"
)

var ErrorHourWorked = errors.New("Error: the worker cannot have worked less than 80 hours per month")

func calculateSalary(hourWorked, pricePerHour float64) (salary float64, err error) {
	if hourWorked < 80 {
		err = fmt.Errorf("%w", ErrorHourWorked)
		return
	}

	salary = pricePerHour * hourWorked
	if salary >= 150000 {
		salary += -((salary * 10) / 100)
	}
	return
}

func main() {
	hourWorked, pricePerHour := 1.0, 300.0

	salary, err := calculateSalary(hourWorked, pricePerHour)

	if err != nil {
		if errors.Is(err, ErrorHourWorked) {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Unknow error")
		}
	}
	fmt.Println("This salary is: ", salary)
}
