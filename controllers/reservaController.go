package controllers

import (
	"Restaurante/services"
	"Restaurante/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AgregarReserva(w http.ResponseWriter, r *http.Request) {
	var reserva services.Reserva
	if err := json.NewDecoder(r.Body).Decode(&reserva); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	id, err := services.AgregarReserva(db, reserva)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Reserva agregada correctamente",
	})
}

func AnalizarReserva(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fechaReserva := params["fechaReserva"]

	db := utils.GetDB()
	analisis, err := services.AnalizarReserva(db, fechaReserva)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(analisis)
}
