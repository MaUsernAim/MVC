package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() {
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		log.Print("error open db", err)
		panic("FATAL ERROR")
	}
	defer db.Close()
	createTable(db)
}

func addDataLocate(db *sql.DB, obj locateFile) {
	query := `INSERT INTO local(name ,typePoint ,locate ,inCloud ,locateInCloud) VALUES(?,?,?,?,?)`
	_, err := db.Exec(query, obj.Name, obj.TypePoint, obj.InCloud, obj.Locate, obj.LocateInCloud)
	if err != nil {
		log.Print("error add data in table locate ", err)
		return
	}
}
func addDataCloud(db *sql.DB, obj cloud) { //db
	query := `INSERT INTO cloud(name,locate,type)`
	_, err := db.Exec(query, obj.nameC, obj.locateC, obj.typeC)
	if err != nil {
		log.Print("error add data to table cloud")
		return
	}
}
func getDataLocal(db *sql.DB) (obj locateFile) { //db
	defer log.Print("last log")
	log.Print("log1 start ")
	query := `
	SELECT * FROM local
	`
	err := db.QueryRow(query).Scan(&obj.Name, &obj.TypePoint, &obj.Locate, &obj.InCloud, &obj.LocateInCloud)
	log.Print("log 2 g", obj) //successfully
	if err != nil {
		log.Print("error get deta from first table", err)
		return
	}
	return obj
}
func getDataCloud(db *sql.DB) (obj cloud) { //db
	query := `
	SELECT * FROM local
	`
	err := db.QueryRow(query).Scan(&obj.nameC, &obj.locateC, &obj.typeC)
	if err != nil {
		log.Print("error get data from second table", err)
		return
	}
	return obj
}
func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS local(
	name TEXT NOT NULL,
	typePoint TEXT ,
	locate TEXT,
	inCloud TEXT,
	locateInCloud TEXT
	)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Print("error make table first ", err)
		return
	}
	query = `CREATE TABLE IF NOT EXISTS cloud(
	name TEXT NOT NULL,
	locate TEXT NOT NULL,
	type TEXT
	)`
	_, err = db.Exec(query)
	if err != nil {
		log.Print("error make table second ", err)
		return
	}
	log.Print("table sreate +")
}
