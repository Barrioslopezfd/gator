# GATOR
A CLI tool to fetch, manage, and display posts from feed URLs using Go and PostgreSQL.

## Getting Started
To get started, you'll need:
- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

## Install GATOR! 
```bash
go install github.com/Barrioslopezfd/gator@latest
```

Once installed, GATOR allows you to:
- Fetch posts from feed URLs.
- Store them in a PostgreSQL database.
- Display them in a user-friendly CLI format.

### Configuration
The login information and database URL will be stored in a `.gatorconfig.json` file located in your home directory. 
Here’s an example configuration:
```json
{
   "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
   "current_user_name": "This is where we will store the name of the currently logged-in user"
}
```
Make sure to replace username and password with your PostgreSQL credentials.

## Available Commands

### Help! 
- If you ever need access to this section just tyme `gator help`!

### User Management

- `gator register <username>`: Register a new user and set them as the logged-in user.
- `gator login <username>`: Log in as an existing user.
- `gator reset`: Erase all users.
- `gator users`: Display all registered users.

### Feed Management

- `gator addfeed <username> <Feed_URL>`: Add a new feed URL.
- `gator feeds`: Show all available feeds.
- `gator follow <Feed_URL>`: Follow a specific feed as the logged-in user.
- `gator following`: Display all feeds followed by the logged-in user.
- `gator unfollow <Feed_URL>`: Unfollow a specific feed.

### Post Aggregation

- `gator agg <Time_between_requests>`: Aggregate posts from all added feeds into the database.

#### Important 

- The format is #s/m/h, where # is a number and s/m/h are seconds, minutes and hours respectively.
- The program will fail if you dont write the time in this format.

### Browsing Posts

- `gator browse`: Display the content of posts.

## Example Usage
### Registering a User

```bash 
gator register alice
```

### Adding and Following a Feed

```bash
gator addfeed https://example.com/feed
gator follow https://example.com/feed
```

### Aggregating Posts

```bash
gator agg 1m
```
- m means minutes and the format is number m as you have seen
- This is the minutes between requests to fetch posts



# Notes

    - Ensure your PostgreSQL database is set up and accessible via
      the `db_url` in your configuration file.
    - Run `gator reset` with caution as it will delete all user data.

## 🤝 Contributing

### You can clone the repo

```bash
git clone https://github.com/Barrioslopezfd/gator@latest
cd gator
```

