package controllers

import (
	"Restaurante/services"
	"Restaurante/utils"
	"encoding/json"
	"net/http"
)

func AgregarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente services.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	id, err := services.AgregarCliente(db, cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Cliente agregado correctamente",
	})
}

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	clientes, err := services.ListarClientes(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clientes)
}
