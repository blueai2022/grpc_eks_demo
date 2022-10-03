-- name: CreateApiAccount :one
INSERT INTO api_accounts (
  username,
  is_active,
  service_type,
  plan_name,
  credit_balance,
  last_use_at,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetApiAccount :one
SELECT * FROM api_accounts
WHERE id = $1 LIMIT 1;

-- name: GetActiveApiAccount :one
SELECT * FROM api_accounts
WHERE 
    username = $1 AND
    is_active = true AND
    service_type = $2
LIMIT 1;

-- name: DebitApiAccountBalance :one
UPDATE api_accounts
SET credit_balance = credit_balance - 1
WHERE id = $1
RETURNING *;
