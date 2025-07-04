-- name: CreateCircuit :one
INSERT INTO circuits (name, circuit_name, start_date, end_date, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, circuit_name, start_date, end_date, status, created_at, updated_at;

-- name: GetCircuitByID :one
SELECT id, name, circuit_name, start_date, end_date, status, created_at, updated_at
FROM circuits
WHERE id = $1;

-- name: ListCircuits :many
SELECT id, name, circuit_name, start_date, end_date, status, created_at, updated_at
FROM circuits
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountCircuits :one
SELECT COUNT(*)
FROM circuits;

-- name: UpdateCircuit :one
UPDATE circuits
SET name = $2, circuit_name = $3, start_date = $4, end_date = $5, status = $6, updated_at = NOW()   
WHERE id = $1
RETURNING id, name, circuit_name, start_date, end_date, status, created_at, updated_at;

-- name: DeleteCircuit :exec
UPDATE circuits
SET status = 0, updated_at = NOW()
WHERE id = $1;

-- name: GetActiveCircuits :many
SELECT id, name, circuit_name, start_date, end_date, status, created_at, updated_at
FROM circuits
WHERE status = 1
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;