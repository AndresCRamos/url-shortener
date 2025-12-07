-- name: GetShortenLink :one
SELECT
id, original_url, created_at, views
FROM SHORTEN_LINKS
WHERE short_code = ?
LIMIT 1;

-- name: CreateShortenLink :one
INSERT INTO SHORTEN_LINKS
(short_code, original_url)
VALUES(
    ?, ?
)
RETURNING *;

-- name: IncrementViews :exec
UPDATE shorten_links
SET views = views + 1
WHERE short_code = ?;

