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
	defer statement.Close()
	statement.Exec()
}

func addSiteDatabase(app, site string) {
	statement, _ := sites.Prepare("INSERT INTO site (app, url) VALUES (?, ?)")
	defer statement.Close()
	statement.Exec(app, site)
	checkInput()

}

func delSiteDatabase(app string) {
	statement, _ := sites.Prepare("DELETE FROM site WHERE app == ?;")
	defer statement.Close()
	statement.Exec(app)
	checkInput()
}

func getAllSites() []Site {
	fmt.Println("Getting all sites.")
	rows, _ := sites.Query("SELECT * FROM site")
	var sites []Site

	for rows.Next() {
		var site Site
		rows.Scan(&site.ID, &site.APP, &site.URL)
		sites = append(sites, site)
	}
	return sites
}

func clearDatabaseTable() {
	statement1, _ := sites.Prepare("DROP TABLE IF EXISTS site")
	statement2, _ := sites.Prepare("CREATE TABLE IF NOT EXISTS site (id INTEGER PRIMARY KEY, app TEXT, url TEXT)")
	defer statement1.Close()
	defer statement2.Close()
	statement1.Exec()
	statement2.Exec()
	fmt.Println("DATABASE CLEAR FINISHED.")
	checkInput()
}

func checkIfSitesPresent() bool {
	sites := getAllSites()
	if len(sites) == 0 {
		return false
	} else {
		return true
	}
}
