-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  username,
  email,
  secret_code
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: VerifyEmail :exec
UPDATE verify_emails 
SET is_used = true 
WHERE id = $1 
  AND secret_code = $2 
  AND is_used = false 
  AND expired_at > now();

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE id = $1 LIMIT 1;

-- name: ListVerifyEmails :many
SELECT * FROM verify_emails
WHERE username = $1
ORDER BY created_at DESC;

-- name: DeleteExpiredVerifyEmails :exec
DELETE FROM verify_emails
WHERE expired_at < now();

-- name: UpdateUserEmailVerified :exec
UPDATE users 
SET is_email_verified = $2 
WHERE username = $1;
