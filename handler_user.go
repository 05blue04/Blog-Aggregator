package main

import (
	"context"
	"fmt"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err != nil {
		return fmt.Errorf("user does not exist in database")
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting current_user %v", err)
	}

	fmt.Println("Username was set in config")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	u, err := s.db.CreateUser(context.Background(), params)

	if err != nil {
		return err
	}
	fmt.Printf("User %v was created with %v", u.Name, u)
	s.cfg.SetUser(u.Name)

	return nil
}
