package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Cliente struct {
	Nombre   string
	Telefono string
	Email    string
}

func AgregarCliente(db *sql.DB, cliente Cliente) error {
	// Consulta para agregar un cliente
	query := `INSERT INTO Clientes (nombre, telefono, email) VALUES (?,?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, cliente.Nombre, cliente.Telefono, cliente.Email)
	if err != nil {
		log.Printf("Error al insertar cliente: %v", err)
		return err
	}
	fmt.Println("Cliente agregado correctamente.")
	return nil
}
