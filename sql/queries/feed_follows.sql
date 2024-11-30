-- name: CreateFeedFollow :many
with inserted_feed_follow as(
    insert into feed_follow (id, created_at, updated_at, user_id, feed_id)
    values (
        $1,
        $2,
        $3,
        $4,
        $5
    )
        returning *
)

select
    inserted_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
from inserted_feed_follow
inner join feeds on inserted_feed_follow.feed_id = feeds.id
inner join users on inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
select 
    feed_follow.*,
    feeds.name,
    users.name
    from feed_follow
inner join users on users.id = feed_follow.user_id
inner join feeds on feeds.id = feed_follow.feed_id
where users.name = $1;

-- name: DeleteFeedFollowByUrl :exec
DELETE FROM feed_follow
USING feeds
WHERE feeds.id = feed_follow.feed_id
AND feeds.url = $1
AND feed_follow.user_id = $2;

