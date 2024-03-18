package config

import (
	"database/sql"
	"fmt"
	"log"
	"reviewsApp/server/helpers"

	_ "github.com/lib/pq" //Postgres golang driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "reviewsDB"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)

	// url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// pql.conf.Postgres.User,
	// pql.conf.Postgres.Password,
	// pql.conf.Postgres.Host,
	// pql.conf.Postgres.Port,
	// pql.conf.Postgres.DB)

	db, err := sql.Open("postgres", sqlInfo)
	helpers.ThrowIfError(err)

	err = db.Ping()
	helpers.ThrowIfError(err)

	log.Printf("Connected to database!")
	//log.Info().Msg("Connected to database!!")

	return db
}
