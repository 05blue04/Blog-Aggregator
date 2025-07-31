package main

import (
	"context"

	"github.com/05blue04/Blog-Aggregator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, u database.User) error,
) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		u, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			return err
		}

		return handler(s, cmd, u)
	}
}
