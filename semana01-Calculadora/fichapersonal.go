package main

import "fmt"

func main() {

	var nombre string
	var edad int
	var estudio string
	var semestre int
	var promedio float64

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)

	fmt.Println("Ingrese su carrera:")
	fmt.Scanln(&estudio)

	fmt.Println("Ingrese su semestre:")
	fmt.Scanln(&semestre)

	fmt.Println("Ingrese su promedio:")
	fmt.Scanln(&promedio)

	fmt.Printf("Soy %s tengo %d años\nEstudio %s, semestre %d\nMi promedio es %.2f\n",
		nombre, edad, estudio, semestre, promedio)
}