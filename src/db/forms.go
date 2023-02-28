package db

import (
	"database/sql"
	"fmt"
	"nocode/src/models"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	var (
		host    = os.Getenv("SQL_HOST")
		port, _ = strconv.Atoi(os.Getenv("SQL_PORT"))
		user    = os.Getenv("SQL_USER")
		pass    = os.Getenv("SQL_PASS")
		dbname  = os.Getenv("DBNAME")
	)
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	var err error
	db, err = sql.Open("postgres", psqlConn)
	CheckError(err)

	// defer db.Close()
	err = db.Ping()
	CheckError(err)

	fmt.Println("connection successfull")
}

func QueryAll() ([]models.Form, error) {

	db := GetDB()
	rows, err := db.Query("SELECT id, name, gender, age FROM forms")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	defer rows.Close()
	forms := make([]models.Form, 0)
	for rows.Next() {
		var form models.Form
		if err := rows.Scan(&form.Id, &form.Name, &form.Gender, &form.Age); err != nil {
			return nil, fmt.Errorf("error scanning form row: %v", err)
		}
		fmt.Println(form)
		forms = append(forms, form)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading form rows: %v", err)
	}

	return forms, nil
}

func GetDB() *sql.DB {
	if db == nil {
		InitDb()
		//CheckError(fmt.Errorf("database not initialized"))
	}
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
