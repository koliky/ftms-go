package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var databaseUrl string

func SetupDatabase(dbUrl, config string) {
	databaseUrl = dbUrl
	if config == "create" {
		createTable()
	}
}

func getConnection() *sql.DB {
	dbConnect, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("can not connect database : %v", err)
	}
	return dbConnect
}

func createTable() {
	db := getConnection()
	defer db.Close()
	dataByte, err := ioutil.ReadFile("./script.sql")
	checkErr(err)
	strList := strings.Split(string(dataByte), ";")
	for i := 0; i < len(strList)-1; i++ {
		_, err := db.Exec(strList[i])
		checkErr(err)
	}
	fmt.Println("create database.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
