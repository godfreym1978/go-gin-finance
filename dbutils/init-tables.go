/* File: init-tables.go

Description:
 Initiates the connection to databases and also creates the backend table for mysql to persist the records in the table
*/

package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(orders)
	if driverError != nil {
		log.Println(driverError)
	}

	// Create train table
	_, statementError := statement.Exec()

	if statementError != nil {
		log.Println("Table already exists!")
	}

	log.Println("All tables created/initialized successfully!")
}
