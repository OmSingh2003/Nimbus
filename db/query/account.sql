-- name: CreateAccount :one
INSERT INTO account (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;