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
	version = "V0.01"
	creator = "Daan S"
	purpose = "To remove unwanted security risks created by having to search the web for software repeatedly."
)

var (
	allCommands [][]string
)

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
	database, _ := sql.Open("sqlite3", "./csmsl.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS site (id INTEGER PRIMARY KEY, app TEXT, url TEXT)")
	statement.Exec()
}

func initHelpCommand() {
	rawCommands := []string{
		"info", "Info shows information about the software version, creator and purpose.", "i",
		"help", "Help displays all the available commands in the application.", "h",
		"start", "Starts the downloading and monitoring process.", "begin",
		"stop", "Stops the downloading and monitoring process.", "end",
		"exit", "Quits the application entirely.", "e",
	}

	for x := 0; x <= (len(rawCommands) - 1); {
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
		fmt.Println("\nAvailable commands:")
		for x := range allCommands {
			fmt.Println(allCommands[x][0] + ": " + allCommands[x][1])
		}
		fmt.Println()
		checkInput()
	case allCommands[4][0], allCommands[4][2]: //exit
		fmt.Println("\nExiting. Press the 'ENTER' key to close the application.")
		fmt.Scanln()
		os.Exit(0)

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
