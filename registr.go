package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/itency/blog_aggregator/internal/database"
)

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("requires a username")
	}

	userName := cmd.args[0]

	createUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	user, err := s.db.CreateUser(context.Background(), createUser)
	if err != nil {
		return fmt.Errorf("user already registered: %v", err)
	}

	s.cfg.SetUser(userName)

	fmt.Printf("user %s created and registered\n", user.Name)
	fmt.Println(user)

	return nil
}
