-- name: CreateUser :one
INSERT INTO accounts( Name, email, Role)
VALUES($1, $2, $3)
RETURNING *;