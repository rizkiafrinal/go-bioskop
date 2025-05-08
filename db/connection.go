package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func Connect() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=Aviat12 dbname=bioskop_db sslmode=disable"
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	if err = Conn.Ping(); err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}

	fmt.Println("Koneksi ke database berhasil.")
}
