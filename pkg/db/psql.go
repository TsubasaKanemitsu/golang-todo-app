package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/config"
)

func NewPSQL(c config.Postgres) *sql.DB {
	postgresURL := GetPostgresURL(c.DBName, c.Host, c.Port, c.User, c.Pass, c.Sslmode)
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Println("failed to open conenection..")
		panic(err)
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(2 * time.Second)

	return db
}

func GetPostgresURL(dbName, host, port, user, pass, sslMode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		host, port, user, dbName, sslMode, pass,
	)
}
