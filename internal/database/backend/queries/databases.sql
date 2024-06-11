-- name: CreateDatabase :one
INSERT INTO databases (name, descr, dsn, token)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetDatabaseByID :one
SELECT * FROM databases
WHERE database_id = $1 LIMIT 1;

-- name: UpdateDatabase :one
UPDATE databases
SET 
    name = $2,
    descr = $3,
    dsn = $4,
    token = $5
WHERE database_id = $1 RETURNING *;