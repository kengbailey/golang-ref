/*
	Kenneth Bailey
	4/1/18

	Postgresql and Go

	Notes:
		- postgres module being used: github.com/lib/pq
		- go get github.com/lib/pq
		- We must ping database after opening a connection in order to connect to db

*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

var db *sql.DB
var dbConfig config

func main() {

	// config
	getDbConfig()

	// connect to database
	initDb()

}

func getDbConfig() {
	file, err := ioutil.ReadFile("./postgres_config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &dbConfig)
}

func initDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port,
		dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to Database!!")
}
