package main

import (
	"database/sql"
	"log"

	"github.com/I-Maged/ecom/cmd/api"
	"github.com/I-Maged/ecom/config"
	"github.com/I-Maged/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.ENVS.DBUser,
		Passwd:               config.ENVS.DBPassword,
		Addr:                 config.ENVS.DBAddress,
		DBName:               config.ENVS.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
