package main

import (
	"fmt"
	"os"
)

func initDir() {
	err := os.Mkdir("Downloaded Software Setups", 0777)
	if err != nil {
		fmt.Println("Downloaded Software Setups directory already exists - or something went wrong.")
	}
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
		"cleardb", "Clears the entire DB table", "clrdb",
	}

	for x := 0; x <= (len(rawCommands) - 3); {
		singleCommand := []string{rawCommands[x], rawCommands[x+1], rawCommands[x+2]}
		allCommands = append(allCommands, singleCommand)
		x += 3
	}
}
