/*
package main

import (

	"Restaurante/services"
	"Restaurante/utils"
	"fmt"
	"log"
	"os"

)

	func main() {
		// Establece la conexión con la base de datos
		db, err := utils.ConectarBaseDeDatos()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Opciones del sistema
		for {
			fmt.Print("Bienvenido!")
			fmt.Print("*******************************************\n")
			fmt.Print("Elige una opcion\n")

			fmt.Print("*******************************************\n")
			fmt.Println("1. Agregar una mesa")
			fmt.Println("2. Agregar un cliente")
			fmt.Println("3. Agregar una reserva")
			fmt.Println("4. Analizar patrones de reserva")
			fmt.Println("5. Salir del sistema")
			fmt.Println("******************************************")
			fmt.Print("Ingrese su opcion: ")
			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				fmt.Println("Agregar una mesa")
				fmt.Printf("Ingrese el numero de mesa: ")
				var numero int
				fmt.Scanln(&numero)
				fmt.Printf("Ingrese la capacidad de la mesa: ")
				var capacidad int
				fmt.Scanln(&capacidad)
				nuevaMesa := services.Mesa{
					NumeroDeMesa: numero,
					Capacidad:    capacidad,
					Disponible:   1,
				}
				services.AgregarMesa(db, nuevaMesa)
			case 2:
				fmt.Println("Agregar un cliente")
				fmt.Printf("Ingrese el nombre del cliente: ")
				var nombre string
				fmt.Scanln(&nombre)
				fmt.Printf("Ingrese el telefono del cliente: ")
				var telefono string
				fmt.Scanln(&telefono)
				fmt.Printf("Ingrese el email del cliente: ")
				var email string
				fmt.Scanln(&email)
				nuevoCliente := services.Cliente{
					Nombre:   nombre,
					Telefono: telefono,
					Email:    email,
				}
				services.AgregarCliente(db, nuevoCliente)
			case 3:
				fmt.Println("Agregar reserva")
				fmt.Printf("Ingrese el ID de la mesa: ")
				var mesaId int
				fmt.Scanln(&mesaId)
				fmt.Printf("Ingrese el ID del cliente: ")
				var clienteId int
				fmt.Scanln(&clienteId)
				fmt.Printf("Ingrese la fecha de reserva: ")
				var fecha string
				fmt.Scanln(&fecha)
				fmt.Printf("Ingrese la hora de reserva: ")
				var hora string
				fmt.Scanln(&hora)
				fmt.Printf("Ingrese la cantidad de comensales: ")
				var cantidad int
				fmt.Scanln(&cantidad)
				nuevaReserva := services.Reserva{
					ClienteID:          clienteId,
					MesaID:             mesaId,
					FechaReserva:       fecha,
					HoraReserva:        hora,
					NumeroDeComensales: cantidad,
				}
				services.AgregarReserva(db, nuevaReserva)
				services.ActualizarMesa(db, 1, mesaId)
			case 4:
				fmt.Println("Analizar patrones de reserva")
				fmt.Printf("Ingrese la fecha a analizar: ")
				var fecha string
				fmt.Scanln(&fecha)
				services.AnalizarReserva(db, fecha)
			case 5:
				os.Exit(0)
			default:
				fmt.Println("Opcion invalida.")
			}
		}
	}
*/
package main

import (
	"Restaurante/routers"
	"Restaurante/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
