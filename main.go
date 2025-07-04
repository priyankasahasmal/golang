package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/priyankasahasmal/golang/.gen/bootcamp_db_qc7b/public/model"
	"github.com/priyankasahasmal/golang/.gen/bootcamp_db_qc7b/public/table"

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

func FetchAllUsersQuery(tx *sql.Tx, pointerErr *error) []FetchAllUsersOutput {
    if *pointerErr != nil {
        return []FetchAllUsersOutput{}
    }

    type destType struct {
        model.User // ✅ Not Users
    }
    var dest []destType

    stmt := postgres.SELECT(
        table.Users.AllColumns, // ✅ Must match actual table
    ).FROM(table.Users)

    err := stmt.Query(tx, &dest)
    if err != nil {
        *pointerErr = err
        return []FetchAllUsersOutput{}
    }

    return funk.Map(dest, func(item destType) FetchAllUsersOutput {
        return FetchAllUsersOutput{
            Id:    int(item.User.UserID),
            Email: item.User.Email,
            Name:  utils.GetIfNotNilString(item.User.Name),
        }
    }).([]FetchAllUsersOutput)
}
