/*
package services

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
*/
package services

import (
	"database/sql"
	"log"
)

type Cliente struct {
	Nombre   string `json:"nombre"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
}

func AgregarCliente(db *sql.DB, cliente Cliente) (int64, error) {
	query := `INSERT INTO Clientes (nombre, telefono, email) VALUES (?,?,?)`
	result, err := db.Exec(query, cliente.Nombre, cliente.Telefono, cliente.Email)
	if err != nil {
		log.Printf("Error al insertar cliente: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID del cliente: %v", err)
		return 0, err
	}

	log.Printf("Cliente agregado correctamente con ID: %d", id)
	return id, nil
}

func ListarClientes(db *sql.DB) ([]Cliente, error) {
	query := `SELECT nombre, telefono, email FROM Clientes`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar clientes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var clientes []Cliente
	for rows.Next() {
		var cliente Cliente
		if err := rows.Scan(&cliente.Nombre, &cliente.Telefono, &cliente.Email); err != nil {
			log.Printf("Error al leer filas: %v", err)
			continue
		}
		clientes = append(clientes, cliente)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return clientes, nil
}
