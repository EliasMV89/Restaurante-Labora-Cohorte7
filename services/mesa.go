package services

import (
	"database/sql"
	"log"
)

type Mesa struct {
	NumeroDeMesa int `json:"numero_de_mesa"`
	Capacidad    int `json:"capacidad"`
	Disponible   int `json:"disponible"`
}

func AgregarMesa(db *sql.DB, mesa Mesa) (int64, error) {
	query := `INSERT INTO Mesas (NumeroDeMesa, Capacidad, Disponible) VALUES (?, ?, 1)`
	result, err := db.Exec(query, mesa.NumeroDeMesa, mesa.Capacidad)
	if err != nil {
		log.Printf("Error al agregar Mesa: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID de la mesa: %v", err)
		return 0, err
	}

	log.Printf("Mesa agregada correctamente con ID: %d", id)
	return id, nil
}

func ActualizarMesa(db *sql.DB, disponible, mesaId int) error {
	query := `UPDATE Mesas SET Disponible = ? WHERE ID = ?`
	_, err := db.Exec(query, disponible, mesaId)
	if err != nil {
		log.Printf("Error al actualizar mesa: %v", err)
		return err
	}
	log.Printf("Mesa actualizada correctamente con ID: %d", mesaId)
	return nil
}

func ListarMesas(db *sql.DB) ([]Mesa, error) {
	query := `SELECT NumeroDeMesa, Capacidad, Disponible FROM Mesas`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar mesas: %v", err)
		return nil, err
	}
	defer rows.Close()

	var mesas []Mesa
	for rows.Next() {
		var mesa Mesa
		if err := rows.Scan(&mesa.NumeroDeMesa, &mesa.Capacidad, &mesa.Disponible); err != nil {
			log.Printf("Error al leer filas: %v", err)
			continue
		}
		mesas = append(mesas, mesa)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return mesas, nil
}
