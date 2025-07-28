package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

type RSSFeed struct {
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resource RSSFeed

	err = xml.Unmarshal(data, &resource)

	if err != nil {
		return nil, err
	}

	return nil, nil //place holder for neow
}
