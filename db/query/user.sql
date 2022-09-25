-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  email,
  agency,
  app_contact,
  app_contact_email
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
