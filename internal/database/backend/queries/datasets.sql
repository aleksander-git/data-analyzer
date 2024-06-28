-- name: CreateDataset :one
INSERT INTO datasets (user_id, dataset_status_id, database_id, table_name)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetDatasetByID :one
SELECT * FROM datasets
WHERE dataset_id = $1 LIMIT 1;

-- name: UpdateDataset :one
UPDATE datasets
SET 
  table_name = $2,
  is_archieved = $3,
  archieved_at = $4,
  dataset_status_id = $5,
  database_id = $6
WHERE dataset_id = $1 RETURNING *;

