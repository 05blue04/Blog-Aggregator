package main

import (
	"context"
	"fmt"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	u, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
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

func handlerFollows(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	url := cmd.args[0]

	u, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
	}

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

	fmt.Printf("User %s is now following feed %s", resource.UserName, resource.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s *(no arguments expected)", cmd.name)
	}

	u, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), u.ID)
	return nil
}
