package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	cfg "github.com/ledepede1/url-shortener/pkg/config"
)

var dbUrl = cfg.DBConnectionString

type Database struct {
	*sql.DB
}

func EstablishDBCon() (Database, error) {

	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		fmt.Print(err.Error())
		log.Fatalf("Failed to establish connection to: %s", dbUrl)
	}

	return Database{db}, nil
}
