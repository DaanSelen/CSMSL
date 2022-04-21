package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	version = "V0.1"
	creator = "Daan S"
	purpose = "To remove unwanted security risks created by having to search the web for software repeatedly."
)

var (
	quit        = make(chan struct{})
	allCommands [][]string
	activeIndex int
)

type Site struct {
	ID  int    `json:"id"`
	APP string `json:"app"`
	URL string `json:"url"`
}

func initAll() {
	fmt.Println("Welcome to CeldServ MSL!")
	initHelpCommand()
	initDB()
}

func getAddSiteInfo() {
	fmt.Println("Enter the application's name you wish to add.")
	var inputApp string
	fmt.Scanln(&inputApp)
	fmt.Println("Enter the full site URL you wish to add.")
	var inputSite string
	fmt.Scanln(&inputSite)
	check := checkIfHttps(inputSite)
	fmt.Println()
	if check {
		addSiteDatabase(inputApp, inputSite)
	} else {
		fmt.Println("Insecure site, try again.")
		checkInput()
	}
}

func getDelSiteInfo() {
	fmt.Println("Enter the application's name you wish to delete")
	var inputApp string
	fmt.Scanln(&inputApp)
	fmt.Println()
	delSiteDatabase(inputApp)
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
		"exit", "Quits the application entirely.", "quit",
		"addsite", "Gives the ability to add a site to the database.", "adds",
		"deletesite", "Gives the ability to add a site to the database.", "dels",
		"showall", "Shows all the sites stored in the database.", "seeall",
	}

	for x := 0; x <= (len(rawCommands) - 3); {
		singleCommand := []string{rawCommands[x], rawCommands[x+1], rawCommands[x+2]}
		allCommands = append(allCommands, singleCommand)
		x += 3
	}
}

func startProcess() {
	quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				scanLinks()
				time.Sleep(10 * time.Second)
			}
		}
	}()
	checkInput()
}

func scanLinks() {
	fmt.Println("10 Secondns")
}

func stopProcess() {
	close(quit)
	checkInput()
}

func checkInput() {
	fmt.Println("Enter a command to begin/continue.")
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
	case allCommands[2][0], allCommands[2][2]: //start
		fmt.Println("STARTING PROCESS")
		startProcess()
	case allCommands[3][0], allCommands[3][2]: //stop
		fmt.Println("STOPPING PROCESS")
		stopProcess()
	case allCommands[4][0], allCommands[4][2]: //exit
		fmt.Println("\nExiting. Press the 'ENTER' key to close the application.")
		fmt.Scanln()
		os.Exit(0)
	case allCommands[5][0], allCommands[5][2]: //addsite
		getAddSiteInfo()
	case allCommands[6][0], allCommands[6][2]: //deletesite
		getDelSiteInfo()
	case allCommands[7][0], allCommands[7][2]: //showall
		getAllSites()
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
