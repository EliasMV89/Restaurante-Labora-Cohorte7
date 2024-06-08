/*
package services

import (

	"database/sql"
	"fmt"
	"log"

)

	type Reserva struct {
		ClienteID          int
		MesaID             int
		FechaReserva       string
		HoraReserva        string
		NumeroDeComensales int
	}

// Agregar una reserva a la tabla reservas

	func AgregarReserva(db *sql.DB, reserva Reserva) error {
		// Consulta para agregar reserva
		query := `INSERT INTO Reservas (ClienteID, MesaID, FechaReserva, HoraReserva, NumeroDeComensales) VALUES (?,?,?,?,?)`
		// Ejecuta la consulta
		_, err := db.Exec(query, reserva.ClienteID, reserva.MesaID, reserva.FechaReserva, reserva.HoraReserva, reserva.NumeroDeComensales)
		if err != nil {
			log.Printf("Error al insertar reserva: %v", err)
			return err
		}
		fmt.Println("Reserva agregada correctamente.")
		return nil
	}

	func AnalizarReserva(db *sql.DB, fechaReserva string) error {
		// Consulta para analizar reservas
		query := `SELECT DAYNAME(FechaReserva) AS DiaDeLaSemana, HOUR(HoraReserva) AS HoraDelDia, NumeroDeComensales, COUNT(*) AS Frecuencia FROM Reservas WHERE FechaReserva = ? GROUP BY DiaDeLaSemana, HoraDelDia, NumeroDeComensales ORDER BY Frecuencia DESC`
		// Ejecuta la consulta
		rows, err := db.Query(query, fechaReserva)
		if err != nil {
			log.Printf("Error al leer filas: %v", err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var DiaDeLaSemana string
			var HoraDelDia int
			var NumeroDeComensales int
			var Frecuencia int

			err := rows.Scan(&DiaDeLaSemana, &HoraDelDia, &NumeroDeComensales, &Frecuencia)
			if err != nil {
				log.Printf("Error al iterar fila: %v", err)
				return err
			}
			fmt.Printf("DiaDeLaSemana: %s, HoraDelDia: %d, NumeroDeComensales: %d, Frecuencia: %d\n", DiaDeLaSemana, HoraDelDia, NumeroDeComensales, Frecuencia)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Error al iterar sobre filas: %v", err)
			return err
		}

		return nil
	}
*/
package services

import (
	"database/sql"
	"log"
)

type Reserva struct {
	ClienteID          int    `json:"cliente_id"`
	MesaID             int    `json:"mesa_id"`
	FechaReserva       string `json:"fecha_reserva"`
	HoraReserva        string `json:"hora_reserva"`
	NumeroDeComensales int    `json:"numero_de_comensales"`
}

func AgregarReserva(db *sql.DB, reserva Reserva) (int64, error) {
	query := `INSERT INTO Reservas (ClienteID, MesaID, FechaReserva, HoraReserva, NumeroDeComensales) VALUES (?,?,?,?,?)`
	result, err := db.Exec(query, reserva.ClienteID, reserva.MesaID, reserva.FechaReserva, reserva.HoraReserva, reserva.NumeroDeComensales)
	if err != nil {
		log.Printf("Error al insertar reserva: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID de la reserva: %v", err)
		return 0, err
	}

	log.Printf("Reserva agregada correctamente con ID: %d", id)
	return id, nil
}

func AnalizarReserva(db *sql.DB, fechaReserva string) ([]map[string]interface{}, error) {
	query := `SELECT DAYNAME(FechaReserva) AS DiaDeLaSemana, HOUR(HoraReserva) AS HoraDelDia, NumeroDeComensales, COUNT(*) AS Frecuencia FROM Reservas WHERE FechaReserva = ? GROUP BY DiaDeLaSemana, HoraDelDia, NumeroDeComensales ORDER BY Frecuencia DESC`
	rows, err := db.Query(query, fechaReserva)
	if err != nil {
		log.Printf("Error al leer filas: %v", err)
		return nil, err
	}
	defer rows.Close()

	var analisis []map[string]interface{}
	for rows.Next() {
		var diaDeLaSemana string
		var horaDelDia int
		var numeroDeComensales int
		var frecuencia int

		err := rows.Scan(&diaDeLaSemana, &horaDelDia, &numeroDeComensales, &frecuencia)
		if err != nil {
			log.Printf("Error al iterar fila: %v", err)
			continue
		}

		analisis = append(analisis, map[string]interface{}{
			"DiaDeLaSemana":      diaDeLaSemana,
			"HoraDelDia":         horaDelDia,
			"NumeroDeComensales": numeroDeComensales,
			"Frecuencia":         frecuencia,
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar sobre filas: %v", err)
		return nil, err
	}

	return analisis, nil
}
