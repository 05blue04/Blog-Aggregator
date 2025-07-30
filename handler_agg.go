package main

import (
	"context"
	"fmt"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), feed)

	if err != nil {
		return err
	}

	printRSS(rss)

	return nil
}

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
