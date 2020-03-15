package db_config

import (
	"day_frame/tasks"
	"github.com/go-pg/pg/v9"
	"log"
	"os"
)

//Connecting to the databse
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "414702",
		Addr:     "localhost:5432",
		Database: "day_frame",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	tasks.CreateTaskTable(db)

	tasks.InitializeDB(db)

	return db
}
