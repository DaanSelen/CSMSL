package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sites *sql.DB
)

func initDB() {
	sites, _ = sql.Open("sqlite3", "./csmsl.db")
	statement, _ := sites.Prepare("CREATE TABLE IF NOT EXISTS site (id INTEGER PRIMARY KEY, app TEXT, url TEXT)")
	statement.Exec()
	statement.Close()
}

func addSiteDatabase(app, site string) {
	statement, _ := sites.Prepare("INSERT INTO site (app, url) VALUES (?, ?)")
	statement.Exec(app, site)
	checkInput()

}

func delSiteDatabase(app string) {
	statement, _ := sites.Prepare("DELETE FROM site WHERE app == ?;")
	statement.Exec(app)
	checkInput()
}

func getAllSites() {
	fmt.Println("Getting all sites.")
	rows, _ := sites.Query("SELECT * FROM site")
	var site Site
	for rows.Next() {
		rows.Scan(&site.ID, &site.APP, &site.URL)
		fmt.Println(site)
	}
	checkInput()
}
