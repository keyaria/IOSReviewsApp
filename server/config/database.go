package config

import (
	"database/sql"
	"log"
	"os"
	"reviewsApp/server/helpers"

	_ "github.com/lib/pq" // Postgres golang driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "reviewsDB"
)

func DatabaseConnection() *sql.DB {
	//	sqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	helpers.ThrowIfError(err)

	err = db.Ping()
	helpers.ThrowIfError(err)

	log.Printf("Connected to database!")

	return db
}
