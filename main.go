package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

const (
	HOST     = "dpg-d1ibu73ipnbc73bfqdrg-a.oregon-postgres.render.com"
	PORT     = "5432"
	USERNAME = "bootcamp_db_qc7b_user"
	PASSWORD = "25Zq6NeZxYyLS855V7c9dqbZjKNWU2ZH"
	DBNAME   = "bootcamp_db_qc7b"
)

func GetPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		HOST, PORT, USERNAME, PASSWORD, DBNAME)
}

func CreateDbObject() error {
	var err error

	DB, err = sql.Open("postgres", GetPsqlInfo())
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %w", err)
	}

	fmt.Println("Connected successfully")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(1 * time.Hour)

	return nil
}

func main() {
	if err := CreateDbObject(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Continue with your logic here
}

