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

func handleAgg (_ *state, _ command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed.Channel)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil) 
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}
	
	var feed *RSSFeed

	err = xml.Unmarshal(body,  &feed)
	if err != nil {
		return nil, err
	}
	unescapeRSSFeed(feed)
	return feed, err
}


func unescapeRSSFeed(feed *RSSFeed){
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Link = html.UnescapeString(feed.Channel.Link)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Link = html.UnescapeString(feed.Channel.Item[i].Link)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
		feed.Channel.Item[i].PubDate = html.UnescapeString(feed.Channel.Item[i].PubDate)
	}
}
