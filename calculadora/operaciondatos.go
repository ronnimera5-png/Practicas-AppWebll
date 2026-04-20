package main

import "fmt"

func main() {

	// Constantes
	const p1 float64 = 49.99
	const p2 float64 = 29.99
	const p3 float64 = 15.50

	var total, promedio, totalConDescuento float64

	// Cálculos
	total = p1 + p2 + p3
	promedio = total / 3
	totalConDescuento = total * 0.85

	// Salida
	fmt.Printf("Producto 1: $%.2f\n", p1)
	fmt.Printf("Producto 2: $%.2f\n", p2)
	fmt.Printf("Producto 3: $%.2f\n", p3)
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Printf("Promedio: $%.2f\n", promedio)
	fmt.Printf("Total con 15%% descuento: $%.2f\n", totalConDescuento)
	/*
	   import "fmt" func main() { var p1, p2, p3 float64 var total, promedio, descuento float64 // Ingreso de datos fmt.Println("Ingrese el precio del primer producto:") fmt.Scanln(&p1) fmt.Println("Ingrese el precio del segundo producto:") fmt.Scanln(&p2) fmt.Println("Ingrese el precio del tercer producto:") fmt.Scanln(&p3) // Cálculos total = p1 + p2 + p3 promedio = total / 3 descuento = total * 0.15 totalConDescuento := total - descuento // Salida de resultados fmt.Println("\n----- RESULTADOS -----") fmt.Printf("Total: %.2f\n", total) fmt.Printf("Promedio: %.2f\n", promedio) fmt.Printf("Descuento (15%%): %.2f\n", descuento) fmt.Printf("Total con descuento: %.2f\n", totalConDescuento) }
	*/
}
