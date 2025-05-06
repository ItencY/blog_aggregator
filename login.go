package main

import "fmt"

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("handler expects a single argument - the username")
	}

	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error loging in user: %s, %v", userName, err)
	}

	fmt.Printf("user %s logged in\n", userName)
	return nil
}
