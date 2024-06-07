package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Mesa struct {
	NumeroDeMesa int
	Capacidad    int
	Disponible   int
}

// Funcion para agregar un registro a la tabla Mesas
func AgregarMesa(db *sql.DB, mesa Mesa) error {
	// Consulta para agregar registro
	query := `INSERT INTO Mesas (NumeroDeMesa, Capacidad, Disponible) VALUES (?, ?, 1)`
	// Ejecuta la consulta
	_, err := db.Exec(query, mesa.NumeroDeMesa, mesa.Capacidad)

	if err != nil {
		log.Printf("Error al agregar Mesa: %v", err)
		return err
	}
	fmt.Println("Mesa agregada correctamente.")
	return nil
}

// Actualizar la disponiblidad de una mesa (antes y despues de reservar)
func ActualizarMesa(db *sql.DB, disponible, mesaId int) error {
	// Consulta para actualizar la disponibilidad
	query := `UPDATE Mesas SET Disponible = ? WHERE ID = ?`
	// Ejecuta la consulta
	_, err := db.Exec(query, disponible, mesaId)
	if err != nil {
		log.Printf("Error al actualizar el mesa: %v", err)
		return err
	}
	fmt.Println("Mesa actualizada correctamente.")
	return nil
}
