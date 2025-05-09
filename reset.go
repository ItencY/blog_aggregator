package main

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("failed resset: %v", err)
	}
	fmt.Println("reset successful")
	return nil
}
