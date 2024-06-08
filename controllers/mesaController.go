package controllers

import (
	"Restaurante/services"
	"Restaurante/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AgregarMesa(w http.ResponseWriter, r *http.Request) {
	var mesa services.Mesa
	if err := json.NewDecoder(r.Body).Decode(&mesa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	id, err := services.AgregarMesa(db, mesa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Mesa agregada correctamente",
	})
}

func ActualizarMesa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mesaID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de la mesa debe ser un número entero", http.StatusBadRequest)
		return
	}

	var disponibilidad struct {
		Disponible int `json:"disponible"`
	}
	if err := json.NewDecoder(r.Body).Decode(&disponibilidad); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	err = services.ActualizarMesa(db, disponibilidad.Disponible, mesaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Mesa actualizada correctamente",
	})
}

func ListarMesas(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	mesas, err := services.ListarMesas(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(mesas)
}
