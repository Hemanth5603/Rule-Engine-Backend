package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var POSTGRES_DB *sql.DB
var POSTGRES_CONNECTION_STRING string

func init() {
	InitializePostgresSQL()
}

func InitializePostgresSQL() {
	var err error
	USER := "postgres"
	PASS := "iitt-admin"
	HOST := "iitt-db.c34oscioyvnc.eu-north-1.rds.amazonaws.com"
	DBNAME := "zeotap"
	PORT := "5432"

	POSTGRES_CONNECTION_STRING = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", HOST, USER, PASS, DBNAME, PORT)

	POSTGRES_DB, err = sql.Open("pgx", POSTGRES_CONNECTION_STRING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to POSTGRES database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("conncted to database")
	}
	POSTGRES_DB.SetMaxIdleConns(10)

}
