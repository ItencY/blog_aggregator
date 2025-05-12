package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/itency/blog_aggregator/internal/config"
	"github.com/itency/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Step 1: Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Step 2: Connection DB
	db, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Step 3: Create new *database.Queries
	dbQueries := database.New(db)

	// Step 3: Store the config in a new State instance

	state := &State{
		cfg: &cfg,
		db:  dbQueries,
	}

	// Step 4: Create new instance commands with an initialized map of handler functions
	commands := Commands{validCommands: make(map[string]func(*State, Command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)

	// Step 5: Check the command-line arguments passed in by the user
	if len(os.Args) < 2 {
		fmt.Println("error: no command passed")
		os.Exit(1)
	}

	// Step 6: Split the command-line arguments into the command name and the arguments slice to create a command instance
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

	// Step 7: Read the config file again
	updateCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	// Step 8: Print the config
	fmt.Printf("Updated Config:\nDB URL: %s\nCurrent User: %s\n", updateCfg.DBConnection, updateCfg.UserName)
}
