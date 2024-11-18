-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (username, email)
VALUES (?, ?);

-- name: UpdateUser :exec
UPDATE users
SET username = ?,
    email = ?,
    updated_at = NOW()
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?; 