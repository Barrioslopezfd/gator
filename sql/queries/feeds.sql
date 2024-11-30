-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetAlmostFeed :many
SELECT feeds.name, feeds.url, users.name
    FROM feeds
JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedByUrl :one
select * from feeds
where url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $1,last_fetched_at=$1
WHERE id = $2;

-- name: GetNextFeedToFetch :one
SELECT * from feeds
order by feeds.last_fetched_at NULLS FIRST
LIMIT 1;
