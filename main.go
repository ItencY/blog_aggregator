package main

import (
	"fmt"
	"log"
	"os"

	"github.com/itency/blog_aggregator/internal/config"
)

func main() {
	// Step 1: Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Step 2: Store the config in a new State instance

	state := &State{
		cfg: &cfg,
	}

	// Step 3: Create new instance commands with an initialized map of handler functions
	commands := Commands{validCommands: make(map[string]func(*State, Command) error)}
	commands.register("login", handlerLogin)

	// Step 4: Check the command-line arguments passed in by the user
	if len(os.Args) < 2 {
		fmt.Println("error: no command passed")
		os.Exit(1)
	}

	// Step 5: Split the command-line arguments into the command name and the arguments slice to create a command instance
	var cmdArgs []string
	if len(os.Args) == 2 {
		cmdArgs = []string{}
	} else {
		cmdArgs = os.Args[2:]
	}

	err = commands.run(state, Command{name: os.Args[1], args: cmdArgs})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Step 6: Read the config file again
	updateCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	// Step 4: Print the config
	fmt.Printf("Updated Config:\nDB URL: %s\nCurrent User: %s\n", updateCfg.DBConnection, updateCfg.UserName)
}
