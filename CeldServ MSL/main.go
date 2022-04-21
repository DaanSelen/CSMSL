package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/jasonlvhit/gocron"
	_ "github.com/mattn/go-sqlite3"
)

const (
	version = "V0.1"
	creator = "Daan S"
	purpose = "To remove unwanted security risks created by having to search the web for software repeatedly."
)

var (
	sites       *sql.DB
	allCommands [][]string
)

type Site struct {
	ID  int    `json:"id"`
	APP string `json:"app"`
	URL string `json:"url"`
}

func myTask() {
	fmt.Println("This task will run periodically")
}
func executeCronJob() {
	gocron.Every(1).Minute().Do(myTask)
	<-gocron.Start()
}

func initAll() {
	fmt.Println("Welcome to CeldServ MSL!")
	initHelpCommand()
	initDB()
}

func initDB() {
	sites, _ = sql.Open("sqlite3", "./csmsl.db")
	statement, _ := sites.Prepare("CREATE TABLE IF NOT EXISTS site (id INTEGER PRIMARY KEY, app TEXT, url TEXT)")
	statement.Exec()
	statement.Close()
	rows, _ := sites.Query("SELECT * FROM site")
	var site Site
	for rows.Next() {
		rows.Scan(&site.ID, &site.APP, &site.URL)
		fmt.Println(site)
	}
}

func getAddSiteInfo() {
	fmt.Println("Enter the application's name you wish to add.")
	var inputApp string
	fmt.Scanln(&inputApp)
	fmt.Println("Enter the full site URL you wish to add.")
	var inputSite string
	fmt.Scanln(&inputSite)
	check := checkIfHttps(inputSite)
	if check {
		addSiteDatabase(inputApp, inputSite)
	}
}

func addSiteDatabase(app, site string) {
	statement, _ := sites.Prepare("INSERT INTO site (app, url) VALUES (?, ?)")
	statement.Exec(app, site)

}

func getDelSiteInfo() {
	fmt.Println("Enter the application's name you wish to delete")
	var inputApp string
	fmt.Scanln(&inputApp)
}

func checkIfHttps(candidate string) bool {
	var check bool
	if strings.Contains(candidate, "https://") {
		check = true
	} else {
		check = false
	}
	return check
}

func initHelpCommand() {
	rawCommands := []string{
		"info", "Info shows information about the software version, creator and purpose.", "i",
		"help", "Help displays all the available commands in the application.", "h",
		"start", "Starts the downloading and monitoring process.", "begin",
		"stop", "Stops the downloading and monitoring process.", "end",
		"exit", "Quits the application entirely.", "e",
		"addsite", "Gives the ability to add a site to the database.", "adds",
		"deletesite", "Gives the ability to add a site to the database.", "dels",
	}

	for x := 0; x <= (len(rawCommands) - 3); {
		singleCommand := []string{rawCommands[x], rawCommands[x+1], rawCommands[x+2]}
		allCommands = append(allCommands, singleCommand)
		x += 3
	}
}

func checkInput() {
	fmt.Println("Enter a command to begin.")
	input := ""
	switch fmt.Scanln(&input); strings.ToLower(input) {
	case allCommands[0][0], allCommands[0][2]: //info
		fmt.Println("\nCurrent Version:", version+"\nCreator:", creator+"\nPurpose:", purpose)
		fmt.Println()
		checkInput()
	case allCommands[1][0], allCommands[1][2]: //help
		fmt.Println("\nAvailable commands with expanation and aliases:")
		for x := range allCommands {
			fmt.Println(allCommands[x][0]+": "+allCommands[x][1], "Aliases:", allCommands[x][2])
		}
		fmt.Println()
		checkInput()
	case allCommands[4][0], allCommands[4][2]: //exit
		fmt.Println("\nExiting. Press the 'ENTER' key to close the application.")
		fmt.Scanln()
		os.Exit(0)
	case allCommands[5][0], allCommands[5][2]: //addsite
		getAddSiteInfo()
	case allCommands[6][0], allCommands[6][2]: //deletesite
		getDelSiteInfo()
	default: //unrecognised
		fmt.Println("Unrecognised command, try again.")
		checkInput()
	}
}

func main() {
	initAll()
	fmt.Println()
	checkInput()
}
