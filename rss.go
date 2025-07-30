package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator") //common practice to identify program to server

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

	resource.Channel.Title = html.UnescapeString(resource.Channel.Title)
	resource.Channel.Description = html.UnescapeString(resource.Channel.Title)

	for i := range resource.Channel.Item {
		resource.Channel.Item[i].Title = html.UnescapeString(resource.Channel.Item[i].Title)
		resource.Channel.Item[i].Description = html.UnescapeString(resource.Channel.Item[i].Description)
	}

	return &resource, nil
}

func printRSS(r *RSSFeed) {
	fmt.Printf("Title: %s\nLink: %s\nDescription: %s\n", r.Channel.Title, r.Channel.Link, r.Channel.Description)

	for i, item := range r.Channel.Item {
		fmt.Printf("Item %d:\nTitle: %s\nLink: %s\nDescription:%s\nPubDate:%s\n", i, item.Title, item.Link, item.Description, item.PubDate)
	}
}
