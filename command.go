package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	validCommands map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	cmds, ok := c.validCommands[cmd.name]
	if !ok {
		return fmt.Errorf("invalid command")
	}

	return cmds(s, cmd)
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.validCommands[name] = f
}
