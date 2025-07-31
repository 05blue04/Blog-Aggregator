package main

import (
	"context"
	"fmt"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollows(s *state, cmd command, u database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	url := cmd.args[0]

	f, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    f.ID,
	}

	resource, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("User %s is now following feed %s\n", resource.UserName, resource.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, u database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s *(no arguments expected)", cmd.name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), u.ID)
	if err != nil {
		return err
	}

	for _, name := range following {
		fmt.Println(name)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, u database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <Feed url>", cmd.name)
	}

	url := cmd.args[0]

	params := database.DeleteFeedFollowParams{
		UserID: u.ID,
		Url:    url,
	}
	err := s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}
