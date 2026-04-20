package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, n := range clientes {
		if n.ID == id {
			return i
		}
	}
	return -1
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes)")
		return
	}
	for _, c := range clientes {
		fmt.Printf("ID: %d | Nombre: %s | Carrera: %s | Saldo: $%.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func EliminarClientes(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, n := range productos {
		if n.ID == id {
			return i
		}
	}
	return -1
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS REGISTRADOS ===")
	if len(productos) == 0 {
		fmt.Println("(no hay productos)")
		return
	}
	for _, c := range productos {
		fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Stock: %d | Categoria: %s\n",
			c.ID, c.Nombre, c.Precio, c.Stock, c.Categoria)
	}
}

func EliminarProductos(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", id)
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func DescontarSaldo(cliente *Cliente, monto float64) error {
	if cliente.Saldo < monto {
		return fmt.Errorf("saldo insuficiente")
	}
	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if producto.Stock < cantidad {
		return fmt.Errorf("stock insuficiente")
	}
	producto.Stock -= cantidad
	return nil
}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {

	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		productos[idxP].Stock += cantidad
		return pedidos, err
	}

	nuevo := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevo)
	return pedidos, nil
}

// ================== HELPERS ==================

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}

func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return n
}

func leerFloat(lector *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	f, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return -1
	}
	return f
}

// ================== MENÚ ==================

func mostrarMenu() {
	fmt.Println("\n===== SISTEMA =====")
	fmt.Println("1. Listar clientes")
	fmt.Println("2. Listar productos")
	fmt.Println("3. Agregar cliente")
	fmt.Println("4. Agregar producto")
	fmt.Println("5. Registrar pedido")
	fmt.Println("6. Ver pedidos de cliente")
	fmt.Println("0. Salir")
}

// ================== REPORTE ==================

func PedidosDeCliente(pedidos []Pedido, clientes []Cliente, productos []Producto, clienteID int) {

	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		fmt.Println("Cliente no existe")
		return
	}

	fmt.Println("\nPedidos de:", clientes[idxC].Nombre)

	total := 0.0
	encontrados := 0

	for _, p := range pedidos {
		if p.ClienteID == clienteID {

			idxP := BuscarProductoPorID(productos, p.ProductoID)
			if idxP == -1 {
				continue
			}

			fmt.Printf("Pedido %d | Producto: %s | Cantidad: %d | Total: %.2f | Fecha: %s\n",
				p.ID,
				productos[idxP].Nombre,
				p.Cantidad,
				p.Total,
				p.Fecha)

			total += p.Total
			encontrados++
		}
	}

	if encontrados == 0 {
		fmt.Println("No tiene pedidos")
		return
	}

	fmt.Println("Total gastado:", total)
}

// ================== MAIN ==================

func main() {

	lector := bufio.NewReader(os.Stdin)

	clientes := []Cliente{
		{1, "Juan Perez", "TI", 50},
		{2, "Maria Lopez", "Civil", 75},
		{3, "Carlos Ruiz", "Industrial", 100},
	}

	productos := []Producto{
		{1, "Coca Cola", 1.50, 10, "Bebida"},
		{2, "Sandwich", 2.50, 5, "Almuerzo"},
		{3, "Galletas", 1.00, 20, "Snack"},
		{4, "Jugo", 1.25, 8, "Bebida"},
	}

	var pedidos []Pedido

	for {
		mostrarMenu()
		op := leerEntero(lector, "Opcion: ")

		switch op {

		case 1:
			ListarClientes(clientes)

		case 2:
			ListarProductos(productos)

		case 3:
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo: ")

			clientes = AgregarCliente(clientes, Cliente{id, nombre, carrera, saldo})

		case 4:
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio: ")
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoria: ")
			categoria := leerLinea(lector)

			productos = AgregarProducto(productos, Producto{id, nombre, precio, stock, categoria})

		case 5:
			cid := leerEntero(lector, "Cliente ID: ")
			pid := leerEntero(lector, "Producto ID: ")
			cant := leerEntero(lector, "Cantidad: ")
			fmt.Print("Fecha: ")
			fecha := leerLinea(lector)

			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, cid, pid, cant, fecha)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Pedido registrado")
			}

		case 6:
			id := leerEntero(lector, "Cliente ID: ")
			PedidosDeCliente(pedidos, clientes, productos, id)

		case 0:
			fmt.Println("Hasta luego")
			return

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
