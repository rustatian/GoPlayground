package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	dStub "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func main() {
	// Create and use an existing database instance.
	db, err := sql.Open("postgres", "postgres://postgres:1@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create driver instance from db.
	// Check each driver if it supports the WithInstance function.
	// `import "github.com/mattes/migrate/database/postgres"`
	instance, err := dStub.WithInstance(db, &dStub.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.NewWithDatabaseInstance("/Users/0xdev/go/src/github.com/ValeryPiashchynski/GoPlayground/migrations/1_AddInventoryTables.up.sql", "postgres", instance)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err = migrate.New("/Users/0xdev/go/src/github.com/ValeryPiashchynski/GoPlayground/migrations/1_AddInventoryTables.up.sql", "postgres://postgres:1@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
