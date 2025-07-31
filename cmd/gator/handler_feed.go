package main

import (
	"context"
	"fmt"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, u database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    u.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)

	if err != nil {
		return err
	}

	params2 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    params.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), params2)
	if err != nil {
		return err
	}

	fmt.Printf("Feed :%+v\n", feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s *(no arguments expected)", cmd.name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		userName, err := s.db.GetUserbyID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("feed %d\nName:%s\nURL:%s\nUserId:%s\n", i, feed.Name, feed.Url, userName)

	}

	return nil
}
