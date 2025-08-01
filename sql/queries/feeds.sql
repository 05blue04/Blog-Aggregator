-- name: CreateFeed :one
INSERT INTO feeds(id,name,url,created_at,updated_at,user_id)
VALUES($1, $2, $3 ,$4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * FROM feeds 
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $1, last_fetched_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;
