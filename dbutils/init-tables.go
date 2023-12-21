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
	log.Println("All tables created/initialized successfully!")
}
