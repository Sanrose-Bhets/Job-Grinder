-- name: CreateApplications :one
INSERT INTO applications( department, status, companyid, positionid, interviewId, userid)
VALUES($1, $2, $3,$4, $5, $6)
RETURNING *;