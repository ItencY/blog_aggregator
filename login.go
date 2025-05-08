package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("handler expects a single argument - the username")
	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("username %s not registered with database", userName)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error logging in user %s: %v", userName, err)
	}

	fmt.Printf("user %s logged in\n", userName)
	return nil
}
