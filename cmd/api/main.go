package main

import (
	"database/sql"
	"fmt"
	"log"
	"msg-inventory/data"
	"net/http"
	"time"
)

const port = "13780"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting inventory serivce on port: %s\n", port)

	// Connect to the database
	conn := connectToDB()
	if conn == nil {
		log.Panic("Could not connect to the database")
	}

	// Setup AppConfig
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// Start the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := "host=db-postgres port=5432 dbname=postgres user=postgres password=P@ssw0rd sslmode=disable timezone=UTC connect_timeout=5"

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres is unavailable ...")
			count++
		} else {
			log.Println("Postgres is available ...")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue

	}

}
