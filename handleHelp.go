package main 

import "fmt"

func handleHelp() {
	fmt.Println("register <username>: Register a new user and set them as the logged-in user.")
	fmt.Println("login <username>: Log in as an existing user.")
	fmt.Println("reset: Erase all users.:")
	fmt.Println("users: Display all registered users.")
	fmt.Println("addfeed <username> <Feed_URL>: Add a new feed URL.")
	fmt.Println("feeds: Show all available feeds.")
	fmt.Println("follow <Feed_URL>: Follow a specific feed as the logged-in user.")
	fmt.Println("following: Display all feeds followed by the logged-in user.")
	fmt.Println("unfollow <Feed_URL>: Unfollow a specific feed.")
	fmt.Println("agg <Time_between_requests>: Aggregate posts from all added feeds into the database.")
	fmt.Println("The format is #s/m/h, where # is a number and s/m/h are seconds, minutes and hours respectively.")
	fmt.Println("The program will fail if you dont write the time in this format.")
	fmt.Println("browse: Display the content of posts.")
}
