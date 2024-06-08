package routers

import (
	"Restaurante/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Definir rutas
	router.HandleFunc("/clientes", controllers.AgregarCliente).Methods("POST")
	router.HandleFunc("/mesas", controllers.AgregarMesa).Methods("POST")
	router.HandleFunc("/mesas", controllers.ActualizarMesa).Methods("PUT")
	router.HandleFunc("/reservas", controllers.AgregarReserva).Methods("POST")
	router.HandleFunc("/reservas/analizar", controllers.AnalizarReserva).Methods("POST")

	return router
}
