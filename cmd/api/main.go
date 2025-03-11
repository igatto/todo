package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/igatto/todo/internal/store"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	app := &application{
		store: store,
	}
	app.run()
}
