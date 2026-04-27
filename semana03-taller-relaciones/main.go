package main

import (
	"errors"
	"fmt"
	"semana3_taller_relaciones/internal/cafeteria"
)

func main() {

	//Crear repo usando la interfaz
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Carlos", Carrera: "Sistemas", Saldo: 50})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Xavier", Carrera: "Industrial", Saldo: 30})

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Batido", Precio: 2.5, Stock: 10, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Sandwich", Precio: 5.0, Stock: 5, Categoria: "Comida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Jugo", Precio: 3.0, Stock: 8, Categoria: "Bebida"})

	// Buscar cliente que existe
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	// Buscar cliente que NO existe
	c, err = repo.ObtenerCliente(99)
	if err != nil {
		// Bonus: verificar tipo de error
		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("Cliente no existe:", err)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	// 5. Listar productos
	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	productos := repo.ListarProductos()
	for _, p := range productos {
		fmt.Printf("[%d] %s - $%.2f (Stock: %d)\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}

	// 6. Mostrar Pedido con Cliente y Producto completos
	cliente, _ := repo.ObtenerCliente(1)
	producto, _ := repo.ObtenerProducto(1)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: producto,
		Cantidad: 2,
		Total:    float64(2) * producto.Precio,
		Fecha:    "2026-04-27",
	}

	fmt.Println("\n=== PEDIDO ===")
	fmt.Printf("Pedido ID: %d\n", pedido.ID)
	fmt.Printf("Cliente: %s (%s)\n", pedido.Cliente.Nombre, pedido.Cliente.Carrera)
	fmt.Printf("Producto: %s\n", pedido.Producto.Nombre)
	fmt.Printf("Cantidad: %d\n", pedido.Cantidad)
	fmt.Printf("Total: $%.2f\n", pedido.Total)
}
