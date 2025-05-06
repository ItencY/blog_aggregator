package main

import (
	"fmt"
	"log"

	"github.com/itency/blog_aggregator/internal/config"
)

func main() {
	// Step 1: Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Step 2: Set the current user to "Alex" and write to the config file
	if err := cfg.SetUser("Alex"); err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	// Step 3: Read the config file again
	updateCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	// Step 4: Print the config
	fmt.Printf("Updated Config:\nDB URL: %s\nCurrent User: %s\n", updateCfg.DBConnection, updateCfg.UserName)
}
