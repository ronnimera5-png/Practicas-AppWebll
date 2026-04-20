package main

import "fmt"

func main() {

	var nombre string
	var nota1, nota2 float64
	var promedio float64

	// Entrada de datos
	fmt.Println("¿Cómo te llamas?")
	fmt.Scanln(&nombre)

	fmt.Println("Ingresa tu nota 1:")
	fmt.Scanln(&nota1)

	fmt.Println("Ingresa tu nota 2:")
	fmt.Scanln(&nota2)

	// Cálculo
	promedio = (nota1 + nota2) / 2

	// Salida
	fmt.Printf("%s, tu promedio es: %.2f\n", nombre, promedio)

	if promedio >= 7 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}
}
