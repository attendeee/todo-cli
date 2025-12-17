package repository

import (
	"fmt"
	"log"
	"task/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func initTables() {
	if _, err := Db().Exec(taskCreateTable); err != nil {
		panic(fmt.Errorf("Unable to create table task: %s\n", err))
	}

}

var db *sqlx.DB

func InitDb() {
	var err error
	db, err = sqlx.Open("sqlite3", config.GetSqlPath())
	if err != nil {
		log.Fatal(fmt.Errorf("Database init error: %s\n", err))
	}

	initTables()

}

func Db() *sqlx.DB {
	if db == nil {
		InitDb()
	}
	return db
}
