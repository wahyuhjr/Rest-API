-- name: GetExecutionTimes :many
SELECT * FROM execution_time;

-- name: CreateExecutionTime :one
INSERT INTO execution_time (parameter,test,value, deviation)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExecutionTimeByID :one
SELECT * FROM execution_time WHERE id = $1;

-- name: UpdateExecutionTime :one
UPDATE execution_time
SET parameter = $2, test = $4, value = $3,  deviation = $5
WHERE id = $1
RETURNING *;

-- name: DeleteExecutionTime :one
DELETE FROM execution_time WHERE id = $1
RETURNING *;