package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	exec, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("error command %v does not exist", cmd.name)
	}

	err := exec(s, cmd)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
