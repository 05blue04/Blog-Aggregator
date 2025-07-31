package main

import (
	"context"
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
