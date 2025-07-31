package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/05blue04/Blog-Aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	if len(cmd.args) != 0 && len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <Limit>*(optional argument)", cmd.name)
	}

	var (
		limit int
		err   error
	)

	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("couldn't convert %s to integer", cmd.args[0])
		}
	} else {
		limit = 2
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for i := 0; i < len(posts); i++ {
		printPost(posts[i])
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Printf("Title: %s\nLink: %s\nDescription:%v\nPubDate:%v\n", post.Title, post.Url, post.Description, post.PublishedAt)
}
