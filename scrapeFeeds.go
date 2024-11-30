package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Barrioslopezfd/gator/internal/database"
	"github.com/google/uuid"
)


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markedFeed := database.MarkFeedFetchedParams{
		UpdatedAt:	time.Now(),
		ID:			nextFeed.ID,
	}
	s.db.MarkFeedFetched(context.Background(), markedFeed)

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	for _, item:= range feed.Channel.Item {
		desc := sql.NullString{
			String:		item.Description,
			Valid:		true,
		}
		ntime:=sql.NullTime{}
		t, err:=time.Parse(time.DateTime, item.PubDate)
		if err == nil {
			ntime = sql.NullTime{
				Time:	t,
				Valid:	true,
			}
		}
		post := database.CreatePostParams {
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: ntime,
			FeedID:      nextFeed.ID,
		}
		s.db.CreatePost(context.Background(), post)
	}
	fmt.Println("Done")
	return nil
}
