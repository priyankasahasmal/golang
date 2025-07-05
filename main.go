package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	funk "github.com/thoas/go-funk"
    jet "github.com/go-jet/jet/v2/postgres"

    "github.com/priyankasahasmal/bootcamp_db_qc7b/public/model"
	"github.com/priyankasahasmal/bootcamp_db_qc7b/public/table"
	"github.com/priyankasahasmal/bootcamp_db_qc7b/utils"
)

var DB *sql.DB

const (
	HOST     = "dpg-d1ibu73ipnbc73bfqdrg-a.oregon-postgres.render.com"
	PORT     = "5432"
	USERNAME = "bootcamp_db_qc7b_user"
	PASSWORD = "25Zq6NeZxYyLS855V7c9dqbZjKNWU2ZH"
	DBNAME   = "bootcamp_db_qc7b"
)

type FetchAllUsersOutput struct {
	Id    int
	Email string
	Name  string
}

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

	if err = DB.Ping(); err != nil {
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
		return nil
	}

	var dest []model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users)

	err := stmt.Query(tx, &dest)
	if err != nil {
		*pointerErr = err
		return nil
	}

	return funk.Map(dest, func(user model.Users) FetchAllUsersOutput {
		return FetchAllUsersOutput{
			Id:    int(user.UsersID),
			Email: user.Email,
			Name:  utils.GetIfNotNilString(user.Name),
		}
	}).([]FetchAllUsersOutput)
}

func GetAllUsers() {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("DB transaction error:", err)
	}
	defer tx.Rollback()

	var queryErr error
	users := FetchAllUsersQuery(tx, &queryErr)
	if queryErr != nil {
		log.Println("Error fetching users:", queryErr)
		return
	}

	for _, u := range users {
		fmt.Printf("User: %+v\n", u)
	}
}

func main() {
	if err := CreateDbObject(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	GetAllUsers()
}


