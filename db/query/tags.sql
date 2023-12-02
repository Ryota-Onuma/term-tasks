-- name: GetTag :one
SELECT * FROM tags
WHERE tags.id = ? LIMIT 1;


-- name: ListTags :many
SELECT * FROM tags
WHERE tags.id = ? LIMIT 1;

