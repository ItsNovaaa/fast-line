-- name: CreateUser :one
INSERT INTO users (email, password, first_name, last_name, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, first_name, last_name, phone, is_active, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, password, first_name, last_name, phone, is_active, created_at, updated_at
FROM users
WHERE email = $1 AND is_active = TRUE;

-- name: GetUserByID :one
SELECT id, email, first_name, last_name, phone, is_active, created_at, updated_at
FROM users
WHERE id = $1 AND is_active = TRUE;

-- name: ListUsers :many
SELECT id, email, first_name, last_name, phone, is_active, created_at, updated_at
FROM users
WHERE is_active = TRUE
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*)
FROM users
WHERE is_active = TRUE;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3, phone = $4, updated_at = NOW()
WHERE id = $1 AND is_active = TRUE
RETURNING id, email, first_name, last_name, phone, is_active, created_at, updated_at;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2, updated_at = NOW()
WHERE id = $1 AND is_active = TRUE;

-- name: DeleteUser :exec
UPDATE users
SET is_active = FALSE, updated_at = NOW()
WHERE id = $1;

-- name: SearchUsers :many
SELECT id, email, first_name, last_name, phone, is_active, created_at, updated_at
FROM users
WHERE is_active = TRUE
  AND (
    first_name ILIKE '%' || $1 || '%' OR
    last_name ILIKE '%' || $1 || '%' OR
    email ILIKE '%' || $1 || '%'
  )
ORDER BY created_at DESC;