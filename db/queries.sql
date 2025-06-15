-- name: CreateGodo :one
INSERT INTO godos (text, done) VALUES (?, ?) RETURNING *;

-- name: GetGodo :one
SELECT * FROM godos WHERE id = ? LIMIT 1;

-- name: ListGodos :many
SELECT * FROM godos ORDER BY id;

-- name: UpdateGodoDone :exec
UPDATE godos SET done = ? WHERE id = ?
