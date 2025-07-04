-- name: CreateAccount :one
INSERT INTO account (
  owner, balance, currency, account_number
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountByNumber :one
SELECT * FROM account
WHERE account_number = $1 LIMIT 1;

SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR UPDATE;



-- name: ListAccounts :many
SELECT * FROM account
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :exec
UPDATE account
SET balance = $2
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

-- name: AddAccountBalance :one
UPDATE account
SET balance = balance + $2
WHERE id = $1
RETURNING *;
