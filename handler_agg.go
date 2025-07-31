package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/05blue04/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}

	log.Println("Found a feed to fetch!")
	params := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		log.Printf("Couldn't mark feed %s as fetched :%v", feed.Name, err)
		return
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	err = saveRSS(s, rss, feed.ID)
	if err != nil {
		log.Printf("Couldn't save a post to db :%v\n", err)
	}
	// printRSS(rss) uncomment to view posts being aggregated

}

func saveRSS(s *state, r *RSSFeed, feedId uuid.UUID) error {
	fmt.Printf("Title: %s\nLink: %s\nDescription: %s\n", r.Channel.Title, r.Channel.Link, r.Channel.Description)

	for _, item := range r.Channel.Item {
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullString{
				String: item.PubDate,
				Valid:  true,
			},
			FeedID: feedId,
		}
		_, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	return nil
}
