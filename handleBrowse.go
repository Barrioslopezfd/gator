package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Barrioslopezfd/gator/internal/database"
)


func handleBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1 {
		specifiedLimit, err := strconv.Atoi(cmd.arguments[0]); 
		if err == nil {
			limit = specifiedLimit
		} else {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.PublishedAt) 
		fmt.Println(post.FeedName)
		fmt.Println(post.Title)
		fmt.Println(post.Description.String)
		fmt.Println(post.Url)
	}

	return nil
}
