-- name: CreateCompany :one
INSERT INTO companies( companyName, companyLocation, companyWebsite, companyEmail, HiringManagerName)
VALUES($1, $2, $3, $4, $5)
RETURNING *;