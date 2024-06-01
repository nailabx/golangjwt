package main

import (
	"fmt"
	mysqCf "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nailabx/golangjwt/config"
	"github.com/nailabx/golangjwt/db"
	"log"
	"os"
)

func main() {
	db, err := db.NewMySQLStorage(mysqCf.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("DB config %v", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Error creating driver %v", err)
	}

	fullPath := fmt.Sprintf("%s/cmd/migrate/migrations", os.Getenv("PWD"))
	log.Println("Running migrations from", fullPath)
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", fullPath),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance %v", err)
	}
	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		log.Println("Running migration up")
		if err := m.Up(); err != nil {
			log.Fatalf("Error running migration up %v", err)
		}
	}
	if cmd == "down" {
		log.Println("Running migration down")
		if err := m.Down(); err != nil {
			log.Fatalf("Error running migration down %v", err)
		}
	}

}
