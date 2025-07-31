-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
	INSERT INTO feeds_follows(id,created_at,updated_at,user_id,feed_id)
    VALUES($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT inserted_feed_follow.*, feeds.name AS feed_name, users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name FROM feeds
INNER JOIN feeds_follows ON feeds.id = feeds_follows.feed_id
WHERE feeds_follows.user_id = $1;


-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows
WHERE feeds_follows.user_id = $1 AND feeds_follows.feed_id = (SELECT id FROM feeds WHERE url = $2);